package auth

import (
	"database/sql"
)

func loginUser(db *sql.DB, user, pass string) (id int, hashedPassword string, err error) {
	err = db.QueryRow("SELECT id, password FROM administrators WHERE username = $1", user).Scan(&id, &hashedPassword)
	if err != nil {
		return 0, "", err
	}

	return id, hashedPassword, nil
}

func registerUser(db *sql.DB, user, pass string) (err error) {
	_, err = db.Exec("INSERT INTO administrators (username, password) VALUES ($1, $2)", user, pass)
	if err != nil {
		return err
	}
	return nil
}
