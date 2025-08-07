package service

import (
	"context"

	pb "myblog/api/v1/comment"
	"myblog/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type CommentService struct {
	us  *biz.CommentUseCase
	log *log.Helper
	pb.UnimplementedCommentServiceServer
}

func NewCommentService(us *biz.CommentUseCase, logger log.Logger) *CommentService {
	return &CommentService{us: us, log: log.NewHelper(logger)}
}

func (s *CommentService) CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CreateCommentReply, error) {
	s.log.Infof("Input data is:%v", req)
	coments, err := s.us.CreateComment(ctx, &biz.Comment{
		Name:    req.Content,
		Content: req.Content,
		Article: req.Article,
	})
	reply := &pb.CreateCommentReply{}
	for _, c := range coments {
		reply.Comments = append(reply.Comments, &pb.Comment{
			Id:      c.Id,
			Name:    c.Name,
			Content: c.Content,
		})
	}
	return reply, err
}
func (s *CommentService) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*pb.DeleteCommentReply, error) {
	err := s.us.DeleteComment(ctx, req.Id)
	return &pb.DeleteCommentReply{}, err
}
func (s *CommentService) GetComments(ctx context.Context, req *pb.GetCommentRequest) (*pb.GetCommentReply, error) {
	comments, err := s.us.GetComment(ctx, req.Article)
	reply := &pb.GetCommentReply{}
	for _, c := range comments {
		reply.Comments = append(reply.Comments, &pb.Comment{
			Id:      c.Id,
			Name:    c.Name,
			Content: c.Content,
		})
	}
	return reply, err
}
