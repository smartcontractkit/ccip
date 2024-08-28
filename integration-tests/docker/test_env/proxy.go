package test_env

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog"

	"github.com/smartcontractkit/chainlink-testing-framework/docker/test_env"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_onramp"
)

type proxy struct {
	targetHTTP     string
	targetWS       string
	port           int
	logger         zerolog.Logger
	chainID        string
	logFile        *os.File
	txFile         *os.File
	upgrader       websocket.Upgrader
	sendRawTxCount int
}

func newProxy(rpcProvider test_env.RpcProvider, chainID string) test_env.RpcProvider {
	logFile, err := os.Create(fmt.Sprintf("proxy_%s.log", chainID))
	if err != nil {
		log.Fatal("Failed to create log file:", err)
	}
	txFile, err := os.Create(fmt.Sprintf("tx_%s.json", chainID))
	if err != nil {
		log.Fatal("Failed to create tx file:", err)
	}
	p := &proxy{
		targetHTTP: rpcProvider.PublicHttpUrls()[0],
		targetWS:   rpcProvider.PublicWsUrls()[0],
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin:     func(_ *http.Request) bool { return true },
		},
		logger:  zerolog.New(logFile).With().Timestamp().Logger(),
		logFile: logFile,
		txFile:  txFile,
		chainID: chainID,
		port:    8080,
	}

	if chainID == "2337" {
		p.port = 8081
	}

	go func() {
		defer logFile.Close()
		defer txFile.Close()
		// flush log file every 5 seconds
		for {
			time.Sleep(5 * time.Second)
			err := logFile.Sync()
			if err != nil {
				p.logger.Error().Err(err).Msg("Failed to sync log file")
			}
		}
	}()

	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", p.handler)
		p.logger.Info().Str("Source", fmt.Sprintf("0.0.0.0:%d", p.port)).Str("HTTP Target", p.targetHTTP).Str("WS Target", p.targetWS).Msg("Starting Proxy Server")
		p.logger.Error().Err(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", p.port), mux)).Msg("Failed to start proxy server")
	}()

	return test_env.NewRPCProvider(
		[]string{fmt.Sprintf("http://host.docker.internal:%d/", p.port)},
		[]string{fmt.Sprintf("ws://host.docker.internal:%d/", p.port)},
		[]string{fmt.Sprintf("http://localhost:%d/", p.port)},
		[]string{fmt.Sprintf("ws://localhost:%d/", p.port)},
	)
}

func (p *proxy) handler(w http.ResponseWriter, r *http.Request) {
	// Check if it's an upgrade request for a WebSocket connection
	if websocket.IsWebSocketUpgrade(r) {
		p.handleWebSocket(w, r)
	} else {
		p.handleHttp(w, r)
	}
}

// HTTP proxy function
func (p *proxy) handleHttp(w http.ResponseWriter, r *http.Request) {
	targetURL, err := url.Parse(p.targetHTTP)
	if err != nil {
		log.Fatalf("Failed to parse target URL: %v", err)
	}

	r.Host = targetURL.Host
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		p.logger.Error().Err(err).Msg("Failed to read request body")
		return
	}
	r.Body = io.NopCloser(strings.NewReader(string(body)))

	p.logger.Info().Str("Body", string(body)).Msg("HTTP Request")

	// Write transactions to file
	var bodyArray []map[string]any
	var bodyMap map[string]any
	if string(body)[0] == '[' {
		err = json.Unmarshal(body, &bodyArray)
		if err != nil {
			p.logger.Error().Err(err).Str("Raw Body", string(body)).Msg("Failed to unmarshal request body")
		}
		for _, b := range bodyArray {
			if method, ok := b["method"]; ok {
				if isValidMethod(method.(string)) {
					_, err := p.txFile.WriteString(string(body) + ",\n")
					if err != nil {
						p.logger.Error().Err(err).Str("Body", string(body)).Msg("Failed to write transaction to file")
					}
				}
			}
		}
	} else {
		err = json.Unmarshal(body, &bodyMap)
		if err != nil {
			p.logger.Error().Err(err).Str("Raw Body", string(body)).Msg("Failed to unmarshal request body")
		}
		if method, ok := bodyMap["method"]; ok {
			bodyMap["timestamp"] = time.Now().String()
			if isValidMethod(method.(string)) {
				bodyMap["sendRawTxCount"] = p.sendRawTxCount
				p.sendRawTxCount++

				params := bodyMap["params"].([]any)
				paramsStr := strings.TrimPrefix(params[0].(string), "0x")
				txBytes, err := hex.DecodeString(paramsStr)
				if err != nil {
					p.logger.Error().Err(err).Str("Body", string(body)).Msg("Failed to decode transaction")
				}

				var tx types.Transaction
				err = rlp.DecodeBytes(txBytes, &tx)
				if err != nil {
					p.logger.Error().Err(err).Str("Body", string(body)).Msg("Failed to decode transaction")
				}
				bodyMap["to"] = tx.To().Hex()
				bodyMap["nonce"] = tx.Nonce()
				bodyMap["value"] = tx.Value()
				bodyMap["hash"] = tx.Hash().Hex()

				if len(tx.Data()) > 4 {
					var parsedABI abi.ABI
					if p.chainID == "1337" {
						parsedABI, err = abi.JSON(strings.NewReader(evm_2_evm_onramp.EVM2EVMOnRampABI))
					} else {
						parsedABI, err = abi.JSON(strings.NewReader(evm_2_evm_offramp.EVM2EVMOffRampABI))
					}
					if err != nil {
						p.logger.Error().Err(err).Str("Body", string(body)).Msg("Failed to parse ABI")
					}

					method, err := parsedABI.MethodById(tx.Data()[:4])
					if err != nil {
						p.logger.Error().Err(err).Str("Body", string(body)).Msg("Failed to get method by ID")
						bodyMap["method"] = err.Error()
					} else {
						bodyMap["method"] = method.Name
						inputs, err := method.Inputs.UnpackValues(tx.Data()[4:])
						if err != nil {
							p.logger.Error().Err(err).Str("Body", string(body)).Msg("Failed to unpack inputs")
							bodyMap["inputs"] = err.Error()
						} else {
							bodyMap["inputs"] = inputs
						}
					}
				}

				jsonStr, err := json.Marshal(bodyMap)
				if err != nil {
					p.logger.Error().Err(err).Str("Body", string(body)).Msg("Failed to marshal request body")
				}
				_, err = p.txFile.WriteString(string(jsonStr) + ",\n")
				if err != nil {
					p.logger.Error().Err(err).Str("Body", string(body)).Msg("Failed to write transaction to file")
				}
			}
		}
	}

	proxy.ServeHTTP(w, r)
}

func isValidMethod(meth string) bool {
	switch meth {
	case "eth_sendRawTransaction":
		return true
	case "eth_subscribe":
		return false
	case "eth_unsubscribe":
		return false
	case "eth_getLogs":
		return false
	case "eth_getBlockByNumber":
		return false
	case "eth_getBalance":
		return false
	case "eth_chainId":
		return false
	case "eth_getTransactionCount":
		return false
	case "web3_clientVersion":
		return false
	case "eth_getTransactionReceipt":
		return false
	case "eth_getBlockByHash":
		return false
	default:
		return false
	}
}

func (p *proxy) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	clientConn, err := p.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}
	defer clientConn.Close()

	// Handle close messages from the client
	clientConn.SetCloseHandler(func(code int, text string) error {
		p.logger.Debug().Int("Code", code).Str("Text", text).Msg("Client requested close")
		return clientConn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	})

	// Connect to the blockchain node's WebSocket endpoint
	nodeConn, _, err := websocket.DefaultDialer.Dial(p.targetWS, nil)
	if err != nil {
		p.logger.Error().Err(err).Msg("Failed to dial node WS")
		return
	}
	defer nodeConn.Close()

	// Start proxying WebSocket messages
	done := make(chan struct{})
	var once sync.Once
	go p.proxyWebSocketMessages(clientConn, nodeConn, done, &once)
	go p.proxyWebSocketMessages(nodeConn, clientConn, done, &once)

	// Wait until done
	<-done
	clientConn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseAbnormalClosure, ""))
	nodeConn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseAbnormalClosure, ""))
}

func (p *proxy) proxyWebSocketMessages(src, dest *websocket.Conn, done chan struct{}, once *sync.Once) {
	defer func() {
		once.Do(func() {
			close(done)
		})
	}()
	for {
		messageType, message, err := src.ReadMessage()
		if err != nil {
			p.logger.Error().Err(err).
				Int("Message Type", messageType).
				Str("RemoteAddr", src.RemoteAddr().String()).
				Str("Body", string(message)).
				Msg("Failed to read WS Request")
			return
		}

		p.logger.Info().Str("Message", string(message)).Msg("WS Request")

		// Forward the WebSocket message to the destination
		err = dest.WriteMessage(messageType, message)
		if err != nil {
			p.logger.Error().Err(err).
				Int("Message Type", messageType).
				Str("RemoteAddr", src.RemoteAddr().String()).
				Str("Body", string(message)).
				Msg("Failed to Write WS Request")
			return
		}
	}
}
