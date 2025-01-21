/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// ProformaError represents a ProformaError struct.
type ProformaError struct {
    // The error is base if it is not directly associated with a single attribute.
    Subscription         *BaseStringError       `json:"subscription,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ProformaError,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p ProformaError) String() string {
    return fmt.Sprintf(
    	"ProformaError[Subscription=%v, AdditionalProperties=%v]",
    	p.Subscription, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ProformaError.
// It customizes the JSON marshaling process for ProformaError objects.
func (p ProformaError) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(p.AdditionalProperties,
        "subscription"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(p.toMap())
}

// toMap converts the ProformaError object to a map representation for JSON marshaling.
func (p ProformaError) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, p.AdditionalProperties)
    if p.Subscription != nil {
        structMap["subscription"] = p.Subscription.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ProformaError.
// It customizes the JSON unmarshaling process for ProformaError objects.
func (p *ProformaError) UnmarshalJSON(input []byte) error {
    var temp tempProformaError
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "subscription")
    if err != nil {
    	return err
    }
    p.AdditionalProperties = additionalProperties
    
    p.Subscription = temp.Subscription
    return nil
}

// tempProformaError is a temporary struct used for validating the fields of ProformaError.
type tempProformaError  struct {
    Subscription *BaseStringError `json:"subscription,omitempty"`
}
