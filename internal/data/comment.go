package data

import (
	"context"
	"myblog/internal/biz"
)

type CommentRepo struct {
	data *Data
}

func NewCommentRepo(data *Data) biz.CommentRepo {
	return &CommentRepo{data: data}
}

func (cr *CommentRepo) CreateComment(ctx context.Context, c *biz.Comment) ([]*biz.Comment, error) {
	//创建评论
	_, err := cr.data.db.Comment.Create().
		SetID("111111").
		SetName(c.Name).
		SetContent(c.Content).
		SetArticleID(c.Article).
		Save(ctx)
	//查询文章下的评论
	comments := make([]*biz.Comment, 0)
	cs, error := cr.data.db.Comment.Query().Select(c.Article).All(ctx)
	if err != nil {
		panic(error)
	}
	for _, c := range cs {
		comments = append(comments, &biz.Comment{
			Id:       c.ID,
			Name:     c.Name,
			Content:  c.Content,
			CreateAt: c.CreateAt,
			UpdateAt: c.UpdateAt,
		})
	}
	return comments, err
}

func (cr *CommentRepo) DeleteComment(ctx context.Context, id string) error {
	return cr.data.db.Comment.DeleteOneID(id).Exec(ctx)
}

func (cr *CommentRepo) GetComments(ctx context.Context, article string) ([]*biz.Comment, error) {
	comments := make([]*biz.Comment, 0)
	cs, err := cr.data.db.Comment.Query().Select(article).All(ctx)
	for _, c := range cs {
		comments = append(comments, &biz.Comment{
			Id:       c.ID,
			Name:     c.Name,
			Content:  c.Content,
			CreateAt: c.CreateAt,
			UpdateAt: c.UpdateAt,
		})
	}
	return comments, err
}
