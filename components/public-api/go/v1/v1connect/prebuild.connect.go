// Copyright (c) 2025 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: gitpod/v1/prebuild.proto

package v1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/gitpod-io/gitpod/components/public-api/go/v1"
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
	// PrebuildServiceName is the fully-qualified name of the PrebuildService service.
	PrebuildServiceName = "gitpod.v1.PrebuildService"
)

// PrebuildServiceClient is a client for the gitpod.v1.PrebuildService service.
type PrebuildServiceClient interface {
	StartPrebuild(context.Context, *connect_go.Request[v1.StartPrebuildRequest]) (*connect_go.Response[v1.StartPrebuildResponse], error)
	CancelPrebuild(context.Context, *connect_go.Request[v1.CancelPrebuildRequest]) (*connect_go.Response[v1.CancelPrebuildResponse], error)
	GetPrebuild(context.Context, *connect_go.Request[v1.GetPrebuildRequest]) (*connect_go.Response[v1.GetPrebuildResponse], error)
	ListPrebuilds(context.Context, *connect_go.Request[v1.ListPrebuildsRequest]) (*connect_go.Response[v1.ListPrebuildsResponse], error)
	WatchPrebuild(context.Context, *connect_go.Request[v1.WatchPrebuildRequest]) (*connect_go.ServerStreamForClient[v1.WatchPrebuildResponse], error)
	// ListOrganizationPrebuilds lists all prebuilds of an organization
	ListOrganizationPrebuilds(context.Context, *connect_go.Request[v1.ListOrganizationPrebuildsRequest]) (*connect_go.Response[v1.ListOrganizationPrebuildsResponse], error)
}

// NewPrebuildServiceClient constructs a client for the gitpod.v1.PrebuildService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPrebuildServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) PrebuildServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &prebuildServiceClient{
		startPrebuild: connect_go.NewClient[v1.StartPrebuildRequest, v1.StartPrebuildResponse](
			httpClient,
			baseURL+"/gitpod.v1.PrebuildService/StartPrebuild",
			opts...,
		),
		cancelPrebuild: connect_go.NewClient[v1.CancelPrebuildRequest, v1.CancelPrebuildResponse](
			httpClient,
			baseURL+"/gitpod.v1.PrebuildService/CancelPrebuild",
			opts...,
		),
		getPrebuild: connect_go.NewClient[v1.GetPrebuildRequest, v1.GetPrebuildResponse](
			httpClient,
			baseURL+"/gitpod.v1.PrebuildService/GetPrebuild",
			opts...,
		),
		listPrebuilds: connect_go.NewClient[v1.ListPrebuildsRequest, v1.ListPrebuildsResponse](
			httpClient,
			baseURL+"/gitpod.v1.PrebuildService/ListPrebuilds",
			opts...,
		),
		watchPrebuild: connect_go.NewClient[v1.WatchPrebuildRequest, v1.WatchPrebuildResponse](
			httpClient,
			baseURL+"/gitpod.v1.PrebuildService/WatchPrebuild",
			opts...,
		),
		listOrganizationPrebuilds: connect_go.NewClient[v1.ListOrganizationPrebuildsRequest, v1.ListOrganizationPrebuildsResponse](
			httpClient,
			baseURL+"/gitpod.v1.PrebuildService/ListOrganizationPrebuilds",
			opts...,
		),
	}
}

// prebuildServiceClient implements PrebuildServiceClient.
type prebuildServiceClient struct {
	startPrebuild             *connect_go.Client[v1.StartPrebuildRequest, v1.StartPrebuildResponse]
	cancelPrebuild            *connect_go.Client[v1.CancelPrebuildRequest, v1.CancelPrebuildResponse]
	getPrebuild               *connect_go.Client[v1.GetPrebuildRequest, v1.GetPrebuildResponse]
	listPrebuilds             *connect_go.Client[v1.ListPrebuildsRequest, v1.ListPrebuildsResponse]
	watchPrebuild             *connect_go.Client[v1.WatchPrebuildRequest, v1.WatchPrebuildResponse]
	listOrganizationPrebuilds *connect_go.Client[v1.ListOrganizationPrebuildsRequest, v1.ListOrganizationPrebuildsResponse]
}

// StartPrebuild calls gitpod.v1.PrebuildService.StartPrebuild.
func (c *prebuildServiceClient) StartPrebuild(ctx context.Context, req *connect_go.Request[v1.StartPrebuildRequest]) (*connect_go.Response[v1.StartPrebuildResponse], error) {
	return c.startPrebuild.CallUnary(ctx, req)
}

// CancelPrebuild calls gitpod.v1.PrebuildService.CancelPrebuild.
func (c *prebuildServiceClient) CancelPrebuild(ctx context.Context, req *connect_go.Request[v1.CancelPrebuildRequest]) (*connect_go.Response[v1.CancelPrebuildResponse], error) {
	return c.cancelPrebuild.CallUnary(ctx, req)
}

// GetPrebuild calls gitpod.v1.PrebuildService.GetPrebuild.
func (c *prebuildServiceClient) GetPrebuild(ctx context.Context, req *connect_go.Request[v1.GetPrebuildRequest]) (*connect_go.Response[v1.GetPrebuildResponse], error) {
	return c.getPrebuild.CallUnary(ctx, req)
}

// ListPrebuilds calls gitpod.v1.PrebuildService.ListPrebuilds.
func (c *prebuildServiceClient) ListPrebuilds(ctx context.Context, req *connect_go.Request[v1.ListPrebuildsRequest]) (*connect_go.Response[v1.ListPrebuildsResponse], error) {
	return c.listPrebuilds.CallUnary(ctx, req)
}

// WatchPrebuild calls gitpod.v1.PrebuildService.WatchPrebuild.
func (c *prebuildServiceClient) WatchPrebuild(ctx context.Context, req *connect_go.Request[v1.WatchPrebuildRequest]) (*connect_go.ServerStreamForClient[v1.WatchPrebuildResponse], error) {
	return c.watchPrebuild.CallServerStream(ctx, req)
}

// ListOrganizationPrebuilds calls gitpod.v1.PrebuildService.ListOrganizationPrebuilds.
func (c *prebuildServiceClient) ListOrganizationPrebuilds(ctx context.Context, req *connect_go.Request[v1.ListOrganizationPrebuildsRequest]) (*connect_go.Response[v1.ListOrganizationPrebuildsResponse], error) {
	return c.listOrganizationPrebuilds.CallUnary(ctx, req)
}

// PrebuildServiceHandler is an implementation of the gitpod.v1.PrebuildService service.
type PrebuildServiceHandler interface {
	StartPrebuild(context.Context, *connect_go.Request[v1.StartPrebuildRequest]) (*connect_go.Response[v1.StartPrebuildResponse], error)
	CancelPrebuild(context.Context, *connect_go.Request[v1.CancelPrebuildRequest]) (*connect_go.Response[v1.CancelPrebuildResponse], error)
	GetPrebuild(context.Context, *connect_go.Request[v1.GetPrebuildRequest]) (*connect_go.Response[v1.GetPrebuildResponse], error)
	ListPrebuilds(context.Context, *connect_go.Request[v1.ListPrebuildsRequest]) (*connect_go.Response[v1.ListPrebuildsResponse], error)
	WatchPrebuild(context.Context, *connect_go.Request[v1.WatchPrebuildRequest], *connect_go.ServerStream[v1.WatchPrebuildResponse]) error
	// ListOrganizationPrebuilds lists all prebuilds of an organization
	ListOrganizationPrebuilds(context.Context, *connect_go.Request[v1.ListOrganizationPrebuildsRequest]) (*connect_go.Response[v1.ListOrganizationPrebuildsResponse], error)
}

// NewPrebuildServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPrebuildServiceHandler(svc PrebuildServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/gitpod.v1.PrebuildService/StartPrebuild", connect_go.NewUnaryHandler(
		"/gitpod.v1.PrebuildService/StartPrebuild",
		svc.StartPrebuild,
		opts...,
	))
	mux.Handle("/gitpod.v1.PrebuildService/CancelPrebuild", connect_go.NewUnaryHandler(
		"/gitpod.v1.PrebuildService/CancelPrebuild",
		svc.CancelPrebuild,
		opts...,
	))
	mux.Handle("/gitpod.v1.PrebuildService/GetPrebuild", connect_go.NewUnaryHandler(
		"/gitpod.v1.PrebuildService/GetPrebuild",
		svc.GetPrebuild,
		opts...,
	))
	mux.Handle("/gitpod.v1.PrebuildService/ListPrebuilds", connect_go.NewUnaryHandler(
		"/gitpod.v1.PrebuildService/ListPrebuilds",
		svc.ListPrebuilds,
		opts...,
	))
	mux.Handle("/gitpod.v1.PrebuildService/WatchPrebuild", connect_go.NewServerStreamHandler(
		"/gitpod.v1.PrebuildService/WatchPrebuild",
		svc.WatchPrebuild,
		opts...,
	))
	mux.Handle("/gitpod.v1.PrebuildService/ListOrganizationPrebuilds", connect_go.NewUnaryHandler(
		"/gitpod.v1.PrebuildService/ListOrganizationPrebuilds",
		svc.ListOrganizationPrebuilds,
		opts...,
	))
	return "/gitpod.v1.PrebuildService/", mux
}

// UnimplementedPrebuildServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPrebuildServiceHandler struct{}

func (UnimplementedPrebuildServiceHandler) StartPrebuild(context.Context, *connect_go.Request[v1.StartPrebuildRequest]) (*connect_go.Response[v1.StartPrebuildResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.PrebuildService.StartPrebuild is not implemented"))
}

func (UnimplementedPrebuildServiceHandler) CancelPrebuild(context.Context, *connect_go.Request[v1.CancelPrebuildRequest]) (*connect_go.Response[v1.CancelPrebuildResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.PrebuildService.CancelPrebuild is not implemented"))
}

func (UnimplementedPrebuildServiceHandler) GetPrebuild(context.Context, *connect_go.Request[v1.GetPrebuildRequest]) (*connect_go.Response[v1.GetPrebuildResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.PrebuildService.GetPrebuild is not implemented"))
}

func (UnimplementedPrebuildServiceHandler) ListPrebuilds(context.Context, *connect_go.Request[v1.ListPrebuildsRequest]) (*connect_go.Response[v1.ListPrebuildsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.PrebuildService.ListPrebuilds is not implemented"))
}

func (UnimplementedPrebuildServiceHandler) WatchPrebuild(context.Context, *connect_go.Request[v1.WatchPrebuildRequest], *connect_go.ServerStream[v1.WatchPrebuildResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.PrebuildService.WatchPrebuild is not implemented"))
}

func (UnimplementedPrebuildServiceHandler) ListOrganizationPrebuilds(context.Context, *connect_go.Request[v1.ListOrganizationPrebuildsRequest]) (*connect_go.Response[v1.ListOrganizationPrebuildsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.PrebuildService.ListOrganizationPrebuilds is not implemented"))
}
