package databaserepo

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"time"
	"webapp/pkg/data"
	"webapp/pkg/repository"

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
var testRepo repository.DatabaseRepo

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

	// create the database repository
	testRepo = &PostgresDBRepo{DB: testDB}

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

// ----------------------------------------------------------------------------
func TestPostgresDBRepo_InsertUser(t *testing.T) {
	testUser := data.User{
		FirstName: "Charlie",
		LastName:  "Harpoon",
		Email:     "harpoon@whale.com",
		Password:  "secret",
		IsAdmin:   1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	id, err := testRepo.InsertUser(testUser)
	if err != nil {
		t.Errorf("insert user returned error: %s", err)
	}

	if id != 1 {
		t.Errorf("database is supposed to be empty, received id %d, expected 1", id)
	}
}

// ----------------------------------------------------------------------------
func TestPostgresDBRepo_AllUsers(t *testing.T) {
	users, err := testRepo.AllUsers()
	if err != nil {
		t.Errorf("failed to get all users: %s", err)
	}

	if len(users) != 1 {
		t.Errorf("received more users than the expected 1: %d", len(users))
	}

	if users[0].FirstName != "Charlie" {
		t.Errorf("username %s is not 'Charlie'", users[0].FirstName)
	}

	testUser := data.User{
		FirstName: "Ron",
		LastName:  "Ronsom",
		Email:     "ron@whale.com",
		Password:  "secret",
		IsAdmin:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, _ = testRepo.InsertUser(testUser)

	users, err = testRepo.AllUsers()
	if err != nil {
		t.Errorf("failed to get all users: %s", err)
	}

	if len(users) != 2 {
		t.Errorf("received more users than the expected 2: %d", len(users))
	}

	if users[0].FirstName != "Charlie" {
		t.Errorf("username %s is not 'Charlie'", users[0].FirstName)
	}

	if users[1].FirstName != "Ron" {
		t.Errorf("username %s is not 'Ron'", users[0].FirstName)
	}

}

// ----------------------------------------------------------------------------
func TestPostgresDBRepo_GetUser(t *testing.T) {
	usr, err := testRepo.GetUser(1)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if usr.Email != "harpoon@whale.com" {
		t.Errorf("user 1 has the wrong email address (%s), expected \"harpoon@whale.com\"", usr.Email)
	}

	usr, err = testRepo.GetUser(3)
	if err == nil {
		t.Errorf("expected an error when trying to find user 3, got a user with email: %s", usr.Email)
	}
}

// ----------------------------------------------------------------------------
func TestPostgresDBRepo_GetUserByEmail(t *testing.T) {
	usr, err := testRepo.GetUserByEmail("ron@whale.com")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if usr.ID != 2 {
		t.Errorf("user has the wrong ID (%d), expected 2.", usr.ID)
	}

	usr, err = testRepo.GetUserByEmail("shorty@stuff.com")
	if err == nil {
		t.Errorf("received a user unexpectedly for shorty@stuff.com: ID %d", usr.ID)
	}
}

// ----------------------------------------------------------------------------
func TestPostgresDBRepo_UpdateUser(t *testing.T) {
	user, err := testRepo.GetUser(2)
	if err != nil {
		t.Errorf("unable to get user with id 2: %s", err)
	}

	user.FirstName = "Ronald"
	err = testRepo.UpdateUser(*user)
	if err != nil {
		t.Errorf("unable to update user: %s", err)
	}

	user, err = testRepo.GetUser(2)
	if err != nil {
		t.Errorf("unable to read updated user with ID 2: %s", err)
	}

	if user.FirstName != "Ronald" {
		t.Errorf("error updating user record, changes not persisted")
	}
}

// ----------------------------------------------------------------------------
func TestProgresDBRepo_DeleteUser(t *testing.T) {
	err := testRepo.DeleteUser(2)
	if err != nil {
		t.Errorf("error deleting user ID 2: %s", err)
	}

	_, err = testRepo.GetUser(2)
	if err == nil {
		t.Errorf("expected an error getting user with ID 2.")
	}
}

// ----------------------------------------------------------------------------
func TestPostgresDBRepo_ResetPassword(t *testing.T) {
	err := testRepo.ResetPassword(1, "seeker")
	if err != nil {
		t.Errorf("unexpected error during password reset: %s", err)
	}

	user, err := testRepo.GetUser(1)
	if err != nil {
		t.Errorf("failed to load user with id 1: %s", err)
	}

	matches, err := user.PasswordMatches("seeker")
	if err != nil {
		t.Errorf("error matching passwords: %s", err)
	}

	if !matches {
		t.Error("expected passwords to match")
	}

	err = testRepo.ResetPassword(2, "dope")
	if err != nil {
		t.Error("expected an error resetting password for non-existent user")
	}
}

// ----------------------------------------------------------------------------
func TestPostgresDBRepo_InsertUserImage(t *testing.T) {

}
