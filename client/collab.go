package client

import (
	"discovery_service/config"
	pb "discovery_service/genproto/comments"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CommentsClient(podcastID string) (*pb.CommentsClient, error) {
	cfg := config.Load()
	conn, err := grpc.NewClient(cfg.COLLABORATION_SERVICE_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.New("failed to connect to the address: " + err.Error())
	}
	defer conn.Close()

	c := pb.NewCommentsClient(conn)

	return &c, nil
}
