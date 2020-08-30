package endpoint

import (
	"context"
	"github.com/MasterVentures/hackaton-ms-swingby/pkg/service"
	"github.com/go-kit/kit/endpoint"
)

//MakeCreateTransactionEndpoint returns an endpoint that invokes CreateTransaction on the service
func MakeCreateRequestEndpoint(svc service.WalletService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (resp interface{}, err error) {
		req, ok := r.(service.CreateReqRequest)
		if !ok {
			//TODO: handle request parsing errors (BadRequest?)
			return nil, nil
		}
		id, err := svc.CreateRequest(ctx)
		return service.CreateReqResponse{ID: id, Err: err.Error()}, nil
	}
}
