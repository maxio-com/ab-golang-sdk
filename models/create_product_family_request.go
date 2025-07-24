// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CreateProductFamilyRequest represents a CreateProductFamilyRequest struct.
type CreateProductFamilyRequest struct {
    ProductFamily        CreateProductFamily    `json:"product_family"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CreateProductFamilyRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateProductFamilyRequest) String() string {
    return fmt.Sprintf(
    	"CreateProductFamilyRequest[ProductFamily=%v, AdditionalProperties=%v]",
    	c.ProductFamily, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreateProductFamilyRequest.
// It customizes the JSON marshaling process for CreateProductFamilyRequest objects.
func (c CreateProductFamilyRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "product_family"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateProductFamilyRequest object to a map representation for JSON marshaling.
func (c CreateProductFamilyRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["product_family"] = c.ProductFamily.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateProductFamilyRequest.
// It customizes the JSON unmarshaling process for CreateProductFamilyRequest objects.
func (c *CreateProductFamilyRequest) UnmarshalJSON(input []byte) error {
    var temp tempCreateProductFamilyRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "product_family")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.ProductFamily = *temp.ProductFamily
    return nil
}

// tempCreateProductFamilyRequest is a temporary struct used for validating the fields of CreateProductFamilyRequest.
type tempCreateProductFamilyRequest  struct {
    ProductFamily *CreateProductFamily `json:"product_family"`
}

func (c *tempCreateProductFamilyRequest) validate() error {
    var errs []string
    if c.ProductFamily == nil {
        errs = append(errs, "required field `product_family` is missing for type `Create Product Family Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
