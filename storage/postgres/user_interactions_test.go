package postgres

import (
	"database/sql"
	pb "discovery_service/genproto/user_interactions"
	"testing"

	"github.com/google/uuid"
)

func TestLikeEpisodeOfPodcast(t *testing.T) {
	ids := pb.InteractEpisode{
		UserId:          "69106980-0bf6-44bd-9cb8-7cfe79383edf",
		PodcastId:       "da66f014-d648-40a3-91ac-fc0958b8688e",
		EpisodeId:       "0f627d19-1862-4949-935a-56f2979966d1",
		InteractionType: "like",
	}

	db, err := ConnectDB()
	if err != nil {
		t.Error(err)
	}
	got, err := NewUserInterRepo(db).LikeEpisodeOfPodcast(&ids)
	if err != nil || err == sql.ErrNoRows {
		t.Error(err)
	}
	if _, err := uuid.Parse(got.Id); err != nil {
		t.Error(err)
	}
}

func TestDeleteLikeFromEpisodeOfPodcast(t *testing.T) {
	ids := pb.DeleteLike{
		UserId:    "69106980-0bf6-44bd-9cb8-7cfe79383edf",
		PodcastId: "da66f014-d648-40a3-91ac-fc0958b8688e",
		EpisodeId: "0f627d19-1862-4949-935a-56f2979966d1",
	}

	db, err := ConnectDB()
	if err != nil {
		t.Error(err)
	}
	got, err := NewUserInterRepo(db).DeleteLikeFromEpisodeOfPodcast(&ids)
	if err != nil || err == sql.ErrNoRows {
		t.Error(err)
	}

	if !got.Success {
		t.Errorf("Cannot dislike")
	}
}

func TestListenEpisodeOfPodcast(t *testing.T) {
	ids := pb.InteractEpisode{
		UserId:          "69106980-0bf6-44bd-9cb8-7cfe79383edf",
		PodcastId:       "da66f014-d648-40a3-91ac-fc0958b8688e",
		EpisodeId:       "0f627d19-1862-4949-935a-56f2979966d1",
		InteractionType: "listen",
	}

	db, err := ConnectDB()
	if err != nil {
		t.Error(err)
	}
	got, err := NewUserInterRepo(db).ListenEpisodeOfPodcast(&ids)
	if err != nil || err == sql.ErrNoRows {
		t.Error(err)
	}

	if _, err := uuid.Parse(got.Id); err != nil {
		t.Error(err)
	}
}
