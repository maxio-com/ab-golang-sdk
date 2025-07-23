// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// DelayedCancellationResponse represents a DelayedCancellationResponse struct.
type DelayedCancellationResponse struct {
    Message              *string                `json:"message,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for DelayedCancellationResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (d DelayedCancellationResponse) String() string {
    return fmt.Sprintf(
    	"DelayedCancellationResponse[Message=%v, AdditionalProperties=%v]",
    	d.Message, d.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for DelayedCancellationResponse.
// It customizes the JSON marshaling process for DelayedCancellationResponse objects.
func (d DelayedCancellationResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(d.AdditionalProperties,
        "message"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(d.toMap())
}

// toMap converts the DelayedCancellationResponse object to a map representation for JSON marshaling.
func (d DelayedCancellationResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, d.AdditionalProperties)
    if d.Message != nil {
        structMap["message"] = d.Message
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DelayedCancellationResponse.
// It customizes the JSON unmarshaling process for DelayedCancellationResponse objects.
func (d *DelayedCancellationResponse) UnmarshalJSON(input []byte) error {
    var temp tempDelayedCancellationResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "message")
    if err != nil {
    	return err
    }
    d.AdditionalProperties = additionalProperties
    
    d.Message = temp.Message
    return nil
}

// tempDelayedCancellationResponse is a temporary struct used for validating the fields of DelayedCancellationResponse.
type tempDelayedCancellationResponse  struct {
    Message *string `json:"message,omitempty"`
}
