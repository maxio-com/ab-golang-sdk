// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CancellationRequest represents a CancellationRequest struct.
type CancellationRequest struct {
    Subscription         CancellationOptions    `json:"subscription"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CancellationRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CancellationRequest) String() string {
    return fmt.Sprintf(
    	"CancellationRequest[Subscription=%v, AdditionalProperties=%v]",
    	c.Subscription, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CancellationRequest.
// It customizes the JSON marshaling process for CancellationRequest objects.
func (c CancellationRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "subscription"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CancellationRequest object to a map representation for JSON marshaling.
func (c CancellationRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["subscription"] = c.Subscription.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CancellationRequest.
// It customizes the JSON unmarshaling process for CancellationRequest objects.
func (c *CancellationRequest) UnmarshalJSON(input []byte) error {
    var temp tempCancellationRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "subscription")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Subscription = *temp.Subscription
    return nil
}

// tempCancellationRequest is a temporary struct used for validating the fields of CancellationRequest.
type tempCancellationRequest  struct {
    Subscription *CancellationOptions `json:"subscription"`
}

func (c *tempCancellationRequest) validate() error {
    var errs []string
    if c.Subscription == nil {
        errs = append(errs, "required field `subscription` is missing for type `Cancellation Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
