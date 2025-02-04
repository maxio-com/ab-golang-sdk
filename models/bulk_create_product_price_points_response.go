/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// BulkCreateProductPricePointsResponse represents a BulkCreateProductPricePointsResponse struct.
type BulkCreateProductPricePointsResponse struct {
    PricePoints          []ProductPricePoint    `json:"price_points,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for BulkCreateProductPricePointsResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (b BulkCreateProductPricePointsResponse) String() string {
    return fmt.Sprintf(
    	"BulkCreateProductPricePointsResponse[PricePoints=%v, AdditionalProperties=%v]",
    	b.PricePoints, b.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for BulkCreateProductPricePointsResponse.
// It customizes the JSON marshaling process for BulkCreateProductPricePointsResponse objects.
func (b BulkCreateProductPricePointsResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(b.AdditionalProperties,
        "price_points"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(b.toMap())
}

// toMap converts the BulkCreateProductPricePointsResponse object to a map representation for JSON marshaling.
func (b BulkCreateProductPricePointsResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, b.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "price_points")
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
