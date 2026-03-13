// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ScheduledRenewalConfigurationItemRequest represents a ScheduledRenewalConfigurationItemRequest struct.
type ScheduledRenewalConfigurationItemRequest struct {
    RenewalConfigurationItem ScheduledRenewalConfigurationItemRequestRenewalConfigurationItem `json:"renewal_configuration_item"`
    AdditionalProperties     map[string]interface{}                                           `json:"_"`
}

// String implements the fmt.Stringer interface for ScheduledRenewalConfigurationItemRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s ScheduledRenewalConfigurationItemRequest) String() string {
    return fmt.Sprintf(
    	"ScheduledRenewalConfigurationItemRequest[RenewalConfigurationItem=%v, AdditionalProperties=%v]",
    	s.RenewalConfigurationItem, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ScheduledRenewalConfigurationItemRequest.
// It customizes the JSON marshaling process for ScheduledRenewalConfigurationItemRequest objects.
func (s ScheduledRenewalConfigurationItemRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "renewal_configuration_item"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the ScheduledRenewalConfigurationItemRequest object to a map representation for JSON marshaling.
func (s ScheduledRenewalConfigurationItemRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["renewal_configuration_item"] = s.RenewalConfigurationItem.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ScheduledRenewalConfigurationItemRequest.
// It customizes the JSON unmarshaling process for ScheduledRenewalConfigurationItemRequest objects.
func (s *ScheduledRenewalConfigurationItemRequest) UnmarshalJSON(input []byte) error {
    var temp tempScheduledRenewalConfigurationItemRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "renewal_configuration_item")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.RenewalConfigurationItem = *temp.RenewalConfigurationItem
    return nil
}

// tempScheduledRenewalConfigurationItemRequest is a temporary struct used for validating the fields of ScheduledRenewalConfigurationItemRequest.
type tempScheduledRenewalConfigurationItemRequest  struct {
    RenewalConfigurationItem *ScheduledRenewalConfigurationItemRequestRenewalConfigurationItem `json:"renewal_configuration_item"`
}

func (s *tempScheduledRenewalConfigurationItemRequest) validate() error {
    var errs []string
    if s.RenewalConfigurationItem == nil {
        errs = append(errs, "required field `renewal_configuration_item` is missing for type `Scheduled Renewal Configuration Item Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
