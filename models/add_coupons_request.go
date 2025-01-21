/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// AddCouponsRequest represents a AddCouponsRequest struct.
type AddCouponsRequest struct {
    Codes                []string               `json:"codes,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AddCouponsRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AddCouponsRequest) String() string {
    return fmt.Sprintf(
    	"AddCouponsRequest[Codes=%v, AdditionalProperties=%v]",
    	a.Codes, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AddCouponsRequest.
// It customizes the JSON marshaling process for AddCouponsRequest objects.
func (a AddCouponsRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "codes"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AddCouponsRequest object to a map representation for JSON marshaling.
func (a AddCouponsRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Codes != nil {
        structMap["codes"] = a.Codes
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AddCouponsRequest.
// It customizes the JSON unmarshaling process for AddCouponsRequest objects.
func (a *AddCouponsRequest) UnmarshalJSON(input []byte) error {
    var temp tempAddCouponsRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "codes")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.Codes = temp.Codes
    return nil
}

// tempAddCouponsRequest is a temporary struct used for validating the fields of AddCouponsRequest.
type tempAddCouponsRequest  struct {
    Codes []string `json:"codes,omitempty"`
}
