package server

import (
	"context"
	ar "myblog/api/v1/article"
	us "myblog/api/v1/authn"
	co "myblog/api/v1/comment"
	tg "myblog/api/v1/tag"
	"myblog/internal/conf"
	"myblog/internal/service"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	kjwt "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/golang-jwt/jwt/v5"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
	c *conf.Server,
	article *service.ArticleService,
	comment *service.CommentService,
	tag *service.TagService,
	user *service.UserService,
	logger log.Logger,
) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			selector.Server(
				kjwt.Server(func(t *jwt.Token) (interface{}, error) {
					return []byte("serverkey"), nil
				}),
			).Match(func(ctx context.Context, operation string) bool {
				return !strings.HasPrefix("/api.v1.user.UserService/Login", operation)
			}).Build(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	ar.RegisterArticleServiceServer(srv, article)
	co.RegisterCommentServiceServer(srv, comment)
	tg.RegisterTagServiceServer(srv, tag)
	us.RegisterUserServiceServer(srv, user)
	return srv
}
