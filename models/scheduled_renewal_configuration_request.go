// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ScheduledRenewalConfigurationRequest represents a ScheduledRenewalConfigurationRequest struct.
type ScheduledRenewalConfigurationRequest struct {
    RenewalConfiguration ScheduledRenewalConfigurationRequestBody `json:"renewal_configuration"`
    AdditionalProperties map[string]interface{}                   `json:"_"`
}

// String implements the fmt.Stringer interface for ScheduledRenewalConfigurationRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s ScheduledRenewalConfigurationRequest) String() string {
    return fmt.Sprintf(
    	"ScheduledRenewalConfigurationRequest[RenewalConfiguration=%v, AdditionalProperties=%v]",
    	s.RenewalConfiguration, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ScheduledRenewalConfigurationRequest.
// It customizes the JSON marshaling process for ScheduledRenewalConfigurationRequest objects.
func (s ScheduledRenewalConfigurationRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "renewal_configuration"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the ScheduledRenewalConfigurationRequest object to a map representation for JSON marshaling.
func (s ScheduledRenewalConfigurationRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["renewal_configuration"] = s.RenewalConfiguration.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ScheduledRenewalConfigurationRequest.
// It customizes the JSON unmarshaling process for ScheduledRenewalConfigurationRequest objects.
func (s *ScheduledRenewalConfigurationRequest) UnmarshalJSON(input []byte) error {
    var temp tempScheduledRenewalConfigurationRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "renewal_configuration")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.RenewalConfiguration = *temp.RenewalConfiguration
    return nil
}

// tempScheduledRenewalConfigurationRequest is a temporary struct used for validating the fields of ScheduledRenewalConfigurationRequest.
type tempScheduledRenewalConfigurationRequest  struct {
    RenewalConfiguration *ScheduledRenewalConfigurationRequestBody `json:"renewal_configuration"`
}

func (s *tempScheduledRenewalConfigurationRequest) validate() error {
    var errs []string
    if s.RenewalConfiguration == nil {
        errs = append(errs, "required field `renewal_configuration` is missing for type `Scheduled Renewal Configuration Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
