package postgres

import "database/sql"

type UserInterRepo struct {
	Db *sql.DB
}

func NewUserInterRepo(db *sql.DB) *UserInterRepo {
	return &UserInterRepo{Db: db}
}
