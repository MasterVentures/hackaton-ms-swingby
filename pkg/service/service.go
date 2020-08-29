// Package service holds business logic and
// service layer middleware
package service

import "context"

// BasicService is the interface that expose the service
// available methods and it's the core of the service
type BasicService interface {
	//This is just an example
	CreateTransaction(ctx context.Context, buyer, seller string) (string, error)
}

// concrete service, dependencies go in here
type basicService struct{}

// basicService inplements BasicService
func (bs *basicService) CreateTransaction(ctx context.Context, buyer, seller string) (id string, err error) {
	//business logic goes here
	return id, nil
}

// NewBasicService is a factory for basicService
// dependencies are passed as parameters
func NewBasicService() BasicService {
	return &basicService{}
}

// New returns a BasicService with all of the expected middleware wired in.
func New(middleware []Middleware) BasicService {
	var svc BasicService = NewBasicService()
	for _, mdw := range middleware {
		svc = mdw(svc)
	}
	return svc
}
