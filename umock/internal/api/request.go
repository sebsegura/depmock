package api

import (
	"bytes"
	"context"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func buildHTTPRequest(c *gin.Context) (*http.Request, error) {
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return nil, err
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	req, err := http.NewRequest(c.Request.Method, c.Request.RequestURI, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}
	req.Header = c.Request.Header.Clone()
	req.RemoteAddr = c.Request.RemoteAddr

	return req, nil
}

func validateRequest(c *gin.Context, apiRouter routers.Router) error {
	req, err := buildHTTPRequest(c)
	if err != nil {
		return err
	}

	route, _, err := apiRouter.FindRoute(req)
	if err != nil {
		return err
	}

	requestValidationInput := &openapi3filter.RequestValidationInput{
		Request:     req,
		Route:       route,
		PathParams:  getPathParams(c),
		QueryParams: c.Request.URL.Query(),
		Options: &openapi3filter.Options{
			AuthenticationFunc: func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
				return nil
			},
		},
	}

	if err = openapi3filter.ValidateRequest(c, requestValidationInput); err != nil {
		return fmt.Errorf("request validation failed: %v", err)
	}

	return nil
}

func getPathParams(c *gin.Context) map[string]string {
	params := make(map[string]string)
	for _, param := range c.Params {
		params[param.Key] = param.Value
	}
	return params
}
