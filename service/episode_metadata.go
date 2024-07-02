package service

import (
	"context"
	"database/sql"
	pbCollaboration "discovery_service/genproto/collaborations"
	pbComment "discovery_service/genproto/comments"
	pb "discovery_service/genproto/episode_metadata"
	pbPodcast "discovery_service/genproto/podcasts"
	"discovery_service/pkg"
	"discovery_service/storage/postgres"
	"log"
)

type EpisodeMetadata struct {
	pb.UnimplementedEpisodeMetadataServer
	EpisodeMetadataRepo  *postgres.EpisodeMetadataRepo
	ClientCollaborations pbCollaboration.CollaborationsClient
	ClientPodcasts       pbPodcast.PodcastsClient
	ClientComment        pbComment.CommentsClient
}

func GetArgumentOfEpisodeMetadate(db *sql.DB) (*postgres.EpisodeMetadataRepo, pbCollaboration.CollaborationsClient, pbPodcast.PodcastsClient, pbComment.CommentsClient) {
	episodeMetadataRepo := postgres.NewEpisodeMetadataRepo(db)
	ClientCollaborations, err := pkg.CreateCollaborationsClient()
	if err != nil {
		log.Println(err)
	}
	ClientPodcasts, err := pkg.CreatePodcastClient()
	if err != nil {
		log.Println(err)
	}
	ClientComment, err := pkg.CreateCommentsClient()
	if err != nil {
		log.Println(err)
	}

	return episodeMetadataRepo, ClientCollaborations, ClientPodcasts, ClientComment
}

func NewEpisodeMetadata(db *sql.DB) *EpisodeMetadata {
	episodeMetadataRepo, ClientCollaborations, ClientPodcasts, ClientComment := GetArgumentOfEpisodeMetadate(db)
	return &EpisodeMetadata{
		EpisodeMetadataRepo:  episodeMetadataRepo,
		ClientCollaborations: ClientCollaborations,
		ClientPodcasts:       ClientPodcasts,
		ClientComment:        ClientComment,
	}
}

func (e *EpisodeMetadata) GetRecommendedPodcasts(ctx context.Context, userId *pb.ID) (*pb.Podcasts, error) {
	podcastsIdUserWatched, err := e.EpisodeMetadataRepo.GetPodcastsIdUserWatched(userId)
	if err != nil {
		return nil, err
	}

	// podcasts Id of Recommentded podcasts
	podcastsId, err := e.ClientCollaborations.GetAllPodcastsUsersWorkedOn(ctx, &pbCollaboration.PodcastsId{PodcastsId: podcastsIdUserWatched.PodcastsId})
	if err != nil {
		return nil, err
	}

	podcasts, err := e.EpisodeMetadataRepo.GetRecommendedPodcasts(&podcastsId.PodcastsId)
	if err != nil {
		return nil, err
	}

	for i := range podcasts.Podcasts {
		id := pbPodcast.ID{Id: podcasts.Podcasts[i].PodcastId}
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
		e.ClientComment.CountComments(ctx, &pbComment.CountFilter{PodcastId: id.Id})
	}
	return podcasts, nil
}
