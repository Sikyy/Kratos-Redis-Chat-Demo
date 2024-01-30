package data

import (
	"Redis-chat/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
)

type userRepo struct {
	client *redis.Client
	data   *Data
	log    *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	client := data.rdb
	if client == nil {
		log.NewHelper(logger).Error("Redis client is nil")
		return nil
	}

	return &userRepo{
		client: client,
		data:   data,
		log:    log.NewHelper(logger),
	}
}

// Subscribe 订阅主题
func (r *userRepo) Subscribe(ctx context.Context, s *biz.Subscribe) error {
	r.data.rdb.Subscribe(ctx, s.Topic)
	return nil
}

// SendMessage 发送消息
func (r *userRepo) SendMessage(ctx context.Context, m *biz.Message) error {
	r.data.rdb.Publish(ctx, m.Topic, m.Message)
	return nil
}

// Login 登录
func (r *userRepo) Login(ctx context.Context, u *biz.User) (*redis.IntCmd, error) {
	reult := r.data.rdb.SAdd(ctx, u.Setname, u.Username)
	if reult.Err() != nil {
		return nil, reult.Err()
	}
	return reult, nil
}

// Logout 登出
func (r *userRepo) Logout(ctx context.Context, u *biz.User) (*redis.IntCmd, error) {
	reult := r.data.rdb.SRem(ctx, u.Setname, u.Username)
	if reult.Err() != nil {
		return nil, reult.Err()
	}
	return reult, nil
}

// GetSetPeopleNum 获取集合内人数
func (r *userRepo) GetSetPeopleNum(ctx context.Context, s *biz.Set) (*redis.IntCmd, error) {
	reult := r.data.rdb.SCard(ctx, s.Setname)
	return reult, nil
}

// GetSetPeople 获取集合内人员
func (r *userRepo) GetSetPeople(ctx context.Context, s *biz.Set) (*redis.StringSliceCmd, error) {
	result := r.data.rdb.SMembers(ctx, s.Setname)
	return result, nil
}
