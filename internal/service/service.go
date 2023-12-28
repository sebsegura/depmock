package service

import (
	"context"
	"sebsegura/sample-lambda/pkg/logger"
)

type Service struct{}

type CreditRequest struct {
	UserID      string `json:"user_uuid"`
	Amount      string `json:"amount"`
	OperationID string `json:"operation_id"`
}

func New() *Service {
	return &Service{}
}

func (svc *Service) Credit(ctx context.Context, in *CreditRequest) error {
	log := logger.Logger(ctx)
	log.Debug("hello world")
	return nil
}
