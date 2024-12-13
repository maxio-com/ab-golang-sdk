/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CreateProductPricePointRequest represents a CreateProductPricePointRequest struct.
type CreateProductPricePointRequest struct {
    PricePoint           CreateProductPricePoint `json:"price_point"`
    AdditionalProperties map[string]interface{}  `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreateProductPricePointRequest.
// It customizes the JSON marshaling process for CreateProductPricePointRequest objects.
func (c CreateProductPricePointRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "price_point"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateProductPricePointRequest object to a map representation for JSON marshaling.
func (c CreateProductPricePointRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["price_point"] = c.PricePoint.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateProductPricePointRequest.
// It customizes the JSON unmarshaling process for CreateProductPricePointRequest objects.
func (c *CreateProductPricePointRequest) UnmarshalJSON(input []byte) error {
    var temp tempCreateProductPricePointRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "price_point")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.PricePoint = *temp.PricePoint
    return nil
}

// tempCreateProductPricePointRequest is a temporary struct used for validating the fields of CreateProductPricePointRequest.
type tempCreateProductPricePointRequest  struct {
    PricePoint *CreateProductPricePoint `json:"price_point"`
}

func (c *tempCreateProductPricePointRequest) validate() error {
    var errs []string
    if c.PricePoint == nil {
        errs = append(errs, "required field `price_point` is missing for type `Create Product Price Point Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
