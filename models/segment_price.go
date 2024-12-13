/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// SegmentPrice represents a SegmentPrice struct.
type SegmentPrice struct {
    Id                   *int                   `json:"id,omitempty"`
    ComponentId          *int                   `json:"component_id,omitempty"`
    StartingQuantity     *int                   `json:"starting_quantity,omitempty"`
    EndingQuantity       Optional[int]          `json:"ending_quantity"`
    UnitPrice            *string                `json:"unit_price,omitempty"`
    PricePointId         *int                   `json:"price_point_id,omitempty"`
    FormattedUnitPrice   *string                `json:"formatted_unit_price,omitempty"`
    SegmentId            *int                   `json:"segment_id,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SegmentPrice.
// It customizes the JSON marshaling process for SegmentPrice objects.
func (s SegmentPrice) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "id", "component_id", "starting_quantity", "ending_quantity", "unit_price", "price_point_id", "formatted_unit_price", "segment_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SegmentPrice object to a map representation for JSON marshaling.
func (s SegmentPrice) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Id != nil {
        structMap["id"] = s.Id
    }
    if s.ComponentId != nil {
        structMap["component_id"] = s.ComponentId
    }
    if s.StartingQuantity != nil {
        structMap["starting_quantity"] = s.StartingQuantity
    }
    if s.EndingQuantity.IsValueSet() {
        if s.EndingQuantity.Value() != nil {
            structMap["ending_quantity"] = s.EndingQuantity.Value()
        } else {
            structMap["ending_quantity"] = nil
        }
    }
    if s.UnitPrice != nil {
        structMap["unit_price"] = s.UnitPrice
    }
    if s.PricePointId != nil {
        structMap["price_point_id"] = s.PricePointId
    }
    if s.FormattedUnitPrice != nil {
        structMap["formatted_unit_price"] = s.FormattedUnitPrice
    }
    if s.SegmentId != nil {
        structMap["segment_id"] = s.SegmentId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SegmentPrice.
// It customizes the JSON unmarshaling process for SegmentPrice objects.
func (s *SegmentPrice) UnmarshalJSON(input []byte) error {
    var temp tempSegmentPrice
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "component_id", "starting_quantity", "ending_quantity", "unit_price", "price_point_id", "formatted_unit_price", "segment_id")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Id = temp.Id
    s.ComponentId = temp.ComponentId
    s.StartingQuantity = temp.StartingQuantity
    s.EndingQuantity = temp.EndingQuantity
    s.UnitPrice = temp.UnitPrice
    s.PricePointId = temp.PricePointId
    s.FormattedUnitPrice = temp.FormattedUnitPrice
    s.SegmentId = temp.SegmentId
    return nil
}

// tempSegmentPrice is a temporary struct used for validating the fields of SegmentPrice.
type tempSegmentPrice  struct {
    Id                 *int          `json:"id,omitempty"`
    ComponentId        *int          `json:"component_id,omitempty"`
    StartingQuantity   *int          `json:"starting_quantity,omitempty"`
    EndingQuantity     Optional[int] `json:"ending_quantity"`
    UnitPrice          *string       `json:"unit_price,omitempty"`
    PricePointId       *int          `json:"price_point_id,omitempty"`
    FormattedUnitPrice *string       `json:"formatted_unit_price,omitempty"`
    SegmentId          *int          `json:"segment_id,omitempty"`
}
