package context

import (
	"context"
	httppb "github.com/ducketlab/mingo/pb/http"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type key int

const defaultKey = key(1)

func WithContext(req *http.Request, rctx *ReqContext) *http.Request {
	ctx := context.WithValue(req.Context(), defaultKey, rctx)
	return req.WithContext(ctx)
}

func GetContext(req *http.Request) *ReqContext {
	if v, ok := req.Context().Value(defaultKey).(*ReqContext); ok {
		return v
	}

	return new(ReqContext)
}

type ReqContext struct {
	Entry    *httppb.Entry
	Ps       httprouter.Params
	AuthInfo interface{}
}
