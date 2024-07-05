package postgres

import (
	pb "discovery_service/genproto/episode_metadata"
	"log"
	"reflect"
	"testing"
)

func DB() *EpisodeMetadataRepo {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	return NewEpisodeMetadataRepo(db)
}

// func TestCreateEpisodeMetaData(t *testing.T) {
// 	e := DB()
// 	err := e.CreateEpisodeMetaData(&pb.EpisodeMetadata{
// 		EpisodeId: "1ce9e59b-d985-4a28-ab46-801fb8102760",
// 		PodcastId: "04c70d01-930c-45b0-9e6e-161fc5f39175",
// 		Genre: "philosophy",
// 		Tags: []string{"Epistemology", "Logic"},
// 	})

// 	if err != nil {
// 		t.Errorf("error while creating episode metadata: %v\n", err)
// 	}
// }

func TestGetTrendingPodcasts(t *testing.T) {
	e := DB()
	res, err := e.GetTrendingPodcasts(&pb.Pagination{})
	if err != nil {
		t.Errorf("error while getting trending podcasts: %v", err)
	}

	exp := pb.Podcasts{
		Podcasts: []*pb.Podcast{
			{PodcastId: "04c70d01-930c-45b0-9e6e-161fc5f39175", Genre: "philosophy"},
			{PodcastId: "323e4567-e89b-12d3-a456-426614174002", Genre: "education", ListenCount: 75, LikeCount: 25},
			{PodcastId: "923e4567-e89b-12d3-a456-426614174008", Genre: "religion", ListenCount: 85, LikeCount: 40},
			{PodcastId: "523e4567-e89b-12d3-a456-426614174004", Genre: "politics", ListenCount: 90, LikeCount: 45},
			{PodcastId: "223e4567-e89b-12d3-a456-426614174001", Genre: "technology", ListenCount: 100, LikeCount: 50},
			{PodcastId: "723e4567-e89b-12d3-a456-426614174006", Genre: "health", ListenCount: 95, LikeCount: 55},
			{PodcastId: "623e4567-e89b-12d3-a456-426614174005", Genre: "business", ListenCount: 110, LikeCount: 60},
			{PodcastId: "a23e4567-e89b-12d3-a456-426614174009", Genre: "technology", ListenCount: 105, LikeCount: 65},
			{PodcastId: "823e4567-e89b-12d3-a456-426614174007", Genre: "sports", ListenCount: 130, LikeCount: 70},
			{PodcastId: "b23e4567-e89b-12d3-a456-426614174010", Genre: "education", ListenCount: 125, LikeCount: 75},
			{PodcastId: "423e4567-e89b-12d3-a456-426614174003", Genre: "philosophy", ListenCount: 120, LikeCount: 80},
		},
	}

	if !reflect.DeepEqual(res, &exp) {
		t.Errorf("error:\ngot: %v\nwant: %v\n", res, &exp)
	}
}

func TestGetRecommendedPodcasts(t *testing.T) {
	e := DB()
	res, err := e.GetRecommendedPodcasts("823e4567-e89b-12d3-a456-426614174007",
		&pb.Pagination{})
	if err != nil {
		t.Errorf("error while getting recommended podcasts: %v", err)
	}

	exp := pb.Podcasts{
		Podcasts: []*pb.Podcast{
			{
				PodcastId:   "823e4567-e89b-12d3-a456-426614174007",
				Genre:       "sports",
				Tags:        []string{"Football", "Basketball"},
				ListenCount: 130,
				LikeCount:   70,
			},
		},
	}

	if !reflect.DeepEqual(res, &exp) {
		t.Errorf("error:\ngot: %v\nwant: %v\n", res, &exp)
	}
}

func TestGetPodcastsByGenre(t *testing.T) {
	e := DB()
	res, err := e.GetPodcastsByGenre(&pb.Filter{
		Genres:     []string{"sports", "education"},
		Pagination: &pb.Pagination{Limit: 1, Offset: 1}})
	if err != nil {
		t.Errorf("error while getting podcasts by genre: %v", err)
	}

	exp := pb.Podcasts{
		Podcasts: []*pb.Podcast{
			{
				PodcastId:   "823e4567-e89b-12d3-a456-426614174007",
				Genre:       "sports",
				Tags:        []string{"Football", "Basketball"},
				ListenCount: 1,
				LikeCount:   0,
			},
		},
	}

	if !reflect.DeepEqual(res, exp.Podcasts) {
		t.Errorf("error:\ngot: %v\nwant: %v\n", res, exp.Podcasts)
	}
}

func TestGetListensAndLikes(t *testing.T) {
	e := DB()
	res1, res2, err := e.GetListensAndLikes("123e4567-e89b-12d3-a456-426614174000")
	if err != nil {
		t.Errorf("error while getting listens and likes: %v", err)
	}

	exp1, exp2 := 100, 50

	if res1 != exp1 || res2 != exp2 {
		t.Errorf("error:\ngot: %d and %d\nwant: %d and %d\n", res1, res2, exp1, exp2)
	}
}

func TestGetPodcastIDs(t *testing.T) {
	e := DB()
	res, err := e.GetPodcastIDs()
	if err != nil {
		t.Errorf("error while getting podcast ids: %v", err)
	}

	exp := []string{
		"223e4567-e89b-12d3-a456-426614174001",
		"323e4567-e89b-12d3-a456-426614174002",
		"423e4567-e89b-12d3-a456-426614174003",
		"523e4567-e89b-12d3-a456-426614174004",
		"623e4567-e89b-12d3-a456-426614174005",
		"723e4567-e89b-12d3-a456-426614174006",
		"823e4567-e89b-12d3-a456-426614174007",
		"a23e4567-e89b-12d3-a456-426614174009",
		"b23e4567-e89b-12d3-a456-426614174010",
		"923e4567-e89b-12d3-a456-426614174008",
		"04c70d01-930c-45b0-9e6e-161fc5f39175",
	}

	if !reflect.DeepEqual(res, exp) {
		t.Errorf("error:\ngot: %v\nwant: %v\n", res, exp)
	}
}

func TestGetPodcastsIdUserWatched(t *testing.T) {
	e := DB()
	res, err := e.GetPodcastsIdUserWatched(&pb.ID{Id: "123e4567-e89b-12d3-a456-426614174100"})
	if err != nil {
		t.Errorf("error while getting listened podcast ids: %v", err)
	}

	exp := pb.PodcastsId{
		PodcastsId: []string{
			"223e4567-e89b-12d3-a456-426614174001",
			"523e4567-e89b-12d3-a456-426614174004",
			"923e4567-e89b-12d3-a456-426614174008",
			"a23e4567-e89b-12d3-a456-426614174009",
		},
	}

	if !reflect.DeepEqual(res, &exp) {
		t.Errorf("error:\ngot: %v\nwant: %v\n", res, &exp)
	}
}
