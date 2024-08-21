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

// ProductPricePointResponse represents a ProductPricePointResponse struct.
type ProductPricePointResponse struct {
    PricePoint           ProductPricePoint `json:"price_point"`
    AdditionalProperties map[string]any    `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ProductPricePointResponse.
// It customizes the JSON marshaling process for ProductPricePointResponse objects.
func (p ProductPricePointResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the ProductPricePointResponse object to a map representation for JSON marshaling.
func (p ProductPricePointResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
    structMap["price_point"] = p.PricePoint.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ProductPricePointResponse.
// It customizes the JSON unmarshaling process for ProductPricePointResponse objects.
func (p *ProductPricePointResponse) UnmarshalJSON(input []byte) error {
    var temp tempProductPricePointResponse
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
    
    p.AdditionalProperties = additionalProperties
    p.PricePoint = *temp.PricePoint
    return nil
}

// tempProductPricePointResponse is a temporary struct used for validating the fields of ProductPricePointResponse.
type tempProductPricePointResponse  struct {
    PricePoint *ProductPricePoint `json:"price_point"`
}

func (p *tempProductPricePointResponse) validate() error {
    var errs []string
    if p.PricePoint == nil {
        errs = append(errs, "required field `price_point` is missing for type `Product Price Point Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
