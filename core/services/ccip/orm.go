package ccip

import (
	"context"
	"fmt"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/sqlutil"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"
)

type GasPrice struct {
	SourceChainSelector uint64
	GasPrice            *assets.Wei
	CreatedAt           time.Time
}

type GasPriceUpdate struct {
	SourceChainSelector uint64
	GasPrice            *assets.Wei
}

type TokenPrice struct {
	TokenAddr  string
	TokenPrice *assets.Wei
	CreatedAt  time.Time
}

type TokenPriceUpdate struct {
	TokenAddr  string
	TokenPrice *assets.Wei
}

type ORM interface {
	GetGasPricesByDestChain(ctx context.Context, destChainSelector uint64) ([]GasPrice, error)
	GetTokenPricesByDestChain(ctx context.Context, destChainSelector uint64) ([]TokenPrice, error)

	InsertGasPricesForDestChain(ctx context.Context, destChainSelector uint64, gasPrices []GasPriceUpdate) error
	InsertTokenPricesForDestChain(ctx context.Context, destChainSelector uint64, tokenPrices []TokenPriceUpdate) error
}

type orm struct {
	ds sqlutil.DataSource
}

var _ ORM = (*orm)(nil)

func NewORM(ds sqlutil.DataSource) (ORM, error) {
	if ds == nil {
		return nil, fmt.Errorf("datasource to CCIP NewORM cannot be nil")
	}

	return &orm{
		ds: ds,
	}, nil
}

func (o *orm) GetGasPricesByDestChain(ctx context.Context, destChainSelector uint64) ([]GasPrice, error) {
	var gasPrices []GasPrice
	stmt := `
		SELECT source_chain_selector, gas_price, created_at
		FROM ccip.observed_gas_prices
		WHERE chain_selector = $1
		ORDER BY source_chain_selector, created_at DESC;
	`
	err := o.ds.SelectContext(ctx, &gasPrices, stmt, destChainSelector)
	if err != nil {
		return nil, err
	}

	return gasPrices, nil
}

func (o *orm) GetTokenPricesByDestChain(ctx context.Context, destChainSelector uint64) ([]TokenPrice, error) {
	var tokenPrices []TokenPrice
	stmt := `
		SELECT token_addr, token_price, created_at
		FROM ccip.observed_token_prices
		WHERE chain_selector = $1
		ORDER BY token_addr;
	`
	err := o.ds.SelectContext(ctx, &tokenPrices, stmt, destChainSelector)
	if err != nil {
		return nil, err
	}

	return tokenPrices, nil
}

func (o *orm) InsertGasPricesForDestChain(ctx context.Context, destChainSelector uint64, gasPrices []GasPriceUpdate) error {
	if len(gasPrices) == 0 {
		return nil
	}

	uniqueGasUpdates := make(map[string]GasPriceUpdate)
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
		ON CONFLICT (chain_selector, source_chain_selector)
		DO UPDATE SET gas_price = EXCLUDED.gas_price, created_at = EXCLUDED.created_at;`
	_, err := o.ds.NamedExecContext(ctx, stmt, insertData)
	if err != nil {
		err = fmt.Errorf("error inserting gas prices %w", err)
	}
	return err
}

func (o *orm) InsertTokenPricesForDestChain(ctx context.Context, destChainSelector uint64, tokenPrices []TokenPriceUpdate) error {
	if len(tokenPrices) == 0 {
		return nil
	}

	uniqueTokenPrices := make(map[string]TokenPriceUpdate)
	for _, tokenPrice := range tokenPrices {
		key := fmt.Sprintf("%s-%d", tokenPrice.TokenAddr, destChainSelector)
		uniqueTokenPrices[key] = tokenPrice
	}

	insertData := make([]map[string]interface{}, 0, len(uniqueTokenPrices))
	for _, price := range uniqueTokenPrices {
		insertData = append(insertData, map[string]interface{}{
			"chain_selector": destChainSelector,
			"token_addr":     price.TokenAddr,
			"token_price":    price.TokenPrice,
		})
	}

	stmt := `INSERT INTO ccip.observed_token_prices (chain_selector, token_addr, token_price, created_at)
		VALUES (:chain_selector, :token_addr, :token_price, statement_timestamp())
		ON CONFLICT (chain_selector, token_addr) 
		DO UPDATE SET token_price = EXCLUDED.token_price, created_at = EXCLUDED.created_at;`
	_, err := o.ds.NamedExecContext(ctx, stmt, insertData)
	if err != nil {
		err = fmt.Errorf("error inserting token prices %w", err)
	}
	return err
}
