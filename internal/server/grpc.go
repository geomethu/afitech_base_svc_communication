package server

import (
	"github.com/geomethu/afitech_base_svc_communication/config"
	"google.golang.org/grpc"
)

//GRPC server structure
type GRPCServer struct { Server *grpc.Server }

func NewGRPCServer(cfg*config.Config) *GRPCServer {
	opts := []grpc.ServerOption{}
	return &GRPCServer{ Server: grpc.NewServer(opts...)}
}