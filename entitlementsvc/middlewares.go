package entitlementsvc

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type Middleware func(Service) Service

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next   Service
	logger log.Logger
}

func (mw loggingMiddleware) GetEntitlement(ctx context.Context, id string) (p Entitlement, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "GetEntitlement", "id", id, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.GetEntitlement(ctx, id)
}
