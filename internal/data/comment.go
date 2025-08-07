package data

import (
	"context"
	"myblog/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type CommentRepo struct {
	data *Data
	log  *log.Helper
}

func NewCommentRepo(data *Data, logger log.Logger) biz.CommentRepo {
	return &CommentRepo{data: data, log: log.NewHelper(logger)}
}

func (cr *CommentRepo) CreateComment(ctx context.Context, c *biz.Comment) ([]*biz.Comment, error) {
	//创建评论
	_, err1 := cr.data.db.Comment.Create().
		SetID("111111").
		SetName(c.Name).
		SetContent(c.Content).
		SetArticleID(c.Article).
		Save(ctx)
	if err1 != nil {
		//可以不使用log.Helper,直接使用框架默认初始化好的log.DefaultLogger实例，keyvals奇数为标签，偶数为值
		log.DefaultLogger.Log(log.LevelError, "result:", "创建评论失败")
	}
	//查询文章下的评论
	comments := make([]*biz.Comment, 0)
	cs, err2 := cr.data.db.Comment.Query().Select(c.Article).All(ctx)
	if err2 != nil {
		cr.log.Errorf("查询文章评论失败：%v", err2)
		return nil, err2
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
	return comments, nil
}

func (cr *CommentRepo) DeleteComment(ctx context.Context, id string) error {
	err := cr.data.db.Comment.DeleteOneID(id).Exec(ctx)
	if err != nil {
		cr.log.Errorf("删除评论失败：%v", err)
		return err
	}
	return nil
}

func (cr *CommentRepo) GetComments(ctx context.Context, article string) ([]*biz.Comment, error) {
	comments := make([]*biz.Comment, 0)
	cs, err := cr.data.db.Comment.Query().Select(article).All(ctx)
	if err != nil {
		cr.log.Errorf("查询文章评论失败：%v", err)
		return nil, err
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
	return comments, nil
}
