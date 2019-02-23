package api

import (
	"fmt"
	"go-grpc-sample/server/config"
	"go-grpc-sample/server/Gretter"
	"go-grpc-sample/server/MyMath"
	"go-grpc-sample/server/pb"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//NewGrpcServer : create grpc service
func NewGrpcServer(conf config.Config, wg *sync.WaitGroup) (*grpc.Server, error) {
	lis, err := net.Listen("tcp", conf.GrpcEndpoint)
	if err != nil {
		fmt.Printf("failed to initializa TCP listen: %v", err)
		return nil, err
	}
	// defer lis.Close()

	server := grpc.NewServer()
	g := greeter.Greeter{}
	m := mymath.MyMath{}
	pbSample.RegisterGreeterServer(server, &g)
	pbSample.RegisterMyMathServer(server, &m)
	reflection.Register(server)
	fmt.Printf("GRPC serve(%s) ... \n", conf.GrpcEndpoint)
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := server.Serve(lis); err != nil {
			fmt.Printf("Failed to serve grpc: %v\n", err)
		}

		fmt.Printf("end of grpc serve go routine\n")
	}()

	return server, nil
}
