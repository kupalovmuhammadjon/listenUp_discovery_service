package client

import (
	"context"
	"discovery_service/config"
	pb "discovery_service/genproto/comments"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RetrieveComments(podcastID string) (*pb.AllComments, error) {
	cfg := config.Load()
	conn, err := grpc.NewClient(cfg.COLLABORATION_SERVICE_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.New("failed to connect to the address: " + err.Error())
	}
	defer conn.Close()

	c := pb.NewCommentsClient(conn)

	comments, err := c.GetCommentsByPodcastId(context.Background(), &pb.ID{Id: podcastID})
	if err != nil {
		return nil, err
	}

	return comments, nil
}
