package server

import (
	ar "myblog/api/v1/article"
	co "myblog/api/v1/comment"
	tg "myblog/api/v1/tag"
	"myblog/internal/conf"
	"myblog/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, article *service.ArticleService, comment *service.CommentService, tag *service.TagService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
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
	return srv
}
