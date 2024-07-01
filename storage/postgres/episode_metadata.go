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

	rows, err := e.Db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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

func (e *EpisodeMetadataRepo) SearchPodcast(titles *pb.Title) ([]string, error) {
	rows, err := e.Db.Query("select podcast_id from episode_metadata where deleted_at = null")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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
