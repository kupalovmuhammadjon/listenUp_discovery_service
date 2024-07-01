package postgres

import "database/sql"

type EpisodeMetadataRepo struct {
	Db *sql.DB
}

func NewEpisodeMetadataRepo(db *sql.DB) *EpisodeMetadataRepo {
	return &EpisodeMetadataRepo{Db: db}
}
