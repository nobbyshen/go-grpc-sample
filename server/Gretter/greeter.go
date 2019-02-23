package greeter

import (
	"fmt"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"go-grpc-sample/server/pb"
)

// Greeter is GRRc Service to implement proto interface
type Greeter struct {
}

// SayHello implements in Greeter
func (receiver *Greeter) SayHello(ctx context.Context, in *pbSample.HelloRequest) (*pbSample.HelloReply, error) {
	fmt.Printf("input value:%s \n", in.Name)
	reply := pbSample.HelloReply{Name: "reply input:"+in.Name}
	return &reply, nil
}

// GetApps blabla
func (receiver *Greeter) GetApps(ctx context.Context, in *empty.Empty) (*pbSample.HelloReply, error) {
	reply := pbSample.HelloReply{Name: "result"}
	return &reply, nil
}
