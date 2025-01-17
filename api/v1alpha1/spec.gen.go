// Package v1alpha1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.15.0 DO NOT EDIT.
package v1alpha1

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
