// Package endpoint describes domain specific types
// here goes: infra operations (metrics), data validation
// and any other cross-domain concern
package endpoint

import (
	"github.com/MasterVentures/hackaton-ms-swingby/pkg/service"
	"github.com/go-kit/kit/endpoint"
)

//Endpoints collects all endpoints that compose a profile service
type Endpoints struct {
	CreateTransactionEndpoint endpoint.Endpoint

	//add endpoints here
}

// New returns a Endpoints struct that wraps the provided service
// and wires in all the passed endpoint middleware
func New(svc service.WalletService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		CreateTransactionEndpoint: MakeCreateTransactionEndpoint(svc),
	}

	for _, m := range mdw["CreateTransaction"] {
		eps.CreateTransactionEndpoint = m(eps.CreateTransactionEndpoint)
	}

	return eps
}
