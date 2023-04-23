package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/Jiang-Gianni/go-docker-test/client"
	"github.com/Jiang-Gianni/go-docker-test/proto"
	// "github.com/Jiang-Gianni/go-docker-test/client"
)

func main() {

	/*
		client := client.New("http://localhost:3000")
		price, err := client.FetchPrice(context.Background(), "ET")
		if err!=nil {
			fmt.Println("Error ",err)
		}
		fmt.Printf("%+v\n", price)
		return

	*/

	var (
		jsonAddr = flag.String("json", ":3000", "listen address of the json transport")
		grpcAddr = flag.String("grpc", ":4000", "listen address of the grpc transport")
		//	grpcAddr = flag.String("grpc", "go-docker-test-production.up.railway.app:4000", "listen address of the grpc transport")
		svc = NewLogginService(NewMetricService(&priceFetcher{}))
		ctx = context.Background()
	)

	flag.Parse()

	fmt.Println("JSON address ", *jsonAddr)
	fmt.Println("grpc adderss ", *grpcAddr)

	grpcClient, err := client.NewGRPCClient(*grpcAddr)

	if err != nil {
		log.Fatal(err)
	}
	go func() {
		time.Sleep(3 * time.Second)
		resp, err := grpcClient.FetchPrice(ctx, &proto.PriceRequest{Ticker: "BTC"})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v\n", resp)
	}()

	go makeGRPCServerAndRun(*grpcAddr, svc)

	jsonServer := NewJSONAPIServer(*jsonAddr, svc)
	jsonServer.Run()

}
