package repo

import (
	"database/sql"
	"zydukas-rankings/models"
)

type UserRepo struct {
	db *sql.DB
}

func CreateUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) CreateUser(user *models.CreateUserPayload) error {
	query := `INSERT INTO users (username) VALUES ($1) RETURNING id;`

	if err := u.db.QueryRow(query, user.Username); err != nil {
		return err.Err()
	}

	return nil
}
