// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: merge.proto

package merge

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"math"
)

import (
	"context"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Merge service

type MergeService interface {
	TestMerge(ctx context.Context, in *TestMergeReq, opts ...client.CallOption) (*TestMergeRes, error)
}

type mergeService struct {
	c    client.Client
	name string
}

func NewMergeService(name string, c client.Client) MergeService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "merge"
	}
	return &mergeService{
		c:    c,
		name: name,
	}
}

func (c *mergeService) TestMerge(ctx context.Context, in *TestMergeReq, opts ...client.CallOption) (*TestMergeRes, error) {
	req := c.c.NewRequest(c.name, "Merge.TestMerge", in)
	out := new(TestMergeRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Merge service

type MergeHandler interface {
	TestMerge(context.Context, *TestMergeReq, *TestMergeRes) error
}

func RegisterMergeHandler(s server.Server, hdlr MergeHandler, opts ...server.HandlerOption) error {
	type merge interface {
		TestMerge(ctx context.Context, in *TestMergeReq, out *TestMergeRes) error
	}
	type Merge struct {
		merge
	}
	h := &mergeHandler{hdlr}
	return s.Handle(s.NewHandler(&Merge{h}, opts...))
}

type mergeHandler struct {
	MergeHandler
}

func (h *mergeHandler) TestMerge(ctx context.Context, in *TestMergeReq, out *TestMergeRes) error {
	return h.MergeHandler.TestMerge(ctx, in, out)
}
