package main

import (
	"fmt"
	"google.golang.org/grpc"
)
import "github.com/authzed/grpcutil"

func main() {
	sysCertOption, err := grpcutil.WithSystemCerts(grpcutil.SkipVerifyCA)

	var opts []grpc.DialOption
	fmt.Println("Hello, World!")
}
