package ccip

import (
	"context"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/smartcontractkit/chainlink-common/pkg/sqlutil"
)

var (
	sqlLatencyBuckets = []float64{
		float64(1 * time.Millisecond),
		float64(5 * time.Millisecond),
		float64(10 * time.Millisecond),
		float64(20 * time.Millisecond),
		float64(30 * time.Millisecond),
		float64(40 * time.Millisecond),
		float64(50 * time.Millisecond),
		float64(60 * time.Millisecond),
		float64(70 * time.Millisecond),
		float64(80 * time.Millisecond),
		float64(90 * time.Millisecond),
		float64(100 * time.Millisecond),
		float64(200 * time.Millisecond),
		float64(300 * time.Millisecond),
		float64(400 * time.Millisecond),
		float64(500 * time.Millisecond),
		float64(750 * time.Millisecond),
		float64(1 * time.Second),
		float64(2 * time.Second),
		float64(5 * time.Second),
	}
	ccipQueryDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "ccip_orm_query_duration",
		Buckets: sqlLatencyBuckets,
	}, []string{"query", "destChainSelector"})
	ccipQueryDatasets = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "ccip_orm_dataset_size",
	}, []string{"query", "destChainSelector"})
	ccipOrmRowsInserted = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "ccip_orm_rows_affected",
	}, []string{"query", "destChainSelector"})
)

type observedORM struct {
	ORM
	queryDuration *prometheus.HistogramVec
	datasetSize   *prometheus.GaugeVec
	rowsInserted  *prometheus.GaugeVec
}

var _ ORM = (*observedORM)(nil)

func NewObservedORM(ds sqlutil.DataSource) (*observedORM, error) {
	delegate, err := NewORM(ds)
	if err != nil {
		return nil, err
	}

	return &observedORM{
		ORM:           delegate,
		queryDuration: ccipQueryDuration,
		datasetSize:   ccipQueryDatasets,
		rowsInserted:  ccipOrmRowsInserted,
	}, nil
}

func (o *observedORM) GetGasPricesByDestChain(ctx context.Context, destChainSelector uint64) ([]GasPrice, error) {
	return withObservedQueryAndResults(o, "GetGasPricesByDestChain", destChainSelector, func() ([]GasPrice, error) {
		return o.ORM.GetGasPricesByDestChain(ctx, destChainSelector)
	})
}

func (o *observedORM) GetTokenPricesByDestChain(ctx context.Context, destChainSelector uint64) ([]TokenPrice, error) {
	return withObservedQueryAndResults(o, "GetTokenPricesByDestChain", destChainSelector, func() ([]TokenPrice, error) {
		return o.ORM.GetTokenPricesByDestChain(ctx, destChainSelector)
	})
}

func (o *observedORM) InsertGasPricesForDestChain(ctx context.Context, destChainSelector uint64, jobId int32, gasPrices []GasPriceUpdate) (int64, error) {
	return withObservedQueryAndRowsAffected(o, "InsertGasPricesForDestChain", destChainSelector, func() (int64, error) {
		return o.ORM.InsertGasPricesForDestChain(ctx, destChainSelector, jobId, gasPrices)
	})
}

func (o *observedORM) InsertTokenPricesForDestChain(ctx context.Context, destChainSelector uint64, jobId int32, tokenPrices []TokenPriceUpdate) (int64, error) {
	return withObservedQueryAndRowsAffected(o, "InsertTokenPricesForDestChain", destChainSelector, func() (int64, error) {
		return o.ORM.InsertTokenPricesForDestChain(ctx, destChainSelector, jobId, tokenPrices)
	})
}

func (o *observedORM) ClearGasPricesByDestChain(ctx context.Context, destChainSelector uint64, expireSec int) (int64, error) {
	return withObservedQueryAndRowsAffected(o, "ClearGasPricesByDestChain", destChainSelector, func() (int64, error) {
		return o.ORM.ClearGasPricesByDestChain(ctx, destChainSelector, expireSec)
	})
}

func (o *observedORM) ClearTokenPricesByDestChain(ctx context.Context, destChainSelector uint64, expireSec int) (int64, error) {
	return withObservedQueryAndRowsAffected(o, "ClearTokenPricesByDestChain", destChainSelector, func() (int64, error) {
		return o.ORM.ClearTokenPricesByDestChain(ctx, destChainSelector, expireSec)
	})
}

func withObservedQueryAndRowsAffected(o *observedORM, queryName string, chainSelector uint64, query func() (int64, error)) (int64, error) {
	rowsAffected, err := withObservedQuery(o, queryName, chainSelector, query)
	if err == nil {
		o.rowsInserted.
			WithLabelValues(queryName, strconv.FormatUint(chainSelector, 10)).
			Set(float64(rowsAffected))
	}
	return rowsAffected, err
}

func withObservedQueryAndResults[T any](o *observedORM, queryName string, chainSelector uint64, query func() ([]T, error)) ([]T, error) {
	results, err := withObservedQuery(o, queryName, chainSelector, query)
	if err == nil {
		o.datasetSize.
			WithLabelValues(queryName, strconv.FormatUint(chainSelector, 10)).
			Set(float64(len(results)))
	}
	return results, err
}

func withObservedQuery[T any](o *observedORM, queryName string, chainSelector uint64, query func() (T, error)) (T, error) {
	queryStarted := time.Now()
	defer func() {
		o.queryDuration.
			WithLabelValues(queryName, strconv.FormatUint(chainSelector, 10)).
			Observe(float64(time.Since(queryStarted)))
	}()
	return query()
}
