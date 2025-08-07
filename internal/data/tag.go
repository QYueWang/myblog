package data

import (
	"context"
	"myblog/internal/biz"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type TagRepo struct {
	data *Data
	//没有log.Helper对象，全部使用框架默认初始化好的log.DefaultLogger实例
}

func NewTagRepo(data *Data) biz.TagRepo {
	return &TagRepo{data: data}
}

func (tr *TagRepo) CreateTag(ctx context.Context, name string) error {
	if _, err := tr.data.db.Tag.Create().SetID("22222").SetName(name).Save(ctx); err != nil {
		log.DefaultLogger.Log(log.LevelError, "创建标签失败")
		return err
	}

	return nil
}

func (tr *TagRepo) UpdateTag(ctx context.Context, id, name string) error {
	if _, err := tr.data.db.Tag.UpdateOneID(id).SetName(name).SetUpdateAt(time.Now()).Save(ctx); err != nil {
		log.DefaultLogger.Log(log.LevelError, "更新标签失败")
		return err
	}
	return nil
}

func (tr *TagRepo) DeleteTag(ctx context.Context, id string) error {
	if err := tr.data.db.Tag.DeleteOneID(id).Exec(ctx); err != nil {
		log.DefaultLogger.Log(log.LevelError, "删除标签失败")
		return err
	}
	return nil
}

func (tr *TagRepo) GetTag(ctx context.Context, article string) ([]*biz.Tag, error) {
	tags := make([]*biz.Tag, 0)
	ts, err := tr.data.db.Tag.Query().Select(article).All(ctx)
	if err != nil {
		log.DefaultLogger.Log(log.LevelError, "查询标签失败")
		return nil, err
	}
	for _, t := range ts {
		tags = append(tags, &biz.Tag{
			Id:       t.ID,
			Name:     t.Name,
			CreateAt: t.CreateAt,
			UpdateAt: t.UpdateAt,
		})
	}
	return tags, nil
}

func (tr *TagRepo) ListTag(ctx context.Context) ([]*biz.Tag, error) {
	tags := make([]*biz.Tag, 0)
	ts, err := tr.data.db.Tag.Query().All(ctx)
	if err != nil {
		log.DefaultLogger.Log(log.LevelError, "查询标签失败")
		return nil, err
	}
	for _, t := range ts {
		tags = append(tags, &biz.Tag{
			Id:       t.ID,
			Name:     t.Name,
			CreateAt: t.CreateAt,
			UpdateAt: t.UpdateAt,
		})
	}
	return tags, nil
}
