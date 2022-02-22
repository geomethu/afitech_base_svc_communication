package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/geomethu/afitech/svc_communication/config"
	"github.com/geomethu/afitech/svc_communication/internal/delivery/grpc"
	"github.com/geomethu/afitech/svc_communication/internal/server"
	"github.com/geomethu/afitech/svc_communication/pkg/log"
)


func Run(cfg *config.Config) {


	grpcHandler := grpc.NewHandler()
	srv := server.NewServer(cfg, grpcHandler)


	go srv.Run()

	// Exit application
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGTERM, syscall.SIGINT)
	<-exit

	//Stoping server
	srv.Stop()

	log.Logger.Fatal("Application Stopped")

}