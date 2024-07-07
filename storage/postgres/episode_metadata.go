package postgres

import (
	"database/sql"
	pb "discovery_service/genproto/episode_metadata"
	"fmt"

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

func (e *EpisodeMetadataRepo) GetTrendingPodcasts(p *pb.Pagination) (*pb.Podcasts, error) {
	query := `
    SELECT
        podcast_id,
        genre,
        SUM(listen_count) as total_listen_count,
        SUM(like_count) as total_like_count
    FROM
        episode_metadata
    WHERE
        created_at BETWEEN current_timestamp - interval '3 months' AND current_timestamp
        AND deleted_at IS NULL
    GROUP BY
        podcast_id, genre
    ORDER BY
        total_like_count DESC, total_listen_count DESC
    LIMIT $1
    OFFSET $2
    `
	rows, err := e.Db.Query(query, p.Limit, p.Offset)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	podcasts := &pb.Podcasts{
		Podcasts: []*pb.Podcast{},
	}

	for rows.Next() {
		p := &pb.Podcast{}
		err := rows.Scan(&p.PodcastId, &p.Genre, &p.ListenCount, &p.LikeCount)
		if err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}
		podcasts.Podcasts = append(podcasts.Podcasts, p)
	}

	// Check for errors encountered during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	if len(podcasts.Podcasts) == 0 {
		return &pb.Podcasts{Podcasts: []*pb.Podcast{}}, nil
	}

	return podcasts, nil
}

func (e *EpisodeMetadataRepo) GetRecommendedPodcasts(podcastId string,
	p *pb.Pagination) (*pb.Podcasts, error) {
	query := `
	select
		podcast_id, 
		array_agg(genre) as genre, 
		sum(listen_count) as listen_count, 
		sum(like_count) as like_count
	from
		episode_metadata
	where
		podcast_id = $1
	group by
		podcast_id
	limit $2
	offset $3
	  `

	rows, err := e.Db.Query(query, podcastId, p.Limit, p.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	podcasts := pb.Podcasts{}
	for rows.Next() {
		var podcast pb.Podcast
		err := rows.Scan(&podcast.PodcastId, pq.Array(&podcast.Genre),
			&podcast.ListenCount, &podcast.LikeCount)

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

func (e *EpisodeMetadataRepo) GetPodcastsByGenre(f *pb.Filter) ([]*pb.Podcast, error) {
	query := `
	WITH tags_agg AS (
		SELECT
			em.podcast_id,
			string_agg(distinct em.genre::text, ',') as genre,
			COALESCE(string_agg(distinct tag, ','), '') as tags
		FROM
			episode_metadata em
		LEFT JOIN
			unnest(em.tags) as tag ON em.podcast_id IS NOT NULL
		WHERE
			em.genre::text = any($1) AND em.deleted_at IS NULL
		GROUP BY
			em.podcast_id
	)
	SELECT
		ta.podcast_id,
		ta.genre,
		ta.tags,
		COALESCE(SUM(CASE WHEN ui.interaction_type = 'listen' THEN 1 ELSE 0 END), 0) AS listen_count,
		COALESCE(SUM(CASE WHEN ui.interaction_type = 'like' THEN 1 ELSE 0 END), 0) AS like_count
	FROM
		tags_agg ta
	LEFT JOIN
		user_interactions ui ON ta.podcast_id = ui.podcast_id AND ui.deleted_at IS NULL
	GROUP BY
		ta.podcast_id, ta.genre, ta.tags
	LIMIT $2
	OFFSET $3`

	rows, err := e.Db.Query(query, pq.Array(f.Genres), f.Pagination.Limit, f.Pagination.Offset)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var podcasts []*pb.Podcast
	for rows.Next() {
		var id string
		var listen, like int
		var tags, genre string
		err := rows.Scan(&id, &genre, &tags, &listen, &like)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		podcasts = append(podcasts, &pb.Podcast{
			PodcastId:   id,
			Genre:       splitNonEmptyString(genre, ','),
			Tags:        splitNonEmptyString(tags, ','),
			ListenCount: int64(listen),
			LikeCount:   int64(like),
		})
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return podcasts, nil
}

// splitNonEmptyString splits a string by a separator and removes empty strings
func splitNonEmptyString(s string, sep rune) []string {
	fields := strings.FieldsFunc(s, func(c rune) bool {
		return c == sep
	})
	result := []string{}
	for _, field := range fields {
		if field != "" {
			result = append(result, field)
		}
	}
	return result
}

func (e *EpisodeMetadataRepo) GetListensAndLikes(id string) (int, int, error) {
	query := `
	select
		listen_count,
		like_count
	from
		episode_metadata
	where
		deleted_at is null and episode_id = $1`

	var listens, likes int
	err := e.Db.QueryRow(query, id).Scan(&listens, &likes)
	if err != nil {
		return -1, -1, err
	}

	return listens, likes, nil
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

func (e *EpisodeMetadataRepo) GetPodcastsIdUserWatched(id *pb.ID) (
	*pb.PodcastsId, error) {
	queryToTakePodcastsId := `
	select 
		distinct podcast_id
	from 
		user_interactions
	where 
		user_id = $1`

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
