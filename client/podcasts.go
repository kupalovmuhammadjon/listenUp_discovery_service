package client

import (
	"context"
	"discovery_service/config"
	pbE "discovery_service/genproto/episodes"
	pbP "discovery_service/genproto/podcasts"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RetrieveTitles(podcastIDs []string) ([]*pbP.Podcast, map[string]*pbE.Episodes, error) {
	cfg := config.Load()
	conn, err := grpc.NewClient(cfg.PODCAST_SERVICE_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, errors.New("failed to connect to the address: " + err.Error())
	}
	defer conn.Close()

	p := pbP.NewPodcastsClient(conn)
	e := pbE.NewEpisodesServiceClient(conn)

	podcasts, err := GetPodcasts(p, podcastIDs)
	if err != nil {
		return nil, nil, errors.New("failed to fetch podcast titles: " + err.Error())
	}

	episodes, err := GetEpisodes(e, podcastIDs)
	if err != nil {
		return nil, nil, errors.New("failed to fetch episode titles: " + err.Error())
	}

	return podcasts, episodes, nil
}

func GetPodcasts(p pbP.PodcastsClient, ids []string) ([]*pbP.Podcast, error) {
	var res []*pbP.Podcast

	for _, v := range ids {
		podcast, err := p.GetPodcastById(context.Background(), &pbP.ID{Id: v})
		if err != nil {
			return nil, err
		}
		res = append(res, podcast)
	}

	return res, nil
}

func GetEpisodes(e pbE.EpisodesServiceClient, ids []string) (map[string]*pbE.Episodes, error) {
	res := make(map[string]*pbE.Episodes, 0)

	for _, v := range ids {
		episodes, err := e.GetEpisodesByPodcastId(context.Background(), &pbE.ID{Id: v})
		if err != nil {
			return nil, err
		}
		res[v] = episodes
	}

	return res, nil
}
