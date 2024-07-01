package pkg

import (
	"discovery_service/config"
	pbC "discovery_service/genproto/collaborations"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreateCollaborationsClient() (pbC.CollaborationsClient, error) {
	config := config.Load()
	conn, err := grpc.NewClient(config.DISCOVERY_SERVICE_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.New("failed to connect to the address: " + err.Error())
	}
	defer conn.Close()

	client := pbC.NewCollaborationsClient(conn)
	return client, nil
}
