package pkg

import (
	"discovery_service/config"
	pbCollaboration "discovery_service/genproto/collaborations"
	pbComment "discovery_service/genproto/comments"
	pbPodcast "discovery_service/genproto/podcasts"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

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
