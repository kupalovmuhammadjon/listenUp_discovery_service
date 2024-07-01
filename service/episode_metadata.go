package service

import (
	"context"
	"database/sql"
	pbC "discovery_service/genproto/collaborations"
	pb "discovery_service/genproto/episode_metadata"
	pbP "discovery_service/genproto/podcasts"
	"discovery_service/storage/postgres"
)

type EpisodeMetadata struct {
	pb.UnimplementedEpisodeMetadataServer
	EpisodeMetadataRepo  *postgres.EpisodeMetadataRepo
	ClientCollaborations pbC.CollaborationsClient
	ClientPodcasts       pbP.PodcastsClient
}

func NewEpisodeMetadata(db *sql.DB, ClientCollaborations pbC.CollaborationsClient) *EpisodeMetadata {
	episodeMetadataRepo := postgres.NewEpisodeMetadataRepo(db)
	return &EpisodeMetadata{
		EpisodeMetadataRepo:  episodeMetadataRepo,
		ClientCollaborations: ClientCollaborations,
	}
}

func (e *EpisodeMetadata) GetRecommendedPodcasts(ctx context.Context, userId *pb.ID) (*pb.Podcasts, error) {
	podcastsIdUserWatched, err := e.EpisodeMetadataRepo.GetPodcastsIdUserWatched(userId)
	if err != nil {
		return nil, err
	}

	// podcasts Id of Recommentded podcasts
	podcastsId, err := e.ClientCollaborations.GetAllPodcastsUsersWorkedOn(ctx, &pbC.PodcastsId{PodcastsId: podcastsIdUserWatched.PodcastsId})

	podcasts, err := e.EpisodeMetadataRepo.GetRecommendedPodcasts(&podcastsId.PodcastsId)
	if err != nil {
		return nil, err
	}

	for i := range podcasts.Podcasts {
		id := pbP.ID{Id: podcasts.Podcasts[i].PodcastId}
		additial, err := e.ClientPodcasts.GetPodcastById(ctx, &id)
		if err != nil {
			return nil, err
		}

		// from podcast proto
		// title, created_at, updated_at
		podcasts.Podcasts[i].PodcastTitle = additial.Title
		podcasts.Podcasts[i].CreatedAt = additial.CreatedAt
		podcasts.Podcasts[i].UpdatedAt = additial.UpdatedAt

		// from comments
		// comment_count
	}
	return podcasts, nil
}
