// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// ScheduledRenewalConfigurationItemResponse represents a ScheduledRenewalConfigurationItemResponse struct.
type ScheduledRenewalConfigurationItemResponse struct {
    ScheduledRenewalConfigurationItem *ScheduledRenewalConfigurationItem `json:"scheduled_renewal_configuration_item,omitempty"`
    AdditionalProperties              map[string]interface{}             `json:"_"`
}

// String implements the fmt.Stringer interface for ScheduledRenewalConfigurationItemResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s ScheduledRenewalConfigurationItemResponse) String() string {
    return fmt.Sprintf(
    	"ScheduledRenewalConfigurationItemResponse[ScheduledRenewalConfigurationItem=%v, AdditionalProperties=%v]",
    	s.ScheduledRenewalConfigurationItem, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ScheduledRenewalConfigurationItemResponse.
// It customizes the JSON marshaling process for ScheduledRenewalConfigurationItemResponse objects.
func (s ScheduledRenewalConfigurationItemResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "scheduled_renewal_configuration_item"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the ScheduledRenewalConfigurationItemResponse object to a map representation for JSON marshaling.
func (s ScheduledRenewalConfigurationItemResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.ScheduledRenewalConfigurationItem != nil {
        structMap["scheduled_renewal_configuration_item"] = s.ScheduledRenewalConfigurationItem.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ScheduledRenewalConfigurationItemResponse.
// It customizes the JSON unmarshaling process for ScheduledRenewalConfigurationItemResponse objects.
func (s *ScheduledRenewalConfigurationItemResponse) UnmarshalJSON(input []byte) error {
    var temp tempScheduledRenewalConfigurationItemResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "scheduled_renewal_configuration_item")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.ScheduledRenewalConfigurationItem = temp.ScheduledRenewalConfigurationItem
    return nil
}

// tempScheduledRenewalConfigurationItemResponse is a temporary struct used for validating the fields of ScheduledRenewalConfigurationItemResponse.
type tempScheduledRenewalConfigurationItemResponse  struct {
    ScheduledRenewalConfigurationItem *ScheduledRenewalConfigurationItem `json:"scheduled_renewal_configuration_item,omitempty"`
}
