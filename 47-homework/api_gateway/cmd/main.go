package main

import (
	"fmt"
	"model/api_gateway/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("error --> ", err.Error())
		return
	}

	defer conn.Close()

	r := api.NewRouter(conn)
	if err := r.Run(); err != nil {
		panic(err)
	}

}
