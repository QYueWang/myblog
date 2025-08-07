package data

import (
	"context"
	"myblog/internal/conf"
	"myblog/internal/data/ent"

	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewArticleRepo, NewCommentRepo, NewTagRepo)

// Data
type Data struct {
	db *ent.Client
}

// NewData
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	db, err := ent.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		logger.Log(log.LevelError, err)
	}
	//执行数据迁移
	if err := db.Schema.Create(context.Background()); err != nil {
		logger.Log(log.LevelError, err)
	}

	cleanup := func() {
		if err := db.Close(); err != nil {
			logger.Log(log.LevelError, err)
		}
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db}, cleanup, nil
}
