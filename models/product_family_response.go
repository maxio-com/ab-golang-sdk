/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// ProductFamilyResponse represents a ProductFamilyResponse struct.
type ProductFamilyResponse struct {
    ProductFamily        *ProductFamily `json:"product_family,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ProductFamilyResponse.
// It customizes the JSON marshaling process for ProductFamilyResponse objects.
func (p ProductFamilyResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the ProductFamilyResponse object to a map representation for JSON marshaling.
func (p ProductFamilyResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
    if p.ProductFamily != nil {
        structMap["product_family"] = p.ProductFamily.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ProductFamilyResponse.
// It customizes the JSON unmarshaling process for ProductFamilyResponse objects.
func (p *ProductFamilyResponse) UnmarshalJSON(input []byte) error {
    var temp tempProductFamilyResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "product_family")
    if err != nil {
    	return err
    }
    
    p.AdditionalProperties = additionalProperties
    p.ProductFamily = temp.ProductFamily
    return nil
}

// tempProductFamilyResponse is a temporary struct used for validating the fields of ProductFamilyResponse.
type tempProductFamilyResponse  struct {
    ProductFamily *ProductFamily `json:"product_family,omitempty"`
}
