/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// ComponentPrice represents a ComponentPrice struct.
type ComponentPrice struct {
    Id                   *int                   `json:"id,omitempty"`
    ComponentId          *int                   `json:"component_id,omitempty"`
    StartingQuantity     *int                   `json:"starting_quantity,omitempty"`
    EndingQuantity       Optional[int]          `json:"ending_quantity"`
    UnitPrice            *string                `json:"unit_price,omitempty"`
    PricePointId         *int                   `json:"price_point_id,omitempty"`
    FormattedUnitPrice   *string                `json:"formatted_unit_price,omitempty"`
    SegmentId            Optional[int]          `json:"segment_id"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ComponentPrice,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ComponentPrice) String() string {
    return fmt.Sprintf(
    	"ComponentPrice[Id=%v, ComponentId=%v, StartingQuantity=%v, EndingQuantity=%v, UnitPrice=%v, PricePointId=%v, FormattedUnitPrice=%v, SegmentId=%v, AdditionalProperties=%v]",
    	c.Id, c.ComponentId, c.StartingQuantity, c.EndingQuantity, c.UnitPrice, c.PricePointId, c.FormattedUnitPrice, c.SegmentId, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ComponentPrice.
// It customizes the JSON marshaling process for ComponentPrice objects.
func (c ComponentPrice) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "id", "component_id", "starting_quantity", "ending_quantity", "unit_price", "price_point_id", "formatted_unit_price", "segment_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ComponentPrice object to a map representation for JSON marshaling.
func (c ComponentPrice) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Id != nil {
        structMap["id"] = c.Id
    }
    if c.ComponentId != nil {
        structMap["component_id"] = c.ComponentId
    }
    if c.StartingQuantity != nil {
        structMap["starting_quantity"] = c.StartingQuantity
    }
    if c.EndingQuantity.IsValueSet() {
        if c.EndingQuantity.Value() != nil {
            structMap["ending_quantity"] = c.EndingQuantity.Value()
        } else {
            structMap["ending_quantity"] = nil
        }
    }
    if c.UnitPrice != nil {
        structMap["unit_price"] = c.UnitPrice
    }
    if c.PricePointId != nil {
        structMap["price_point_id"] = c.PricePointId
    }
    if c.FormattedUnitPrice != nil {
        structMap["formatted_unit_price"] = c.FormattedUnitPrice
    }
    if c.SegmentId.IsValueSet() {
        if c.SegmentId.Value() != nil {
            structMap["segment_id"] = c.SegmentId.Value()
        } else {
            structMap["segment_id"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ComponentPrice.
// It customizes the JSON unmarshaling process for ComponentPrice objects.
func (c *ComponentPrice) UnmarshalJSON(input []byte) error {
    var temp tempComponentPrice
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "component_id", "starting_quantity", "ending_quantity", "unit_price", "price_point_id", "formatted_unit_price", "segment_id")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Id = temp.Id
    c.ComponentId = temp.ComponentId
    c.StartingQuantity = temp.StartingQuantity
    c.EndingQuantity = temp.EndingQuantity
    c.UnitPrice = temp.UnitPrice
    c.PricePointId = temp.PricePointId
    c.FormattedUnitPrice = temp.FormattedUnitPrice
    c.SegmentId = temp.SegmentId
    return nil
}

// tempComponentPrice is a temporary struct used for validating the fields of ComponentPrice.
type tempComponentPrice  struct {
    Id                 *int          `json:"id,omitempty"`
    ComponentId        *int          `json:"component_id,omitempty"`
    StartingQuantity   *int          `json:"starting_quantity,omitempty"`
    EndingQuantity     Optional[int] `json:"ending_quantity"`
    UnitPrice          *string       `json:"unit_price,omitempty"`
    PricePointId       *int          `json:"price_point_id,omitempty"`
    FormattedUnitPrice *string       `json:"formatted_unit_price,omitempty"`
    SegmentId          Optional[int] `json:"segment_id"`
}
