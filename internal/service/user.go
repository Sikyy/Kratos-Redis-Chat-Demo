package service

import (
	v1 "Redis-chat/api/chat/v1"
	"context"
	"fmt"
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
		Username: sd.Username,
		Topic:    sd.Topic,
		Success:  true,
		Message:  "发送消息成功,消息内容为" + sd.Message + "发送主题为" + sd.Topic + "发送者为" + sd.Username,
	}, nil
}

func (s *ChatService) Login(ctx context.Context, req *v1.LoginRequest) (reply *v1.LoginReply, err error) {
	u, err := s.uc.Login(ctx, req.Setname, req.Username)
	if err != nil {
		return nil, err
	}
	return &v1.LoginReply{
		Username: u.Username,
		Setname:  u.Setname,
		Success:  true,
		Message:  "用户" + u.Username + "登录成功" + "集合：" + u.Setname + "内已录入",
	}, nil
}

func (s *ChatService) Logout(ctx context.Context, req *v1.LogoutRequest) (reply *v1.LogoutReply, err error) {
	u, err := s.uc.Logout(ctx, req.Setname, req.Username)
	if err != nil {
		return nil, err
	}
	return &v1.LogoutReply{
		Username: u.Username,
		Setname:  u.Setname,
		Success:  true,
		Message:  "用户" + u.Username + "登出成功" + "集合：" + u.Setname + "内已删除",
	}, nil
}

func (s *ChatService) GetSetPeopleNum(ctx context.Context, req *v1.GetSetPeopleNumRequest) (reply *v1.GetSetPeopleNumReply, err error) {
	set, err := s.uc.GetSetPeopleNum(ctx, req.Setname)
	if err != nil {
		return nil, err
	}
	return &v1.GetSetPeopleNumReply{
		Setname: set.Setname,
		Num:     set.SetPeopleNum,
		Message: fmt.Sprintf("集合%s内人数为%d", set.Setname, set.SetPeopleNum),
	}, nil
}

func (s *ChatService) GetSetPeople(ctx context.Context, req *v1.GetSetPeopleRequest) (reply *v1.GetSetPeopleReply, err error) {
	set, err := s.uc.GetSetPeople(ctx, req.Setname)
	if err != nil {
		return nil, err
	}
	return &v1.GetSetPeopleReply{
		Setname: set.Setname,
		People:  set.SetPeople,
		Message: fmt.Sprintf("集合%s内人员为%s", set.Setname, set.SetPeople),
	}, nil
}
