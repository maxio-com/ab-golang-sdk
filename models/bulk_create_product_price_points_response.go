/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// BulkCreateProductPricePointsResponse represents a BulkCreateProductPricePointsResponse struct.
type BulkCreateProductPricePointsResponse struct {
    PricePoints          []ProductPricePoint `json:"price_points,omitempty"`
    AdditionalProperties map[string]any      `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for BulkCreateProductPricePointsResponse.
// It customizes the JSON marshaling process for BulkCreateProductPricePointsResponse objects.
func (b BulkCreateProductPricePointsResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(b.toMap())
}

// toMap converts the BulkCreateProductPricePointsResponse object to a map representation for JSON marshaling.
func (b BulkCreateProductPricePointsResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, b.AdditionalProperties)
    if b.PricePoints != nil {
        structMap["price_points"] = b.PricePoints
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BulkCreateProductPricePointsResponse.
// It customizes the JSON unmarshaling process for BulkCreateProductPricePointsResponse objects.
func (b *BulkCreateProductPricePointsResponse) UnmarshalJSON(input []byte) error {
    var temp tempBulkCreateProductPricePointsResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "price_points")
    if err != nil {
    	return err
    }
    
    b.AdditionalProperties = additionalProperties
    b.PricePoints = temp.PricePoints
    return nil
}

// tempBulkCreateProductPricePointsResponse is a temporary struct used for validating the fields of BulkCreateProductPricePointsResponse.
type tempBulkCreateProductPricePointsResponse  struct {
    PricePoints []ProductPricePoint `json:"price_points,omitempty"`
}
