package openapi

import (
	"context"
	"github.com/getkin/kin-openapi/openapi3"
)

// LoadSpec load and validate OpenAPI3 specification
func LoadSpec(ctx context.Context, path string) (*openapi3.T, error) {
	loader := &openapi3.Loader{Context: ctx, IsExternalRefsAllowed: true}
	doc, err := loader.LoadFromFile(path)
	if err != nil {
		return nil, err
	}

	if err = doc.Validate(loader.Context); err != nil {
		return nil, err
	}

	return doc, nil
}
