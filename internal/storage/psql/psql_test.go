package psql_test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/errrov/urlshortener/internal/model"
	"github.com/errrov/urlshortener/internal/shorten"
	"github.com/errrov/urlshortener/internal/storage/psql"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestConnection(t *testing.T) {
	d := psql.InitConnectionInfo()
	shortened := model.Shortened{Identifier: "xddd", Original: "https://www.google.com"}
	connectionStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", d.User, d.Password, d.Host, d.Port, d.Name)
	cfg, err := pgxpool.ParseConfig(connectionStr)
	if err != nil {
		log.Panic(err)
	}
	dpPool, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		log.Panic(err)
	}
	defer dpPool.Close()
	_, err = dpPool.Exec(context.Background(), "INSERT INTO shorturl (originalurl, id) VALUES ($1, $2) ON CONFLICT (id) DO NOTHING", shortened.Original, shortened.Identifier)
	if err != nil {
		log.Panic(err)
	}
	log.Println(shortened)
}

func TestAddInStoragePSQL(t *testing.T) {
	d := psql.ConnectionInfo{
		User:     "postgres",
		Password: "your_password",
		Host:     "localhost",
		Port:     "5432",
		Name:     "shourturl",
	}
	storage := psql.NewPsql(d)
	s := shorten.NewService(storage)
	customInput := model.UserInput{OriginalURL:"https://www.google.com",Identifier:"cursed"}
	_, err := s.AddShortenLinkToStorage(customInput)
	if err != nil {
		log.Panic(err)
	}

}
