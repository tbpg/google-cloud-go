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

package generativelanguage

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"net/url"

	generativelanguagepb "cloud.google.com/go/ai/generativelanguage/apiv1alpha/generativelanguagepb"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newTextClientHook clientHook

// TextCallOptions contains the retry settings for each method of TextClient.
type TextCallOptions struct {
	GenerateText    []gax.CallOption
	EmbedText       []gax.CallOption
	BatchEmbedText  []gax.CallOption
	CountTextTokens []gax.CallOption
	GetOperation    []gax.CallOption
	ListOperations  []gax.CallOption
}

func defaultTextGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("generativelanguage.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("generativelanguage.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("generativelanguage.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://generativelanguage.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultTextCallOptions() *TextCallOptions {
	return &TextCallOptions{
		GenerateText:    []gax.CallOption{},
		EmbedText:       []gax.CallOption{},
		BatchEmbedText:  []gax.CallOption{},
		CountTextTokens: []gax.CallOption{},
		GetOperation:    []gax.CallOption{},
		ListOperations:  []gax.CallOption{},
	}
}

func defaultTextRESTCallOptions() *TextCallOptions {
	return &TextCallOptions{
		GenerateText:    []gax.CallOption{},
		EmbedText:       []gax.CallOption{},
		BatchEmbedText:  []gax.CallOption{},
		CountTextTokens: []gax.CallOption{},
		GetOperation:    []gax.CallOption{},
		ListOperations:  []gax.CallOption{},
	}
}

// internalTextClient is an interface that defines the methods available from Generative Language API.
type internalTextClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GenerateText(context.Context, *generativelanguagepb.GenerateTextRequest, ...gax.CallOption) (*generativelanguagepb.GenerateTextResponse, error)
	EmbedText(context.Context, *generativelanguagepb.EmbedTextRequest, ...gax.CallOption) (*generativelanguagepb.EmbedTextResponse, error)
	BatchEmbedText(context.Context, *generativelanguagepb.BatchEmbedTextRequest, ...gax.CallOption) (*generativelanguagepb.BatchEmbedTextResponse, error)
	CountTextTokens(context.Context, *generativelanguagepb.CountTextTokensRequest, ...gax.CallOption) (*generativelanguagepb.CountTextTokensResponse, error)
	GetOperation(context.Context, *longrunningpb.GetOperationRequest, ...gax.CallOption) (*longrunningpb.Operation, error)
	ListOperations(context.Context, *longrunningpb.ListOperationsRequest, ...gax.CallOption) *OperationIterator
}

// TextClient is a client for interacting with Generative Language API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// API for using Generative Language Models (GLMs) trained to generate text.
//
// Also known as Large Language Models (LLM)s, these generate text given an
// input prompt from the user.
type TextClient struct {
	// The internal transport-dependent client.
	internalClient internalTextClient

	// The call options for this service.
	CallOptions *TextCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *TextClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *TextClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *TextClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GenerateText generates a response from the model given an input message.
func (c *TextClient) GenerateText(ctx context.Context, req *generativelanguagepb.GenerateTextRequest, opts ...gax.CallOption) (*generativelanguagepb.GenerateTextResponse, error) {
	return c.internalClient.GenerateText(ctx, req, opts...)
}

// EmbedText generates an embedding from the model given an input message.
func (c *TextClient) EmbedText(ctx context.Context, req *generativelanguagepb.EmbedTextRequest, opts ...gax.CallOption) (*generativelanguagepb.EmbedTextResponse, error) {
	return c.internalClient.EmbedText(ctx, req, opts...)
}

// BatchEmbedText generates multiple embeddings from the model given input text in a
// synchronous call.
func (c *TextClient) BatchEmbedText(ctx context.Context, req *generativelanguagepb.BatchEmbedTextRequest, opts ...gax.CallOption) (*generativelanguagepb.BatchEmbedTextResponse, error) {
	return c.internalClient.BatchEmbedText(ctx, req, opts...)
}

// CountTextTokens runs a model’s tokenizer on a text and returns the token count.
func (c *TextClient) CountTextTokens(ctx context.Context, req *generativelanguagepb.CountTextTokensRequest, opts ...gax.CallOption) (*generativelanguagepb.CountTextTokensResponse, error) {
	return c.internalClient.CountTextTokens(ctx, req, opts...)
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *TextClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	return c.internalClient.GetOperation(ctx, req, opts...)
}

// ListOperations is a utility method from google.longrunning.Operations.
func (c *TextClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	return c.internalClient.ListOperations(ctx, req, opts...)
}

// textGRPCClient is a client for interacting with Generative Language API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type textGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing TextClient
	CallOptions **TextCallOptions

	// The gRPC API client.
	textClient generativelanguagepb.TextServiceClient

	operationsClient longrunningpb.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewTextClient creates a new text service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// API for using Generative Language Models (GLMs) trained to generate text.
//
// Also known as Large Language Models (LLM)s, these generate text given an
// input prompt from the user.
func NewTextClient(ctx context.Context, opts ...option.ClientOption) (*TextClient, error) {
	clientOpts := defaultTextGRPCClientOptions()
	if newTextClientHook != nil {
		hookOpts, err := newTextClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := TextClient{CallOptions: defaultTextCallOptions()}

	c := &textGRPCClient{
		connPool:         connPool,
		textClient:       generativelanguagepb.NewTextServiceClient(connPool),
		CallOptions:      &client.CallOptions,
		logger:           internaloption.GetLogger(opts),
		operationsClient: longrunningpb.NewOperationsClient(connPool),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *textGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *textGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *textGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type textRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing TextClient
	CallOptions **TextCallOptions

	logger *slog.Logger
}

// NewTextRESTClient creates a new text service rest client.
//
// API for using Generative Language Models (GLMs) trained to generate text.
//
// Also known as Large Language Models (LLM)s, these generate text given an
// input prompt from the user.
func NewTextRESTClient(ctx context.Context, opts ...option.ClientOption) (*TextClient, error) {
	clientOpts := append(defaultTextRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultTextRESTCallOptions()
	c := &textRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
		logger:      internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	return &TextClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultTextRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://generativelanguage.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://generativelanguage.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://generativelanguage.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://generativelanguage.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableNewAuthLibrary(),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *textRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *textRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *textRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *textGRPCClient) GenerateText(ctx context.Context, req *generativelanguagepb.GenerateTextRequest, opts ...gax.CallOption) (*generativelanguagepb.GenerateTextResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "model", url.QueryEscape(req.GetModel()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GenerateText[0:len((*c.CallOptions).GenerateText):len((*c.CallOptions).GenerateText)], opts...)
	var resp *generativelanguagepb.GenerateTextResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.textClient.GenerateText, req, settings.GRPC, c.logger, "GenerateText")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *textGRPCClient) EmbedText(ctx context.Context, req *generativelanguagepb.EmbedTextRequest, opts ...gax.CallOption) (*generativelanguagepb.EmbedTextResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "model", url.QueryEscape(req.GetModel()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).EmbedText[0:len((*c.CallOptions).EmbedText):len((*c.CallOptions).EmbedText)], opts...)
	var resp *generativelanguagepb.EmbedTextResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.textClient.EmbedText, req, settings.GRPC, c.logger, "EmbedText")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *textGRPCClient) BatchEmbedText(ctx context.Context, req *generativelanguagepb.BatchEmbedTextRequest, opts ...gax.CallOption) (*generativelanguagepb.BatchEmbedTextResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "model", url.QueryEscape(req.GetModel()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).BatchEmbedText[0:len((*c.CallOptions).BatchEmbedText):len((*c.CallOptions).BatchEmbedText)], opts...)
	var resp *generativelanguagepb.BatchEmbedTextResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.textClient.BatchEmbedText, req, settings.GRPC, c.logger, "BatchEmbedText")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *textGRPCClient) CountTextTokens(ctx context.Context, req *generativelanguagepb.CountTextTokensRequest, opts ...gax.CallOption) (*generativelanguagepb.CountTextTokensResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "model", url.QueryEscape(req.GetModel()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CountTextTokens[0:len((*c.CallOptions).CountTextTokens):len((*c.CallOptions).CountTextTokens)], opts...)
	var resp *generativelanguagepb.CountTextTokensResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.textClient.CountTextTokens, req, settings.GRPC, c.logger, "CountTextTokens")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *textGRPCClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.operationsClient.GetOperation, req, settings.GRPC, c.logger, "GetOperation")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *textGRPCClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListOperations[0:len((*c.CallOptions).ListOperations):len((*c.CallOptions).ListOperations)], opts...)
	it := &OperationIterator{}
	req = proto.Clone(req).(*longrunningpb.ListOperationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*longrunningpb.Operation, string, error) {
		resp := &longrunningpb.ListOperationsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = executeRPC(ctx, c.operationsClient.ListOperations, req, settings.GRPC, c.logger, "ListOperations")
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetOperations(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// GenerateText generates a response from the model given an input message.
func (c *textRESTClient) GenerateText(ctx context.Context, req *generativelanguagepb.GenerateTextRequest, opts ...gax.CallOption) (*generativelanguagepb.GenerateTextResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1alpha/%v:generateText", req.GetModel())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "model", url.QueryEscape(req.GetModel()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GenerateText[0:len((*c.CallOptions).GenerateText):len((*c.CallOptions).GenerateText)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &generativelanguagepb.GenerateTextResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "GenerateText")
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// EmbedText generates an embedding from the model given an input message.
func (c *textRESTClient) EmbedText(ctx context.Context, req *generativelanguagepb.EmbedTextRequest, opts ...gax.CallOption) (*generativelanguagepb.EmbedTextResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1alpha/%v:embedText", req.GetModel())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "model", url.QueryEscape(req.GetModel()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).EmbedText[0:len((*c.CallOptions).EmbedText):len((*c.CallOptions).EmbedText)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &generativelanguagepb.EmbedTextResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "EmbedText")
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// BatchEmbedText generates multiple embeddings from the model given input text in a
// synchronous call.
func (c *textRESTClient) BatchEmbedText(ctx context.Context, req *generativelanguagepb.BatchEmbedTextRequest, opts ...gax.CallOption) (*generativelanguagepb.BatchEmbedTextResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1alpha/%v:batchEmbedText", req.GetModel())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "model", url.QueryEscape(req.GetModel()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).BatchEmbedText[0:len((*c.CallOptions).BatchEmbedText):len((*c.CallOptions).BatchEmbedText)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &generativelanguagepb.BatchEmbedTextResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "BatchEmbedText")
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// CountTextTokens runs a model’s tokenizer on a text and returns the token count.
func (c *textRESTClient) CountTextTokens(ctx context.Context, req *generativelanguagepb.CountTextTokensRequest, opts ...gax.CallOption) (*generativelanguagepb.CountTextTokensResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1alpha/%v:countTextTokens", req.GetModel())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "model", url.QueryEscape(req.GetModel()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).CountTextTokens[0:len((*c.CallOptions).CountTextTokens):len((*c.CallOptions).CountTextTokens)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &generativelanguagepb.CountTextTokensResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "CountTextTokens")
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *textRESTClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1alpha/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "GetOperation")
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// ListOperations is a utility method from google.longrunning.Operations.
func (c *textRESTClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	it := &OperationIterator{}
	req = proto.Clone(req).(*longrunningpb.ListOperationsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*longrunningpb.Operation, string, error) {
		resp := &longrunningpb.ListOperationsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		baseUrl, err := url.Parse(c.endpoint)
		if err != nil {
			return nil, "", err
		}
		baseUrl.Path += fmt.Sprintf("/v1alpha/%v/operations", req.GetName())

		params := url.Values{}
		params.Add("$alt", "json;enum-encoding=int")
		if req.GetFilter() != "" {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req.GetPageSize() != 0 {
			params.Add("pageSize", fmt.Sprintf("%v", req.GetPageSize()))
		}
		if req.GetPageToken() != "" {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		hds := append(c.xGoogHeaders, "Content-Type", "application/json")
		headers := gax.BuildHeaders(ctx, hds...)
		e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			if settings.Path != "" {
				baseUrl.Path = settings.Path
			}
			httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
			if err != nil {
				return err
			}
			httpReq.Header = headers

			buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "ListOperations")
			if err != nil {
				return err
			}
			if err := unm.Unmarshal(buf, resp); err != nil {
				return err
			}

			return nil
		}, opts...)
		if e != nil {
			return nil, "", e
		}
		it.Response = resp
		return resp.GetOperations(), resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}
