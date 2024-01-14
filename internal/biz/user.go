package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Username string
}

type UserRepo interface {
	//在 Redis 中创建消费者组
	CreateConsumerGroup(ctx context.Context, groupName string) error
	//在 Redis 中添加消费者到组
	//AddConsumerToGroup(ctx context.Context, groupName string, u *User) error
}

type UserUsecase struct {
	ur  UserRepo
	log *log.Helper
}

func NewUserUsecase(ur UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{ur: ur, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) CreateConsumerGroup(ctx context.Context, groupName string) error {
	err := uc.ur.CreateConsumerGroup(ctx, groupName)
	if err != nil {
		return err
	}
	return nil
}
