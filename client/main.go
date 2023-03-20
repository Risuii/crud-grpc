package main

import (
	"context"
	pb "crud-grpc/proto"
	"crud-grpc/server/model"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
)

func addDataUser(client pb.UserServiceClient, data model.User) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	s := &pb.AddUserReq{
		User: &pb.User{
			Id:    int64(data.ID),
			Name:  data.Name,
			Age:   int64(data.Age),
			Email: data.Email,
		},
	}

	newUser, err := client.AddUser(ctx, s)
	if err != nil {
		log.Fatalln("error when add user", err)
	}

	fmt.Println(newUser)
}

func getAllDataUser(client pb.UserServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s := &pb.GetAllReq{}

	user, err := client.GetAll(ctx, s)
	if err != nil {
		log.Fatalln("error when get user", err)
	}

	fmt.Println(user)
}

func getOneUser(client pb.UserServiceClient, id int) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	s := &pb.GetUserReq{
		Id: int64(id),
	}

	data, err := client.FindUserByID(ctx, s)
	if err != nil {
		log.Fatalln("error when get one user")
	}

	fmt.Println(data)
}

func updateUser(client pb.UserServiceClient, id int, data model.User) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	s := &pb.UpdateReq{
		UserUpdate: &pb.User{
			Id:    int64(id),
			Name:  data.Name,
			Age:   int64(data.Age),
			Email: data.Email,
		},
	}

	updateUser, err := client.UpdateUser(ctx, s)
	if err != nil {
		log.Fatalln("error when update user", err)
	}

	fmt.Println(updateUser)
}

func deleteUser(client pb.UserServiceClient, id int) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	s := &pb.DeleteUserReq{
		Id: int64(id),
	}

	data, err := client.DeleteUser(ctx, s)
	if err != nil {
		log.Fatalln("error when delete one user", err)
	}

	fmt.Println(data)
}

func main() {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(":8080", opts...)
	if err != nil {
		log.Fatalln("error in dial", err.Error())
	}

	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	data := model.User{
		ID:    1,
		Name:  "test-1",
		Age:   25,
		Email: "test-1@test.com",
	}

	data2 := model.User{
		ID:    2,
		Name:  "test-2",
		Age:   30,
		Email: "test-2@test.com",
	}

	data3 := model.User{
		ID:    3,
		Name:  "test-3",
		Age:   35,
		Email: "test-3@test.com",
	}

	// dataUpdate := model.User{
	// 	Name:  "test-update",
	// 	Age:   40,
	// 	Email: "test-update@test.com",
	// }

	addDataUser(client, data)
	addDataUser(client, data2)
	addDataUser(client, data3)

	// getAllDataUser(client)
	// getOneUser(client, 2)

	// updateUser(client, 1, dataUpdate)

	// deleteUser(client, 2)
}
