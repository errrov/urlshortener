package psql

import (
	"context"
	"fmt"
	"log"

	"os"

	"github.com/errrov/urlshortener/internal/model"
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
	Pgxpool          *pgxpool.Pool
}

type Dbrespones struct {
	OriginalURL string
	Identifier  string
}

const query = `CREATE TABLE IF NOT EXISTS shorturl
(
    id CHAR (10) PRIMARY KEY,
    originalurl TEXT
)`

func NewPsql(d ConnectionInfo) *Postgresql {
	log.Println("Creating new PSQL connection")
	var resPsql Postgresql
	resPsql.ConnectionString = d
	resPsql.connect()
	log.Println("Connected")
	resPsql.EnsureTableExists()
	return &resPsql

}

func InitConnectionInfo() ConnectionInfo {
	User := os.Getenv("POSTGRES_USER")
	Password := os.Getenv("POSTGRES_PASSWORD")
	Host := "postgres"
	Port := "5432"
	Name := os.Getenv("POSTGRES_DB")
	return ConnectionInfo{User: User, Password: Password, Host: Host, Port: Port, Name: Name}
}

func (d *Postgresql) connect() {
	connectionStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", d.ConnectionString.User, d.ConnectionString.Password, d.ConnectionString.Host, d.ConnectionString.Port, d.ConnectionString.Name)
	cfg, err := pgxpool.ParseConfig(connectionStr)
	if err != nil {
		log.Panic(err)
	}
	dpPool, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		log.Panic(err)
	}
	d.Pgxpool = dpPool
}

func (d *Postgresql) EnsureTableExists() {
	if _, err := d.Pgxpool.Exec(context.Background(), query); err != nil {
		log.Panic(err)
	}
}

func (d *Postgresql) Add(shortened model.Shortened) (*model.Shortened, error) {
	_, err := d.Pgxpool.Exec(context.Background(), "INSERT INTO shorturl (originalurl, id) VALUES ($1, $2)", shortened.Original, shortened.Identifier)
	if err != nil {
		return nil, model.ErrIdExist
	}
	return &shortened, nil
}

func (d *Postgresql) Get(identifier string) (*model.Shortened, error) {
	var db_response Dbrespones
	err := d.Pgxpool.QueryRow(context.Background(), "select * FROM shorturl WHERE id=$1", identifier).Scan(&db_response.OriginalURL, &db_response.Identifier)
	if err != nil {
		return nil, model.ErrNotFound
	}
	shortening := model.Shortened{Identifier: db_response.Identifier, Original: db_response.OriginalURL}
	return &shortening, nil
}
