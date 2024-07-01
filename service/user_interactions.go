package service

import (
	"context"
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

func (u *UserInterService) LikeEpisodeOfPodcast(ctx context.Context, req *pb.InteractEpisode) (*pb.ID, error) {
	resp, err := u.Repo.LikeEpisodeOfPodcast(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *UserInterService) DeleteLikeFromEpisodeOfPodcast(ctx context.Context, req *pb.DeleteLike) (*pb.Success, error) {
	resp, err := u.Repo.DeleteLikeFromEpisodeOfPodcast(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *UserInterService) ListenEpisodeOfPodcast(ctx context.Context, req *pb.InteractEpisode) (*pb.ID, error) {
	resp, err := u.Repo.ListenEpisodeOfPodcast(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
