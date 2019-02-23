package mymath
import (
	"fmt"
	"math"
	"golang.org/x/net/context"
	"go-grpc-sample/server/pb"
)

type MyMath struct {
}

func (receiver *MyMath) SquareRoot(ctx context.Context, in *pbSample.MathInput) (*pbSample.MathResult, error) {
	fmt.Printf("Input square root of %v", in.Para1)
	var data float64
	data = float64(in.Para1)
	root := math.Sqrt(data) 
	ret := pbSample.MathResult{
		Result: float32(root),
	}
	return &ret, nil
}
