package endpoint

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
)

// LogginMiddleware returns an endpoint middleware that logs
// the duration of each invocation
func LogginMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req interface{}) (resp interface{}, err error) {
			defer func(begin time.Time) {
				logger.Log(
					"layer", "endpoint",
					"took", time.Since(begin),
				)
			}(time.Now())

			return next(ctx, req)
		}
	}
}

// InstrumentingMiddleware returns an endpoint middleware that records
// the duration of each invocation to the passed histogram. The middleware adds
// a single field: "success", which is "true" if no error is returned, and
// "false" otherwise.
func InstrumentingMiddleware(duration metrics.Histogram) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req interface{}) (resp interface{}, err error) {
			defer func(begin time.Time) {
				duration.With("success", fmt.Sprint(err != nil)).Observe(time.Since(begin).Seconds())
			}(time.Now())

			return next(ctx, req)
		}
	}
}
