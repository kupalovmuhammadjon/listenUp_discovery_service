package service

import (
	"context"
	"database/sql"
	"discovery_service/config"
	pbe "discovery_service/genproto/episodes"
	pbp "discovery_service/genproto/podcasts"
	pbu "discovery_service/genproto/user"
	pb "discovery_service/genproto/user_interactions"
	"discovery_service/pkg"
	"discovery_service/storage/postgres"
	"fmt"
)

type UserInterService struct {
	pb.UnimplementedUserInteractionsServer
	Repo          *postgres.UserInterRepo
	UserClient    pbu.UserManagementClient
	PodcastClient pbp.PodcastsClient
	EpisodeClient pbe.EpisodesServiceClient
}

func NewUserInterService(db *sql.DB) *UserInterService {
	cfg := config.Load()
	return &UserInterService{
		Repo:          postgres.NewUserInterRepo(db),
		UserClient:    pkg.CreateUserManagementClient(cfg),
		PodcastClient: pkg.CreatePodcastsClient(cfg),
		EpisodeClient: pkg.CreateEpisodesClient(cfg),
	}
}

func (u *UserInterService) LikeEpisodeOfPodcast(ctx context.Context, req *pb.InteractEpisode) (*pb.ID, error) {

	exists, err := u.UserClient.ValidateUserId(ctx, &pbu.ID{Id: req.UserId})
	if err != nil || !exists.Success {
		return nil, fmt.Errorf("error while validating user id %v", err)
	}
	s, err := u.PodcastClient.ValidatePodcastId(ctx, &pbp.ID{Id: req.PodcastId})
	if err != nil || !s.Success {
		return nil, fmt.Errorf("error while validating podcast id %v", err)
	}
	e, err := u.EpisodeClient.ValidateEpisodeId(ctx, &pbe.ID{Id: req.EpisodeId})
	if err != nil || !e.Success {
		return nil, fmt.Errorf("error while validating episode id %v", err)
	}

	resp, err := u.Repo.LikeEpisodeOfPodcast(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *UserInterService) ValidateUserInteractionId(ctx context.Context, id *pb.ID) (*pb.Success, error) {
	exists, err := u.Repo.ValidateUserInteractionId(id.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Success{Success: exists}, err
}

func (u *UserInterService) DeleteLikeFromEpisodeOfPodcast(ctx context.Context, req *pb.DeleteLike) (*pb.Success, error) {

	exists, err := u.UserClient.ValidateUserId(ctx, &pbu.ID{Id: req.UserId})
	if err != nil || !exists.Success {
		return nil, fmt.Errorf("error while validating user id %v", err)
	}
	s, err := u.PodcastClient.ValidatePodcastId(ctx, &pbp.ID{Id: req.PodcastId})
	if err != nil || !s.Success {
		return nil, fmt.Errorf("error while validating podcast id %v", err)
	}
	e, err := u.EpisodeClient.ValidateEpisodeId(ctx, &pbe.ID{Id: req.EpisodeId})
	if err != nil || !e.Success {
		return nil, fmt.Errorf("error while validating episode id %v", err)
	}

	resp, err := u.Repo.DeleteLikeFromEpisodeOfPodcast(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *UserInterService) ListenEpisodeOfPodcast(ctx context.Context, req *pb.InteractEpisode) (*pb.ID, error) {

	exists, err := u.UserClient.ValidateUserId(ctx, &pbu.ID{Id: req.UserId})
	if err != nil || !exists.Success {
		return nil, fmt.Errorf("error while validating user id %v", err)
	}
	s, err := u.PodcastClient.ValidatePodcastId(ctx, &pbp.ID{Id: req.PodcastId})
	if err != nil || !s.Success {
		return nil, fmt.Errorf("error while validating podcast id %v", err)
	}
	e, err := u.EpisodeClient.ValidateEpisodeId(ctx, &pbe.ID{Id: req.EpisodeId})
	if err != nil || !e.Success {
		return nil, fmt.Errorf("error while validating episode id %v", err)
	}

	resp, err := u.Repo.ListenEpisodeOfPodcast(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
