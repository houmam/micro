// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/micro/micro/bot/proto/bot.proto

/*
Package go_micro_bot is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/micro/bot/proto/bot.proto

It has these top-level messages:
	HelpRequest
	HelpResponse
	ExecRequest
	ExecResponse
*/
package go_micro_bot

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Command service

type CommandClient interface {
	Help(ctx context.Context, in *HelpRequest, opts ...client.CallOption) (*HelpResponse, error)
	Exec(ctx context.Context, in *ExecRequest, opts ...client.CallOption) (*ExecResponse, error)
}

type commandClient struct {
	c           client.Client
	serviceName string
}

func NewCommandClient(serviceName string, c client.Client) CommandClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.bot"
	}
	return &commandClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *commandClient) Help(ctx context.Context, in *HelpRequest, opts ...client.CallOption) (*HelpResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Command.Help", in)
	out := new(HelpResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandClient) Exec(ctx context.Context, in *ExecRequest, opts ...client.CallOption) (*ExecResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Command.Exec", in)
	out := new(ExecResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Command service

type CommandHandler interface {
	Help(context.Context, *HelpRequest, *HelpResponse) error
	Exec(context.Context, *ExecRequest, *ExecResponse) error
}

func RegisterCommandHandler(s server.Server, hdlr CommandHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Command{hdlr}, opts...))
}

type Command struct {
	CommandHandler
}

func (h *Command) Help(ctx context.Context, in *HelpRequest, out *HelpResponse) error {
	return h.CommandHandler.Help(ctx, in, out)
}

func (h *Command) Exec(ctx context.Context, in *ExecRequest, out *ExecResponse) error {
	return h.CommandHandler.Exec(ctx, in, out)
}
