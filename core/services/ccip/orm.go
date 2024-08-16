package ccip

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/sqlutil"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

type GasPrice struct {
	SourceChainSelector uint64
	GasPrice            *assets.Wei
}

type TokenPrice struct {
	TokenAddr  string
	TokenPrice *assets.Wei
}

type ORM interface {
	GetGasPricesByDestChain(ctx context.Context, destChainSelector uint64) ([]GasPrice, error)
	GetTokenPricesByDestChain(ctx context.Context, destChainSelector uint64) ([]TokenPrice, error)

	UpsertGasPricesForDestChain(ctx context.Context, destChainSelector uint64, gasPrices []GasPrice) (int64, error)
	UpsertTokenPricesForDestChain(ctx context.Context, destChainSelector uint64, tokenPrices []TokenPrice, interval time.Duration) (int64, error)
}

type tokenPriceRow struct {
	TokenAddr  string
	TokenPrice *assets.Wei
	CreatedAt  time.Time
}

type orm struct {
	ds   sqlutil.DataSource
	lggr logger.Logger
}

var _ ORM = (*orm)(nil)

func NewORM(ds sqlutil.DataSource, lggr logger.Logger) (ORM, error) {
	if ds == nil {
		return nil, fmt.Errorf("datasource to CCIP NewORM cannot be nil")
	}

	return &orm{
		ds:   ds,
		lggr: lggr,
	}, nil
}

func (o *orm) GetGasPricesByDestChain(ctx context.Context, destChainSelector uint64) ([]GasPrice, error) {
	var gasPrices []GasPrice
	stmt := `
		SELECT source_chain_selector, gas_price
		FROM ccip.observed_gas_prices
		WHERE chain_selector = $1;
	`
	o.withAnalyze(ctx, "GetGasPricesByDestChain", stmt, destChainSelector)
	err := o.ds.SelectContext(ctx, &gasPrices, stmt, destChainSelector)
	if err != nil {
		return nil, err
	}

	return gasPrices, nil
}

func (o *orm) GetTokenPricesByDestChain(ctx context.Context, destChainSelector uint64) ([]TokenPrice, error) {
	var tokenPrices []TokenPrice
	stmt := `
		SELECT token_addr, token_price
		FROM ccip.observed_token_prices
		WHERE chain_selector = $1;
	`
	o.withAnalyze(ctx, "GetTokenPricesByDestChain", stmt, destChainSelector)
	err := o.ds.SelectContext(ctx, &tokenPrices, stmt, destChainSelector)
	if err != nil {
		return nil, err
	}
	return tokenPrices, nil
}

func (o *orm) UpsertGasPricesForDestChain(ctx context.Context, destChainSelector uint64, gasPrices []GasPrice) (int64, error) {
	if len(gasPrices) == 0 {
		return 0, nil
	}

	uniqueGasUpdates := make(map[string]GasPrice)
	for _, gasPrice := range gasPrices {
		key := fmt.Sprintf("%d-%d", gasPrice.SourceChainSelector, destChainSelector)
		uniqueGasUpdates[key] = gasPrice
	}

	insertData := make([]map[string]interface{}, 0, len(uniqueGasUpdates))
	for _, price := range uniqueGasUpdates {
		insertData = append(insertData, map[string]interface{}{
			"chain_selector":        destChainSelector,
			"source_chain_selector": price.SourceChainSelector,
			"gas_price":             price.GasPrice,
		})
	}

	stmt := `INSERT INTO ccip.observed_gas_prices (chain_selector, source_chain_selector, gas_price, created_at)
		VALUES (:chain_selector, :source_chain_selector, :gas_price, statement_timestamp())
		ON CONFLICT (source_chain_selector, chain_selector)
		DO UPDATE SET gas_price = EXCLUDED.gas_price, created_at = EXCLUDED.created_at;`

	result, err := o.ds.NamedExecContext(ctx, stmt, insertData)
	if err != nil {
		return 0, fmt.Errorf("error inserting gas prices %w", err)
	}
	return result.RowsAffected()
}

// UpsertTokenPricesForDestChain inserts or updates only relevant token prices.
// In order to reduce locking an unnecessary writes to the table, we start with fetching current prices.
// If price for a token doesn't change or was updated recently we don't include that token to the upsert query.
// We don't run in TX intentionally, because we don't want to lock the table and conflicts are resolved on the insert level
func (o *orm) UpsertTokenPricesForDestChain(ctx context.Context, destChainSelector uint64, tokenPrices []TokenPrice, interval time.Duration) (int64, error) {
	if len(tokenPrices) == 0 {
		return 0, nil
	}

	tokensToUpdate, err := o.pickOnlyRelevantTokensForUpdate(ctx, destChainSelector, tokenPrices, interval)
	if err != nil || len(tokensToUpdate) == 0 {
		return 0, err
	}

	insertData := make([]map[string]interface{}, 0, len(tokensToUpdate))
	for _, price := range tokensToUpdate {
		insertData = append(insertData, map[string]interface{}{
			"chain_selector": destChainSelector,
			"token_addr":     price.TokenAddr,
			"token_price":    price.TokenPrice,
		})
	}

	stmt := `INSERT INTO ccip.observed_token_prices (chain_selector, token_addr, token_price, created_at)
		VALUES (:chain_selector, :token_addr, :token_price, statement_timestamp())
		ON CONFLICT (token_addr, chain_selector) 
		DO UPDATE SET token_price = EXCLUDED.token_price, created_at = EXCLUDED.created_at;`
	result, err := o.ds.NamedExecContext(ctx, stmt, insertData)
	if err != nil {
		return 0, fmt.Errorf("error inserting token prices %w", err)
	}
	return result.RowsAffected()
}

// pickOnlyRelevantTokensForUpdate returns only tokens that need to be updated. Multiple jobs can be updating the same tokens,
// in order to reduce table locking and redundant upserts we start with reading the table and checking which tokens are eligible for update.
// A token is eligible for update when:
// * price has changed and created_at date is older than the interval
// * it's not present in the result set from the db query (e.g. it's a new token)
// Therefore if there are no price changes for a single token we won't update the state of the database
func (o *orm) pickOnlyRelevantTokensForUpdate(
	ctx context.Context,
	destChainSelector uint64,
	tokenPrices []TokenPrice,
	interval time.Duration,
) ([]TokenPrice, error) {
	stmt := `
		SELECT token_addr, token_price, created_at
		FROM ccip.observed_token_prices
		WHERE chain_selector = $1 and token_addr = ANY($2);
	`

	o.withAnalyze(ctx, "pickOnlyRelevantTokensForUpdate", stmt, destChainSelector, tokenAddrsToBytes(tokenPrices))

	var dbTokenPrices []tokenPriceRow
	if err := o.ds.SelectContext(ctx, &dbTokenPrices, stmt, destChainSelector, tokenAddrsToBytes(tokenPrices)); err != nil {
		return nil, err
	}

	updateThrehsold := time.Now().Add(-interval)
	tokenPricesByAddr := toTokensByAddress(tokenPrices)
	dbTokenPricesByAddr := toTokenRowsByAddress(dbTokenPrices)
	tokenPricesToUpdate := make([]TokenPrice, 0, len(tokenPrices))
	for addr, price := range tokenPricesByAddr {
		dbToken, ok := dbTokenPricesByAddr[addr]
		eligibleForUpdate := false
		if !ok ||
			(dbToken.CreatedAt.Before(updateThrehsold) && !dbToken.TokenPrice.Equal(price)) {
			tokenPricesToUpdate = append(tokenPricesToUpdate, TokenPrice{addr, price})
			eligibleForUpdate = true
		}
		o.lggr.Debugw("Token price database update", "eligibleForUpdate", eligibleForUpdate, "token", addr, "price", price)
	}
	return tokenPricesToUpdate, nil
}

func toTokensByAddress(tokens []TokenPrice) map[string]*assets.Wei {
	tokensByAddr := make(map[string]*assets.Wei, len(tokens))
	for _, tk := range tokens {
		tokensByAddr[tk.TokenAddr] = tk.TokenPrice
	}
	return tokensByAddr
}

func toTokenRowsByAddress(tokens []tokenPriceRow) map[string]tokenPriceRow {
	tokensByAddr := make(map[string]tokenPriceRow, len(tokens))
	for _, tk := range tokens {
		tokensByAddr[tk.TokenAddr] = tk
	}
	return tokensByAddr
}

func tokenAddrsToBytes(tokens []TokenPrice) [][]byte {
	addrs := make([][]byte, 0, len(tokens))
	for _, tk := range tokens {
		addrs = append(addrs, []byte(tk.TokenAddr))
	}
	return addrs
}

func (o *orm) withAnalyze(ctx context.Context, queryName string, query string, args ...interface{}) {
	query = "EXPLAIN (ANALYZE, BUFFERS) " + query

	var response []string
	err := o.ds.SelectContext(ctx, &response, query, args...)
	if err != nil {
		return
	}
	if len(response) > 0 {
		o.lggr.Infow("Analyze query", "query", queryName, "response", strings.Join(response, "\n"))
	}
}
