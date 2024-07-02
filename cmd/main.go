package main

import (
	"discovery_service/config"
	pbe "discovery_service/genproto/episode_metadata"
	pbu "discovery_service/genproto/user_interactions"
	"discovery_service/service"
	"discovery_service/storage/postgres"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", config.Load().DISCOVERY_SERVICE_PORT)
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	pbe.RegisterEpisodeMetadataServer(server, service.NewEpisodeMetadata(db))
	pbu.RegisterUserInteractionsServer(server, service.NewUserInterService(db))

	fmt.Printf("server is listening on port %s", config.Load().DISCOVERY_SERVICE_PORT)
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
