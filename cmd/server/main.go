package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"os"
	"os/signal"
	"simple-grpc-crud/config"
	"simple-grpc-crud/database"
	"simple-grpc-crud/internal/gateway"
	"simple-grpc-crud/internal/repository"
	"simple-grpc-crud/internal/server"
	"simple-grpc-crud/internal/service"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Load config from env
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to database
	db, err := database.Connect(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(ctx)

	grpcServer := grpc.NewServer()
	defer grpcServer.GracefulStop()

	srv := server.NewGrpcServer(grpcServer)

	clientRepository := repository.NewClientStorage(db)
	clientService := service.NewClientService(clientRepository)

	go func() {
		if err = srv.Start(cfg, clientService); err != nil {
			log.Fatal(err)
		}
	}()

	grpcGateway := gateway.NewGrpcGateway(grpcServer)
	go func() {
		if err = grpcGateway.Start(ctx, cfg); err != nil {
			log.Fatal(err)
		}
	}()

	// Graceful shutdown
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGINT)
	<-c

	// Shutting down app
	log.Println("shutting down app")
}
