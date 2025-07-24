// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// UpsertPrepaidConfigurationRequest represents a UpsertPrepaidConfigurationRequest struct.
type UpsertPrepaidConfigurationRequest struct {
    PrepaidConfiguration UpsertPrepaidConfiguration `json:"prepaid_configuration"`
    AdditionalProperties map[string]interface{}     `json:"_"`
}

// String implements the fmt.Stringer interface for UpsertPrepaidConfigurationRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpsertPrepaidConfigurationRequest) String() string {
    return fmt.Sprintf(
    	"UpsertPrepaidConfigurationRequest[PrepaidConfiguration=%v, AdditionalProperties=%v]",
    	u.PrepaidConfiguration, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UpsertPrepaidConfigurationRequest.
// It customizes the JSON marshaling process for UpsertPrepaidConfigurationRequest objects.
func (u UpsertPrepaidConfigurationRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "prepaid_configuration"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpsertPrepaidConfigurationRequest object to a map representation for JSON marshaling.
func (u UpsertPrepaidConfigurationRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    structMap["prepaid_configuration"] = u.PrepaidConfiguration.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpsertPrepaidConfigurationRequest.
// It customizes the JSON unmarshaling process for UpsertPrepaidConfigurationRequest objects.
func (u *UpsertPrepaidConfigurationRequest) UnmarshalJSON(input []byte) error {
    var temp tempUpsertPrepaidConfigurationRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "prepaid_configuration")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.PrepaidConfiguration = *temp.PrepaidConfiguration
    return nil
}

// tempUpsertPrepaidConfigurationRequest is a temporary struct used for validating the fields of UpsertPrepaidConfigurationRequest.
type tempUpsertPrepaidConfigurationRequest  struct {
    PrepaidConfiguration *UpsertPrepaidConfiguration `json:"prepaid_configuration"`
}

func (u *tempUpsertPrepaidConfigurationRequest) validate() error {
    var errs []string
    if u.PrepaidConfiguration == nil {
        errs = append(errs, "required field `prepaid_configuration` is missing for type `Upsert Prepaid Configuration Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
