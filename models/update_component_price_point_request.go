package models

import (
    "encoding/json"
)

// UpdateComponentPricePointRequest represents a UpdateComponentPricePointRequest struct.
type UpdateComponentPricePointRequest struct {
    PricePoint *UpdateComponentPricePoint `json:"price_point,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateComponentPricePointRequest.
// It customizes the JSON marshaling process for UpdateComponentPricePointRequest objects.
func (u *UpdateComponentPricePointRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateComponentPricePointRequest object to a map representation for JSON marshaling.
func (u *UpdateComponentPricePointRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    if u.PricePoint != nil {
        structMap["price_point"] = u.PricePoint
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateComponentPricePointRequest.
// It customizes the JSON unmarshaling process for UpdateComponentPricePointRequest objects.
func (u *UpdateComponentPricePointRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        PricePoint *UpdateComponentPricePoint `json:"price_point,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    u.PricePoint = temp.PricePoint
    return nil
}
