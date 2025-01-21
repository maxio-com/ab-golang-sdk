/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// EBBEvent represents a EBBEvent struct.
type EBBEvent struct {
    Chargify             *ChargifyEBB           `json:"chargify,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for EBBEvent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (e EBBEvent) String() string {
    return fmt.Sprintf(
    	"EBBEvent[Chargify=%v, AdditionalProperties=%v]",
    	e.Chargify, e.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for EBBEvent.
// It customizes the JSON marshaling process for EBBEvent objects.
func (e EBBEvent) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(e.AdditionalProperties,
        "chargify"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(e.toMap())
}

// toMap converts the EBBEvent object to a map representation for JSON marshaling.
func (e EBBEvent) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, e.AdditionalProperties)
    if e.Chargify != nil {
        structMap["chargify"] = e.Chargify.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EBBEvent.
// It customizes the JSON unmarshaling process for EBBEvent objects.
func (e *EBBEvent) UnmarshalJSON(input []byte) error {
    var temp tempEBBEvent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "chargify")
    if err != nil {
    	return err
    }
    e.AdditionalProperties = additionalProperties
    
    e.Chargify = temp.Chargify
    return nil
}

// tempEBBEvent is a temporary struct used for validating the fields of EBBEvent.
type tempEBBEvent  struct {
    Chargify *ChargifyEBB `json:"chargify,omitempty"`
}
