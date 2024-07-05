package service

import (
	"context"
	"database/sql"
	"discovery_service/config"
	pbc "discovery_service/genproto/collaborations"
	pbcom "discovery_service/genproto/comments"
	pb "discovery_service/genproto/episode_metadata"
	pbe "discovery_service/genproto/episodes"
	pbp "discovery_service/genproto/podcasts"
	"discovery_service/pkg"

	"discovery_service/storage/postgres"
)

type EpisodeMetadataService struct {
	pb.UnimplementedEpisodeMetadataServer
	Repo          *postgres.EpisodeMetadataRepo
	PodcastClient pbp.PodcastsClient
	EpisodeClient pbe.EpisodesServiceClient
	CollabClient  pbc.CollaborationsClient
	CommentClient pbcom.CommentsClient
}

func GetArgumentOfEpisodeMetadate(db *sql.DB) (*postgres.EpisodeMetadataRepo, pbp.PodcastsClient, pbe.EpisodesServiceClient, pbc.CollaborationsClient, pbcom.CommentsClient) {
	cfg := config.Load()
	episodeMetadataRepo := postgres.NewEpisodeMetadataRepo(db)
	ClientPodcasts := pkg.CreatePodcastsClient(cfg)

	ClientEpisodes := pkg.CreateEpisodesClient(cfg)

	ClientCollaborations := pkg.CreateCollaborationsClient(cfg)

	ClientComments := pkg.CreateCommentsClient(cfg)

	return episodeMetadataRepo, ClientPodcasts, ClientEpisodes, ClientCollaborations, ClientComments
}

func NewEpisodeMetadata(db *sql.DB) *EpisodeMetadataService {
	episodeMetadataRepo, clientPodcasts, clientEpisodes, clientCollaborations, clientComments := GetArgumentOfEpisodeMetadate(db)
	return &EpisodeMetadataService{
		Repo:          episodeMetadataRepo,
		PodcastClient: clientPodcasts,
		EpisodeClient: clientEpisodes,
		CollabClient:  clientCollaborations,
		CommentClient: clientComments,
	}
}

func (e *EpisodeMetadataService) CreateEpisodeMetaData(ctx context.Context, episode *pb.EpisodeMetadata) (*pb.Void, error) {
	err := e.Repo.CreateEpisodeMetaData(episode)

	return &pb.Void{}, err
}

func (e *EpisodeMetadataService) GetTrendingPodcasts(ctx context.Context, req *pb.Pagination) (*pb.Podcasts, error) {
	podcasts, err := e.Repo.GetTrendingPodcasts(req)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(podcasts.Podcasts); i++ {
		p, err := e.PodcastClient.GetPodcastById(context.Background(), &pbp.ID{Id: podcasts.Podcasts[i].PodcastId})
		if err != nil {
			return nil, err
		}

		count, err := e.CommentClient.CountComments(context.Background(), &pbcom.CountFilter{PodcastId: podcasts.Podcasts[i].PodcastId})
		if err != nil {
			return nil, err
		}

		podcasts.Podcasts[i].PodcastTitle = p.Title
		podcasts.Podcasts[i].CommentCount = count.Count
	}

	return podcasts, nil
}

func (e *EpisodeMetadataService) GetRecommendedPodcasts(ctx context.Context, req *pb.IdPage) (*pb.Podcasts, error) {
	podcastsIdUserWatched, err := e.Repo.GetPodcastsIdUserWatched(&pb.ID{Id: req.Id})
	if err != nil {
		return nil, err
	}

	// podcasts Id of Recommentded podcasts
	podcastsId, err := e.CollabClient.GetAllPodcastsUsersWorkedOn(ctx, &pbc.PodcastsId{PodcastsId: podcastsIdUserWatched.PodcastsId})
	if err != nil {
		return nil, err
	}

	podcasts, err := e.Repo.GetRecommendedPodcasts(podcastsId.PodcastsId[0], req.Pagination)
	if err != nil {
		return nil, err
	}

	for i := range podcasts.Podcasts {
		id := pbp.ID{Id: podcasts.Podcasts[i].PodcastId}
		additial, err := e.PodcastClient.GetPodcastById(ctx, &id)
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
		e.CommentClient.CountComments(ctx, &pbcom.CountFilter{PodcastId: id.Id})
	}
	return podcasts, nil
}

func (e *EpisodeMetadataService) GetPodcastsByGenre(ctx context.Context, req *pb.Filter) (*pb.Podcasts, error) {
	podcastsInfo, err := e.Repo.GetPodcastsByGenre(req)
	if err != nil {
		return nil, err
	}

	var resp []*pb.Podcast
	for _, p := range podcastsInfo {
		pod, err := e.PodcastClient.GetPodcastById(context.Background(), &pbp.ID{Id: p.PodcastId})
		if err != nil {
			return nil, err
		}

		episodes, err := e.EpisodeClient.GetEpisodesByPodcastId(context.Background(), &pbe.Filter{Id: p.PodcastId})
		if err != nil {
			return nil, err
		}

		var commentCount int64
		for _, ep := range episodes.Episodes {
			count, err := e.CommentClient.CountComments(context.Background(), &pbcom.CountFilter{EpisodeId: ep.Id, PodcastId: p.PodcastId})
			if err != nil {
				return nil, err
			}
			commentCount += count.Count
		}

		resp = append(resp, &pb.Podcast{
			PodcastId:    p.PodcastId,
			PodcastTitle: pod.Title,
			Genre:        p.Genre,
			Tags:         p.Tags,
			CommentCount: commentCount,
			ListenCount:  p.ListenCount,
			LikeCount:    p.LikeCount,
			CreatedAt:    pod.CreatedAt,
			UpdatedAt:    pod.UpdatedAt,
		})
	}

	return &pb.Podcasts{Podcasts: resp}, nil
}

func (e *EpisodeMetadataService) SearchEpisode(ctx context.Context, req *pb.Title) (*pb.Episode, error) {
	resp, err := e.EpisodeClient.SearchEpisodeByTitle(context.Background(), &pbe.Title{Title: req.EpisodeTitle})
	if err != nil {
		return nil, err
	}

	listens, likes, err := e.Repo.GetListensAndLikes(resp.Id)
	if err != nil {
		return nil, err
	}

	comments, err := e.CommentClient.CountComments(context.Background(), &pbcom.CountFilter{
		EpisodeId: resp.Id,
		PodcastId: resp.PodcastId,
	})
	if err != nil {
		return nil, err
	}

	return &pb.Episode{
		Id:           resp.Id,
		PodcastId:    resp.PodcastId,
		UserId:       resp.UserId,
		Title:        resp.Title,
		FileAudio:    resp.FileAudio,
		Description:  resp.Description,
		Duration:     resp.Duration,
		Genre:        resp.Genre,
		Tags:         resp.Tags,
		ListenCount:  int64(listens),
		LikeCount:    int64(likes),
		CommentCount: comments.Count,
		CreatedAt:    resp.CreatedAt,
		UpdatedAt:    resp.UpdatedAt,
	}, nil
}
