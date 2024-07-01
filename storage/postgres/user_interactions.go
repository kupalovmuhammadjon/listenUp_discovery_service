package postgres

import (
	"database/sql"
	pb "discovery_service/genproto/user_interactions"
)

type UserInterRepo struct {
	Db *sql.DB
}

func NewUserInterRepo(db *sql.DB) *UserInterRepo {
	return &UserInterRepo{Db: db}
}

func (u *UserInterRepo) LikeEpisodeOfPodcast(ids *pb.InteractEpisode) (*pb.ID, error) {
	query := `insert into user_interactions(user_id, podcast_id, episode_id, interaction_type)
	values ($1, $2, $3, $4) returning id`

	tr, err := u.Db.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			tr.Rollback()
		} else {
			tr.Commit()
		}
	}()

	var id string
	row := tr.QueryRow(query, ids.UserId, ids.PodcastId, ids.EpisodeId, ids.InteractionType)
	err = row.Scan(&id)
	if err != nil {
		return nil, err
	}

	return &pb.ID{Id: id}, nil
}

func (u *UserInterRepo) DeleteLikeFromEpisodeOfPodcast(ids *pb.DeleteLike) (*pb.Success, error) {
	query := `delete from user_interactions
	where deleted_at = null and user_id = $1 and podcast_id = $2 and episode_id = $3`

	tr, err := u.Db.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			tr.Rollback()
		} else {
			tr.Commit()
		}
	}()

	_, err = tr.Exec(query, ids.UserId, ids.PodcastId, ids.EpisodeId)
	if err != nil {
		return &pb.Success{Success: false}, err
	}

	return &pb.Success{Success: true}, nil
}

func (u *UserInterRepo) ListenEpisodeOfPodcast(ids *pb.InteractEpisode) (*pb.ID, error) {
	query := `insert into user_interactions(user_id, podcast_id, episode_id, interaction_type)
	values ($1, $2, $3, $4) returning id`

	tr, err := u.Db.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			tr.Rollback()
		} else {
			tr.Commit()
		}
	}()

	var id string
	row := tr.QueryRow(query, ids.UserId, ids.PodcastId, ids.EpisodeId, ids.InteractionType)
	err = row.Scan(&id)
	if err != nil {
		return nil, err
	}

	return &pb.ID{Id: id}, nil
}
