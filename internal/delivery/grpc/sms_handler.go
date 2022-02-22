package grpc

import "github.com/geomethu/afitech_base_svc_communication/internal/delivery/grpc/v1/pb"

type SmsHandler struct {
	
	pb.UnimplementedCommunicationServiceServer
}

func NewSmsHandler() *SmsHandler {
	return &SmsHandler{

	}
}