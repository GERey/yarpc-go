// Code generated by protoc-gen-yarpc-go
// source: internal/crossdock/crossdockpb/crossdock.proto
// DO NOT EDIT!

// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package crossdockpb

import (
	"context"
	"reflect"

	"github.com/gogo/protobuf/proto"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/encoding/x/protobuf"
	"go.uber.org/yarpc/yarpcproto"
)

// EchoYarpcClient is the yarpc client-side interface for the Echo service.
type EchoYarpcClient interface {
	Echo(context.Context, *Ping, ...yarpc.CallOption) (*Pong, error)
}

// NewEchoYarpcClient builds a new yarpc client for the Echo service.
func NewEchoYarpcClient(clientConfig transport.ClientConfig, options ...protobuf.ClientOption) EchoYarpcClient {
	return &_EchoYarpcCaller{protobuf.NewClient("uber.yarpc.internal.crossdock.Echo", clientConfig, options...)}
}

// EchoYarpcServer is the yarpc server-side interface for the Echo service.
type EchoYarpcServer interface {
	Echo(context.Context, *Ping) (*Pong, error)
}

// BuildEchoYarpcProcedures prepares an implementation of the Echo service for yarpc registration.
func BuildEchoYarpcProcedures(server EchoYarpcServer) []transport.Procedure {
	handler := &_EchoYarpcHandler{server}
	return protobuf.BuildProcedures(
		"uber.yarpc.internal.crossdock.Echo",
		map[string]transport.UnaryHandler{
			"Echo": protobuf.NewUnaryHandler(handler.Echo, newEcho_EchoYarpcRequest),
		},
		map[string]transport.OnewayHandler{},
	)
}

type _EchoYarpcCaller struct {
	client protobuf.Client
}

func (c *_EchoYarpcCaller) Echo(ctx context.Context, request *Ping, options ...yarpc.CallOption) (*Pong, error) {
	responseMessage, err := c.client.Call(ctx, "Echo", request, newEcho_EchoYarpcResponse, options...)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*Pong)
	if !ok {
		return nil, protobuf.CastError(emptyEcho_EchoYarpcResponse, responseMessage)
	}
	return response, err
}

type _EchoYarpcHandler struct {
	server EchoYarpcServer
}

func (h *_EchoYarpcHandler) Echo(ctx context.Context, requestMessage proto.Message) (proto.Message, error) {
	var request *Ping
	var ok bool
	if requestMessage != nil {
		request, ok = requestMessage.(*Ping)
		if !ok {
			return nil, protobuf.CastError(emptyEcho_EchoYarpcRequest, requestMessage)
		}
	}
	response, err := h.server.Echo(ctx, request)
	if response == nil {
		return nil, err
	}
	return response, err
}

func newEcho_EchoYarpcRequest() proto.Message {
	return &Ping{}
}

func newEcho_EchoYarpcResponse() proto.Message {
	return &Pong{}
}

var (
	emptyEcho_EchoYarpcRequest  = &Ping{}
	emptyEcho_EchoYarpcResponse = &Pong{}
)

// OnewayYarpcClient is the yarpc client-side interface for the Oneway service.
type OnewayYarpcClient interface {
	Echo(context.Context, *Token, ...yarpc.CallOption) (yarpc.Ack, error)
}

// NewOnewayYarpcClient builds a new yarpc client for the Oneway service.
func NewOnewayYarpcClient(clientConfig transport.ClientConfig, options ...protobuf.ClientOption) OnewayYarpcClient {
	return &_OnewayYarpcCaller{protobuf.NewClient("uber.yarpc.internal.crossdock.Oneway", clientConfig, options...)}
}

// OnewayYarpcServer is the yarpc server-side interface for the Oneway service.
type OnewayYarpcServer interface {
	Echo(context.Context, *Token) error
}

// BuildOnewayYarpcProcedures prepares an implementation of the Oneway service for yarpc registration.
func BuildOnewayYarpcProcedures(server OnewayYarpcServer) []transport.Procedure {
	handler := &_OnewayYarpcHandler{server}
	return protobuf.BuildProcedures(
		"uber.yarpc.internal.crossdock.Oneway",
		map[string]transport.UnaryHandler{},
		map[string]transport.OnewayHandler{
			"Echo": protobuf.NewOnewayHandler(handler.Echo, newOneway_EchoYarpcRequest),
		},
	)
}

type _OnewayYarpcCaller struct {
	client protobuf.Client
}

func (c *_OnewayYarpcCaller) Echo(ctx context.Context, request *Token, options ...yarpc.CallOption) (yarpc.Ack, error) {
	return c.client.CallOneway(ctx, "Echo", request, options...)
}

type _OnewayYarpcHandler struct {
	server OnewayYarpcServer
}

func (h *_OnewayYarpcHandler) Echo(ctx context.Context, requestMessage proto.Message) error {
	var request *Token
	var ok bool
	if requestMessage != nil {
		request, ok = requestMessage.(*Token)
		if !ok {
			return protobuf.CastError(emptyOneway_EchoYarpcRequest, requestMessage)
		}
	}
	return h.server.Echo(ctx, request)
}

func newOneway_EchoYarpcRequest() proto.Message {
	return &Token{}
}

func newOneway_EchoYarpcResponse() proto.Message {
	return &yarpcproto.Oneway{}
}

var (
	emptyOneway_EchoYarpcRequest  = &Token{}
	emptyOneway_EchoYarpcResponse = &yarpcproto.Oneway{}
)

func init() {
	yarpc.RegisterClientBuilder(
		func(clientConfig transport.ClientConfig, structField reflect.StructField) EchoYarpcClient {
			return NewEchoYarpcClient(clientConfig, protobuf.ClientBuilderOptions(clientConfig, structField)...)
		},
	)
	yarpc.RegisterClientBuilder(
		func(clientConfig transport.ClientConfig, structField reflect.StructField) OnewayYarpcClient {
			return NewOnewayYarpcClient(clientConfig, protobuf.ClientBuilderOptions(clientConfig, structField)...)
		},
	)
}
