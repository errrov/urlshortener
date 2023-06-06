package main

import (
	"context"
	"flag"
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
		d = psql.InitConnectionInfo()
		shorteningStorage = psql.NewPsql(d)

	} else {
		shorteningStorage = in_memory.NewInMemory()
	}
	shortenService := shorten.NewService(shorteningStorage)
	srv := server.New(shortenService)
	port := ":8080"
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	go func() {
		if err := http.ListenAndServe(port, srv); err == http.ErrServerClosed {
			log.Fatalf("Server running error: %v", err)
		}
	}()

	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.E.Shutdown(ctx); err != nil {
		srv.E.Logger.Fatal(err)
	}
}
