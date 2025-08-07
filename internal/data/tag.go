package data

import (
	"context"
	"myblog/internal/biz"
	"time"
)

type TagRepo struct {
	data *Data
}

func NewTagRepo(data *Data) biz.TagRepo {
	return &TagRepo{data: data}
}

func (tr *TagRepo) CreateTag(ctx context.Context, name string) error {
	_, err := tr.data.db.Tag.Create().SetID("22222").SetName(name).Save(ctx)
	return err
}

func (tr *TagRepo) UpdateTag(ctx context.Context, id, name string) error {
	_, err := tr.data.db.Tag.UpdateOneID(id).SetName(name).SetUpdateAt(time.Now()).Save(ctx)
	return err
}

func (tr *TagRepo) DeleteTag(ctx context.Context, id string) error {
	err := tr.data.db.Tag.DeleteOneID(id).Exec(ctx)
	return err
}

func (tr *TagRepo) GetTag(ctx context.Context, article string) ([]*biz.Tag, error) {
	tags := make([]*biz.Tag, 0)
	ts, err := tr.data.db.Tag.Query().Select(article).All(ctx)
	for _, t := range ts {
		tags = append(tags, &biz.Tag{
			Id:       t.ID,
			Name:     t.Name,
			CreateAt: t.CreateAt,
			UpdateAt: t.UpdateAt,
		})
	}
	return tags, err
}

func (tr *TagRepo) ListTag(ctx context.Context) ([]*biz.Tag, error) {
	tags := make([]*biz.Tag, 0)
	ts, err := tr.data.db.Tag.Query().All(ctx)
	for _, t := range ts {
		tags = append(tags, &biz.Tag{
			Id:       t.ID,
			Name:     t.Name,
			CreateAt: t.CreateAt,
			UpdateAt: t.UpdateAt,
		})
	}
	return tags, err
}
