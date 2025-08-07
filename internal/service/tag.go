package service

import (
	"context"

	pb "myblog/api/v1/tag"
	"myblog/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type TagService struct {
	uc  *biz.TagUseCase
	log *log.Helper
	pb.UnimplementedTagServiceServer
}

func NewTagService(uc *biz.TagUseCase, logger log.Logger) *TagService {
	return &TagService{uc: uc, log: log.NewHelper(logger)}
}

func (s *TagService) CreateTag(ctx context.Context, req *pb.CreateTagRequest) (*pb.CreateTagReply, error) {
	s.log.Infof("Input data is:%v", req)
	err := s.uc.CreateTag(ctx, req.Name)
	return &pb.CreateTagReply{}, err
}
func (s *TagService) UpdateTag(ctx context.Context, req *pb.UpdateTagRequest) (*pb.UpdateTagReply, error) {
	s.log.Infof("Input data is:%v", req)
	err := s.uc.UpdateTag(ctx, req.Id, req.Name)
	return &pb.UpdateTagReply{}, err
}
func (s *TagService) DeleteTag(ctx context.Context, req *pb.DeleteTagRequest) (*pb.DeleteTagReply, error) {
	err := s.uc.DeleteTag(ctx, req.Id)
	return &pb.DeleteTagReply{}, err
}
func (s *TagService) GetTag(ctx context.Context, req *pb.GetTagRequest) (*pb.GetTagReply, error) {
	tags, err := s.uc.GetTag(ctx, req.Article)
	reply := &pb.GetTagReply{}
	for _, tag := range tags {
		reply.Tags = append(reply.Tags, &pb.Tag{
			Id:   tag.Id,
			Name: tag.Name,
		})
	}
	return reply, err
}
func (s *TagService) ListTag(ctx context.Context, req *pb.ListTagRequest) (*pb.ListTagReply, error) {
	tags, err := s.uc.ListTag(ctx)
	reply := &pb.ListTagReply{}
	for _, tag := range tags {
		reply.Tags = append(reply.Tags, &pb.Tag{
			Id:   tag.Id,
			Name: tag.Name,
		})
	}
	return reply, err
}
