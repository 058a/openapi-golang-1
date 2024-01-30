// Package oapicodegen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package oapicodegen

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// BadRequestResponse defines model for BadRequestResponse.
type BadRequestResponse struct {
	Message string `json:"message"`
}

// InboundItem defines model for InboundItem.
type InboundItem struct {
	Id   openapi_types.UUID `json:"id" validate:"required"`
	Name string             `json:"name" validate:"required,gte=1,lt=100"`
}

// InboundItemId defines model for InboundItemId.
type InboundItemId struct {
	Id openapi_types.UUID `json:"id" validate:"required"`
}

// InboundItemName defines model for InboundItemName.
type InboundItemName struct {
	Name string `json:"name" validate:"required,gte=1,lt=100"`
}

// NewInboundItem defines model for NewInboundItem.
type NewInboundItem = InboundItemName

// BadRequest defines model for BadRequest.
type BadRequest = BadRequestResponse

// Created defines model for Created.
type Created = InboundItem

// OK defines model for OK.
type OK = InboundItem

// PostInboundItemJSONRequestBody defines body for PostInboundItem for application/json ContentType.
type PostInboundItemJSONRequestBody = NewInboundItem

// PutInboundItemJSONRequestBody defines body for PutInboundItem for application/json ContentType.
type PutInboundItemJSONRequestBody = NewInboundItem

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Create Inbound Item
	// (POST /inbound/items)
	PostInboundItem(ctx echo.Context) error
	// Delete Inbound Item
	// (DELETE /inbound/items/{inbounditemId})
	DeleteInboundItem(ctx echo.Context, inbounditemId InboundItemId) error
	// Update Inbound Item
	// (PUT /inbound/items/{inbounditemId})
	PutInboundItem(ctx echo.Context, inbounditemId InboundItemId) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PostInboundItem converts echo context to params.
func (w *ServerInterfaceWrapper) PostInboundItem(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostInboundItem(ctx)
	return err
}

// DeleteInboundItem converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteInboundItem(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "inbounditemId" -------------
	var inbounditemId InboundItemId

	err = runtime.BindStyledParameterWithLocation("simple", false, "inbounditemId", runtime.ParamLocationPath, ctx.Param("inbounditemId"), &inbounditemId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter inbounditemId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteInboundItem(ctx, inbounditemId)
	return err
}

// PutInboundItem converts echo context to params.
func (w *ServerInterfaceWrapper) PutInboundItem(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "inbounditemId" -------------
	var inbounditemId InboundItemId

	err = runtime.BindStyledParameterWithLocation("simple", false, "inbounditemId", runtime.ParamLocationPath, ctx.Param("inbounditemId"), &inbounditemId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter inbounditemId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PutInboundItem(ctx, inbounditemId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/inbound/items", wrapper.PostInboundItem)
	router.DELETE(baseURL+"/inbound/items/:inbounditemId", wrapper.DeleteInboundItem)
	router.PUT(baseURL+"/inbound/items/:inbounditemId", wrapper.PutInboundItem)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8xVy27bOhD9FWHuXdKW3KQbAVk0fQBCACdI0VWQBSuNHRbiI+QoTWDw3wtSimPJSmK3",
	"RpudODyc15kzWkGppdEKFTnIV2DRGa0cxsMpry7xtkFH4VRqRajiJzemFiUnoVX6w2kVbK68QcnD1/8W",
	"F5DDf+mT67S9demTy8suEnjvGVToSitM8Ah5CJw8RvYMPlrkhNXBkijUd92oqiCUY9Efw3kGhSK0itdf",
	"0d6h/WyttsF5H/8ISlpU0sI8g7mmLyHQ9pO5pqS98gzOz/5WZednEIwduk/xmo98BcZqg5ZEOwYSnePL",
	"eEEPBiEHR1aoZfRl8bYRNpBztQZex849ZRJKquvzBeRXO6dfhNa8jJ7jz169g7BFtV2KiLaFtpIT5NA0",
	"ogI2qIrB/URzIyalrnCJaoL3ZPmE+DK6uOO1qDiFB+va/bAToho2Yc7lSGtVZ5X8XshGQj7LMgZSqO70",
	"57mxJeHJjNV0Msuy7URjAiHVQTN/h7JYor+OMYRa6DgwgmqMGomwJOCSDxcFMLhD69qxnE2zaRaUoA0q",
	"bgTkcDTNpkfAwHC6iaWlovWQCkIZLUa3e2lMvMlmPIiObdRTGAq40I42i21bgo5OdfVwMC0Ox7PferIN",
	"RsPGvn2XzZ5zusalG9vpOMtex28scc/g/S5PxrZe3BuNlNw+PNPlgOizlK66o4hy9C1bNRJu8/Yp2l/m",
	"rcX0mTPccomE1sVRFcFXGBpgnbiglwMMOWD7b9awmvz1Fnc7NDbs30Db8evQ9b/jcKSNtdgzMM2IjL6Z",
	"6nUZNfSGuHgbAt5jCPbW7r+am7FZ8N77XwEAAP//SdNNBjwKAAA=",
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
