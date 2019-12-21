/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package grpc

import (
	"context"
	"fmt"
)

import (
	native_grpc "google.golang.org/grpc"
)

import (
	"github.com/apache/dubbo-go/config"
	"github.com/apache/dubbo-go/protocol"
	"github.com/apache/dubbo-go/protocol/grpc/internal"
	"github.com/apache/dubbo-go/protocol/invocation"
)

// userd grpc-dubbo biz service
func addService() {
	config.SetProviderService(NewGreeterProvider())
}

type greeterProvider struct {
	*greeterProviderBase
}

func NewGreeterProvider() *greeterProvider {
	return &greeterProvider{
		greeterProviderBase: &greeterProviderBase{},
	}
}

func (g *greeterProvider) SayHello(ctx context.Context, req *internal.HelloRequest) (reply *internal.HelloReply, err error) {
	fmt.Printf("req: %v", req)
	return &internal.HelloReply{Message: "this is message from reply"}, nil
}

func (g *greeterProvider) Reference() string {
	return "GrpcGreeterImpl"
}

// code generated by greeter.go
type greeterProviderBase struct {
	proxyImpl protocol.Invoker
}

func (g *greeterProviderBase) SetProxyImpl(impl protocol.Invoker) {
	g.proxyImpl = impl
}

func (g *greeterProviderBase) GetProxyImpl() protocol.Invoker {
	return g.proxyImpl
}

func (g *greeterProviderBase) ServiceDesc() *native_grpc.ServiceDesc {
	return &native_grpc.ServiceDesc{
		ServiceName: "helloworld.Greeter",
		HandlerType: (*internal.GreeterServer)(nil),
		Methods: []native_grpc.MethodDesc{
			{
				MethodName: "SayHello",
				Handler:    _DUBBO_Greeter_SayHello_Handler,
			},
		},
		Streams:  []native_grpc.StreamDesc{},
		Metadata: "helloworld.proto",
	}
}

func _DUBBO_Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor native_grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(internal.HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(DubboGrpcService)

	args := []interface{}{}
	args = append(args, in)
	invo := invocation.NewRPCInvocation("SayHello", args, nil)

	if interceptor == nil {
		result := base.GetProxyImpl().Invoke(invo)
		return result.Result(), result.Error()
	}
	info := &native_grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.GetProxyImpl().Invoke(invo)
		return result.Result(), result.Error()
	}
	return interceptor(ctx, in, info, handler)
}
