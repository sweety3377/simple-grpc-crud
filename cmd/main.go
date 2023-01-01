package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"simple-crud/config"
	"simple-crud/database"
	"simple-crud/internal/repository"
	"simple-crud/internal/server"
	"simple-crud/internal/service"
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

	// Create client repository and service
	clientRepository := repository.NewClientStorage(db)
	clientService := service.NewClientService(clientRepository)

	// Create and run server
	srv := server.NewServer(db, clientService)
	go func() {
		if err = srv.Start(cfg); err != nil {
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
