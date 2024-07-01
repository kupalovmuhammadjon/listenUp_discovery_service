package postgres

import (
	"database/sql"

	pb "discovery_service/genproto/episode_metadata"
)

type EpisodeMetadataRepo struct {
	Db *sql.DB
}

func NewEpisodeMetadataRepo(db *sql.DB) *EpisodeMetadataRepo {
	return &EpisodeMetadataRepo{Db: db}
}

func (e *EpisodeMetadataRepo) GetPodcastsIdUserWatched(id *pb.ID) (*pb.PodcastsId, error) {
	queryToTakePodcastsId := `
		select 
			distinct podcast_id
		from 
			user_interactions
		where user_id = $1`

	podcastsId := []string{}
	rows, err := e.Db.Query(queryToTakePodcastsId, id.Id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		podcastId := ""
		err := rows.Scan(&podcastId)
		if err != nil {
			return nil, err
		}
		podcastsId = append(podcastsId, podcastId)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &pb.PodcastsId{PodcastsId: podcastsId}, nil
}

func (e *EpisodeMetadataRepo) GetRecommendedPodcasts(podcastsId *[]string) (*pb.Podcasts, error) {
	query := `
		select
			podcast_id, 
			array_agg(genre) as genre, 
			array_agg(tags) as tags, 
			sum(listen_count) as listen_count, 
			sum(like_count) as like_count
		from
			episode_metadata
		where
			podcast_id = any($1)
		group by
			podcast_id
		`

	podcasts := pb.Podcasts{}

	rows, err := e.Db.Query(query, *podcastsId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		podcast := pb.Podcast{}
		err := rows.Scan(&podcast.PodcastId, &podcast.Genre, &podcast.Tags, &podcast.ListenCount,
			&podcast.LikeCount)

		if err != nil {
			return nil, err
		}
		podcasts.Podcasts = append(podcasts.Podcasts, &podcast)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &podcasts, nil
}
