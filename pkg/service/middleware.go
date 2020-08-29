package service

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

// Middleware is a decorator for any concrete service that implements BasicService
// any operation that needs to be processed at the service layer
// should be implemented as Middleware
type Middleware func(BasicService) BasicService

type logginMiddleware struct {
	logger log.Logger
	next   BasicService
}

// LogginMiddleware is a factory method that returns Middleware
func LogginMiddleware(logger log.Logger) Middleware {
	return func(svc BasicService) BasicService {
		return logginMiddleware{logger, svc}
	}
}

// logginMiddleware implements BasicService
func (mw logginMiddleware) CreateTransaction(
	ctx context.Context, buyer, seller string,
) (tr string, err error) {

	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "CreateTransaction",
			"buyer", buyer,
			"seller", seller,
			"output", tr,
			"error", err.Error(),
		)
	}(time.Now())

	tr, err = mw.next.CreateTransaction(ctx, buyer, seller)
	return
}
