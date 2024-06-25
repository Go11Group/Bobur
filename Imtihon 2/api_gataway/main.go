package main

import (
	"api_gateway/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	client, err := grpc.NewClient(":8088", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	r := api.InitRouter(client)

	r.Run(":8081")

}
