// Package service holds business logic and
// service layer middleware
package service

import (
	"context"
)

// WalletService is the interface that expose the service
// available methods and it's the core of the service
type WalletService interface {
	//This is just an example
	CreateTransaction(ctx context.Context, input []Input, output []Output, Confirmation int) (string, error)
	CreateWithdrawal(ctx context.Context, Addresses []Addresses) (string, error)
	CreateDeposit(ctx context.Context) (string, error)
	CreateRequest(ctx context.Context) (string, error)
}

// concrete service, dependencies go in here
type walletService struct{}

// walletService implements WalletService
func (bs *walletService) CreateTransaction(ctx context.Context, input []Input, output []Output, Confirmation int) (id string, err error) {
	//business logic goes here
	return id, nil
}

func (bs *walletService) CreateWithdrawal(ctx context.Context, addresses []Addresses) (id string, err error) {
	//business logic goes here
	return id, nil
}

func (bs *walletService) CreateDeposit(ctx context.Context) (id string, err error) {
	//business logic goes here
	return id, nil
}

func (bs *walletService) CreateRequest(ctx context.Context) (id string, err error) {
	//business logic goes here
	return id, nil
}

// NewWalletService is a factory for walletService
// dependencies are passed as parameters
func NewWalletService() WalletService {
	return &walletService{}
}

// New returns a WalletService with all of the expected middleware wired in.
func New(middleware []Middleware) WalletService {
	var svc WalletService = NewWalletService()
	for _, mdw := range middleware {
		svc = mdw(svc)
	}
	return svc
}
