package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// UpsertPrepaidConfigurationRequest represents a UpsertPrepaidConfigurationRequest struct.
type UpsertPrepaidConfigurationRequest struct {
    PrepaidConfiguration UpsertPrepaidConfiguration `json:"prepaid_configuration"`
    AdditionalProperties map[string]any             `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UpsertPrepaidConfigurationRequest.
// It customizes the JSON marshaling process for UpsertPrepaidConfigurationRequest objects.
func (u UpsertPrepaidConfigurationRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UpsertPrepaidConfigurationRequest object to a map representation for JSON marshaling.
func (u UpsertPrepaidConfigurationRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    structMap["prepaid_configuration"] = u.PrepaidConfiguration.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpsertPrepaidConfigurationRequest.
// It customizes the JSON unmarshaling process for UpsertPrepaidConfigurationRequest objects.
func (u *UpsertPrepaidConfigurationRequest) UnmarshalJSON(input []byte) error {
    var temp upsertPrepaidConfigurationRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "prepaid_configuration")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.PrepaidConfiguration = *temp.PrepaidConfiguration
    return nil
}

// TODO
type upsertPrepaidConfigurationRequest  struct {
    PrepaidConfiguration *UpsertPrepaidConfiguration `json:"prepaid_configuration"`
}

func (u *upsertPrepaidConfigurationRequest) validate() error {
    var errs []string
    if u.PrepaidConfiguration == nil {
        errs = append(errs, "required field `prepaid_configuration` is missing for type `Upsert Prepaid Configuration Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
