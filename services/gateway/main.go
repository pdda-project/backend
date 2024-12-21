package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	usersPb "github.com/pdda-project/backend/services/users/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	userServiceHost := os.Getenv("USER_SERVICE_HOST")

	conn, err := grpc.NewClient(
		userServiceHost,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()

	client := usersPb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := usersPb.RegisterUserRequest{
		Email:       "qodar@gmail.co.id",
		Password:    "secretpassword",
		DisplayName: "Nurfian Qodar",
	}
	resp, err := client.RegisterUser(ctx, &req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.GetUid())

	req2 := &usersPb.GetUserRequest{
		Uid: resp.Uid,
	}
	resp2, err := client.GetUser(ctx, req2)
	if err != nil {
		log.Fatal(err)
	}
	j, err := json.Marshal(resp2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(j))
	fmt.Println(resp2)
}
