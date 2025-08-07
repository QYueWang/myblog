package service

import (
	"context"

	pb "myblog/api/v1/article"
	"myblog/internal/biz"
)

type ArticleService struct {
	au *biz.ArticleUseCase
	pb.UnimplementedArticleServiceServer
}

func NewArticleService(au *biz.ArticleUseCase) *ArticleService {
	return &ArticleService{au: au}
}

func (s *ArticleService) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleReply, error) {
	err := s.au.CreateArticle(ctx, &biz.Article{
		Title:   req.Title,
		Content: req.Content,
	})

	return &pb.CreateArticleReply{}, err
}
func (s *ArticleService) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.UpdateArticleReply, error) {
	err := s.au.UpdateArticle(ctx, req.Id, &biz.Article{
		Title:   req.Title,
		Content: req.Content,
	})

	return &pb.UpdateArticleReply{}, err
}
func (s *ArticleService) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.DeleteArticleReply, error) {
	err := s.au.DeleteArticle(ctx, req.Id)
	return &pb.DeleteArticleReply{}, err
}
func (s *ArticleService) GetArticle(ctx context.Context, req *pb.GetArticleRequest) (*pb.GetArticleReply, error) {
	a, err := s.au.GetArticle(ctx, req.Id)
	return &pb.GetArticleReply{Article: &pb.Article{
		Id:      a.Id,
		Title:   a.Title,
		Content: a.Content,
	}}, err
}
func (s *ArticleService) ListArticle(ctx context.Context, req *pb.ListArticleRequest) (*pb.ListArticleReply, error) {
	as, err := s.au.ListArticle(ctx)
	reply := &pb.ListArticleReply{}
	for _, a := range as {
		reply.Articles = append(reply.Articles, &pb.Article{
			Id:      a.Id,
			Title:   a.Title,
			Content: a.Content,
		})
	}
	return reply, err
}
