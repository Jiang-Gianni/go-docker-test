package main

import (
	"context"
	"fmt"
)

// PriceFetcher is an interface that can fetch a price
type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

// priceFetcher implements the PriceFetcher interface
type priceFetcher struct{}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error){
	return  MockPriceFetcher(ctx, ticker)
}

var priceMocks = map[string]float64{
	"BTC": 20.0,
	"ETH": 25.0,
	"GG": 100.0,
}

func MockPriceFetcher(ctx context.Context, ticker string) (float64, error){
	price, ok := priceMocks[ticker]
	if !ok{
		return price, fmt.Errorf("The given ticker (%s) is not supported",ticker)
	}
	return price, nil
}

 
