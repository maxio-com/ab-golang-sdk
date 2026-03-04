// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// ScheduledRenewalConfigurationsResponse represents a ScheduledRenewalConfigurationsResponse struct.
type ScheduledRenewalConfigurationsResponse struct {
    ScheduledRenewalConfigurations []ScheduledRenewalConfiguration `json:"scheduled_renewal_configurations,omitempty"`
    AdditionalProperties           map[string]interface{}          `json:"_"`
}

// String implements the fmt.Stringer interface for ScheduledRenewalConfigurationsResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s ScheduledRenewalConfigurationsResponse) String() string {
    return fmt.Sprintf(
    	"ScheduledRenewalConfigurationsResponse[ScheduledRenewalConfigurations=%v, AdditionalProperties=%v]",
    	s.ScheduledRenewalConfigurations, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ScheduledRenewalConfigurationsResponse.
// It customizes the JSON marshaling process for ScheduledRenewalConfigurationsResponse objects.
func (s ScheduledRenewalConfigurationsResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "scheduled_renewal_configurations"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the ScheduledRenewalConfigurationsResponse object to a map representation for JSON marshaling.
func (s ScheduledRenewalConfigurationsResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.ScheduledRenewalConfigurations != nil {
        structMap["scheduled_renewal_configurations"] = s.ScheduledRenewalConfigurations
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ScheduledRenewalConfigurationsResponse.
// It customizes the JSON unmarshaling process for ScheduledRenewalConfigurationsResponse objects.
func (s *ScheduledRenewalConfigurationsResponse) UnmarshalJSON(input []byte) error {
    var temp tempScheduledRenewalConfigurationsResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "scheduled_renewal_configurations")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.ScheduledRenewalConfigurations = temp.ScheduledRenewalConfigurations
    return nil
}

// tempScheduledRenewalConfigurationsResponse is a temporary struct used for validating the fields of ScheduledRenewalConfigurationsResponse.
type tempScheduledRenewalConfigurationsResponse  struct {
    ScheduledRenewalConfigurations []ScheduledRenewalConfiguration `json:"scheduled_renewal_configurations,omitempty"`
}
