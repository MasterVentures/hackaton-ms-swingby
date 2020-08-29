// Package endpoint describes domain specific types
// here goes: infra operations (metrics), data validation
// and any other cross-domain concern
package endpoint

import (
	"context"

	"github.com/MasterVentures/hackaton-ms-swingby/pkg/service"
	"github.com/go-kit/kit/endpoint"
)

//CreateTransactionRequest defines fields to create transaction
type CreateTransactionRequest struct {
	Buyer  string `json:"buyer"`
	Seller string `json:"seller"`
}

//CreateTransactionResponse defines the request response
type CreateTransactionResponse struct {
	ID  string `json:"id"`
	Err string `json:"err"`
}

//MakeCreateTransactionEndpoint returns an endpoint that invokes CreateTransaction on the service
func MakeCreateTransactionEndpoint(svc service.BasicService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (resp interface{}, err error) {
		req, ok := r.(CreateTransactionRequest)
		if !ok {
			//TODO: handle request parsing errors (BadRequest?)
			return nil, nil
		}
		id, err := svc.CreateTransaction(ctx, req.Buyer, req.Seller)
		resp = CreateTransactionResponse{ID: id, Err: err.Error()}
		return
	}
}

//Endpoints collects all endpoints that compose a profile service
type Endpoints struct {
	CreateTransactionEndpoint endpoint.Endpoint

	//add endpoints here
}

// New returns a Endpoints struct that wraps the provided service
// and wires in all the passed endpoint middleware
func New(svc service.BasicService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		CreateTransactionEndpoint: MakeCreateTransactionEndpoint(svc),
	}

	for _, m := range mdw["CreateTransaction"] {
		eps.CreateTransactionEndpoint = m(eps.CreateTransactionEndpoint)
	}

	return eps
}
