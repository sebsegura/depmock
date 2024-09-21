package api

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/routers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func createHandler(operation *openapi3.Operation, apiRouter routers.Router) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := validateRequest(c, apiRouter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		statusCode, response := generateMockResponse(operation)

		c.JSON(statusCode, response)

		if operation.Callbacks != nil {
			go HandleCallbacks(operation.Callbacks, c)
		}
	}
}
