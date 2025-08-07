package biz

import (
	"context"
	"time"
)

type Article struct {
	Id       string
	Title    string
	Content  string
	CreateAt time.Time
	UpdateAt time.Time
	DeleteAt time.Time
}

type ArticleRepo interface {
	CreateArticle(ctx context.Context, a *Article) error
	UpdateArticle(ctx context.Context, id string, a *Article) error
	DeleteArticle(ctx context.Context, id string) error
	GetArticle(ctx context.Context, id string) (*Article, error)
	ListArticle(ctx context.Context) ([]*Article, error)
}

type ArticleUseCase struct {
	repo ArticleRepo
}

func NewArticleUseCase(repo ArticleRepo) *ArticleUseCase {
	return &ArticleUseCase{repo: repo}
}

func (uc *ArticleUseCase) CreateArticle(ctx context.Context, a *Article) error {
	return uc.repo.CreateArticle(ctx, a)
}

func (uc *ArticleUseCase) UpdateArticle(ctx context.Context, id string, a *Article) error {
	return uc.repo.UpdateArticle(ctx, id, a)
}

func (uc *ArticleUseCase) DeleteArticle(ctx context.Context, id string) error {
	return uc.repo.DeleteArticle(ctx, id)
}

func (uc *ArticleUseCase) GetArticle(ctx context.Context, id string) (*Article, error) {
	return uc.repo.GetArticle(ctx, id)
}

func (uc *ArticleUseCase) ListArticle(ctx context.Context) ([]*Article, error) {
	return uc.repo.ListArticle(ctx)
}
