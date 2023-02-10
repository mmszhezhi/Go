package main

import (
	"Go/grpc2/protobuf/user"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8889", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client := user.NewUserServiceClient(conn)

	req := &user.UserRequest{
		Id: 1,
	}

	response, _ := client.GetUser(context.Background(), req)

	resp, err := json.Marshal(response)

	fmt.Printf("%s", resp)

	fmt.Println("~~~~~~~~~~~~~~~~~~~")
	response2, _ := client.FuckUser(context.Background(), req)

	resp2, err := json.Marshal(response2)

	fmt.Printf("%s", resp2)

}
