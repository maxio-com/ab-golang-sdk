package models

import (
    "encoding/json"
)

// CreatePrepaidComponent represents a CreatePrepaidComponent struct.
type CreatePrepaidComponent struct {
    PrepaidUsageComponent PrepaidUsageComponent `json:"prepaid_usage_component"`
}

// MarshalJSON implements the json.Marshaler interface for CreatePrepaidComponent.
// It customizes the JSON marshaling process for CreatePrepaidComponent objects.
func (c *CreatePrepaidComponent) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreatePrepaidComponent object to a map representation for JSON marshaling.
func (c *CreatePrepaidComponent) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["prepaid_usage_component"] = c.PrepaidUsageComponent.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreatePrepaidComponent.
// It customizes the JSON unmarshaling process for CreatePrepaidComponent objects.
func (c *CreatePrepaidComponent) UnmarshalJSON(input []byte) error {
    temp := &struct {
        PrepaidUsageComponent PrepaidUsageComponent `json:"prepaid_usage_component"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.PrepaidUsageComponent = temp.PrepaidUsageComponent
    return nil
}
