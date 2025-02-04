// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: gitpod/v1/prebuilds.proto

package v1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/gitpod-io/gitpod/public-api/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// PrebuildsServiceName is the fully-qualified name of the PrebuildsService service.
	PrebuildsServiceName = "gitpod.v1.PrebuildsService"
)

// PrebuildsServiceClient is a client for the gitpod.v1.PrebuildsService service.
type PrebuildsServiceClient interface {
	// GetPrebuild retrieves a single rebuild.
	// Errors:
	//
	//	NOT_FOUND if the prebuild_id does not exist
	GetPrebuild(context.Context, *connect_go.Request[v1.GetPrebuildRequest]) (*connect_go.Response[v1.GetPrebuildResponse], error)
	// GetRunningPrebuild returns the prebuild ID of a running prebuild,
	// or NOT_FOUND if there is no prebuild running for the content_url.
	GetRunningPrebuild(context.Context, *connect_go.Request[v1.GetRunningPrebuildRequest]) (*connect_go.Response[v1.GetRunningPrebuildResponse], error)
	// ListenToPrebuildStatus streams status updates for a prebuild. If the prebuild is already
	// in the Done phase, only that single status is streamed.
	ListenToPrebuildStatus(context.Context, *connect_go.Request[v1.ListenToPrebuildStatusRequest]) (*connect_go.ServerStreamForClient[v1.ListenToPrebuildStatusResponse], error)
	// ListenToPrebuildLogs returns the log output of a prebuild.
	// This does NOT include an image build if one happened.
	ListenToPrebuildLogs(context.Context, *connect_go.Request[v1.ListenToPrebuildLogsRequest]) (*connect_go.ServerStreamForClient[v1.ListenToPrebuildLogsResponse], error)
}

// NewPrebuildsServiceClient constructs a client for the gitpod.v1.PrebuildsService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPrebuildsServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) PrebuildsServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &prebuildsServiceClient{
		getPrebuild: connect_go.NewClient[v1.GetPrebuildRequest, v1.GetPrebuildResponse](
			httpClient,
			baseURL+"/gitpod.v1.PrebuildsService/GetPrebuild",
			opts...,
		),
		getRunningPrebuild: connect_go.NewClient[v1.GetRunningPrebuildRequest, v1.GetRunningPrebuildResponse](
			httpClient,
			baseURL+"/gitpod.v1.PrebuildsService/GetRunningPrebuild",
			opts...,
		),
		listenToPrebuildStatus: connect_go.NewClient[v1.ListenToPrebuildStatusRequest, v1.ListenToPrebuildStatusResponse](
			httpClient,
			baseURL+"/gitpod.v1.PrebuildsService/ListenToPrebuildStatus",
			opts...,
		),
		listenToPrebuildLogs: connect_go.NewClient[v1.ListenToPrebuildLogsRequest, v1.ListenToPrebuildLogsResponse](
			httpClient,
			baseURL+"/gitpod.v1.PrebuildsService/ListenToPrebuildLogs",
			opts...,
		),
	}
}

// prebuildsServiceClient implements PrebuildsServiceClient.
type prebuildsServiceClient struct {
	getPrebuild            *connect_go.Client[v1.GetPrebuildRequest, v1.GetPrebuildResponse]
	getRunningPrebuild     *connect_go.Client[v1.GetRunningPrebuildRequest, v1.GetRunningPrebuildResponse]
	listenToPrebuildStatus *connect_go.Client[v1.ListenToPrebuildStatusRequest, v1.ListenToPrebuildStatusResponse]
	listenToPrebuildLogs   *connect_go.Client[v1.ListenToPrebuildLogsRequest, v1.ListenToPrebuildLogsResponse]
}

// GetPrebuild calls gitpod.v1.PrebuildsService.GetPrebuild.
func (c *prebuildsServiceClient) GetPrebuild(ctx context.Context, req *connect_go.Request[v1.GetPrebuildRequest]) (*connect_go.Response[v1.GetPrebuildResponse], error) {
	return c.getPrebuild.CallUnary(ctx, req)
}

// GetRunningPrebuild calls gitpod.v1.PrebuildsService.GetRunningPrebuild.
func (c *prebuildsServiceClient) GetRunningPrebuild(ctx context.Context, req *connect_go.Request[v1.GetRunningPrebuildRequest]) (*connect_go.Response[v1.GetRunningPrebuildResponse], error) {
	return c.getRunningPrebuild.CallUnary(ctx, req)
}

// ListenToPrebuildStatus calls gitpod.v1.PrebuildsService.ListenToPrebuildStatus.
func (c *prebuildsServiceClient) ListenToPrebuildStatus(ctx context.Context, req *connect_go.Request[v1.ListenToPrebuildStatusRequest]) (*connect_go.ServerStreamForClient[v1.ListenToPrebuildStatusResponse], error) {
	return c.listenToPrebuildStatus.CallServerStream(ctx, req)
}

// ListenToPrebuildLogs calls gitpod.v1.PrebuildsService.ListenToPrebuildLogs.
func (c *prebuildsServiceClient) ListenToPrebuildLogs(ctx context.Context, req *connect_go.Request[v1.ListenToPrebuildLogsRequest]) (*connect_go.ServerStreamForClient[v1.ListenToPrebuildLogsResponse], error) {
	return c.listenToPrebuildLogs.CallServerStream(ctx, req)
}

// PrebuildsServiceHandler is an implementation of the gitpod.v1.PrebuildsService service.
type PrebuildsServiceHandler interface {
	// GetPrebuild retrieves a single rebuild.
	// Errors:
	//
	//	NOT_FOUND if the prebuild_id does not exist
	GetPrebuild(context.Context, *connect_go.Request[v1.GetPrebuildRequest]) (*connect_go.Response[v1.GetPrebuildResponse], error)
	// GetRunningPrebuild returns the prebuild ID of a running prebuild,
	// or NOT_FOUND if there is no prebuild running for the content_url.
	GetRunningPrebuild(context.Context, *connect_go.Request[v1.GetRunningPrebuildRequest]) (*connect_go.Response[v1.GetRunningPrebuildResponse], error)
	// ListenToPrebuildStatus streams status updates for a prebuild. If the prebuild is already
	// in the Done phase, only that single status is streamed.
	ListenToPrebuildStatus(context.Context, *connect_go.Request[v1.ListenToPrebuildStatusRequest], *connect_go.ServerStream[v1.ListenToPrebuildStatusResponse]) error
	// ListenToPrebuildLogs returns the log output of a prebuild.
	// This does NOT include an image build if one happened.
	ListenToPrebuildLogs(context.Context, *connect_go.Request[v1.ListenToPrebuildLogsRequest], *connect_go.ServerStream[v1.ListenToPrebuildLogsResponse]) error
}

// NewPrebuildsServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPrebuildsServiceHandler(svc PrebuildsServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/gitpod.v1.PrebuildsService/GetPrebuild", connect_go.NewUnaryHandler(
		"/gitpod.v1.PrebuildsService/GetPrebuild",
		svc.GetPrebuild,
		opts...,
	))
	mux.Handle("/gitpod.v1.PrebuildsService/GetRunningPrebuild", connect_go.NewUnaryHandler(
		"/gitpod.v1.PrebuildsService/GetRunningPrebuild",
		svc.GetRunningPrebuild,
		opts...,
	))
	mux.Handle("/gitpod.v1.PrebuildsService/ListenToPrebuildStatus", connect_go.NewServerStreamHandler(
		"/gitpod.v1.PrebuildsService/ListenToPrebuildStatus",
		svc.ListenToPrebuildStatus,
		opts...,
	))
	mux.Handle("/gitpod.v1.PrebuildsService/ListenToPrebuildLogs", connect_go.NewServerStreamHandler(
		"/gitpod.v1.PrebuildsService/ListenToPrebuildLogs",
		svc.ListenToPrebuildLogs,
		opts...,
	))
	return "/gitpod.v1.PrebuildsService/", mux
}

// UnimplementedPrebuildsServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPrebuildsServiceHandler struct{}

func (UnimplementedPrebuildsServiceHandler) GetPrebuild(context.Context, *connect_go.Request[v1.GetPrebuildRequest]) (*connect_go.Response[v1.GetPrebuildResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.PrebuildsService.GetPrebuild is not implemented"))
}

func (UnimplementedPrebuildsServiceHandler) GetRunningPrebuild(context.Context, *connect_go.Request[v1.GetRunningPrebuildRequest]) (*connect_go.Response[v1.GetRunningPrebuildResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.PrebuildsService.GetRunningPrebuild is not implemented"))
}

func (UnimplementedPrebuildsServiceHandler) ListenToPrebuildStatus(context.Context, *connect_go.Request[v1.ListenToPrebuildStatusRequest], *connect_go.ServerStream[v1.ListenToPrebuildStatusResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.PrebuildsService.ListenToPrebuildStatus is not implemented"))
}

func (UnimplementedPrebuildsServiceHandler) ListenToPrebuildLogs(context.Context, *connect_go.Request[v1.ListenToPrebuildLogsRequest], *connect_go.ServerStream[v1.ListenToPrebuildLogsResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.PrebuildsService.ListenToPrebuildLogs is not implemented"))
}
