package main

import (
	"context"
	"fmt"
)

type metricService struct {
	next PriceFetcher
}

func NewMetricService(next PriceFetcher) PriceFetcher{
	return &metricService{
		next: next,
	}
}

func (s* metricService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	//your metrics
	fmt.Println("Executing metrics")
	return s.next.FetchPrice(ctx, ticker)
}
