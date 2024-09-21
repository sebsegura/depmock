package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Bancar/goala/ulog"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func HandleCallbacks(callbacks openapi3.Callbacks, c *gin.Context) {
	for _, callbackRef := range callbacks {
		callback := callbackRef.Value
		for expr, pathItem := range callback.Map() {
			var (
				callbackURL string
				err         error
			)
			if isExpression(expr) {
				callbackURL, err = resolveExpression(expr, c)
				if err != nil {
					ulog.Error("error resolving callback expression")
					continue
				}
			} else {
				callbackURL = expr
			}

			op := pathItem.Post
			if op == nil {
				continue
			}

			callbackBody := generateCallbackBody(op)

			go sendCallbackRequest(callbackURL, callbackBody)
		}
	}
}

func isExpression(s string) bool {
	return strings.Contains(s, "{$") && strings.Contains(s, "}")
}

func resolveExpression(expression string, c *gin.Context) (string, error) {
	if strings.HasPrefix(expression, "{$") && strings.HasSuffix(expression, "}") {
		exprContent := expression[2 : len(expression)-1]
		parts := strings.SplitN(exprContent, "#", 2)
		if len(parts) != 2 {
			return "", fmt.Errorf("invalid expression format")
		}
		source, pointer := parts[0], parts[1]

		switch source {
		case "request.body":
			var requestBody interface{}
			if err := c.ShouldBindJSON(&requestBody); err != nil {
				return "", fmt.Errorf("error parsing request body: %v", err)
			}
			value, err := extractValueFromJSONPointer(requestBody, pointer)
			if err != nil {
				return "", fmt.Errorf("error extracting value from request body: %v", err)
			}
			if url, ok := value.(string); ok {
				return url, nil
			}
			return "", fmt.Errorf("value at pointer '%s' is not a string", pointer)
		case "request.query":
			queryParams := c.Request.URL.Query()
			paramName := pointer[1:] // Remover el '/' inicial
			values, exists := queryParams[paramName]
			if !exists || len(values) == 0 {
				return "", fmt.Errorf("query parameter '%s' not found", paramName)
			}
			return values[0], nil
		default:
			return "", fmt.Errorf("unsupported expression source: %s", source)
		}
	}
	return "", fmt.Errorf("unsupported expression format")
}

func extractValueFromJSONPointer(data interface{}, pointer string) (interface{}, error) {
	tokens := strings.Split(pointer, "/")[1:]
	current := data
	for _, token := range tokens {
		if m, ok := current.(map[string]interface{}); ok {
			current = m[token]
		} else {
			return nil, fmt.Errorf("cannot access property '%s' on non-object", token)
		}
	}
	return current, nil
}

func generateCallbackBody(op *openapi3.Operation) any {
	body := op.RequestBody.Value
	for _, mediaType := range body.Content {
		if mediaType.Schema != nil {
			return generate(mediaType.Schema.Value)
		}
	}
	return nil
}

func sendCallbackRequest(url string, body any) {
	jsonData, err := json.Marshal(body)
	if err != nil {
		ulog.Error("error serializing callback body")
		return
	}

	_, err = http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		ulog.Error("error sending callback request")
		return
	}
	ulog.Debug("callback request successfully sent")
}
