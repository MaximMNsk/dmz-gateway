package main

import (
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type DmzServer struct {
	GRPC   *grpc.Server
	Logger zerolog.Logger
}

func main() {
}
