package client

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"sebsegura/sample-lambda/pkg/lambda"
	"testing"
)

func TestClient_Authenticate(t *testing.T) {
	var (
		mockServer *httptest.Server
		tests      = []struct {
			name       string
			respStatus int
			respBody   TokenResponse
			wantErr    assert.ErrorAssertionFunc
		}{
			{
				name:       "should authenticate a user",
				respStatus: http.StatusOK,
				respBody: TokenResponse{
					AccessToken: "1234",
				},
				wantErr: assert.NoError,
			},
			{
				name:       "should return error if got invalid response",
				respStatus: http.StatusUnauthorized,
				respBody:   TokenResponse{},
				wantErr:    assert.Error,
			},
		}
	)
	defer func() {
		mockServer.Close()
	}()

	ctx, _ := lambda.WithLogger(context.TODO())

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				b, _ := json.Marshal(tt.respBody)
				w.WriteHeader(tt.respStatus)
				_, _ = w.Write(b)
			}))

			c := client{
				baseURL: mockServer.URL,
				svc:     &http.Client{},
			}

			res, err := c.Authenticate(ctx, &AuthRequest{UserID: "123"})

			tt.wantErr(t, err)
			if err == nil {
				assert.Equal(t, tt.respBody.AccessToken, res.JWT)
			}
		})
	}
}

func TestClient_GetOperations(t *testing.T) {
	var (
		mockServer *httptest.Server
		tests      = []struct {
			name       string
			respStatus int
			respBody   OperationList
			wantErr    assert.ErrorAssertionFunc
		}{
			{
				name:       "should return operation list",
				respStatus: http.StatusOK,
				respBody: OperationList{
					{
						ID: 2,
					},
				},
				wantErr: assert.NoError,
			},
			{
				name:       "should return error if fails",
				respStatus: http.StatusBadRequest,
				respBody:   OperationList{},
				wantErr:    assert.Error,
			},
		}
	)
	defer func() {
		mockServer.Close()
	}()

	ctx, _ := lambda.WithLogger(context.TODO())

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mux := http.NewServeMux()

			mux.HandleFunc("/oauth/token", func(w http.ResponseWriter, r *http.Request) {
				auth := TokenResponse{
					AccessToken: "123",
				}
				b, _ := json.Marshal(auth)
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write(b)
			})

			mux.HandleFunc("/api/operations/getOperations.json", func(w http.ResponseWriter, r *http.Request) {
				b, _ := json.Marshal(tt.respBody)
				w.WriteHeader(tt.respStatus)
				_, _ = w.Write(b)
			})

			mockServer = httptest.NewServer(mux)

			c := client{
				baseURL: mockServer.URL,
				svc:     &http.Client{},
			}

			res, err := c.GetOperations(ctx, &GetOperationsRequest{
				UserID:      "1",
				OperationID: "2",
			})

			tt.wantErr(t, err)
			if err == nil {
				assert.Equal(t, int64(2), res.Operation.ID)
			}
		})
	}
}
