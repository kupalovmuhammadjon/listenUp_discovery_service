package service

import (
	"database/sql"
	pb "discovery_service/genproto/user_interactions"
	"discovery_service/storage/postgres"
)

type UserInterService struct {
	pb.UnimplementedUserInteractionsServer
	Repo *postgres.UserInterRepo
}

func NewUserInterService(db *sql.DB) *UserInterService {
	return &UserInterService{
		Repo: postgres.NewUserInterRepo(db),
	}
}
