package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"syscall"

	grpcserver "google.golang.org/grpc"
)

var GrpcAddr = ":9093"

func main() {
	inmemRepo := repository.NewInmemRepository()
	svc := service.NewService(inmemRepo)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
		<-sigCh
		cancel()
	}()

	lis, err := net.Listen("tcp", GrpcAddr)
	if err != nil {
		log.Fatalf("[Me]: Failed to listen: %v", err)
	}

	grpServer := grpcserver.NewServer()

	// TODO: initialize our grpc handler implementation

	log.Printf("[Me]: Starting gRPC server Trip service on port %s", lis.Addr().String())

	go func() {
		if err := grpServer.Serve(lis); err != nil {
			log.Printf("[Me]: Failed to serve: %v", err)
			cancel()
		}
	}()

	// wait for the shutdown signal
	<-ctx.Done()
	log.Printf("[Me]: Shutting down the server...")
	grpServer.GracefulStop()
}
