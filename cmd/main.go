package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	//"github.com/errrov/urlshortenerozon/internal/model"
	"os/signal"

	"github.com/errrov/urlshortenerozon/internal/server"
	"github.com/errrov/urlshortenerozon/internal/shorten"
	"github.com/errrov/urlshortenerozon/internal/storage/in_memory"
	"github.com/errrov/urlshortenerozon/internal/storage/psql"
)

func main() {
	var shorteningStorage shorten.Storage
	storageType := flag.String("Memory_type", "in_memory", "type of memory storage, psql for using Postgres / default for in_memory")
	flag.Parse()
	if *storageType == "psql" {
		shorteningStorage := psql.Postgresql{}
		testConnectionString := psql.ConnectionInfo{
			User: "postgres",
			Password: "Counter209688",
			Host: "localhost",
			Port: "5432",
			Name: "shorturl",
		}
		shorteningStorage.ConnectionString = testConnectionString

	} else {
		shorteningStorage = in_memory.NewInMemory()
	}
	shortenService := shorten.NewService(shorteningStorage)
	srv := server.New(shortenService)
	port := ":7000"
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	go func() {
		if err := http.ListenAndServe(port, srv); err == http.ErrServerClosed {
			log.Fatalf("Server running error: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify

	<-quit
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
}
