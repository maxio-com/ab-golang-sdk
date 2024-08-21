/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// ComponentCostDataRateTier represents a ComponentCostDataRateTier struct.
type ComponentCostDataRateTier struct {
    StartingQuantity     *int           `json:"starting_quantity,omitempty"`
    EndingQuantity       Optional[int]  `json:"ending_quantity"`
    Quantity             *string        `json:"quantity,omitempty"`
    UnitPrice            *string        `json:"unit_price,omitempty"`
    Amount               *string        `json:"amount,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ComponentCostDataRateTier.
// It customizes the JSON marshaling process for ComponentCostDataRateTier objects.
func (c ComponentCostDataRateTier) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ComponentCostDataRateTier object to a map representation for JSON marshaling.
func (c ComponentCostDataRateTier) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
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
    if c.Quantity != nil {
        structMap["quantity"] = c.Quantity
    }
    if c.UnitPrice != nil {
        structMap["unit_price"] = c.UnitPrice
    }
    if c.Amount != nil {
        structMap["amount"] = c.Amount
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ComponentCostDataRateTier.
// It customizes the JSON unmarshaling process for ComponentCostDataRateTier objects.
func (c *ComponentCostDataRateTier) UnmarshalJSON(input []byte) error {
    var temp tempComponentCostDataRateTier
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "starting_quantity", "ending_quantity", "quantity", "unit_price", "amount")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.StartingQuantity = temp.StartingQuantity
    c.EndingQuantity = temp.EndingQuantity
    c.Quantity = temp.Quantity
    c.UnitPrice = temp.UnitPrice
    c.Amount = temp.Amount
    return nil
}

// tempComponentCostDataRateTier is a temporary struct used for validating the fields of ComponentCostDataRateTier.
type tempComponentCostDataRateTier  struct {
    StartingQuantity *int          `json:"starting_quantity,omitempty"`
    EndingQuantity   Optional[int] `json:"ending_quantity"`
    Quantity         *string       `json:"quantity,omitempty"`
    UnitPrice        *string       `json:"unit_price,omitempty"`
    Amount           *string       `json:"amount,omitempty"`
}
