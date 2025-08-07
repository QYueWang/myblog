package biz

import (
	"context"
	"time"
)

type Tag struct {
	Id       string
	Name     string
	CreateAt time.Time
	UpdateAt time.Time
	DeleteAt time.Time
}

type TagRepo interface {
	CreateTag(ctx context.Context, name string) error
	UpdateTag(ctx context.Context, id, name string) error
	DeleteTag(ctx context.Context, id string) error
	GetTag(ctx context.Context, article string) ([]*Tag, error)
	ListTag(ctx context.Context) ([]*Tag, error)
}

type TagUseCase struct {
	repo TagRepo
}

func NewTagUseCase(repo TagRepo) *TagUseCase {
	return &TagUseCase{repo: repo}
}

func (uc *TagUseCase) CreateTag(ctx context.Context, name string) error {
	return uc.repo.CreateTag(ctx, name)
}

func (uc *TagUseCase) UpdateTag(ctx context.Context, id, name string) error {
	return uc.repo.UpdateTag(ctx, id, name)
}

func (uc *TagUseCase) DeleteTag(ctx context.Context, id string) error {
	return uc.repo.DeleteTag(ctx, id)
}

func (uc *TagUseCase) GetTag(ctx context.Context, article string) ([]*Tag, error) {
	return uc.repo.GetTag(ctx, article)
}

func (uc *TagUseCase) ListTag(ctx context.Context) ([]*Tag, error) {
	return uc.repo.ListTag(ctx)
}
