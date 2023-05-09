package main

import (
	"fmt"
	"google.golang.org/grpc"
)
import "github.com/authzed/grpcutil"

func main() {
	_, err := grpcutil.WithSystemCerts(grpcutil.SkipVerifyCA)
	if err != nil {
		fmt.Println("foo")
	}
	var opts []grpc.DialOption
	fmt.Printf("Hello, World! opts: %s", opts)
}
