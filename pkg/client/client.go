package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"io"
	"net/http"
	"sebsegura/sample-lambda/pkg/logger"
	"strconv"
	"strings"
	"time"
)

type AuthRequest struct {
	UserID string
}

type AuthResponse struct {
	JWT string
}

type Client interface {
	Authenticate(ctx context.Context, r *AuthRequest) (*AuthResponse, error)
	GetOperations(ctx context.Context, r *GetOperationsRequest) (*GetOperationsResponse, error)
}

type client struct {
	baseURL string
	svc     *http.Client
}

const (
	_url = "https://geuwen2epg3wyaugabjexbnudy0hsmgi.lambda-url.us-east-1.on.aws"
)

func New() Client {
	return &client{
		baseURL: _url,
		svc: &http.Client{
			Timeout: time.Duration(10) * time.Second,
		},
	}
}

func (c *client) Authenticate(ctx context.Context, r *AuthRequest) (*AuthResponse, error) {
	log := logger.Logger(ctx)
	creds := Credentials{
		ClientID:     "123",
		ClientSecret: "123",
		GrantType:    "impersonation",
		User:         r.UserID,
	}

	body, err := json.Marshal(creds)
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s/oauth/token", c.baseURL)
	log.With(zap.String("uri", uri)).Debug("trying to authenticate...")
	req, err := http.NewRequest(http.MethodPost, uri, strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var token TokenResponse
	if err := c.parseResponse(res, &token); err != nil {
		return nil, err
	}

	return &AuthResponse{
		JWT: token.AccessToken,
	}, nil
}

func (c *client) GetOperations(ctx context.Context, in *GetOperationsRequest) (*GetOperationsResponse, error) {
	log := logger.Logger(ctx)
	authRes, err := c.Authenticate(ctx, &AuthRequest{
		UserID: in.UserID,
	})
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s/api/operations/getOperations.json", c.baseURL)
	log.With(zap.String("uri", uri)).Debug("trying to get operation list...")
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authRes.JWT))

	res, err := c.doRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ops OperationList
	if err := c.parseResponse(res, &ops); err != nil {
		return nil, err
	}

	opID, _ := strconv.ParseInt(in.OperationID, 10, 64)

	for _, op := range ops {
		if op.ID == opID {
			log.With(zap.Any("operation", op)).Debug("found operation")
			return &GetOperationsResponse{Operation: op}, nil
		}
	}

	return nil, errors.New("cannot find the operation")
}

func (c *client) doRequest(ctx context.Context, req *http.Request) (*http.Response, error) {
	req = req.WithContext(ctx)

	res, err := c.svc.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *client) parseResponse(res *http.Response, target interface{}) error {
	defer func() {
		_ = res.Body.Close()
	}()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, target)
}
