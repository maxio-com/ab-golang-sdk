package models

import (
    "encoding/json"
)

// Price represents a Price struct.
type Price struct {
    StartingQuantity interface{}           `json:"starting_quantity"`
    EndingQuantity   Optional[interface{}] `json:"ending_quantity"`
    // The price can contain up to 8 decimal places. i.e. 1.00 or 0.0012 or 0.00000065
    UnitPrice        interface{}           `json:"unit_price"`
}

// MarshalJSON implements the json.Marshaler interface for Price.
// It customizes the JSON marshaling process for Price objects.
func (p *Price) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the Price object to a map representation for JSON marshaling.
func (p *Price) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["starting_quantity"] = p.StartingQuantity
    if p.EndingQuantity.IsValueSet() {
        structMap["ending_quantity"] = p.EndingQuantity.Value()
    }
    structMap["unit_price"] = p.UnitPrice
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Price.
// It customizes the JSON unmarshaling process for Price objects.
func (p *Price) UnmarshalJSON(input []byte) error {
    temp := &struct {
        StartingQuantity interface{}           `json:"starting_quantity"`
        EndingQuantity   Optional[interface{}] `json:"ending_quantity"`
        UnitPrice        interface{}           `json:"unit_price"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    p.StartingQuantity = temp.StartingQuantity
    p.EndingQuantity = temp.EndingQuantity
    p.UnitPrice = temp.UnitPrice
    return nil
}
