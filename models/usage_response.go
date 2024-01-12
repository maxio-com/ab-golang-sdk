package models

import (
    "encoding/json"
)

// UsageResponse represents a UsageResponse struct.
type UsageResponse struct {
    Usage Usage `json:"usage"`
}

// MarshalJSON implements the json.Marshaler interface for UsageResponse.
// It customizes the JSON marshaling process for UsageResponse objects.
func (u *UsageResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UsageResponse object to a map representation for JSON marshaling.
func (u *UsageResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["usage"] = u.Usage
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UsageResponse.
// It customizes the JSON unmarshaling process for UsageResponse objects.
func (u *UsageResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Usage Usage `json:"usage"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    u.Usage = temp.Usage
    return nil
}
