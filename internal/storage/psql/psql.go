package psql

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/errrov/urlshortenerozon/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ConnectionInfo struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

type Postgresql struct {
	ConnectionString ConnectionInfo
}

type Dbrespones struct {
	OriginalURL string
	Identifier  string
}

func NewPsql(d ConnectionInfo) *Postgresql {
	return &Postgresql{
		ConnectionString: d,
	}
}

func InitConnectionInfo() ConnectionInfo {
	User := os.Getenv("APP_DB_USERNAME")
	Password := os.Getenv("APP_DB_PASSWORD")
	Host := os.Getenv("APP_DB_HOST")
	Port := "5432"
	Name := os.Getenv("APP_DB_NAME")
	return ConnectionInfo{User: User, Password: Password, Host: Host, Port: Port, Name: Name}
}

func (d *Postgresql) Add(shortened model.Shortened) (*model.Shortened, error) {
	connectionStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", d.ConnectionString.User, d.ConnectionString.Password, d.ConnectionString.Host, d.ConnectionString.Port, d.ConnectionString.Name)
	log.Println("Connection string:", connectionStr)
	cfg, err := pgxpool.ParseConfig(connectionStr)
	if err != nil {
		log.Panic(err)
	}
	dpPool, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		log.Panic(err)
	}
	defer dpPool.Close()
	_, err = dpPool.Exec(context.Background(), "INSERT INTO shorturl (originalurl, id) VALUES ($1, $2)", shortened.Original, shortened.Identifier)
	if err != nil {
		return nil, model.ErrIdExist
	}
	return &shortened, nil
}

func (d *Postgresql) Get(identifier string) (*model.Shortened, error) {
	connectionStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", d.ConnectionString.User, d.ConnectionString.Password, d.ConnectionString.Host, d.ConnectionString.Port, d.ConnectionString.Name)
	cfg, err := pgxpool.ParseConfig(connectionStr)
	if err != nil {
		log.Panic(err)
	}
	dpPool, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		log.Panic(err)
	}
	defer dpPool.Close()
	var db_response Dbrespones
	err = dpPool.QueryRow(context.Background(), "select * FROM shorturl WHERE id=$1", identifier).Scan(&db_response.OriginalURL, &db_response.Identifier)
	if err != nil {
		return nil, model.ErrNotFound
	}
	shortening := model.Shortened{Identifier: db_response.Identifier, Original: db_response.OriginalURL}
	return &shortening, nil
}
