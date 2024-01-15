package service

import (
	"context"

	v1 "Redis-chat/api/chat/v1"
	"Redis-chat/internal/biz"
)

type ChatService struct {
	v1.UnimplementedChatServer

	cc *biz.ChatUsecase
	// uc *biz.UserUsecase
}

func NewChatService(cc *biz.ChatUsecase) *ChatService {
	return &ChatService{cc: cc}
}

func (s *ChatService) CreateStream(ctx context.Context, req *v1.CreateStreamRequest) (reply *v1.CreateStreamReply, err error) {
	str, err := s.cc.CreateStream(ctx, req.Stream, req.Key, req.Value)
	if err != nil {
		return nil, err
	}
	return &v1.CreateStreamReply{
		Stream: str.StreamName,
		Key:    str.Key,
		Value:  str.Value,
	}, nil
}

func (s *ChatService) CreateConsumerGroup(ctx context.Context, req *v1.CreateConsumerGroupRequest) (reply *v1.CreateConsumerGroupReply, err error) {
	cg, err := s.cc.CreateConsumerGroup(ctx, req.Stream, req.Consumergroupname)
	if err != nil {
		return nil, err
	}
	return &v1.CreateConsumerGroupReply{
		Stream:            cg.Stream,
		Consumergroupname: cg.GroupName,
		Success:           true,
		Message:           "创建消费者组成功",
	}, nil
}
