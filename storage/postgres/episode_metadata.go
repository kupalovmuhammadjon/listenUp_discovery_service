package postgres

import (
	"database/sql"
	pb "discovery_service/genproto/episode_metadata"
	"fmt"
)

type EpisodeMetadataRepo struct {
	Db *sql.DB
}

func NewEpisodeMetadataRepo(db *sql.DB) *EpisodeMetadataRepo {
	return &EpisodeMetadataRepo{Db: db}
}

func (e *EpisodeMetadataRepo) CreateEpisodeMetaData(epData *pb.EpisodeMetadata) error {
	query := `
	insert into
	  episode_metadata(
	  episode_id,
	  podcast_id,
	  genre,
	  listen_count,
	  like_count
	  )
	values ($1, $2, $3, $4, $5)`

	_, err := e.Db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (e *EpisodeMetadataRepo) GetTrendingPodcasts() (*pb.Podcasts, error) {
	query := `
	select
	  podcast_id,
	  genre,
	  listen_count,
	  like_count
	from
	  episode_metadata
	where
	  created_at between current_timestamp - interval '3 months' and current_timestamp and deleted_at is null
	order by
	  like_count, listen_count
	`
	rows, err := e.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	podcasts := pb.Podcasts{}
	for rows.Next() {
		p := &pb.Podcast{}
		err := rows.Scan(&p.PodcastId, &p.Genre, &p.ListenCount, &p.LikeCount)
		if err != nil {
			return nil, err
		}
		podcasts.Podcasts = append(podcasts.Podcasts, p)
	}

	return &podcasts, nil
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

func (e *EpisodeMetadataRepo) GetPodcastsByGenre(genres *pb.Genres) ([]pb.Podcast, error) {
	query := `select em.podcast_id,
	array_agg(em.genre) as genres,
	array_agg(em.tags) as tags,
	sum(case when ui.interaction_type = 'listen' then 1 else 0 end) as listen_count,
	sum(case when ui.interaction_type = 'like' then 1 else 0 end) as like_count
	from episode_metadata em
	join user_interactions ui
	on em.episode_id = ui.episode_id and em.podcast_id = ui.podcast_id
	where em.deleted_at is null and ui.deleted_at is null`

	var params []interface{}
	for i, v := range genres.Genres {
		query += fmt.Sprintf(" and genre = $%d", i+1)
		params = append(params, v)
	}
	query += " group by em.podcast_id"

	rows, err := e.Db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var podcasts []pb.Podcast
	for rows.Next() {
		var id, genre string
		var tags []string
		var listen, like int
		err := rows.Scan(&id, &genre, &tags, &listen, &like)
		if err != nil {
			return nil, err
		}
		podcasts = append(podcasts, pb.Podcast{PodcastId: id,
			Genre:       genre,
			Tags:        tags,
			ListenCount: int64(listen),
			LikeCount:   int64(like)})
	}

	return podcasts, nil
}

func (e *EpisodeMetadataRepo) SearchPodcast(titles *pb.Title) ([]pb.Podcast, error) {
	query := `select em.podcast_id,
	array_agg(em.genre) as genres,
	array_agg(em.tags) as tags,
	sum(case when ui.interaction_type = 'listen' then 1 else 0 end) as listen_count,
	sum(case when ui.interaction_type = 'like' then 1 else 0 end) as like_count
	from episode_metadata em
	join user_interactions ui
	on em.episode_id = ui.episode_id and em.podcast_id = ui.podcast_id
	where em.deleted_at is null and ui.deleted_at is null
	group by em.podcast_id`

	rows, err := e.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var podcasts []pb.Podcast
	for rows.Next() {
		var id, genre string
		var tags []string
		var listen, like int
		err := rows.Scan(&id, &genre, &tags, &listen, &like)
		if err != nil {
			return nil, err
		}
		podcasts = append(podcasts, pb.Podcast{PodcastId: id,
			Genre:       genre,
			Tags:        tags,
			ListenCount: int64(listen),
			LikeCount:   int64(like)})
	}

	return podcasts, nil
}

func (e *EpisodeMetadataRepo) GetPodcastIDs() ([]string, error) {
	rows, err := e.Db.Query("select podcast_id from episode_metadata where deleted_at is null")
	if err != nil {
		return nil, err
	}

	var podcastIDs []string
	for rows.Next() {
		var id string
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		podcastIDs = append(podcastIDs, id)
	}

	return podcastIDs, nil
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
