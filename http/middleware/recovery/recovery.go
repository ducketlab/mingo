package recovery

import (
	"fmt"
	"github.com/ducketlab/mingo/exception"
	"github.com/ducketlab/mingo/http/response"
	"github.com/ducketlab/mingo/http/router"
	"github.com/ducketlab/mingo/logger"
	"go.uber.org/zap"
	"log"
	"net/http"
)

const recoveryExplanation = "Something went wrong"

type recovery struct {
	log   logger.WithMetaLogger
	debug bool
}

func (r *recovery) Debug(on bool) {
	r.debug = on
}

func NewWithLogger(l logger.WithMetaLogger) router.Middleware {
	return &recovery{
		log: l,
	}
}

func (r *recovery) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if request := recover(); request != nil {
				msg := fmt.Sprintf("%s. Recovering, but please report this.", recoveryExplanation)
				stack := r.stack()

				r.logf(msg, request, stack)

				if r.debug {
					msg += stack
				}

				response.Failed(writer, exception.NewInternalServerError(msg))
				return
			}
		}()

		next.ServeHTTP(writer, request)
	})
}

func (r *recovery) stack() string {
	return zap.Stack("stack").String
}

func (r *recovery) logf(msg string, request interface{}, stack string) {
	if r.log != nil {
		r.log.Errorw(msg,
			logger.NewAny("panic", request),
			logger.NewAny("stack", stack))
	}
	log.Println(msg, request, stack)
}
