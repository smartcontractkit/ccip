package observability

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/mocks"
)

type MethodCall struct {
	MethodName string
	Arguments  []interface{}
	Returns    []interface{}
}

// TestOnRampObservedMethods tests that all methods of OnRampReader are observed by a wrapper.
// It uses the runtime to detect if the call stack contains the wrapper class.
func TestOnRampObservedMethods(t *testing.T) {

	// The class expected to override the observed methods.
	expectedWrapper := "core/services/ocr2/plugins/ccip/internal/observability.ObservedOnRampReader"

	// Methods not expected to be observed.
	excludedMethods := []string{
		"Address",
		"Close",
	}

	// Defines the overridden method calls to test.
	// Not defining a non-excluded method here will cause the test to fail with an explicit error.
	methodCalls := make(map[string]MethodCall)
	methodCalls["GetDynamicConfig"] = MethodCall{
		MethodName: "GetDynamicConfig",
		Arguments:  []interface{}{testutils.Context(t)},
		Returns:    []interface{}{cciptypes.OnRampDynamicConfig{}, nil},
	}
	methodCalls["GetSendRequestsBetweenSeqNums"] = MethodCall{
		MethodName: "GetSendRequestsBetweenSeqNums",
		Arguments:  []interface{}{testutils.Context(t), uint64(0), uint64(100), true},
		Returns:    []interface{}{nil, nil},
	}
	methodCalls["IsSourceChainHealthy"] = MethodCall{
		MethodName: "IsSourceChainHealthy",
		Arguments:  []interface{}{testutils.Context(t)},
		Returns:    []interface{}{false, nil},
	}
	methodCalls["IsSourceCursed"] = MethodCall{
		MethodName: "IsSourceCursed",
		Arguments:  []interface{}{testutils.Context(t)},
		Returns:    []interface{}{false, nil},
	}
	methodCalls["RouterAddress"] = MethodCall{
		MethodName: "RouterAddress",
		Arguments:  []interface{}{testutils.Context(t)},
		Returns:    []interface{}{cciptypes.Address("0x0"), nil},
	}
	methodCalls["SourcePriceRegistryAddress"] = MethodCall{
		MethodName: "SourcePriceRegistryAddress",
		Arguments:  []interface{}{testutils.Context(t)},
		Returns:    []interface{}{cciptypes.Address("0x0"), nil},
	}

	// Build actual observed object.
	labels = []string{"evmChainID", "plugin", "reader", "function", "success"}
	ph := promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "test_histogram",
	}, labels)
	pg := promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "test_gauge",
	}, labels)
	metric := metricDetails{
		interactionDuration: ph,
		resultSetSize:       pg,
		pluginName:          "test plugin",
		readerName:          "test reader",
		chainId:             1337,
	}
	reader := mocks.NewOnRampReader(t)
	observed := ObservedOnRampReader{reader, metric}

	// Test each method defined in the embedded type.
	observedType := reflect.TypeOf(observed)
	for i := 0; i < observedType.NumMethod(); i++ {

		method := observedType.Method(i)
		t.Run(fmt.Sprintf("observability_wrapper_%s", method.Name), func(t *testing.T) {

			fmt.Printf("Testing method %s\n", method.Name)

			// Skip excluded methods.
			for _, em := range excludedMethods {
				if method.Name == em {
					fmt.Printf("skipping ignored method %s\n", method.Name)
					return
				}
			}

			// Retrieve method call from definition (fail if not present).
			mc := methodCalls[method.Name]
			if mc.MethodName == "" {
				assert.Fail(t, fmt.Sprintf("method %s not defined in methodCalls, please define it or exclude it.", method.Name))
				return
			}

			// Defines mocked wrapped method behavior.
			reader.On(mc.MethodName, mc.Arguments...).Maybe().Return(mc.Returns...).Run(func(args mock.Arguments) {
				var observed = false
				for i := 0; i < 8; i++ {
					pc, _, line, _ := runtime.Caller(i)
					f := runtime.FuncForPC(pc)
					//fmt.Printf("caller %d is %s file %v, line %d\n", i, f.Name(), file, line)
					if strings.Contains(f.Name(), expectedWrapper) {
						fmt.Printf("Method observed by wrapper at line %d\n", line)
						observed = true
						break
					}
				}
				assert.True(t, observed, fmt.Sprintf("method %s not observed by wrapper. Please implement or add to excluded list.", mc.MethodName))
			})

			// Perform call on observed object.
			fmt.Printf("will call %s (%d parameters)\n", mc.MethodName, len(mc.Arguments))
			_, found := reflect.TypeOf(&observed).MethodByName(mc.MethodName)
			if !found {
				assert.Fail(t, fmt.Sprintf("method %s not found", mc.MethodName))
				return
			}
			methodc := reflect.ValueOf(&observed).MethodByName(mc.MethodName)

			// Build call arguments from mc.Arguments.
			callParams := make([]reflect.Value, len(mc.Arguments))
			for i, arg := range mc.Arguments {
				callParams[i] = reflect.ValueOf(arg)
			}
			methodc.Call(callParams)
		})

	}
}
