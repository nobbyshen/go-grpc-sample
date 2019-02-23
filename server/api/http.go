package api

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"go-grpc-sample/server/config"
	"go-grpc-sample/server/pb"
	"net"
	"net/http"
	"sync"
)

// NewHTTPServer is to serve http request
func NewHTTPServer(conf config.Config, wg *sync.WaitGroup) (*http.Server, error) {
	ctx := context.Background()
	// Don't defer, due to ...
	// ctx, cancel = context.WithCancel(ctx)
	// defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	
	err := pbSample.RegisterGreeterHandlerFromEndpoint(ctx, mux, conf.GrpcEndpoint, opts)
	if err != nil {
		fmt.Printf("Failed to register greeter to http")
		return nil, err
	}
	err = pbSample.RegisterMyMathHandlerFromEndpoint(ctx, mux, conf.GrpcEndpoint, opts)
	if err != nil {
		fmt.Printf("Failed to register mymath to http")
		return nil, err
	}
	server := http.Server{
		Addr:    conf.HttpEndpoint,
		Handler: mux,
	}
	lis, err := net.Listen("tcp", conf.HttpEndpoint)
	if err != nil {
		fmt.Printf("failed to initializa TCP listen: %v", err)
		return nil, err
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Printf("Http serve(%s) ...\n", conf.HttpEndpoint)
		if err := server.Serve(lis); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Failed to serve http: %v", err)
		}
	}()

	return &server, nil
}
