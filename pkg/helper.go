package pkg

import (
	"discovery_service/config"
	pbCollaboration "discovery_service/genproto/collaborations"
	pbComment "discovery_service/genproto/comments"
	pbEpisode "discovery_service/genproto/episodes"
	pbPodcast "discovery_service/genproto/podcasts"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreatePodcastClient() (pbPodcast.PodcastsClient, error) {
	config := config.Load()
	conn, err := grpc.NewClient(config.PODCAST_SERVICE_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.New("failed to connect to the address: " + err.Error())
	}
	defer conn.Close()

	client := pbPodcast.NewPodcastsClient(conn)
	return client, nil
}

func CreateEpisodesClient() (pbEpisode.EpisodesServiceClient, error) {
	config := config.Load()
	conn, err := grpc.NewClient(config.EPISODE_SERVICE_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.New("failed to connect to the address: " + err.Error())
	}
	defer conn.Close()

	client := pbEpisode.NewEpisodesServiceClient(conn)
	return client, nil
}

func CreateCollaborationsClient() (pbCollaboration.CollaborationsClient, error) {
	config := config.Load()
	conn, err := grpc.NewClient(config.COLLABORATION_SERVICE_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.New("failed to connect to the address: " + err.Error())
	}
	defer conn.Close()

	client := pbCollaboration.NewCollaborationsClient(conn)
	return client, nil
}

func CreateCommentsClient() (pbComment.CommentsClient, error) {
	config := config.Load()
	conn, err := grpc.NewClient(config.COLLABORATION_SERVICE_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.New("failed to connect to the address: " + err.Error())
	}
	defer conn.Close()

	client := pbComment.NewCommentsClient(conn)
	return client, nil
}
