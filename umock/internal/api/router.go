package api

import (
	"context"
	"github.com/getkin/kin-openapi/routers/legacy"
	"github.com/gin-gonic/gin"
	"sebsegura/umock/internal/openapi"
)

// SetupRouter create and setup new Gin router
func SetupRouter(ctx context.Context, specPath string) (*gin.Engine, error) {
	doc, err := openapi.LoadSpec(ctx, specPath)
	if err != nil {
		return nil, err
	}

	apiRouter, err := legacy.NewRouter(doc)
	if err != nil {
		return nil, err
	}

	r := gin.Default()
	r.Use(LogMiddleware())

	if err = generateRoutes(r, doc, apiRouter); err != nil {
		return nil, err
	}

	return r, nil
}
