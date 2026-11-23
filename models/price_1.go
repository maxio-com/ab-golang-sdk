// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// Price1 represents a Price1 struct.
type Price1 struct {
    StartingQuantity     *int                   `json:"starting_quantity,omitempty"`
    EndingQuantity       Optional[int]          `json:"ending_quantity"`
    UnitPrice            *string                `json:"unit_price,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Price1,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p Price1) String() string {
    return fmt.Sprintf(
    	"Price1[StartingQuantity=%v, EndingQuantity=%v, UnitPrice=%v, AdditionalProperties=%v]",
    	p.StartingQuantity, p.EndingQuantity, p.UnitPrice, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Price1.
// It customizes the JSON marshaling process for Price1 objects.
func (p Price1) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(p.AdditionalProperties,
        "starting_quantity", "ending_quantity", "unit_price"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(p.toMap())
}

// toMap converts the Price1 object to a map representation for JSON marshaling.
func (p Price1) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, p.AdditionalProperties)
    if p.StartingQuantity != nil {
        structMap["starting_quantity"] = p.StartingQuantity
    }
    if p.EndingQuantity.IsValueSet() {
        if p.EndingQuantity.Value() != nil {
            structMap["ending_quantity"] = p.EndingQuantity.Value()
        } else {
            structMap["ending_quantity"] = nil
        }
    }
    if p.UnitPrice != nil {
        structMap["unit_price"] = p.UnitPrice
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Price1.
// It customizes the JSON unmarshaling process for Price1 objects.
func (p *Price1) UnmarshalJSON(input []byte) error {
    var temp tempPrice1
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "starting_quantity", "ending_quantity", "unit_price")
    if err != nil {
    	return err
    }
    p.AdditionalProperties = additionalProperties
    
    p.StartingQuantity = temp.StartingQuantity
    p.EndingQuantity = temp.EndingQuantity
    p.UnitPrice = temp.UnitPrice
    return nil
}

// tempPrice1 is a temporary struct used for validating the fields of Price1.
type tempPrice1  struct {
    StartingQuantity *int          `json:"starting_quantity,omitempty"`
    EndingQuantity   Optional[int] `json:"ending_quantity"`
    UnitPrice        *string       `json:"unit_price,omitempty"`
}
