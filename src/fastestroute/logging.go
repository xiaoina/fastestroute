package main

import (
	"time"

	"github.com/go-kit/kit/log"
)

func loggingMiddleware(logger log.Logger) ServiceMiddleware {
	return func(next RouteService) RouteService {
		return logmw{logger, next}
	}
}

type logmw struct {
	logger log.Logger
	RouteService
}

func (mw logmw) Route(s []string) (output []GoogleMapsResponse, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "route",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.RouteService.Route(s)
	return
}
