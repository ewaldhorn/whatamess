package databaserepo

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

// ----------------------------------------------------------------------------
var (
	host     = "localhost"
	user     = "postgres"
	password = "postgres"
	dbName   = "users_test"
	port     = "5435"
	dsn      = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable timezone=UTC connect_timeout=5"
)

var resource *dockertest.Resource
var pool *dockertest.Pool
var testDB *sql.DB

// ----------------------------------------------------------------------------
func TestMain(m *testing.M) {
	// connect to docker, fail if docker is not running
	tmpPool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("unable to connect to Docker instance: %s", err)
	}

	pool = tmpPool

	// set up docker options
	options := dockertest.RunOptions{
		Repository:   "postgres",
		Tag:          "14.5",
		Env:          []string{"POSTGRES_USER=" + user, "POSTGRES_PASSWORD=" + password, "POSTGRES_DB=" + dbName},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{"5432": {{HostIP: "0.0.0.0", HostPort: port}}}}

	// get docker image
	tmpResource, err := pool.RunWithOptions(&options)
	if err != nil {
		_ = pool.Purge(tmpResource)
		log.Fatalf("could not start resource: %s", err)
	}

	resource = tmpResource

	// start and wait for image
	if err := pool.Retry(func() error {
		var err error
		testDB, err = sql.Open("pgx", fmt.Sprintf(dsn, host, port, user, password, dbName))
		if err != nil {
			log.Println("Error:", err)
			return err
		}
		return testDB.Ping()
	}); err != nil {
		_ = pool.Purge(resource)
		log.Fatalf("could not connect to the database: %s", err)
	}

	// set up database with tables etc.
	err = createTables()
	if err != nil {
		log.Fatalf("error creating tables: %s", err)
	}

	// run tests
	code := m.Run()

	// clean up
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("could not purge resources: %s", err)
	}

	// all done
	os.Exit(code)
}

// ----------------------------------------------------------------------------
func createTables() error {
	// attempt to read the script
	sqlScript, err := os.ReadFile("./testdata/users.sql")
	if err != nil {
		fmt.Println(err)
		return err
	}

	// try to execute it
	_, err = testDB.Exec(string(sqlScript))
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// ----------------------------------------------------------------------------
func Test_PingDB(t *testing.T) {
	err := testDB.Ping()
	if err != nil {
		t.Error("unable to ping database")
	}
}
