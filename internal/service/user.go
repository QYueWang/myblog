package service

import (
	"context"
	pb "myblog/api/v1/authn"
	"myblog/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type UserService struct {
	uc  *biz.UserUseCase
	log *log.Helper
	pb.UnimplementedUserServiceServer
}

func NewUserService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

func (us *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	token, err := us.uc.Login(ctx, req.Account, req.Password)
	return &pb.LoginReply{Token: token}, err
}
func (us *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	err := us.uc.CreateUser(ctx, req.UserName, req.Email, req.Password)
	return &pb.CreateUserReply{}, err
}
func (us *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	err := us.uc.UpdateUser(ctx, req.UserName, req.Email, req.Account)
	return &pb.UpdateUserReply{}, err
}
func (us *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	err := us.uc.DeleteUser(ctx, req.Account)
	return &pb.DeleteUserReply{}, err
}
func (us *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	user, err := us.uc.GetUser(ctx, req.Account)
	reply := &pb.GetUserReply{
		User: &pb.User{
			Account:  user.Account,
			UserName: user.UserName,
			Email:    user.Email,
		},
	}
	return reply, err
}
func (us *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	users, err := us.uc.ListUser(ctx)
	reply := pb.ListUserReply{}
	for _, user := range users {
		reply.Data = append(reply.Data, &pb.User{
			Account:  user.Account,
			UserName: user.UserName,
			Email:    user.Email,
		})
	}
	return &reply, err
}
