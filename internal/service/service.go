package service

import "context"

type Service struct{}

type CreditRequest struct{}

func New() *Service {
	return &Service{}
}

func (svc *Service) Credit(ctx context.Context, in *CreditRequest) error {
	return nil
}
