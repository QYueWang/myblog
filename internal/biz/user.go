package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Account  string
	UserName string
	Email    string
}

type UserRepo interface {
	Login(ctx context.Context, account, password string) (string, error)
	CreateUser(ctx context.Context, userName, email, password string) error
	UpdateUser(ctx context.Context, userName, email, account string) error
	DeleteUser(ctx context.Context, account string) error
	GetUser(ctx context.Context, account string) (*User, error)
	ListUser(ctx context.Context) ([]*User, error)
}

type UserUseCase struct {
	log  log.Helper
	repo UserRepo
}

func NewUserUseCase(reop UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{log: *log.NewHelper(logger), repo: reop}
}

func (uc *UserUseCase) Login(ctx context.Context, account, password string) (string, error) {
	return uc.repo.Login(ctx, account, password)
}

func (uc *UserUseCase) CreateUser(ctx context.Context, userName, email, passpword string) error {
	return uc.repo.CreateUser(ctx, userName, email, passpword)
}

func (uc *UserUseCase) UpdateUser(ctx context.Context, userName, email, account string) error {
	return uc.repo.UpdateUser(ctx, userName, email, account)
}

func (uc *UserUseCase) DeleteUser(ctx context.Context, account string) error {
	return uc.repo.DeleteUser(ctx, account)
}

func (uc *UserUseCase) GetUser(ctx context.Context, account string) (*User, error) {
	return uc.repo.GetUser(ctx, account)
}

func (uc *UserUseCase) ListUser(ctx context.Context) ([]*User, error) {
	return uc.repo.ListUser(ctx)
}
