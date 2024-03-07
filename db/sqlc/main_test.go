package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://pasan:12345@localhost:5432/practice?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M){
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil{
		log.Fatal("cannot connec to the database")
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}