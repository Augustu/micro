// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/Augustu/micro/debug/log/proto/log.proto

package go_micro_debug_log

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/Augustu/go-micro/v2/api"
	client "github.com/Augustu/go-micro/v2/client"
	server "github.com/Augustu/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Log service

func NewLogEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Log service

type LogService interface {
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
}

type logService struct {
	c    client.Client
	name string
}

func NewLogService(name string, c client.Client) LogService {
	return &logService{
		c:    c,
		name: name,
	}
}

func (c *logService) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.name, "Log.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Log service

type LogHandler interface {
	Read(context.Context, *ReadRequest, *ReadResponse) error
}

func RegisterLogHandler(s server.Server, hdlr LogHandler, opts ...server.HandlerOption) error {
	type log interface {
		Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error
	}
	type Log struct {
		log
	}
	h := &logHandler{hdlr}
	return s.Handle(s.NewHandler(&Log{h}, opts...))
}

type logHandler struct {
	LogHandler
}

func (h *logHandler) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.LogHandler.Read(ctx, in, out)
}
