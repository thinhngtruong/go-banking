package db

import (
	"database/sql"
	"github.com/thinhngtruong/go-banking/util"
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
	config, err := util.LoadConfig("../..")

	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatalln("cannot connect to db: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
