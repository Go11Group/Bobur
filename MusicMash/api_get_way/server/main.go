package main

import (
	"api_get_way/api"
	"api_get_way/genproto"
	"log"

	//"github.com/Go11Group/at_lesson/lesson47/api_gateway/api"
	//"github.com/Go11Group/at_lesson/lesson47/api_gateway/genproto"
	//"github.com/grpc/grpc-go/credentials/insecure"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	collaborationClient := genproto.NewCollaborationServiceClient(conn)
	compositionClient := genproto.NewCompositionServiceClient(conn)
	discoveryClient := genproto.NewDiscoveryServiceClient(conn)

	r := api.RouterAPi(compositionClient, collaborationClient, discoveryClient)

	r.Run(":8090")
}
