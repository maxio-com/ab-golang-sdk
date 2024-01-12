package models

import (
    "encoding/json"
)

// PrepaidConfigurationResponse represents a PrepaidConfigurationResponse struct.
type PrepaidConfigurationResponse struct {
    PrepaidConfiguration PrepaidConfiguration `json:"prepaid_configuration"`
}

// MarshalJSON implements the json.Marshaler interface for PrepaidConfigurationResponse.
// It customizes the JSON marshaling process for PrepaidConfigurationResponse objects.
func (p *PrepaidConfigurationResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PrepaidConfigurationResponse object to a map representation for JSON marshaling.
func (p *PrepaidConfigurationResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["prepaid_configuration"] = p.PrepaidConfiguration
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PrepaidConfigurationResponse.
// It customizes the JSON unmarshaling process for PrepaidConfigurationResponse objects.
func (p *PrepaidConfigurationResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        PrepaidConfiguration PrepaidConfiguration `json:"prepaid_configuration"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    p.PrepaidConfiguration = temp.PrepaidConfiguration
    return nil
}
