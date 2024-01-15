package service

import (
	v1 "Redis-chat/api/chat/v1"
	"context"
)

// 订阅主题
func (s *ChatService) Subscribe(ctx context.Context, req *v1.SubscribeRequest) (reply *v1.SubscribeReply, err error) {
	su, err := s.uc.Subscribe(ctx, req.Topic)
	if err != nil {
		return nil, err
	}
	return &v1.SubscribeReply{
		Topic:   su.Topic,
		Message: "订阅成功,订阅主题为" + su.Topic,
	}, nil
}

func (s *ChatService) SendMessage(ctx context.Context, req *v1.SendMessageRequest) (reply *v1.SendMessageReply, err error) {
	sd, err := s.uc.SendMessage(ctx, req.Message, req.Topic, req.Message)
	if err != nil {
		return nil, err
	}
	return &v1.SendMessageReply{
		Username: "siky",
		Topic:    sd.Topic,
		Success:  true,
		Message:  "发送消息成功,消息内容为" + sd.Message + "发送主题为" + sd.Topic + "发送者为" + sd.Username,
	}, nil
}
