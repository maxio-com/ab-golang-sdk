package models

import (
    "encoding/json"
)

// CountResponse represents a CountResponse struct.
type CountResponse struct {
    Count                *int           `json:"count,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CountResponse.
// It customizes the JSON marshaling process for CountResponse objects.
func (c CountResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CountResponse object to a map representation for JSON marshaling.
func (c CountResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Count != nil {
        structMap["count"] = c.Count
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CountResponse.
// It customizes the JSON unmarshaling process for CountResponse objects.
func (c *CountResponse) UnmarshalJSON(input []byte) error {
    var temp countResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "count")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Count = temp.Count
    return nil
}

// countResponse is a temporary struct used for validating the fields of CountResponse.
type countResponse  struct {
    Count *int `json:"count,omitempty"`
}
