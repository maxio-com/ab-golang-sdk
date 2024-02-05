package models

import (
    "encoding/json"
)

// UpsertPrepaidConfigurationRequest represents a UpsertPrepaidConfigurationRequest struct.
type UpsertPrepaidConfigurationRequest struct {
    PrepaidConfiguration UpsertPrepaidConfiguration `json:"prepaid_configuration"`
}

// MarshalJSON implements the json.Marshaler interface for UpsertPrepaidConfigurationRequest.
// It customizes the JSON marshaling process for UpsertPrepaidConfigurationRequest objects.
func (u *UpsertPrepaidConfigurationRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UpsertPrepaidConfigurationRequest object to a map representation for JSON marshaling.
func (u *UpsertPrepaidConfigurationRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["prepaid_configuration"] = u.PrepaidConfiguration.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpsertPrepaidConfigurationRequest.
// It customizes the JSON unmarshaling process for UpsertPrepaidConfigurationRequest objects.
func (u *UpsertPrepaidConfigurationRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        PrepaidConfiguration UpsertPrepaidConfiguration `json:"prepaid_configuration"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    u.PrepaidConfiguration = temp.PrepaidConfiguration
    return nil
}
