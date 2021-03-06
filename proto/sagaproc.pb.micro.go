// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/sagaproc.proto

package sagaproc

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Sagaproc service

func NewSagaprocEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Sagaproc service

type SagaprocService interface {
	HandleOperation(ctx context.Context, in *OperationPayload, opts ...client.CallOption) (*OperationResult, error)
}

type sagaprocService struct {
	c    client.Client
	name string
}

func NewSagaprocService(name string, c client.Client) SagaprocService {
	return &sagaprocService{
		c:    c,
		name: name,
	}
}

func (c *sagaprocService) HandleOperation(ctx context.Context, in *OperationPayload, opts ...client.CallOption) (*OperationResult, error) {
	req := c.c.NewRequest(c.name, "Sagaproc.HandleOperation", in)
	out := new(OperationResult)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Sagaproc service

type SagaprocHandler interface {
	HandleOperation(context.Context, *OperationPayload, *OperationResult) error
}

func RegisterSagaprocHandler(s server.Server, hdlr SagaprocHandler, opts ...server.HandlerOption) error {
	type sagaproc interface {
		HandleOperation(ctx context.Context, in *OperationPayload, out *OperationResult) error
	}
	type Sagaproc struct {
		sagaproc
	}
	h := &sagaprocHandler{hdlr}
	return s.Handle(s.NewHandler(&Sagaproc{h}, opts...))
}

type sagaprocHandler struct {
	SagaprocHandler
}

func (h *sagaprocHandler) HandleOperation(ctx context.Context, in *OperationPayload, out *OperationResult) error {
	return h.SagaprocHandler.HandleOperation(ctx, in, out)
}
