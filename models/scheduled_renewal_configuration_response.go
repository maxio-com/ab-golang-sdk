// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// ScheduledRenewalConfigurationResponse represents a ScheduledRenewalConfigurationResponse struct.
type ScheduledRenewalConfigurationResponse struct {
    ScheduledRenewalConfiguration *ScheduledRenewalConfiguration `json:"scheduled_renewal_configuration,omitempty"`
    AdditionalProperties          map[string]interface{}         `json:"_"`
}

// String implements the fmt.Stringer interface for ScheduledRenewalConfigurationResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s ScheduledRenewalConfigurationResponse) String() string {
    return fmt.Sprintf(
    	"ScheduledRenewalConfigurationResponse[ScheduledRenewalConfiguration=%v, AdditionalProperties=%v]",
    	s.ScheduledRenewalConfiguration, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ScheduledRenewalConfigurationResponse.
// It customizes the JSON marshaling process for ScheduledRenewalConfigurationResponse objects.
func (s ScheduledRenewalConfigurationResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "scheduled_renewal_configuration"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the ScheduledRenewalConfigurationResponse object to a map representation for JSON marshaling.
func (s ScheduledRenewalConfigurationResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.ScheduledRenewalConfiguration != nil {
        structMap["scheduled_renewal_configuration"] = s.ScheduledRenewalConfiguration.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ScheduledRenewalConfigurationResponse.
// It customizes the JSON unmarshaling process for ScheduledRenewalConfigurationResponse objects.
func (s *ScheduledRenewalConfigurationResponse) UnmarshalJSON(input []byte) error {
    var temp tempScheduledRenewalConfigurationResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "scheduled_renewal_configuration")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.ScheduledRenewalConfiguration = temp.ScheduledRenewalConfiguration
    return nil
}

// tempScheduledRenewalConfigurationResponse is a temporary struct used for validating the fields of ScheduledRenewalConfigurationResponse.
type tempScheduledRenewalConfigurationResponse  struct {
    ScheduledRenewalConfiguration *ScheduledRenewalConfiguration `json:"scheduled_renewal_configuration,omitempty"`
}
