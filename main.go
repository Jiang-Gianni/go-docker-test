package main

import (
//	"context"
	"flag"
	"fmt"

//	"github.com/Jiang-Gianni/go-docker-test/client"
)

func main(){

	/*
	client := client.New("http://localhost:3000")
	price, err := client.FetchPrice(context.Background(), "ET")
	if err!=nil {
		fmt.Println("Error ",err)
	}
	fmt.Printf("%+v\n", price)
	return

	*/ 

	listenAddr := flag.String("listenaddr", ":3000", "listen address the server is running")
	flag.Parse()
	svc := NewLogginService(NewMetricService(&priceFetcher{}))

	server := NewJSONAPIServer(*listenAddr, svc)
	server.Run()
	fmt.Println("Hello")
}
