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

// UpdateProductPricePointRequest represents a UpdateProductPricePointRequest struct.
type UpdateProductPricePointRequest struct {
    PricePoint           UpdateProductPricePoint `json:"price_point"`
    AdditionalProperties map[string]any          `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateProductPricePointRequest.
// It customizes the JSON marshaling process for UpdateProductPricePointRequest objects.
func (u UpdateProductPricePointRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateProductPricePointRequest object to a map representation for JSON marshaling.
func (u UpdateProductPricePointRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    structMap["price_point"] = u.PricePoint.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateProductPricePointRequest.
// It customizes the JSON unmarshaling process for UpdateProductPricePointRequest objects.
func (u *UpdateProductPricePointRequest) UnmarshalJSON(input []byte) error {
    var temp tempUpdateProductPricePointRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "price_point")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.PricePoint = *temp.PricePoint
    return nil
}

// tempUpdateProductPricePointRequest is a temporary struct used for validating the fields of UpdateProductPricePointRequest.
type tempUpdateProductPricePointRequest  struct {
    PricePoint *UpdateProductPricePoint `json:"price_point"`
}

func (u *tempUpdateProductPricePointRequest) validate() error {
    var errs []string
    if u.PricePoint == nil {
        errs = append(errs, "required field `price_point` is missing for type `Update Product Price Point Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
