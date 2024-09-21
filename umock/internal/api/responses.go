package api

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	"github.com/go-faker/faker/v4"
	"math/rand"
	"net/http"
	"time"
)

var _statusCodeMap = map[string]int{
	"200": http.StatusOK,
	"201": http.StatusCreated,
	"400": http.StatusBadRequest,
	"401": http.StatusUnauthorized,
	"403": http.StatusForbidden,
	"500": http.StatusInternalServerError,
}

func generateMockResponse(operation *openapi3.Operation) (int, any) {
	for statusCode, ref := range operation.Responses.Map() {
		response := ref.Value
		for _, mediaType := range response.Content {
			// TODO: check encoding
			// only accepts application/json for now
			if mediaType.Schema != nil {
				res := generateResponse(mediaType.Schema.Value)
				status, ok := _statusCodeMap[statusCode]
				if !ok {
					status = http.StatusOK
				}
				return status, res
			}
		}
	}

	return http.StatusOK, nil
}

func generateResponse(schema *openapi3.Schema) gin.H {
	r := make(gin.H)

	for k, v := range schema.Properties {
		if v.Value.Example != nil {
			r[k] = v.Value.Example
		} else {
			r[k] = generate(v.Value)
		}
	}

	return r
}

func generate(schema *openapi3.Schema) any {
	if schema.Type.Is("object") {
		result := make(map[string]any)
		for propName, propSchemaRef := range schema.Properties {
			propSchema := propSchemaRef.Value
			if propSchema == nil {
				continue
			}
			if propSchema.Example != nil {
				result[propName] = propSchema.Example
			} else {
				result[propName] = generate(propSchema)
			}
		}
		return result
	}

	if schema.Type.Is("array") {
		if schema.Items != nil {
			itemSchema := schema.Items.Value
			return []any{generate(itemSchema)}
		}
		return []any{}
	}

	if schema.Type.Is("string") {
		return generateStringByFormat(schema.Format)
	}

	if schema.Type.Is("integer") {
		return rand.Intn(100)
	}

	if schema.Type.Is("number") {
		return rand.Float64() * 100
	}

	if schema.Type.Is("boolean") {
		return rand.Intn(2) == 1
	}

	return nil
}

func generateStringByFormat(format string) string {
	switch format {
	case "date":
		return time.Now().Format("2000-01-02")
	case "date-time":
		return time.Now().Format(time.RFC3339)
	case "email":
		return faker.Email()
	case "uuid":
		return faker.UUIDHyphenated()
	case "jwt":
		return faker.Jwt()
	default:
		return faker.Word()
	}
}
