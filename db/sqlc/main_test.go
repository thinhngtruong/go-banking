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
	dbSource = "postgresql://root:empires@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	//conn, err := sql.Open(dbDriver, dbSource)
	testDB, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatalln("cannot connect to db: ", err)
	}

	//testQueries = New(conn)
	testQueries = New(testDB)

	os.Exit(m.Run())
}
