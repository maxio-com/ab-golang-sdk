/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ChangeChargebackStatusEventData represents a ChangeChargebackStatusEventData struct.
// Example schema for an `change_chargeback_status` event
type ChangeChargebackStatusEventData struct {
    ChargebackStatus     ChargebackStatus `json:"chargeback_status"`
    AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ChangeChargebackStatusEventData.
// It customizes the JSON marshaling process for ChangeChargebackStatusEventData objects.
func (c ChangeChargebackStatusEventData) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ChangeChargebackStatusEventData object to a map representation for JSON marshaling.
func (c ChangeChargebackStatusEventData) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["chargeback_status"] = c.ChargebackStatus
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ChangeChargebackStatusEventData.
// It customizes the JSON unmarshaling process for ChangeChargebackStatusEventData objects.
func (c *ChangeChargebackStatusEventData) UnmarshalJSON(input []byte) error {
    var temp tempChangeChargebackStatusEventData
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "chargeback_status")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.ChargebackStatus = *temp.ChargebackStatus
    return nil
}

// tempChangeChargebackStatusEventData is a temporary struct used for validating the fields of ChangeChargebackStatusEventData.
type tempChangeChargebackStatusEventData  struct {
    ChargebackStatus *ChargebackStatus `json:"chargeback_status"`
}

func (c *tempChangeChargebackStatusEventData) validate() error {
    var errs []string
    if c.ChargebackStatus == nil {
        errs = append(errs, "required field `chargeback_status` is missing for type `Change Chargeback Status Event Data`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
