package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"os/signal"

	"github.com/errrov/urlshortenerozon/internal/server"
	"github.com/errrov/urlshortenerozon/internal/shorten"
	"github.com/errrov/urlshortenerozon/internal/storage/in_memory"
	"github.com/errrov/urlshortenerozon/internal/storage/psql"
)

func main() {
	var shorteningStorage shorten.Storage
	var d psql.ConnectionInfo
	storageType := flag.String("Memory_type", "in_memory", "type of memory storage, psql for using Postgres / default for in_memory")
	flag.Parse()
	if *storageType == "psql" {
		log.Println("Psql")
		d = psql.ConnectionInfo{
			User:     "postgres",
			Password: "your_password",
			Host:     "localhost",
			Port:     "5432",
			Name:     "shourturl",
		}
		shorteningStorage = psql.NewPsql(d)
		log.Println(d)
		

	} else {
		shorteningStorage = in_memory.NewInMemory()
	}
	fmt.Println("HUIH?")
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
