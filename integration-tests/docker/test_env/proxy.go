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
	blockingMsgIDs map[string]time.Time
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
			ReadBufferSize:  2048,
			WriteBufferSize: 2048,
			CheckOrigin:     func(_ *http.Request) bool { return true },
		},
		logger:         zerolog.New(logFile).With().Timestamp().Logger(),
		logFile:        logFile,
		txFile:         txFile,
		chainID:        chainID,
		blockingMsgIDs: map[string]time.Time{},
		port:           8080,
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
		mux.HandleFunc("/msgID", p.newMsgID)
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

func (p *proxy) newMsgID(w http.ResponseWriter, r *http.Request) {
	msgID := r.URL.Query().Get("msgID")
	p.blockingMsgIDs[msgID] = time.Now().Add(8 * time.Minute)
	p.logger.Info().Str("msgID", msgID).Msg("Added msgID to block list")
	w.WriteHeader(http.StatusOK)
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

	toFilter, err := p.filterTxs(body)
	if err != nil {
		p.logger.Error().Err(err).Msg("Failed to filter transactions")
		return
	}

	if toFilter {
		w.WriteHeader(http.StatusOK)
		return
	}
	proxy.ServeHTTP(w, r)
}

func (p *proxy) filterTxs(body []byte) (bool, error) {
	// Write transactions to file
	if p.chainID != "2337" {
		return false, nil
	}
	var bodyMap map[string]any
	if string(body)[0] != '[' {
		err := json.Unmarshal(body, &bodyMap)
		if err != nil {
			return false, fmt.Errorf("Failed to unmarshal body: %w", err)
		}
		bodyMap["timestamp"] = time.Now().Format(time.TimeOnly)
		bodyMap["type"] = "http"
		if method, ok := bodyMap["method"]; ok {
			if isValidMethod(method.(string)) {
				params := bodyMap["params"].([]any)
				paramsStr := ""
				for _, param := range params {
					switch param.(type) {
					case string:
						paramsStr += strings.TrimPrefix(param.(string), "0x")
					case map[string]any:
						for _, v := range param.(map[string]any) {
							paramsStr += v.(string)
						}
					}
				}

				txBytes, err := hex.DecodeString(paramsStr)
				if err != nil {
					return false, fmt.Errorf("Failed to decode transaction bytes: %w", err)
				}

				var tx types.Transaction
				err = rlp.DecodeBytes(txBytes, &tx)
				if err != nil {
					return false, fmt.Errorf("Failed to decode transaction: %w", err)
				}
				bodyMap["to"] = tx.To().Hex()
				bodyMap["nonce"] = tx.Nonce()
				bodyMap["value"] = tx.Value()
				bodyMap["hash"] = tx.Hash().Hex()

				if len(tx.Data()) > 4 {
					parsedABI, err := abi.JSON(strings.NewReader(evm_2_evm_offramp.EVM2EVMOffRampABI))
					if err != nil {
						return false, fmt.Errorf("Failed to parse ABI: %w", err)
					}
					method, err := parsedABI.MethodById(tx.Data()[:4])
					if err != nil {
						bodyMap["method"] = err.Error()
					} else {
						bodyMap["method"] = method.Name
					}

					dataString := hex.EncodeToString(tx.Data())
					blockedMsgIDs := []string{}
					for id := range p.blockingMsgIDs {
						if strings.Contains(dataString, id) {
							blockedMsgIDs = append(blockedMsgIDs, id)
						}
						if len(blockedMsgIDs) > 0 {
							bodyMap["blockedMsgIDs"] = blockedMsgIDs
							bodyMap["progressBlockingMsgIDs"] = p.blockingMsgIDs
						}
					}
				}

				jsonStr, err := json.Marshal(bodyMap)
				if err != nil {
					return false, fmt.Errorf("Failed to marshal body map: %w", err)
				}
				_, err = p.txFile.WriteString(string(jsonStr) + ",\n")
				if err != nil {
					return false, fmt.Errorf("Failed to write transaction to file: %w", err)
				}
			}
		}
		if ids, ok := bodyMap["blockedMsgIDs"]; ok {
			stringIDs := ids.([]string)
			for _, id := range stringIDs {
				if time.Now().Before(p.blockingMsgIDs[id]) {
					return true, nil
				}
			}
		}
	}
	return false, nil
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
	clientConn.SetCloseHandler(func(_ int, _ string) error {
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

		var messageMap map[string]any
		err = json.Unmarshal(message, &messageMap)
		if err != nil {
			return
		}
		if method, ok := messageMap["method"]; ok {
			if isValidMethod(method.(string)) {
				messageMap["timestamp"] = time.Now().Format(time.TimeOnly)
				messageMap["type"] = "ws"
				jsonStr, err := json.Marshal(messageMap)
				if err != nil {
					return
				}
				_, err = p.txFile.WriteString(string(jsonStr) + ",\n")
				if err != nil {
					return
				}
			}
		}

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
