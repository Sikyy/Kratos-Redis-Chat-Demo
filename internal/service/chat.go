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
