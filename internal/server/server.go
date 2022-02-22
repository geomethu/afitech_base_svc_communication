package server

import (
	"net"

	"github.com/geomethu/afitech_svc_communication/config"
	"github.com/geomethu/afitech_svc_communication/internal/delivery/grpc"
	"github.com/geomethu/afitech_svc_communication/pkg/log"
	"go.uber.org/zap"
)

type Server struct {
	listener   *net.Listener
	grpc *GRPCServer
	handler    *grpc.Handler
}

//Create a new server
func NewServer(cfg *config.Config, handler *grpc.Handler) *Server {

	// Server address.
	address := cfg.Server.CommSvcServer.Host + ":" + cfg.Server.CommSvcServer.Port

	return &Server{
		listener:   NewTCPServer(address),
		grpc: NewGRPCServer(cfg),
		handler:    handler,
	}
}

// Creating a new tpc server.
func NewTCPServer(address string) *net.Listener {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		//log.Fatal().Msgf("error creating new tcp server: %s", err.Error())
		log.Logger.Fatal("communnication service", zap.String("NewTCPServer", "error creating new tcp server"))
		return nil
	}

	return &lis
}

// Run grpc server.
func (s *Server) Run() {
	log.Logger.Debug("Running gRPC server...")

	// Registration: grpc handlers.
	s.handler.RegisterHandlers(s.grpc.Server)

	if err := s.grpc.Server.Serve(*s.listener); err != nil {
		log.Logger.Fatal("communication service", zap.String("Run()", "error running grpc server"))
	}
}

// Stop grpc server.
func (s *Server) Stop() {
	log.Logger.Info("Stopping gRPC server...")
	s.grpc.Server.Stop()
}
