package config

import (
	"fmt"
)

type Config struct {
	GrpcEndpoint string
	HttpEndpoint string
}

func init() {
	fmt.Printf(" Init config \n")

}
