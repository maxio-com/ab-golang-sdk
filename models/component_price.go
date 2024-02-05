package models

import (
    "encoding/json"
)

// ComponentPrice represents a ComponentPrice struct.
type ComponentPrice struct {
    Id                 *int          `json:"id,omitempty"`
    ComponentId        *int          `json:"component_id,omitempty"`
    StartingQuantity   *int          `json:"starting_quantity,omitempty"`
    EndingQuantity     Optional[int] `json:"ending_quantity"`
    UnitPrice          *string       `json:"unit_price,omitempty"`
    PricePointId       *int          `json:"price_point_id,omitempty"`
    FormattedUnitPrice *string       `json:"formatted_unit_price,omitempty"`
    SegmentId          Optional[int] `json:"segment_id"`
}

// MarshalJSON implements the json.Marshaler interface for ComponentPrice.
// It customizes the JSON marshaling process for ComponentPrice objects.
func (c *ComponentPrice) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ComponentPrice object to a map representation for JSON marshaling.
func (c *ComponentPrice) toMap() map[string]any {
    structMap := make(map[string]any)
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
        structMap["ending_quantity"] = c.EndingQuantity.Value()
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
        structMap["segment_id"] = c.SegmentId.Value()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ComponentPrice.
// It customizes the JSON unmarshaling process for ComponentPrice objects.
func (c *ComponentPrice) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id                 *int          `json:"id,omitempty"`
        ComponentId        *int          `json:"component_id,omitempty"`
        StartingQuantity   *int          `json:"starting_quantity,omitempty"`
        EndingQuantity     Optional[int] `json:"ending_quantity"`
        UnitPrice          *string       `json:"unit_price,omitempty"`
        PricePointId       *int          `json:"price_point_id,omitempty"`
        FormattedUnitPrice *string       `json:"formatted_unit_price,omitempty"`
        SegmentId          Optional[int] `json:"segment_id"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
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
