package data

import (
	"context"
	"myblog/internal/biz"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type ArticleRepo struct {
	data *Data
	log  *log.Helper
}

func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &ArticleRepo{data: data, log: log.NewHelper(logger)}
}

func (ar *ArticleRepo) CreateArticle(ctx context.Context, a *biz.Article) error {
	_, err := ar.data.db.Article.Create().
		SetID("3115004318").
		SetTitle(a.Title).
		SetContent(a.Content).
		Save(ctx)
	if err != nil {
		ar.log.Errorf("创建文章失败：%v", err)
		return err
	}
	return nil
}

func (ar *ArticleRepo) UpdateArticle(ctx context.Context, id string, a *biz.Article) error {
	_, err := ar.data.db.Article.UpdateOneID(id).
		SetTitle(a.Title).
		SetContent(a.Content).
		SetUpdateAt(time.Now()).
		Save(ctx)
	if err != nil {
		ar.log.Errorf("更新文章失败：%v", err)
		return err
	}
	return nil
}

func (ar *ArticleRepo) DeleteArticle(ctx context.Context, id string) error {
	err := ar.data.db.Article.DeleteOneID(id).Exec(ctx)
	if err != nil {
		ar.log.Errorf("删除文章失败：%v", err)
		return err
	}
	return nil
}

func (ar *ArticleRepo) GetArticle(ctx context.Context, id string) (*biz.Article, error) {
	a, err := ar.data.db.Article.Query().Select(id).First(ctx)
	if err != nil {
		ar.log.Errorf("查询文章失败：%v", err)
		return nil, err
	}
	article := &biz.Article{
		Id:       a.ID,
		Title:    a.Title,
		Content:  a.Content,
		CreateAt: a.CreateAt,
		UpdateAt: a.UpdateAt,
	}
	return article, nil
}

func (ar *ArticleRepo) ListArticle(ctx context.Context) ([]*biz.Article, error) {
	as, err := ar.data.db.Article.Query().All(ctx)
	if err != nil {
		ar.log.Errorf("查询文章失败：%v", err)
		return nil, err
	}

	articles := make([]*biz.Article, 0)
	for _, article := range as {
		articles = append(articles, &biz.Article{
			Id:       article.ID,
			Title:    article.Title,
			Content:  article.Content,
			CreateAt: article.CreateAt,
			UpdateAt: article.UpdateAt,
		})
	}

	return articles, nil
}
