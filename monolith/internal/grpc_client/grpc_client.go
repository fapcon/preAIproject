package grpc_client

import (
	"context"
	"google.golang.org/grpc"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/pkg/grpc/bot"
)

type botServiceGRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewBotServiceGRPCClient(cc grpc.ClientConnInterface) bot.BotServiceGRPCClient {
	return &botServiceGRPCClient{cc}
}

func (c *botServiceGRPCClient) Create(ctx context.Context, in *bot.BotCreateIn, opts ...grpc.CallOption) (*bot.BotOut, error) {
	out := new(bot.BotOut)
	err := c.cc.Invoke(ctx, "/bot.BotServiceGRPC/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServiceGRPCClient) Delete(ctx context.Context, in *bot.BotDeleteIn, opts ...grpc.CallOption) (*bot.BOut, error) {
	out := new(bot.BOut)
	err := c.cc.Invoke(ctx, "/bot.BotServiceGRPC/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServiceGRPCClient) Update(ctx context.Context, in *bot.BotUpdateIn, opts ...grpc.CallOption) (*bot.BOut, error) {
	out := new(bot.BOut)
	err := c.cc.Invoke(ctx, "/bot.BotServiceGRPC/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServiceGRPCClient) Get(ctx context.Context, in *bot.BotGetIn, opts ...grpc.CallOption) (*bot.BotOut, error) {
	out := new(bot.BotOut)
	err := c.cc.Invoke(ctx, "/bot.BotServiceGRPC/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServiceGRPCClient) Toggle(ctx context.Context, in *bot.BotToggleIn, opts ...grpc.CallOption) (*bot.BOut, error) {
	out := new(bot.BOut)
	err := c.cc.Invoke(ctx, "/bot.BotServiceGRPC/Toggle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServiceGRPCClient) Subscribe(ctx context.Context, in *bot.BotSubscribeIn, opts ...grpc.CallOption) (*bot.BOut, error) {
	out := new(bot.BOut)
	err := c.cc.Invoke(ctx, "/bot.BotServiceGRPC/Subscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServiceGRPCClient) Unsubscribe(ctx context.Context, in *bot.BotSubscribeIn, opts ...grpc.CallOption) (*bot.BOut, error) {
	out := new(bot.BOut)
	err := c.cc.Invoke(ctx, "/bot.BotServiceGRPC/Unsubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServiceGRPCClient) List(ctx context.Context, in *bot.BotListIn, opts ...grpc.CallOption) (*bot.BotListOut, error) {
	out := new(bot.BotListOut)
	err := c.cc.Invoke(ctx, "/bot.BotServiceGRPC/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServiceGRPCClient) WebhookSignal(ctx context.Context, in *bot.WebhookSignalIn, opts ...grpc.CallOption) (*bot.WebhookSignalOut, error) {
	out := new(bot.WebhookSignalOut)
	err := c.cc.Invoke(ctx, "/bot.BotServiceGRPC/WebhookSignal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
