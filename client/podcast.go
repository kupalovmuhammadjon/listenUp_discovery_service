package client

import (
	"discovery_service/config"
	pbE "discovery_service/genproto/episodes"
	pbP "discovery_service/genproto/podcasts"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func PodcastsEpisodesClient(podcastIDs []string) (*pbP.PodcastsClient, *pbE.EpisodesServiceClient, error) {
	cfg := config.Load()
	conn, err := grpc.NewClient(cfg.PODCAST_SERVICE_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, errors.New("failed to connect to the address: " + err.Error())
	}
	defer conn.Close()

	p := pbP.NewPodcastsClient(conn)
	e := pbE.NewEpisodesServiceClient(conn)

	return &p, &e, nil
}
