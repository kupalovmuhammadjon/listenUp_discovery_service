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

func (e *EpisodeMetadataRepo) GetPodcastsByGenre(genres *pb.Genres) ([]string, error) {
	query := `select podcast_id
	from episode_metadata
	where deleted_at = null`

	var params []interface{}
	for i, v := range genres.Genres {
		query += fmt.Sprintf(" and genre = $%d", i+1)
		params = append(params, v)
	}

	// rows, err := e.Db.Query(query, params...)
	// if err != nil {
	// 	return nil, err
	// }
	// defer rows.Close()

	// var podcastIDs []string
	// for rows.Next() {
	// 	var id string
	// 	err := rows.Scan(&id)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	podcastIDs = append(podcastIDs, id)
	// }

	return []string{}, nil
}

func (e *EpisodeMetadataRepo) SearchPodcast(titles *pb.Title) ([]string, error) {
	// rows, err := e.Db.Query("select podcast_id from episode_metadata where deleted_at = null")
	// if err != nil {
	// 	return nil, err
	// }
	// defer rows.Close()

	// var podcastIDs []string
	// for rows.Next() {
	// 	var id string
	// 	err := rows.Scan(&id)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	podcastIDs = append(podcastIDs, id)
	// }

	return []string{}, nil
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
	  created_at between current_timestamp - interval '3 months' and current_timestamp
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

func (e *EpisodeMetadataRepo) GetPodcastIDs() ([]string, error) {
	rows, err := e.Db.Query("select podcast_id from episode_metadata where deleted_at = null")
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
