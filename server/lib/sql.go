package lib

import (
	"database/sql"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func createUser(db *sql.DB, username, mail, password string) error {
	if !isValidMail(mail) {
		return errors.New("Invalid mail format")
	}
	if !isValidUsername(username) {
		return errors.New("Invalid username format")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query, err := db.Prepare("INSERT INTO users (username, mail, password) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}
	defer query.Close()

	_, err = query.Exec(username, mail, string(hashedPassword))

	log.Println("Registered user:", username)
	return err
}

func loginUser(db *sql.DB, username, password string) error {
	if !isValidUsername(username) {
		return errors.New("Invalid username format")
	}

	var hashedPassword string
	query := "SELECT password FROM users WHERE username = $1"
	err := db.QueryRow(query, username).Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("User not found")
		}
		return err
	}
	
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return errors.New("Incorrect password")
	}

	log.Println("User logged in:", username)
	return nil
}