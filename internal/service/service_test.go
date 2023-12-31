package service

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"sebsegura/sample-lambda/pkg/client/mocks"
	"testing"
)

func TestService_Credit(t *testing.T) {
	type deps struct {
		ClientMock *mocks.Client
	}

	var (
		tests = []struct {
			name    string
			give    *CreditRequest
			wantErr assert.ErrorAssertionFunc
			on      func(*deps)
		}{
			{
				name: "should perform the action requested",
				give: &CreditRequest{
					UserID:      "1",
					Amount:      "1000",
					OperationID: "2",
				},
				wantErr: assert.NoError,
				on: func(d *deps) {
					d.ClientMock.On("GetOperations", mock.Anything, mock.Anything).Return(nil, nil)
				},
			},
			{
				name: "should return error if fails",
				give: &CreditRequest{
					UserID:      "1",
					Amount:      "1000",
					OperationID: "2",
				},
				wantErr: assert.Error,
				on: func(d *deps) {
					d.ClientMock.On("GetOperations", mock.Anything, mock.Anything).Return(nil, errors.New("error"))
				},
			},
		}
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &deps{
				ClientMock: mocks.NewClient(t),
			}

			svc := New(d.ClientMock)

			if tt.on != nil {
				tt.on(d)
			}

			err := svc.Credit(context.TODO(), tt.give)

			tt.wantErr(t, err)
		})
	}
}
