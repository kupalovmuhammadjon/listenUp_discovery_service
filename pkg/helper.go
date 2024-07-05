package pkg

import (
	"discovery_service/config"
	pbCollaboration "discovery_service/genproto/collaborations"
	pbComment "discovery_service/genproto/comments"
	pbe "discovery_service/genproto/episodes"
	pbp "discovery_service/genproto/podcasts"
	pbu "discovery_service/genproto/user"
	"errors"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreateUserManagementClient(cfg *config.Config) pbu.UserManagementClient {
	conn, err := grpc.NewClient(cfg.USER_SERVICE_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(errors.New("failed to connect to the address: " + err.Error()))
	}
	defer conn.Close()

	u := pbu.NewUserManagementClient(conn)
	return u
}

func CreatePodcastsClient(cfg *config.Config) pbp.PodcastsClient {
	conn, err := grpc.NewClient(cfg.PODCAST_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println("error while connecting podcast service ", err)
	}
	a := pbp.NewPodcastsClient(conn)

	return a
}

func CreateEpisodesClient(cfg *config.Config) pbe.EpisodesServiceClient {
	conn, err := grpc.NewClient(cfg.EPISODE_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println("error while connecting podcast service ", err)
	}
	a := pbe.NewEpisodesServiceClient(conn)

	return a
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
