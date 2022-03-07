package grpc

import (
	auth "github.com/geomethu/afitech_svc_auth/pkg/rpc/pb"
	"github.com/geomethu/afitech_svc_communication/internal/delivery/grpc/v1/pb"
	"google.golang.org/grpc"
)

type Handler struct {
	auth auth.AuthServiceClient
}

func NewHandler() *Handler {
	return &Handler{

	}
}

func (h *Handler) RegisterHandlers(srv *grpc.Server) {
	pb.RegisterCommunicationServiceServer(srv, SmsHandler{})
} 