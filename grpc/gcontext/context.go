package gcontext

import (
	"context"
	"fmt"
	"github.com/ducketlab/mingo/http/request"
	"github.com/rs/xid"
	"google.golang.org/grpc/metadata"
	"net/http"
)

const (
	NamespaceHeader         = "x-rpc-namespace"
	InternalCallTokenHeader = "internal-call-token"
	ClientIdHeader          = "client-id"
	ClientSecretHeader      = "client-secret"
	OauthTokenHeader        = "x-oauth-token"
	RealIpHeader            = "x-real-ip"
	UserAgentHeader         = "user-agent"
	RequestId               = "x-request-id"
)

type GrpcInCtx struct {
	*grpcCtx
}

func NewGrpcInCtx() *GrpcInCtx {
	return &GrpcInCtx{newGrpcCtx(metadata.Pairs())}
}

func GetGrpcInCtx(ctx context.Context) (*GrpcInCtx, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("ctx is not an grpc incoming context")
	}

	return &GrpcInCtx{newGrpcCtx(md)}, nil
}

func (c *GrpcInCtx) Context() context.Context {
	return metadata.NewIncomingContext(context.Background(), c.md)
}

func (c *GrpcInCtx) SetClientCredentials(clientId string, clientSecret string) {
	c.set(ClientIdHeader, clientId)
	c.set(ClientSecretHeader, clientSecret)
}

func (c *GrpcInCtx) GetClientId() string {
	return c.get(ClientIdHeader)
}

func (c *GrpcInCtx) GetClientSecret() string {
	return c.get(ClientSecretHeader)
}

func (c *GrpcInCtx) GetNamespace() string {
	return c.get(NamespaceHeader)
}

func (c *GrpcInCtx) SetAccessToKen(token string) {
	c.set(OauthTokenHeader, token)
}

func (c *GrpcInCtx) GetAccessToKen() string {
	return c.get(OauthTokenHeader)
}

func (c *GrpcInCtx) GetRequestID() string {
	return c.get(RequestId)
}

type GrpcOutCtx struct {
	*grpcCtx
}

func NewGrpcOutCtx() *GrpcOutCtx  {
	return &GrpcOutCtx{newGrpcCtx(metadata.Pairs())}
}

func NewGrpcOutCtxFromIn(in *GrpcInCtx) *GrpcOutCtx  {
	return &GrpcOutCtx{newGrpcCtx(in.md)}
}

func (c *GrpcOutCtx) Context() context.Context {
	return metadata.NewOutgoingContext(context.Background(), c.md)
}

func (c *GrpcOutCtx) SetNamespace(ns string) {
	c.set(NamespaceHeader, ns)
}

func NewGrpcOutCtxFromHttpRequest(r *http.Request) (*GrpcOutCtx, error)  {
	rc := NewGrpcOutCtx()
	rc.SetAccessToken(GetTokenFromHeader(r))
	rc.SetRemoteIp(request.GetRemoteIP(r))
	rc.SetUserAgent(r.UserAgent())

	rid := r.Header.Get(RequestId)
	if rid == "" {
		rid = xid.New().String()
	}

	rc.SetRequestId(rid)

	return rc, nil
}

func NewGrpcInCtxFromHttpRequest(r *http.Request) (*GrpcInCtx, error)  {
	rc := NewGrpcInCtx()
	rc.SetAccessToKen(GetTokenFromHeader(r))
	rc.SetRemoteIp(request.GetRemoteIP(r))
	rc.SetUserAgent(r.UserAgent())
	rid := r.Header.Get(RequestId)
	if rid == "" {
		rid = xid.New().String()
	}

	return rc, nil
}

type grpcCtx struct {
	md metadata.MD
}

func newGrpcCtx(md metadata.MD) *grpcCtx {
	return &grpcCtx{md: md}
}

func (c *grpcCtx) get(key string) string {
	return c.getWithIndex(key, 0)
}

func (c *grpcCtx) getWithIndex(key string, index int) string {
	if val, ok := c.md[key]; ok {
		if len(val) > index {
			return val[index]
		}
	}
	return ""
}

func (c *grpcCtx) set(key string, values ...string) {
	c.md.Set(key, values...)
}

func (c *grpcCtx) SetAccessToken(ak string) {
	c.set(OauthTokenHeader, ak)
}

func (c *grpcCtx) SetRequestId(requestId string) {
	c.set(RequestId, requestId)
}

func (c *grpcCtx) SetRemoteIp(ip string) {
	c.set(RealIpHeader, ip)
}

func (c *grpcCtx) SetUserAgent(ua string) {
	c.set(UserAgentHeader, ua)
}

func GetTokenFromHeader(r *http.Request) string {
	tk := r.Header.Get(OauthTokenHeader)
	if tk != "" {
		return tk
	}

	return r.Header.Get("Authorization")
}

func GetClientCredentialsFromHTTPRequest(r *http.Request) (cid, cs string) {
	cid, cs = r.Header.Get(ClientIdHeader), r.Header.Get(ClientSecretHeader)
	return
}