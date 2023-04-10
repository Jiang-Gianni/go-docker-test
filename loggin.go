package main

import (
	"context"
	"fmt"
	"time"
)

type logginService struct {
	next PriceFetcher	
}

func NewLogginService(next PriceFetcher) PriceFetcher{
	return &logginService{
		next: next,
	}
}

func (s* logginService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	defer func(begin time.Time){
		fmt.Println(
			"requestID", ctx.Value("requestID"), "\n",
			"took", time.Since(begin), "\n",
			"err", err, "\n",
			"ticker", ticker, "\n",
			"price",price,
		)
	}(time.Now())
	return s.next.FetchPrice(ctx, ticker)	
}
