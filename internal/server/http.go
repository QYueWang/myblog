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
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/golang-jwt/jwt/v5"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	c *conf.Server,
	article *service.ArticleService,
	comment *service.CommentService,
	tag *service.TagService,
	user *service.UserService,
	logger log.Logger,
) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			selector.Server(
				kjwt.Server(func(t *jwt.Token) (interface{}, error) {
					return []byte("serverkey"), nil
				}),
			).Match(func(ctx context.Context, operation string) bool {
				log.Log(log.LevelDebug, operation)
				return !strings.HasPrefix("/api.v1.user.UserService/Login", operation)
			}).Build(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	ar.RegisterArticleServiceHTTPServer(srv, article)
	co.RegisterCommentServiceHTTPServer(srv, comment)
	tg.RegisterTagServiceHTTPServer(srv, tag)
	us.RegisterUserServiceHTTPServer(srv, user)
	return srv
}
