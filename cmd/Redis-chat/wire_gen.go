// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"Redis-chat/internal/biz"
	"Redis-chat/internal/conf"
	"Redis-chat/internal/data"
	"Redis-chat/internal/server"
	"Redis-chat/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	client := data.NewRedis(confData)
	dataData, cleanup, err := data.NewData(confData, logger, client)
	if err != nil {
		return nil, nil, err
	}
	chatRepo := data.NewChatRepo(dataData, logger)
	chatUsecase := biz.NewChatUsecase(chatRepo, logger)
	chatService := service.NewChatService(chatUsecase)
	grpcServer := server.NewGRPCServer(confServer, chatService, logger)
	httpServer := server.NewHTTPServer(confServer, chatService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
