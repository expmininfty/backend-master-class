package dbsqlc

import (
	"log"
	"os"
	"testing"

	"github.com/expmininfty/go-master-class/internal/database"
)

var testQueries *Queries
var testDb *database.DB

const dns = "postgresql://root:secret@localhost:54320/simple_bank?sslmode=disable"

func TestMain(m *testing.M) {
	pool, err := database.NewPool(dns)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	testDb = pool
	testQueries = New(testDb)

	os.Exit(m.Run())
}
