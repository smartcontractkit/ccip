package validation

import (
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

type counter[T any] struct {
	data  T
	count int
}

type Validator[T any] interface {
	AddReport(data T) error
	GetValidatedReports() ([]T, error)
}

// validator is a helper object to validate reports for a single chain.
// It keeps track of all reports and determines if they observations are consistent
// with one another and whether they meet the required fChain threshold.
type validator[T any] struct {
	minObservation int
	cache          map[cciptypes.Bytes32]*counter[T]
	idFunc         func(T) [32]byte
}

func NewValidator[T any](min int, idFunc func(T) [32]byte) *validator[T] {
	return &validator[T]{
		minObservation: min,
		cache:          make(map[cciptypes.Bytes32]*counter[T]),
		idFunc:         idFunc,
	}
}

func (cv *validator[T]) AddReport(data T) error {
	//id := sha3.Sum256(data.ToBytes())
	id := cv.idFunc(data)
	if _, ok := cv.cache[id]; ok {
		cv.cache[id].count++
	} else {
		cv.cache[id] = &counter[T]{data: data, count: 1}
	}
	return nil
}

func (cv *validator[T]) GetValidatedReports() ([]T, error) {
	var validated []T
	for _, rc := range cv.cache {
		if rc.count >= cv.minObservation {
			validated = append(validated, rc.data)
		}
	}
	return validated, nil
}
