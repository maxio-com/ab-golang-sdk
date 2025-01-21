/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// ComponentPricePointsResponse represents a ComponentPricePointsResponse struct.
type ComponentPricePointsResponse struct {
    PricePoints          []ComponentPricePoint  `json:"price_points,omitempty"`
    Meta                 *ListPublicKeysMeta    `json:"meta,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ComponentPricePointsResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ComponentPricePointsResponse) String() string {
    return fmt.Sprintf(
    	"ComponentPricePointsResponse[PricePoints=%v, Meta=%v, AdditionalProperties=%v]",
    	c.PricePoints, c.Meta, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ComponentPricePointsResponse.
// It customizes the JSON marshaling process for ComponentPricePointsResponse objects.
func (c ComponentPricePointsResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "price_points", "meta"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ComponentPricePointsResponse object to a map representation for JSON marshaling.
func (c ComponentPricePointsResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.PricePoints != nil {
        structMap["price_points"] = c.PricePoints
    }
    if c.Meta != nil {
        structMap["meta"] = c.Meta.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ComponentPricePointsResponse.
// It customizes the JSON unmarshaling process for ComponentPricePointsResponse objects.
func (c *ComponentPricePointsResponse) UnmarshalJSON(input []byte) error {
    var temp tempComponentPricePointsResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "price_points", "meta")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.PricePoints = temp.PricePoints
    c.Meta = temp.Meta
    return nil
}

// tempComponentPricePointsResponse is a temporary struct used for validating the fields of ComponentPricePointsResponse.
type tempComponentPricePointsResponse  struct {
    PricePoints []ComponentPricePoint `json:"price_points,omitempty"`
    Meta        *ListPublicKeysMeta   `json:"meta,omitempty"`
}
