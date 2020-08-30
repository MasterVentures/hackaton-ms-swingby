package service

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

// Middleware is a decorator for any concrete service that implements WalletService
// any operation that needs to be processed at the service layer
// should be implemented as Middleware
type Middleware func(WalletService) WalletService

type logginMiddleware struct {
	logger log.Logger
	next   WalletService
}

// LogginMiddleware is a factory method that returns Middleware
func LogginMiddleware(logger log.Logger) Middleware {
	return func(svc WalletService) WalletService {
		return logginMiddleware{logger, svc}
	}
}

// logginMiddleware implements WalletService
func (mw logginMiddleware) CreateTransaction(
	ctx context.Context, input []Input, output []Output, confirmations int,
) (tr string, err error) {

	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "CreateTransaction",
			"input", input,
			"output", output,
			"output", tr,
			"error", err.Error(),
		)
	}(time.Now())

	tr, err = mw.next.CreateTransaction(ctx, input, output, confirmations)
	return
}
