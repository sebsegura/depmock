package service

import (
	"context"
	"sebsegura/sample-lambda/pkg/client"
)

type Service struct {
	client client.Client
}

type CreditRequest struct {
	UserID      string `json:"user_uuid"`
	Amount      string `json:"amount"`
	OperationID string `json:"operation_id"`
}

func New(client client.Client) *Service {
	return &Service{
		client: client,
	}
}

func (svc *Service) Credit(ctx context.Context, in *CreditRequest) error {
	_, err := svc.client.GetOperations(ctx, &client.GetOperationsRequest{
		UserID:      in.UserID,
		OperationID: in.OperationID,
	})

	return err
}
