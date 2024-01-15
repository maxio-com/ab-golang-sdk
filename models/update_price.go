package models

import (
    "encoding/json"
)

// UpdatePrice represents a UpdatePrice struct.
type UpdatePrice struct {
    Id               *int    `json:"id,omitempty"`
    EndingQuantity   *int    `json:"ending_quantity,omitempty"`
    UnitPrice        *int    `json:"unit_price,omitempty"`
    Destroy          *string `json:"_destroy,omitempty"`
    StartingQuantity *int    `json:"starting_quantity,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for UpdatePrice.
// It customizes the JSON marshaling process for UpdatePrice objects.
func (u *UpdatePrice) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UpdatePrice object to a map representation for JSON marshaling.
func (u *UpdatePrice) toMap() map[string]any {
    structMap := make(map[string]any)
    if u.Id != nil {
        structMap["id"] = u.Id
    }
    if u.EndingQuantity != nil {
        structMap["ending_quantity"] = u.EndingQuantity
    }
    if u.UnitPrice != nil {
        structMap["unit_price"] = u.UnitPrice
    }
    if u.Destroy != nil {
        structMap["_destroy"] = u.Destroy
    }
    if u.StartingQuantity != nil {
        structMap["starting_quantity"] = u.StartingQuantity
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdatePrice.
// It customizes the JSON unmarshaling process for UpdatePrice objects.
func (u *UpdatePrice) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id               *int    `json:"id,omitempty"`
        EndingQuantity   *int    `json:"ending_quantity,omitempty"`
        UnitPrice        *int    `json:"unit_price,omitempty"`
        Destroy          *string `json:"_destroy,omitempty"`
        StartingQuantity *int    `json:"starting_quantity,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    u.Id = temp.Id
    u.EndingQuantity = temp.EndingQuantity
    u.UnitPrice = temp.UnitPrice
    u.Destroy = temp.Destroy
    u.StartingQuantity = temp.StartingQuantity
    return nil
}
