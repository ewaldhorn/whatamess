package databaserepo

import (
	"database/sql"
	"errors"
	"time"
	"webapp/pkg/data"
)

// ----------------------------------------------------------------------------
type TestingDBRepo struct{}

// ----------------------------------------------------------------------------
func (repo TestingDBRepo) Connection() *sql.DB {
	return nil
}

// ----------------------------------------------------------------------------
// AllUsers returns all users as a slice of *data.User
func (m *TestingDBRepo) AllUsers() ([]*data.User, error) {
	var users []*data.User

	return users, nil
}

// ----------------------------------------------------------------------------
// GetUser returns one user by id
func (m *TestingDBRepo) GetUser(id int) (*data.User, error) {
	var user = data.User{
		ID:        1,
		FirstName: "Bob",
		LastName:  "Nope",
		Email:     "bob@nope.com",
	}

	return &user, nil
}

// ----------------------------------------------------------------------------
// GetUserByEmail returns one user by email address
func (m *TestingDBRepo) GetUserByEmail(email string) (*data.User, error) {
	if email == "admin@example.com" {
		return &data.User{
			ID:        1,
			Email:     "admin@example.com",
			Password:  "$2a$14$ajq8Q7fbtFRQvXpdCq7Jcuy.Rx1h/L4J60Otx.gyNLbAYctGMJ9tK",
			IsAdmin:   1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}, nil
	}

	return nil, errors.New("unknown user")
}

// ----------------------------------------------------------------------------
// UpdateUser updates one user in the database
func (m *TestingDBRepo) UpdateUser(u data.User) error {
	return nil
}

// ----------------------------------------------------------------------------
// DeleteUser deletes one user from the database, by id
func (m *TestingDBRepo) DeleteUser(id int) error {
	return nil
}

// ----------------------------------------------------------------------------
// InsertUser inserts a new user into the database, and returns the ID of the newly inserted row
func (m *TestingDBRepo) InsertUser(user data.User) (int, error) {
	newID := 2

	return newID, nil
}

// ----------------------------------------------------------------------------
// ResetPassword is the method we will use to change a user's password.
func (m *TestingDBRepo) ResetPassword(id int, password string) error {
	return nil
}

// ----------------------------------------------------------------------------
// InsertUserImage inserts a user profile image into the database.
func (m *TestingDBRepo) InsertUserImage(i data.UserImage) (int, error) {
	newID := 1

	return newID, nil
}
