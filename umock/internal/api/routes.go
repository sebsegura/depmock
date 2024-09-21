package api

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/routers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func generateRoutes(router *gin.Engine, doc *openapi3.T, apiRouter routers.Router) error {
	for path, item := range doc.Paths.Map() {
		if err := generateRoute(router, path, item, apiRouter); err != nil {
			return err
		}
	}

	return nil
}

func generateRoute(router *gin.Engine, path string, item *openapi3.PathItem, apiRouter routers.Router) error {
	methods := map[string]*openapi3.Operation{
		http.MethodGet:    item.Get,
		http.MethodPost:   item.Post,
		http.MethodPut:    item.Put,
		http.MethodDelete: item.Delete,
		http.MethodPatch:  item.Patch,
	}

	for method, operation := range methods {
		if operation == nil {
			continue
		}

		h := createHandler(operation, apiRouter)

		if operation.Security != nil {
			router.Handle(method, path, AuthMiddleware(), h)
		} else {
			router.Handle(method, path, h)
		}
	}

	return nil
}
