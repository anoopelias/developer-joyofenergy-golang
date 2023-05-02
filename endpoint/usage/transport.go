package usage

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-kit/kit/endpoint"
	kitlogrus "github.com/go-kit/kit/log/logrus"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/sirupsen/logrus"

	mhttp "joi-energy-golang/http"
	"joi-energy-golang/http/middleware"
	"joi-energy-golang/http/serveroption"
)

func MakeGetUsageHandler(
	logger *logrus.Entry,
) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerBefore(serveroption.ExtractAcceptHeaderIntoContext),
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(kitlogrus.NewLogrusLogger(logger))),
		kithttp.ServerErrorEncoder(middleware.MakeEncodeErrorFunc(logger)),
	}

	mw := endpoint.Chain(
		middleware.MakeAcceptHeaderValidationMiddleware(),
	)

	endpointHandler := kithttp.NewServer(
		mw(func(ctx context.Context, request interface{}) (interface{}, error) {
			return nil, nil
		}),
		decodeSmartMeterIdFromRequest,
		mhttp.EncodeResponse,
		opts...,
	)

	return endpointHandler
}

func decodeSmartMeterIdFromRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return strings.Split(r.URL.Path, "/")[2], nil
}
