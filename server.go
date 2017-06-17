// +build server

package main

import (
	"fmt"
	"net"
	"os"

	"github.com/codedellemc/libstorage-mock/csi"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()

	csi.RegisterControllerServer(grpcServer, &ctrlrServer{})
	csi.RegisterIdentityServer(grpcServer, &identServer{})
	csi.RegisterNodeServer(grpcServer, &nodeServer{})
	csi.RegisterLibStorageServer(grpcServer, &libsServer{})

	protoAddr := os.Getenv("CSI_ENDPOINT")
	if protoAddr == "" {
		protoAddr = "tcp://127.0.0.1:8080"
	}
	proto, addr, err := parseProtoAddr(protoAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: invalid endpoint: %s\n", protoAddr)
		os.Exit(1)
	}

	l, err := net.Listen(proto, addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to listen: %v\n", err)
		os.Exit(1)
	}

	if err := grpcServer.Serve(l); err != nil {
		fmt.Fprintf(os.Stderr, "error: grpc failed: %v\n", err)
		os.Exit(1)
	}
}
