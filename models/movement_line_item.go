package models

import (
    "encoding/json"
)

// MovementLineItem represents a MovementLineItem struct.
type MovementLineItem struct {
    ProductId            *int           `json:"product_id,omitempty"`
    // For Product (or "baseline") line items, this field will have a value of `0`.
    ComponentId          *int           `json:"component_id,omitempty"`
    PricePointId         *int           `json:"price_point_id,omitempty"`
    Name                 *string        `json:"name,omitempty"`
    Mrr                  *int           `json:"mrr,omitempty"`
    MrrMovements         []MRRMovement  `json:"mrr_movements,omitempty"`
    Quantity             *int           `json:"quantity,omitempty"`
    PrevQuantity         *int           `json:"prev_quantity,omitempty"`
    // When `true`, the line item's MRR value will contribute to the `plan` breakout. When `false`, the line item contributes to the `usage` breakout.
    Recurring            *bool          `json:"recurring,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MovementLineItem.
// It customizes the JSON marshaling process for MovementLineItem objects.
func (m MovementLineItem) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MovementLineItem object to a map representation for JSON marshaling.
func (m MovementLineItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    if m.ProductId != nil {
        structMap["product_id"] = m.ProductId
    }
    if m.ComponentId != nil {
        structMap["component_id"] = m.ComponentId
    }
    if m.PricePointId != nil {
        structMap["price_point_id"] = m.PricePointId
    }
    if m.Name != nil {
        structMap["name"] = m.Name
    }
    if m.Mrr != nil {
        structMap["mrr"] = m.Mrr
    }
    if m.MrrMovements != nil {
        structMap["mrr_movements"] = m.MrrMovements
    }
    if m.Quantity != nil {
        structMap["quantity"] = m.Quantity
    }
    if m.PrevQuantity != nil {
        structMap["prev_quantity"] = m.PrevQuantity
    }
    if m.Recurring != nil {
        structMap["recurring"] = m.Recurring
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MovementLineItem.
// It customizes the JSON unmarshaling process for MovementLineItem objects.
func (m *MovementLineItem) UnmarshalJSON(input []byte) error {
    var temp movementLineItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "product_id", "component_id", "price_point_id", "name", "mrr", "mrr_movements", "quantity", "prev_quantity", "recurring")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.ProductId = temp.ProductId
    m.ComponentId = temp.ComponentId
    m.PricePointId = temp.PricePointId
    m.Name = temp.Name
    m.Mrr = temp.Mrr
    m.MrrMovements = temp.MrrMovements
    m.Quantity = temp.Quantity
    m.PrevQuantity = temp.PrevQuantity
    m.Recurring = temp.Recurring
    return nil
}

// TODO
type movementLineItem  struct {
    ProductId    *int          `json:"product_id,omitempty"`
    ComponentId  *int          `json:"component_id,omitempty"`
    PricePointId *int          `json:"price_point_id,omitempty"`
    Name         *string       `json:"name,omitempty"`
    Mrr          *int          `json:"mrr,omitempty"`
    MrrMovements []MRRMovement `json:"mrr_movements,omitempty"`
    Quantity     *int          `json:"quantity,omitempty"`
    PrevQuantity *int          `json:"prev_quantity,omitempty"`
    Recurring    *bool         `json:"recurring,omitempty"`
}
