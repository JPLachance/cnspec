// Code generated by protoc-gen-rangerrpc version DO NOT EDIT.
// source: scan.proto

package scan

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"

	ranger "go.mondoo.com/ranger-rpc"
	"go.mondoo.com/ranger-rpc/metadata"
	jsonpb "google.golang.org/protobuf/encoding/protojson"
	pb "google.golang.org/protobuf/proto"
)

// service interface definition

type Scan interface {
	Run(context.Context, *Job) (*ScanResult, error)
	RunIncognito(context.Context, *Job) (*ScanResult, error)
	Schedule(context.Context, *Job) (*Empty, error)
	RunAdmissionReview(context.Context, *AdmissionReviewJob) (*ScanResult, error)
	GarbageCollectAssets(context.Context, *GarbageCollectOptions) (*Empty, error)
}

// client implementation

type ScanClient struct {
	ranger.Client
	httpclient ranger.HTTPClient
	prefix     string
}

func NewScanClient(addr string, client ranger.HTTPClient, plugins ...ranger.ClientPlugin) (*ScanClient, error) {
	base, err := url.Parse(ranger.SanitizeUrl(addr))
	if err != nil {
		return nil, err
	}

	u, err := url.Parse("./Scan")
	if err != nil {
		return nil, err
	}

	serviceClient := &ScanClient{
		httpclient: client,
		prefix:     base.ResolveReference(u).String(),
	}
	serviceClient.AddPlugins(plugins...)
	return serviceClient, nil
}
func (c *ScanClient) Run(ctx context.Context, in *Job) (*ScanResult, error) {
	out := new(ScanResult)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/Run"}, ""), in, out)
	return out, err
}
func (c *ScanClient) RunIncognito(ctx context.Context, in *Job) (*ScanResult, error) {
	out := new(ScanResult)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/RunIncognito"}, ""), in, out)
	return out, err
}
func (c *ScanClient) Schedule(ctx context.Context, in *Job) (*Empty, error) {
	out := new(Empty)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/Schedule"}, ""), in, out)
	return out, err
}
func (c *ScanClient) RunAdmissionReview(ctx context.Context, in *AdmissionReviewJob) (*ScanResult, error) {
	out := new(ScanResult)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/RunAdmissionReview"}, ""), in, out)
	return out, err
}
func (c *ScanClient) GarbageCollectAssets(ctx context.Context, in *GarbageCollectOptions) (*Empty, error) {
	out := new(Empty)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/GarbageCollectAssets"}, ""), in, out)
	return out, err
}

// server implementation

type ScanServerOption func(s *ScanServer)

func WithUnknownFieldsForScanServer() ScanServerOption {
	return func(s *ScanServer) {
		s.allowUnknownFields = true
	}
}

func NewScanServer(handler Scan, opts ...ScanServerOption) http.Handler {
	srv := &ScanServer{
		handler: handler,
	}

	for i := range opts {
		opts[i](srv)
	}

	service := ranger.Service{
		Name: "Scan",
		Methods: map[string]ranger.Method{
			"Run":                  srv.Run,
			"RunIncognito":         srv.RunIncognito,
			"Schedule":             srv.Schedule,
			"RunAdmissionReview":   srv.RunAdmissionReview,
			"GarbageCollectAssets": srv.GarbageCollectAssets,
		},
	}
	return ranger.NewRPCServer(&service)
}

type ScanServer struct {
	handler            Scan
	allowUnknownFields bool
}

func (p *ScanServer) Run(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req Job
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.Run(ctx, &req)
}
func (p *ScanServer) RunIncognito(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req Job
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.RunIncognito(ctx, &req)
}
func (p *ScanServer) Schedule(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req Job
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.Schedule(ctx, &req)
}
func (p *ScanServer) RunAdmissionReview(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req AdmissionReviewJob
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.RunAdmissionReview(ctx, &req)
}
func (p *ScanServer) GarbageCollectAssets(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req GarbageCollectOptions
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.GarbageCollectAssets(ctx, &req)
}
