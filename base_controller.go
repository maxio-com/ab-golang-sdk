// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package advancedbilling

import (
    "github.com/apimatic/go-core-runtime/https"
)

// callBuilderFactory is an interface that defines a method to get a CallBuilderFactory.
// It allows objects to get a reference to a CallBuilderFactory for creating API call.
type callBuilderFactory interface {
    GetCallBuilder() https.CallBuilderFactory
}

// baseController represents a controller used as a base for other controllers.
// It encapsulates common functionality required by controllers for making API call.
type baseController struct {
    callBuilder    callBuilderFactory
    prepareRequest https.CallBuilderFactory
}

// NewBaseController creates a new instance of baseController.
// It takes a callBuilderFactory as a parameter and returns a pointer to the baseController.
func NewBaseController(cb callBuilderFactory) *baseController {
    baseController := baseController{callBuilder: cb}
    baseController.prepareRequest = baseController.callBuilder.GetCallBuilder()
    return &baseController
}
