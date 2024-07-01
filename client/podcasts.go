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

func RetrieveTitles(podcastIDs []string) (map[string]string, map[string][]string, error) {
	cfg := config.Load()
	conn, err := grpc.NewClient(cfg.PODCAST_SERVICE_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, errors.New("failed to connect to the address: " + err.Error())
	}
	defer conn.Close()

	p := pbP.NewPodcastsClient(conn)
	e := pbE.NewEpisodesServiceClient(conn)

	podcasts, err := GetPodcastTitles(p, podcastIDs)
	if err != nil {
		return nil, nil, errors.New("failed to fetch podcast titles: " + err.Error())
	}

	episodes, err := GetEpisodeTitles(e, podcastIDs)
	if err != nil {
		return nil, nil, errors.New("failed to fetch episode titles: " + err.Error())
	}

	return podcasts, episodes, nil
}

func GetPodcastTitles(p pbP.PodcastsClient, ids []string) (map[string]string, error) {
	var res = make(map[string]string, 0)

	for _, v := range ids {
		podcast, err := p.GetPodcastById(context.Background(), &pbP.ID{Id: v})
		if err != nil {
			return nil, err
		}
		res[v] = podcast.Title
	}

	return res, nil
}

func GetEpisodeTitles(e pbE.EpisodesServiceClient, ids []string) (map[string][]string, error) {
	res := make(map[string][]string, 0)

	for _, v := range ids {
		episodes, err := e.GetEpisodesByPodcastId(context.Background(), &pbE.ID{Id: v})
		if err != nil {
			return nil, err
		}

		for _, e := range episodes.Episodes {
			res[v] = append(res[v], e.Title)
		}
	}

	return res, nil
}
