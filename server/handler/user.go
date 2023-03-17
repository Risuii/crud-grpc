package handler

import (
	"context"
	pb "crud-grpc/proto"
	"crud-grpc/server/model"
	"crud-grpc/server/usecase"
)

type UserServer struct {
	Usecase usecase.UserClass
	pb.UnsafeUserServiceServer
}

func NewUserServer(userUseCase usecase.UserClass) *UserServer {
	return &UserServer{
		Usecase: userUseCase,
	}
}

func (us *UserServer) AddUser(ctx context.Context, user *pb.AddUserReq) (*pb.AddUserResp, error) {
	us.Usecase.AddUser(model.User{
		ID:    int(user.User.GetId()),
		Name:  user.User.GetName(),
		Age:   int(user.User.GetAge()),
		Email: user.User.GetEmail(),
	})

	return &pb.AddUserResp{
		Success: true,
	}, ctx.Err()
}

func (us *UserServer) FindUserByID(ctx context.Context, id *pb.GetUserReq) (*pb.GetUserResp, error) {
	data := us.Usecase.FindUserByID(int(id.GetId()))

	resp := &pb.GetUserResp{
		User: &pb.User{
			Id:    int64(data.ID),
			Name:  data.Name,
			Age:   int64(data.Age),
			Email: data.Email,
		},
	}

	return resp, nil
}

func (us *UserServer) GetAll(ctx context.Context, user *pb.GetAllReq) (*pb.GetAllResp, error) {

	data := us.Usecase.GetAll()

	resp := &pb.GetAllResp{}

	for _, val := range data {
		resp.Users = append(resp.Users, &pb.User{
			Id:    int64(val.ID),
			Name:  val.Name,
			Age:   int64(val.Age),
			Email: val.Email,
		})
	}

	return resp, nil
}

// func (us *UserServer) UpdateUser(ctx context.Context, user *pb.UpdateReq) (*pb.UpdateResp, error) {
// 	return nil, nil
// }

// func (us *UserServer) DeleteUser(ctx context.Context, user *pb.DeleteResp) (*pb.DeleteResp, error) {
// 	return nil, nil
// }
