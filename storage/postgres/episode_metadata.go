package postgres

import (
	"database/sql"
	pb "discovery_service/genproto/episode_metadata"
	"strings"

	"github.com/lib/pq"
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
	  tags
	  )
	values ($1, $2, $3, $4)`

	_, err := e.Db.Exec(query, epData.EpisodeId, epData.PodcastId, epData.Genre, pq.Array(epData.Tags))
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
		array_to_string(array_agg(genre), ',') as genre, 
		array_to_string(array_agg(tags), ',') as tags, 
		sum(listen_count) as listen_count, 
		sum(like_count) as like_count
	  from
		episode_metadata
	  where
		podcast_id = any($1)
	  group by
		podcast_id
	  `

	rows, err := e.Db.Query(query, pq.Array(podcastsId))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	podcasts := pb.Podcasts{}
	for rows.Next() {
		var podcast pb.Podcast
		var tagsString string
		err := rows.Scan(&podcast.PodcastId, &podcast.Genre, &tagsString,
			&podcast.ListenCount, &podcast.LikeCount)

		if err != nil {
			return nil, err
		}

		tagsString = strings.Trim(tagsString, "{}")
		podcast.Tags = strings.Split(tagsString, ",")
		podcasts.Podcasts = append(podcasts.Podcasts, &podcast)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &podcasts, nil
}

func (e *EpisodeMetadataRepo) GetPodcastsByGenre(genres *pb.Genres) ([]*pb.Podcast, error) {
	query := `
	select
		em.podcast_id,
		array_to_string(array_agg(genre), ',') as genre,
		array_to_string(array_agg(tags), ',') as tags,
		sum(case when ui.interaction_type = 'listen' then 1 else 0 end) as listen_count,
		sum(case when ui.interaction_type = 'like' then 1 else 0 end) as like_count
	from
		episode_metadata em
	join
		user_interactions ui
	on
		em.episode_id = ui.episode_id and em.podcast_id = ui.podcast_id
	where
		em.deleted_at is null and ui.deleted_at is null and em.genre::text = any($1)
	group by
		em.podcast_id`

	rows, err := e.Db.Query(query, pq.Array(genres.Genres))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var podcasts []*pb.Podcast
	for rows.Next() {
		var id, genre, tagsString string
		var listen, like int
		err := rows.Scan(&id, &genre, &tagsString, &listen, &like)
		if err != nil {
			return nil, err
		}
		tagsString = strings.Trim(tagsString, "{}")

		podcasts = append(podcasts, &pb.Podcast{PodcastId: id,
			Genre:       genre,
			Tags:        strings.Split(tagsString, ","),
			ListenCount: int64(listen),
			LikeCount:   int64(like)})
	}

	return podcasts, nil
}

func (e *EpisodeMetadataRepo) SearchPodcast(titles *pb.Title) ([]*pb.Podcast, error) {
	query := `
	select
		em.podcast_id,
		array_to_string(array_agg(genre), ',') as genre,
		array_to_string(array_agg(tags), ',') as tags,
		sum(case when ui.interaction_type = 'listen' then 1 else 0 end) as listen_count,
		sum(case when ui.interaction_type = 'like' then 1 else 0 end) as like_count
	from
		episode_metadata em
	join
		user_interactions ui
	on
		em.episode_id = ui.episode_id and em.podcast_id = ui.podcast_id
	where
		em.deleted_at is null and ui.deleted_at is null
	group by
		em.podcast_id`

	rows, err := e.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var podcasts []*pb.Podcast
	for rows.Next() {
		var id, genre, tagsString string
		var listen, like int
		err := rows.Scan(&id, &genre, &tagsString, &listen, &like)
		if err != nil {
			return nil, err
		}
		podcasts = append(podcasts, &pb.Podcast{PodcastId: id,
			Genre:       genre,
			Tags:        strings.Split(tagsString, ","),
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
