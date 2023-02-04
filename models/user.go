package models

import "database/sql"

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

func GetUserByID(db *sql.DB, id string) (User, error) {
	var user User
	err := db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email)

	if err != nil {
		return User{}, err
	}

	return user, err
}