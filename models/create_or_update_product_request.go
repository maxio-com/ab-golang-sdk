// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CreateOrUpdateProductRequest represents a CreateOrUpdateProductRequest struct.
type CreateOrUpdateProductRequest struct {
    Product              CreateOrUpdateProduct  `json:"product"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CreateOrUpdateProductRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateOrUpdateProductRequest) String() string {
    return fmt.Sprintf(
    	"CreateOrUpdateProductRequest[Product=%v, AdditionalProperties=%v]",
    	c.Product, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreateOrUpdateProductRequest.
// It customizes the JSON marshaling process for CreateOrUpdateProductRequest objects.
func (c CreateOrUpdateProductRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "product"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateOrUpdateProductRequest object to a map representation for JSON marshaling.
func (c CreateOrUpdateProductRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["product"] = c.Product.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateOrUpdateProductRequest.
// It customizes the JSON unmarshaling process for CreateOrUpdateProductRequest objects.
func (c *CreateOrUpdateProductRequest) UnmarshalJSON(input []byte) error {
    var temp tempCreateOrUpdateProductRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "product")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Product = *temp.Product
    return nil
}

// tempCreateOrUpdateProductRequest is a temporary struct used for validating the fields of CreateOrUpdateProductRequest.
type tempCreateOrUpdateProductRequest  struct {
    Product *CreateOrUpdateProduct `json:"product"`
}

func (c *tempCreateOrUpdateProductRequest) validate() error {
    var errs []string
    if c.Product == nil {
        errs = append(errs, "required field `product` is missing for type `Create or Update Product Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
