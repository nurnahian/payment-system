package user

import (
	"database/sql"
	"payment-system/internal/common/db"
)

type Repository interface {
	Create(username, hashedPassword string) error
	FindByUsername(username string) (*User, error)
	FindAll() ([]User, error)
}

type userRepo struct {
	db *sql.DB
}

func NewRepository() Repository {
	return &userRepo{db: db.DB}
}

func (r *userRepo) Create(username, hashedPassword string) error {
	_, err := r.db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", username, hashedPassword)
	return err
}

func (r *userRepo) FindByUsername(username string) (*User, error) {
	var u User
	err := r.db.QueryRow("SELECT id, username, password FROM users WHERE username = $1", username).
		Scan(&u.ID, &u.Username, &u.Password)
	return &u, err
}

func (r *userRepo) FindAll() ([]User, error) {
	rows, err := r.db.Query("SELECT id, username FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Username); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
