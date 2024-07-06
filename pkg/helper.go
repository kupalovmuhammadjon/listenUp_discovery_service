package pkg

import (
	"discovery_service/config"
	pbCollab "discovery_service/genproto/collaborations"
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
	conn, err := grpc.NewClient(cfg.USER_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(errors.New("failed to connect to the address: " + err.Error()))
		return nil
	}

	return pbu.NewUserManagementClient(conn)
}

func CreatePodcastsClient(cfg *config.Config) pbp.PodcastsClient {
	conn, err := grpc.NewClient(cfg.PODCAST_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(errors.New("failed to connect to the address: " + err.Error()))
		return nil
	}

	return pbp.NewPodcastsClient(conn)
}

func CreateEpisodesClient(cfg *config.Config) pbe.EpisodesServiceClient {
	conn, err := grpc.NewClient(cfg.PODCAST_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(errors.New("failed to connect to the address: " + err.Error()))
		return nil
	}

	return pbe.NewEpisodesServiceClient(conn)
}

func CreateCollaborationsClient(cfg *config.Config) pbCollab.CollaborationsClient {
	conn, err := grpc.NewClient(cfg.COLLABORATION_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(errors.New("failed to connect to the address: " + err.Error()))
		return nil
	}

	return pbCollab.NewCollaborationsClient(conn)
}

func CreateCommentsClient(cfg *config.Config) pbComment.CommentsClient {
	conn, err := grpc.NewClient(cfg.COLLABORATION_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(errors.New("failed to connect to the address: " + err.Error()))
		return nil
	}

	return pbComment.NewCommentsClient(conn)
}
