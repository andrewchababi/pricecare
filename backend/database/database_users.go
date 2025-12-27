package database

import (
	"database/sql"
	"errors"
	"log"

	"github.com/andrewchababi/pricecare/backend/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(username string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO users (username, password)
		VALUES (?, ?, ?)
	`

	_, err = database.Exec(
		query,
		username,
		string(hashedPassword),
	)

	if err != nil {
		// Optional: handle duplicate username cleanly
		if sqliteErr, ok := err.(interface{ Error() string }); ok {
			if sqliteErr.Error() != "" {
				return errors.New("username already exists")
			}
		}
		return err
	}

	return nil

}

func GetUserByUsername(username string) models.User {
	var user models.User
	q := `
		SELECT username, password 
		FROM users
		WHERE username = ?
		LIMIT 1
	`
	row := database.QueryRow(q, username)

	err := row.Scan(
		&user.Username,
		&user.HashedPassword,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.NullUser()
		}
		log.Printf("Error fetching user by username: %v", err)
		return models.NullUser()
	}

	return user
}
