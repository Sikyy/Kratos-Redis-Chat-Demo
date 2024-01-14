package data

import (
	"Redis-chat/internal/conf"
	"context"
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/types/known/durationpb"
)

func TestNewRedis(t *testing.T) {
	// 创建一个示例的配置对象
	conf := &conf.Data{
		Redis: &conf.Data_Redis{
			Addr:         "localhost:6379",                 // 你的 Redis 地址
			Username:     "",                               // no username set
			Password:     "",                               // no password set
			Db:           0,                                // use default DB
			ReadTimeout:  &durationpb.Duration{Seconds: 5}, // 5秒的持续时间
			WriteTimeout: &durationpb.Duration{Seconds: 5}, // 5秒的持续时间
		},
	}

	// 调用NewRedis函数获取redis.Client
	client := NewRedis(conf)

	// 测试连接
	if err := testConnection(client); err != nil {
		t.Fatalf("Failed to test connection: %v", err)
	}

	// 在此之后，你可以继续执行其他与Redis相关的操作

	// 关闭连接
	if err := client.Close(); err != nil {
		t.Fatalf("Failed to close connection: %v", err)
	}
}

// 测试连接的函数
func testConnection(client *redis.Client) error {
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	fmt.Println("Ping Result:", pong) // 在此处添加打印以检查连接状态
	return nil
}
