// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ScheduledRenewalUpdateRequest represents a ScheduledRenewalUpdateRequest struct.
type ScheduledRenewalUpdateRequest struct {
    RenewalConfigurationItem ScheduledRenewalUpdateRequestRenewalConfigurationItem `json:"renewal_configuration_item"`
    AdditionalProperties     map[string]interface{}                                `json:"_"`
}

// String implements the fmt.Stringer interface for ScheduledRenewalUpdateRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s ScheduledRenewalUpdateRequest) String() string {
    return fmt.Sprintf(
    	"ScheduledRenewalUpdateRequest[RenewalConfigurationItem=%v, AdditionalProperties=%v]",
    	s.RenewalConfigurationItem, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ScheduledRenewalUpdateRequest.
// It customizes the JSON marshaling process for ScheduledRenewalUpdateRequest objects.
func (s ScheduledRenewalUpdateRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "renewal_configuration_item"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the ScheduledRenewalUpdateRequest object to a map representation for JSON marshaling.
func (s ScheduledRenewalUpdateRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["renewal_configuration_item"] = s.RenewalConfigurationItem.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ScheduledRenewalUpdateRequest.
// It customizes the JSON unmarshaling process for ScheduledRenewalUpdateRequest objects.
func (s *ScheduledRenewalUpdateRequest) UnmarshalJSON(input []byte) error {
    var temp tempScheduledRenewalUpdateRequest
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

// tempScheduledRenewalUpdateRequest is a temporary struct used for validating the fields of ScheduledRenewalUpdateRequest.
type tempScheduledRenewalUpdateRequest  struct {
    RenewalConfigurationItem *ScheduledRenewalUpdateRequestRenewalConfigurationItem `json:"renewal_configuration_item"`
}

func (s *tempScheduledRenewalUpdateRequest) validate() error {
    var errs []string
    if s.RenewalConfigurationItem == nil {
        errs = append(errs, "required field `renewal_configuration_item` is missing for type `Scheduled Renewal Update Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
