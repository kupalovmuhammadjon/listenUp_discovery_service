package service

import (
	"database/sql"
	pb "discovery_service/genproto/episode_metadata"
	"discovery_service/storage/postgres"
)

type EpisodeMetadataService struct {
	pb.UnimplementedEpisodeMetadataServer
	Repo *postgres.EpisodeMetadataRepo
}

func NewEpisodeMetadataService(db *sql.DB) *EpisodeMetadataService {
	return &EpisodeMetadataService{
		Repo: postgres.NewEpisodeMetadataRepo(db),
	}
}
