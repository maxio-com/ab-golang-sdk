// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// EventBasedBillingSegmentSingleError represents a EventBasedBillingSegmentSingleError struct.
// The key of the object would be a number (an index in the request array) where the error occurred. In the value object, the key represents the field and the value is an array with error messages. In most cases, this object would contain just one key.
type EventBasedBillingSegmentSingleError struct {
    PricePoint           []string               `json:"price_point"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for EventBasedBillingSegmentSingleError,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (e EventBasedBillingSegmentSingleError) String() string {
    return fmt.Sprintf(
    	"EventBasedBillingSegmentSingleError[PricePoint=%v, AdditionalProperties=%v]",
    	e.PricePoint, e.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for EventBasedBillingSegmentSingleError.
// It customizes the JSON marshaling process for EventBasedBillingSegmentSingleError objects.
func (e EventBasedBillingSegmentSingleError) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(e.AdditionalProperties,
        "price_point"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(e.toMap())
}

// toMap converts the EventBasedBillingSegmentSingleError object to a map representation for JSON marshaling.
func (e EventBasedBillingSegmentSingleError) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, e.AdditionalProperties)
    structMap["price_point"] = e.PricePoint
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EventBasedBillingSegmentSingleError.
// It customizes the JSON unmarshaling process for EventBasedBillingSegmentSingleError objects.
func (e *EventBasedBillingSegmentSingleError) UnmarshalJSON(input []byte) error {
    var temp tempEventBasedBillingSegmentSingleError
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "price_point")
    if err != nil {
    	return err
    }
    e.AdditionalProperties = additionalProperties
    
    e.PricePoint = *temp.PricePoint
    return nil
}

// tempEventBasedBillingSegmentSingleError is a temporary struct used for validating the fields of EventBasedBillingSegmentSingleError.
type tempEventBasedBillingSegmentSingleError  struct {
    PricePoint *[]string `json:"price_point"`
}

func (e *tempEventBasedBillingSegmentSingleError) validate() error {
    var errs []string
    if e.PricePoint == nil {
        errs = append(errs, "required field `price_point` is missing for type `Event Based Billing Segment Single Error`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
