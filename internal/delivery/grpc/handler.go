package grpc

import (
	"github.com/geomethu/afitech_base_svc_communication/internal/delivery/grpc/v1/pb"
	"google.golang.org/grpc"
)

type Handler struct {

}

func NewHandler() *Handler {
	return &Handler{

	}
}

func (h *Handler) RegisterHandlers(srv *grpc.Server) {
	pb.RegisterCommunicationServiceServer(srv, SmsHandler{})
} 