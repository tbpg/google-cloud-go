// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package securitycenter

import (
	"context"
	"io"
	"log/slog"
	"net/http"

	"github.com/googleapis/gax-go/v2/internallog"
	"github.com/googleapis/gax-go/v2/internallog/grpclog"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

const serviceName = "securitycenter.googleapis.com"

// For more information on implementing a client constructor hook, see
// https://github.com/googleapis/google-cloud-go/wiki/Customizing-constructors.
type clientHookParams struct{}
type clientHook func(context.Context, clientHookParams) ([]option.ClientOption, error)

var versionClient string

func getVersionClient() string {
	if versionClient == "" {
		return "UNKNOWN"
	}
	return versionClient
}

// DefaultAuthScopes reports the default set of authentication scopes to use with this package.
func DefaultAuthScopes() []string {
	return []string{
		"https://www.googleapis.com/auth/cloud-platform",
	}
}

func executeHTTPRequestWithResponse(ctx context.Context, client *http.Client, req *http.Request, logger *slog.Logger, body []byte, rpc string) ([]byte, *http.Response, error) {
	logger.DebugContext(ctx, "api request", "serviceName", serviceName, "rpcName", rpc, "request", internallog.HTTPRequest(req, body))
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	logger.DebugContext(ctx, "api response", "serviceName", serviceName, "rpcName", rpc, "response", internallog.HTTPResponse(resp, buf))
	if err = googleapi.CheckResponseWithBody(resp, buf); err != nil {
		return nil, nil, err
	}
	return buf, resp, nil
}

func executeHTTPRequest(ctx context.Context, client *http.Client, req *http.Request, logger *slog.Logger, body []byte, rpc string) ([]byte, error) {
	buf, _, err := executeHTTPRequestWithResponse(ctx, client, req, logger, body, rpc)
	return buf, err
}

func executeStreamingHTTPRequest(ctx context.Context, client *http.Client, req *http.Request, logger *slog.Logger, body []byte, rpc string) (*http.Response, error) {
	logger.DebugContext(ctx, "api request", "serviceName", serviceName, "rpcName", rpc, "request", internallog.HTTPRequest(req, body))
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	logger.DebugContext(ctx, "api response", "serviceName", serviceName, "rpcName", rpc, "response", internallog.HTTPResponse(resp, nil))
	if err = googleapi.CheckResponse(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func executeRPC[I proto.Message, O proto.Message](ctx context.Context, fn func(context.Context, I, ...grpc.CallOption) (O, error), req I, opts []grpc.CallOption, logger *slog.Logger, rpc string) (O, error) {
	var zero O
	logger.DebugContext(ctx, "api request", "serviceName", serviceName, "rpcName", rpc, "request", grpclog.ProtoMessageRequest(ctx, req))
	resp, err := fn(ctx, req, opts...)
	if err != nil {
		return zero, err
	}
	logger.DebugContext(ctx, "api response", "serviceName", serviceName, "rpcName", rpc, "response", grpclog.ProtoMessageResponse(resp))
	return resp, err
}
