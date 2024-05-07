package models

import (
    "encoding/json"
)

// UpdateComponentPricePointRequest represents a UpdateComponentPricePointRequest struct.
type UpdateComponentPricePointRequest struct {
    PricePoint           *UpdateComponentPricePoint `json:"price_point,omitempty"`
    AdditionalProperties map[string]any             `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateComponentPricePointRequest.
// It customizes the JSON marshaling process for UpdateComponentPricePointRequest objects.
func (u UpdateComponentPricePointRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateComponentPricePointRequest object to a map representation for JSON marshaling.
func (u UpdateComponentPricePointRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    if u.PricePoint != nil {
        structMap["price_point"] = u.PricePoint.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateComponentPricePointRequest.
// It customizes the JSON unmarshaling process for UpdateComponentPricePointRequest objects.
func (u *UpdateComponentPricePointRequest) UnmarshalJSON(input []byte) error {
    var temp updateComponentPricePointRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "price_point")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.PricePoint = temp.PricePoint
    return nil
}

// updateComponentPricePointRequest is a temporary struct used for validating the fields of UpdateComponentPricePointRequest.
type updateComponentPricePointRequest  struct {
    PricePoint *UpdateComponentPricePoint `json:"price_point,omitempty"`
}
