package service

import (
	// "context"
	"database/sql"
<<<<<<< HEAD
	// "discovery_service/client"
=======
	"discovery_service/client"
	pbc "discovery_service/genproto/comments"
>>>>>>> origin/Muhammadjon
	pb "discovery_service/genproto/episode_metadata"
	pbp "discovery_service/genproto/podcasts"

	"discovery_service/storage/postgres"
)

type EpisodeMetadataService struct {
	pb.UnimplementedEpisodeMetadataServer
	Repo          *postgres.EpisodeMetadataRepo
	PodcastClient *pbp.PodcastsClient
	CommentClient *pbc.CommentsClient
}

func NewEpisodeMetadataService(db *sql.DB, podcastClient *pbp.PodcastsClient, commentClient *pbc.CommentsClient) *EpisodeMetadataService {
	return &EpisodeMetadataService{
		Repo: postgres.NewEpisodeMetadataRepo(db),
		PodcastClient: podcastClient,
		CommentClient: commentClient,
	}
}

<<<<<<< HEAD
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
=======
func (e *EpisodeMetadataService) GetPodcastsByGenre(ctx context.Context, req *pb.Genres) (*pb.Podcasts, error) {
	ids, err := e.Repo.GetPodcastIDs()
	if err != nil {
		return nil, err
	}

	podcasts, _, err := client.RetrieveTitles(ids)
	if err != nil {
		return nil, err
	}

	// for _, v := range podcasts {

	// }

	return nil, nil
}

func (e *EpisodeMetadataService) CreateEpisodeMetaData(ctx context.Context, episode *pb.EpisodeMetadata) (*pb.Void, error) {
	err := e.Repo.CreateEpisodeMetaData(episode)

	return &pb.Void{}, err
}

func (e *EpisodeMetadataService) GetTrendingPodcasts(context.Context, *pb.Void) (*pb.Podcasts, error) {
	podcasts, err := e.Repo.GetTrendingPodcasts()
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(podcasts.Podcasts); i++ {
		p, err := (*e.PodcastClient).GetPodcastById(context.Background(), &pbp.ID{Id: podcasts.Podcasts[i].PodcastId})
		if err != nil {
			return nil, err
		}

		count, err := (*e.CommentClient).CountComments(context.Background(), &pbc.CountFilter{PodcastId: podcasts.Podcasts[i].PodcastId})
		if err != nil {
			return nil, err
		}

		podcasts.Podcasts[i].PodcastTitle = p.Title
		podcasts.Podcasts[i].CommentCount = count.Count
	}

	return podcasts, nil
}
>>>>>>> origin/Muhammadjon
