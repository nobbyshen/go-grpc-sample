package main

import (
	"fmt"
	"go-grpc-sample/server/api"
	"go-grpc-sample/server/config"
	"sync"
)

func main() {
	conf := config.Config{
		GrpcEndpoint: ":8000",
		HttpEndpoint: ":8001",
	}
	
	wg := &sync.WaitGroup{}
	_, err := api.NewGrpcServer(conf, wg)
	if err != nil {
		fmt.Printf("Failed to create grcp serve")
	}
	_, err = api.NewHTTPServer(conf, wg)
	if err != nil {
		fmt.Printf("Failed to create http serve")
	}
	wg.Wait()
	fmt.Println("Server exited")
}
