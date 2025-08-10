package data

import (
	"context"
	"myblog/internal/biz"
	"myblog/internal/data/ent"
	"myblog/internal/data/ent/user"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v5"
)

type UserRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &UserRepo{data: data, log: log.NewHelper(logger)}
}

func (ur *UserRepo) Login(ctx context.Context, account, password string) (string, error) {
	user, err := ur.data.db.User.Query().Where(user.AccountEQ(account)).Only(ctx)
	if user.Password == password {
		t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{Subject: account})
		token, err := t.SignedString([]byte("serverkey"))
		if err != nil {
			return "Token生成失败", err
		}
		return token, nil
	}
	return "", err
}

func (ur *UserRepo) CreateUser(ctx context.Context, userName, email, password string) error {
	_, err := ur.data.db.User.Create().
		SetAccount("123").
		SetName(userName).
		SetEmail(email).
		SetPassword(password).
		SetCreateAt(time.Now()).
		SetUpdateAt(time.Now()).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepo) UpdateUser(ctx context.Context, userName, email, account string) error {
	err := ur.data.db.User.UpdateOne(&ent.User{
		Account:  account,
		Name:     userName,
		Email:    email,
		UpdateAt: time.Now(),
	}).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepo) DeleteUser(ctx context.Context, account string) error {
	err := ur.data.db.User.DeleteOne(&ent.User{
		Account: account,
	}).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepo) GetUser(ctx context.Context, account string) (*biz.User, error) {
	u, err := ur.data.db.User.Query().Select(account).First(ctx)
	if err != nil {
		return nil, err
	}
	user := &biz.User{
		Account:  u.Account,
		UserName: u.Name,
		Email:    u.Email,
	}
	return user, nil
}

func (ur *UserRepo) ListUser(ctx context.Context) ([]*biz.User, error) {
	us, err := ur.data.db.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	users := make([]*biz.User, 0)
	for _, u := range us {
		users = append(users, &biz.User{
			Account:  u.Account,
			UserName: u.Name,
			Email:    u.Email,
		})
	}
	return users, nil
}
