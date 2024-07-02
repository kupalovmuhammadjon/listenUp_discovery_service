package service

import (
	// "context"
	"database/sql"
	// "discovery_service/client"
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

// func (e *EpisodeMetadataService) GetPodcastsByGenre(ctx context.Context, req *pb.Genres) (*pb.Podcasts, error) {
// 	ids, err := e.Repo.GetPodcastIDs()
// 	if err != nil {
// 		return nil, err
// 	}
	
// 	podcasts, _, err := client.RetrieveTitles(ids)
// 	if err != nil {
// 		return nil, err
// 	}

// 	for _, v := range podcasts {
// 		v.
// 	}

// 	return nil, nil
// }
