package data

import (
	"context"
	"myblog/internal/biz"
	"time"
)

type ArticleRepo struct {
	data *Data
}

func NewArticleRepo(data *Data) biz.ArticleRepo {
	return &ArticleRepo{data: data}
}

func (ar *ArticleRepo) CreateArticle(ctx context.Context, a *biz.Article) error {
	_, err := ar.data.db.Article.Create().
		SetID("3115004318").
		SetTitle(a.Title).
		SetContent(a.Content).
		Save(ctx)
	return err
}

func (ar *ArticleRepo) UpdateArticle(ctx context.Context, id string, a *biz.Article) error {
	_, err := ar.data.db.Article.UpdateOneID(id).
		SetTitle(a.Title).
		SetContent(a.Content).
		SetUpdateAt(time.Now()).
		Save(ctx)
	return err
}

func (ar *ArticleRepo) DeleteArticle(ctx context.Context, id string) error {
	err := ar.data.db.Article.DeleteOneID(id).Exec(ctx)
	return err
}

func (ar *ArticleRepo) GetArticle(ctx context.Context, id string) (*biz.Article, error) {
	a, err := ar.data.db.Article.Query().Select(id).First(ctx)
	article := &biz.Article{
		Id:       a.ID,
		Title:    a.Title,
		Content:  a.Content,
		CreateAt: a.CreateAt,
		UpdateAt: a.UpdateAt,
	}
	return article, err
}

func (ar *ArticleRepo) ListArticle(ctx context.Context) ([]*biz.Article, error) {
	as, err := ar.data.db.Article.Query().All(ctx)
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

	return articles, err
}
