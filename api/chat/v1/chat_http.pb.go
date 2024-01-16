// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.2
// - protoc             v3.21.12
// source: chat/v1/chat.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationChatAddConsumer = "/helloworld.v1.Chat/AddConsumer"
const OperationChatCreateConsumerGroup = "/helloworld.v1.Chat/CreateConsumerGroup"
const OperationChatCreateStream = "/helloworld.v1.Chat/CreateStream"
const OperationChatDelConsumer = "/helloworld.v1.Chat/DelConsumer"
const OperationChatDelConsumerGroup = "/helloworld.v1.Chat/DelConsumerGroup"
const OperationChatGetSetPeople = "/helloworld.v1.Chat/GetSetPeople"
const OperationChatGetSetPeopleNum = "/helloworld.v1.Chat/GetSetPeopleNum"
const OperationChatLogin = "/helloworld.v1.Chat/Login"
const OperationChatLogout = "/helloworld.v1.Chat/Logout"
const OperationChatSayHello = "/helloworld.v1.Chat/SayHello"
const OperationChatSendMessage = "/helloworld.v1.Chat/SendMessage"
const OperationChatSubscribe = "/helloworld.v1.Chat/Subscribe"

type ChatHTTPServer interface {
	AddConsumer(context.Context, *AddConsumerRequest) (*AddConsumerReply, error)
	CreateConsumerGroup(context.Context, *CreateConsumerGroupRequest) (*CreateConsumerGroupReply, error)
	CreateStream(context.Context, *CreateStreamRequest) (*CreateStreamReply, error)
	DelConsumer(context.Context, *DelConsumerRequest) (*DelConsumerReply, error)
	DelConsumerGroup(context.Context, *DelConsumerGroupRequest) (*DelConsumerGroupReply, error)
	GetSetPeople(context.Context, *GetSetPeopleRequest) (*GetSetPeopleReply, error)
	GetSetPeopleNum(context.Context, *GetSetPeopleNumRequest) (*GetSetPeopleNumReply, error)
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	Logout(context.Context, *LogoutRequest) (*LogoutReply, error)
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	// SendMessage发送消息/如果主题不存在，则自动创建主题
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageReply, error)
	// Subscribe订阅主题
	Subscribe(context.Context, *SubscribeRequest) (*SubscribeReply, error)
}

func RegisterChatHTTPServer(s *http.Server, srv ChatHTTPServer) {
	r := s.Route("/")
	r.GET("/test/{name}", _Chat_SayHello0_HTTP_Handler(srv))
	r.POST("/createConsumerGroup", _Chat_CreateConsumerGroup0_HTTP_Handler(srv))
	r.POST("/createStream", _Chat_CreateStream0_HTTP_Handler(srv))
	r.POST("/delConsumerGroup", _Chat_DelConsumerGroup0_HTTP_Handler(srv))
	r.POST("/addConsumer", _Chat_AddConsumer0_HTTP_Handler(srv))
	r.POST("/delConsumer", _Chat_DelConsumer0_HTTP_Handler(srv))
	r.POST("/subscribe", _Chat_Subscribe0_HTTP_Handler(srv))
	r.POST("/sendMessage", _Chat_SendMessage0_HTTP_Handler(srv))
	r.POST("/login", _Chat_Login0_HTTP_Handler(srv))
	r.POST("/logout", _Chat_Logout0_HTTP_Handler(srv))
	r.POST("/getSetPeopleNum", _Chat_GetSetPeopleNum0_HTTP_Handler(srv))
	r.POST("/getSetPeople", _Chat_GetSetPeople0_HTTP_Handler(srv))
}

func _Chat_SayHello0_HTTP_Handler(srv ChatHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChatSayHello)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SayHello(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

func _Chat_CreateConsumerGroup0_HTTP_Handler(srv ChatHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateConsumerGroupRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChatCreateConsumerGroup)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateConsumerGroup(ctx, req.(*CreateConsumerGroupRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateConsumerGroupReply)
		return ctx.Result(200, reply)
	}
}

func _Chat_CreateStream0_HTTP_Handler(srv ChatHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateStreamRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChatCreateStream)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateStream(ctx, req.(*CreateStreamRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateStreamReply)
		return ctx.Result(200, reply)
	}
}

func _Chat_DelConsumerGroup0_HTTP_Handler(srv ChatHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DelConsumerGroupRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChatDelConsumerGroup)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DelConsumerGroup(ctx, req.(*DelConsumerGroupRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DelConsumerGroupReply)
		return ctx.Result(200, reply)
	}
}

func _Chat_AddConsumer0_HTTP_Handler(srv ChatHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddConsumerRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChatAddConsumer)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddConsumer(ctx, req.(*AddConsumerRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AddConsumerReply)
		return ctx.Result(200, reply)
	}
}

func _Chat_DelConsumer0_HTTP_Handler(srv ChatHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DelConsumerRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChatDelConsumer)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DelConsumer(ctx, req.(*DelConsumerRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DelConsumerReply)
		return ctx.Result(200, reply)
	}
}

func _Chat_Subscribe0_HTTP_Handler(srv ChatHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SubscribeRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChatSubscribe)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Subscribe(ctx, req.(*SubscribeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SubscribeReply)
		return ctx.Result(200, reply)
	}
}

func _Chat_SendMessage0_HTTP_Handler(srv ChatHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SendMessageRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChatSendMessage)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SendMessage(ctx, req.(*SendMessageRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SendMessageReply)
		return ctx.Result(200, reply)
	}
}

func _Chat_Login0_HTTP_Handler(srv ChatHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChatLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _Chat_Logout0_HTTP_Handler(srv ChatHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LogoutRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChatLogout)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Logout(ctx, req.(*LogoutRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LogoutReply)
		return ctx.Result(200, reply)
	}
}

func _Chat_GetSetPeopleNum0_HTTP_Handler(srv ChatHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetSetPeopleNumRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChatGetSetPeopleNum)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSetPeopleNum(ctx, req.(*GetSetPeopleNumRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetSetPeopleNumReply)
		return ctx.Result(200, reply)
	}
}

func _Chat_GetSetPeople0_HTTP_Handler(srv ChatHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetSetPeopleRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChatGetSetPeople)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSetPeople(ctx, req.(*GetSetPeopleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetSetPeopleReply)
		return ctx.Result(200, reply)
	}
}

type ChatHTTPClient interface {
	AddConsumer(ctx context.Context, req *AddConsumerRequest, opts ...http.CallOption) (rsp *AddConsumerReply, err error)
	CreateConsumerGroup(ctx context.Context, req *CreateConsumerGroupRequest, opts ...http.CallOption) (rsp *CreateConsumerGroupReply, err error)
	CreateStream(ctx context.Context, req *CreateStreamRequest, opts ...http.CallOption) (rsp *CreateStreamReply, err error)
	DelConsumer(ctx context.Context, req *DelConsumerRequest, opts ...http.CallOption) (rsp *DelConsumerReply, err error)
	DelConsumerGroup(ctx context.Context, req *DelConsumerGroupRequest, opts ...http.CallOption) (rsp *DelConsumerGroupReply, err error)
	GetSetPeople(ctx context.Context, req *GetSetPeopleRequest, opts ...http.CallOption) (rsp *GetSetPeopleReply, err error)
	GetSetPeopleNum(ctx context.Context, req *GetSetPeopleNumRequest, opts ...http.CallOption) (rsp *GetSetPeopleNumReply, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
	Logout(ctx context.Context, req *LogoutRequest, opts ...http.CallOption) (rsp *LogoutReply, err error)
	SayHello(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
	SendMessage(ctx context.Context, req *SendMessageRequest, opts ...http.CallOption) (rsp *SendMessageReply, err error)
	Subscribe(ctx context.Context, req *SubscribeRequest, opts ...http.CallOption) (rsp *SubscribeReply, err error)
}

type ChatHTTPClientImpl struct {
	cc *http.Client
}

func NewChatHTTPClient(client *http.Client) ChatHTTPClient {
	return &ChatHTTPClientImpl{client}
}

func (c *ChatHTTPClientImpl) AddConsumer(ctx context.Context, in *AddConsumerRequest, opts ...http.CallOption) (*AddConsumerReply, error) {
	var out AddConsumerReply
	pattern := "/addConsumer"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationChatAddConsumer))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ChatHTTPClientImpl) CreateConsumerGroup(ctx context.Context, in *CreateConsumerGroupRequest, opts ...http.CallOption) (*CreateConsumerGroupReply, error) {
	var out CreateConsumerGroupReply
	pattern := "/createConsumerGroup"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationChatCreateConsumerGroup))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ChatHTTPClientImpl) CreateStream(ctx context.Context, in *CreateStreamRequest, opts ...http.CallOption) (*CreateStreamReply, error) {
	var out CreateStreamReply
	pattern := "/createStream"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationChatCreateStream))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ChatHTTPClientImpl) DelConsumer(ctx context.Context, in *DelConsumerRequest, opts ...http.CallOption) (*DelConsumerReply, error) {
	var out DelConsumerReply
	pattern := "/delConsumer"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationChatDelConsumer))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ChatHTTPClientImpl) DelConsumerGroup(ctx context.Context, in *DelConsumerGroupRequest, opts ...http.CallOption) (*DelConsumerGroupReply, error) {
	var out DelConsumerGroupReply
	pattern := "/delConsumerGroup"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationChatDelConsumerGroup))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ChatHTTPClientImpl) GetSetPeople(ctx context.Context, in *GetSetPeopleRequest, opts ...http.CallOption) (*GetSetPeopleReply, error) {
	var out GetSetPeopleReply
	pattern := "/getSetPeople"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationChatGetSetPeople))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ChatHTTPClientImpl) GetSetPeopleNum(ctx context.Context, in *GetSetPeopleNumRequest, opts ...http.CallOption) (*GetSetPeopleNumReply, error) {
	var out GetSetPeopleNumReply
	pattern := "/getSetPeopleNum"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationChatGetSetPeopleNum))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ChatHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationChatLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ChatHTTPClientImpl) Logout(ctx context.Context, in *LogoutRequest, opts ...http.CallOption) (*LogoutReply, error) {
	var out LogoutReply
	pattern := "/logout"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationChatLogout))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ChatHTTPClientImpl) SayHello(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/test/{name}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationChatSayHello))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ChatHTTPClientImpl) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...http.CallOption) (*SendMessageReply, error) {
	var out SendMessageReply
	pattern := "/sendMessage"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationChatSendMessage))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ChatHTTPClientImpl) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...http.CallOption) (*SubscribeReply, error) {
	var out SubscribeReply
	pattern := "/subscribe"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationChatSubscribe))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
