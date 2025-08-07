package biz

import (
	"context"
	"time"
)

type Comment struct {
	Id       string
	Name     string
	Content  string
	Article  string
	CreateAt time.Time
	UpdateAt time.Time
	DeleteAt time.Time
}

type CommentRepo interface {
	CreateComment(ctx context.Context, c *Comment) ([]*Comment, error)
	DeleteComment(ctx context.Context, id string) error
	GetComments(ctx context.Context, article string) ([]*Comment, error)
}

type CommentUseCase struct {
	repo CommentRepo
}

func NewCommentUseCase(repo CommentRepo) *CommentUseCase {
	return &CommentUseCase{repo: repo}
}

func (uc *CommentUseCase) CreateComment(ctx context.Context, c *Comment) ([]*Comment, error) {
	return uc.repo.CreateComment(ctx, c)
}

func (uc *CommentUseCase) DeleteComment(ctx context.Context, id string) error {
	return uc.repo.DeleteComment(ctx, id)
}

func (uc *CommentUseCase) GetComment(ctx context.Context, article string) ([]*Comment, error) {
	return uc.repo.GetComments(ctx, article)
}
