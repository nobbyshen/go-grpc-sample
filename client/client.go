package main

import (
	"context"
	"fmt"
	"go-grpc-sample/client/pb"
	"time"

	"google.golang.org/grpc"
)

const address = "localhost:8000"

func main() {
	fmt.Printf("This is GRPC client! \n")

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Failed to connect:%v", err)
	}
	defer conn.Close()
	client := pbSample.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &pbSample.HelloRequest{Name: "hello"})
	if err != nil {
		fmt.Printf("Failed to call sayHello")
	}
	fmt.Printf("Greeting: %s \n", resp.Name)

	math := pbSample.NewMyMathClient(conn)

	resp2, err := math.SquareRoot(ctx, &pbSample.MathInput{
		Para1:64,
	})
	if err != nil {
		fmt.Printf("Failed to call square root")
	}
	fmt.Printf("Squre Root of 64 = %v\n", resp2.Result)
	conn.Close()
}
