package usage

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeUsageEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(string)
		err := validateSmartMeterId(req)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
}
