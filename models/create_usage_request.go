package models

import (
    "encoding/json"
)

// CreateUsageRequest represents a CreateUsageRequest struct.
type CreateUsageRequest struct {
    Usage CreateUsage `json:"usage"`
}

// MarshalJSON implements the json.Marshaler interface for CreateUsageRequest.
// It customizes the JSON marshaling process for CreateUsageRequest objects.
func (c *CreateUsageRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateUsageRequest object to a map representation for JSON marshaling.
func (c *CreateUsageRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["usage"] = c.Usage.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateUsageRequest.
// It customizes the JSON unmarshaling process for CreateUsageRequest objects.
func (c *CreateUsageRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Usage CreateUsage `json:"usage"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.Usage = temp.Usage
    return nil
}
