package service

import (
	"context"

	v1 "Redis-chat/api/chat/v1"
	"Redis-chat/internal/biz"
)

type ChatService struct {
	v1.UnimplementedChatServer

	cc *biz.ChatUsecase
	uc *biz.UserUsecase
}

func NewChatService(cc *biz.ChatUsecase, uc *biz.UserUsecase) *ChatService {
	return &ChatService{cc: cc, uc: uc}
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

func (s *ChatService) DelConsumerGroup(ctx context.Context, req *v1.DelConsumerGroupRequest) (reply *v1.DelConsumerGroupReply, err error) {
	cg, err := s.cc.DelConsumerGroup(ctx, req.Stream, req.Consumergroupname)
	if err != nil {
		return nil, err
	}
	return &v1.DelConsumerGroupReply{
		Stream:            cg.Stream,
		Consumergroupname: cg.GroupName,
		Success:           true,
		Message:           "删除消费者组成功",
	}, nil
}

func (s *ChatService) AddConsumer(ctx context.Context, req *v1.AddConsumerRequest) (reply *v1.AddConsumerReply, err error) {
	c, err := s.cc.AddConsumer(ctx, req.Stream, req.Consumergroupname, req.Consumername)
	if err != nil {
		return nil, err
	}
	return &v1.AddConsumerReply{
		Stream:            c.Stream,
		Consumergroupname: c.GroupName,
		Consumername:      c.ConsumerName,
		Success:           true,
		Message:           "添加消费者成功，消费者名为：" + c.ConsumerName + "，消费者组名为：" + c.GroupName + ",流名为：" + c.Stream,
	}, nil
}

func (s *ChatService) DelConsumer(ctx context.Context, req *v1.DelConsumerRequest) (reply *v1.DelConsumerReply, err error) {
	c, err := s.cc.DelConsumer(ctx, req.Stream, req.Consumergroupname, req.Consumername)
	if err != nil {
		return nil, err
	}
	return &v1.DelConsumerReply{
		Stream:            c.Stream,
		Consumergroupname: c.GroupName,
		Consumername:      c.ConsumerName,
		Success:           true,
		Message:           "删除消费者成功，消费者名为：" + c.ConsumerName + "，消费者组名为：" + c.GroupName + ",流名为：" + c.Stream,
	}, nil
}
