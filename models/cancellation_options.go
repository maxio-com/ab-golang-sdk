/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// CancellationOptions represents a CancellationOptions struct.
type CancellationOptions struct {
    // For your internal use. An indication as to why the subscription is being canceled.
    CancellationMessage  *string                `json:"cancellation_message,omitempty"`
    // The reason code associated with the cancellation. See the list of reason codes associated with your site.
    ReasonCode           *string                `json:"reason_code,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CancellationOptions,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CancellationOptions) String() string {
    return fmt.Sprintf(
    	"CancellationOptions[CancellationMessage=%v, ReasonCode=%v, AdditionalProperties=%v]",
    	c.CancellationMessage, c.ReasonCode, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CancellationOptions.
// It customizes the JSON marshaling process for CancellationOptions objects.
func (c CancellationOptions) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "cancellation_message", "reason_code"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CancellationOptions object to a map representation for JSON marshaling.
func (c CancellationOptions) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.CancellationMessage != nil {
        structMap["cancellation_message"] = c.CancellationMessage
    }
    if c.ReasonCode != nil {
        structMap["reason_code"] = c.ReasonCode
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CancellationOptions.
// It customizes the JSON unmarshaling process for CancellationOptions objects.
func (c *CancellationOptions) UnmarshalJSON(input []byte) error {
    var temp tempCancellationOptions
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "cancellation_message", "reason_code")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.CancellationMessage = temp.CancellationMessage
    c.ReasonCode = temp.ReasonCode
    return nil
}

// tempCancellationOptions is a temporary struct used for validating the fields of CancellationOptions.
type tempCancellationOptions  struct {
    CancellationMessage *string `json:"cancellation_message,omitempty"`
    ReasonCode          *string `json:"reason_code,omitempty"`
}
