// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// EventBasedBillingSegmentError represents a EventBasedBillingSegmentError struct.
type EventBasedBillingSegmentError struct {
	// The key of the object would be a number (an index in the request array) where the error occurred. In the value object, the key represents the field and the value is an array with error messages. In most cases, this object would contain just one key.
	Segments             map[string]interface{} `json:"segments"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for EventBasedBillingSegmentError,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (e EventBasedBillingSegmentError) String() string {
	return fmt.Sprintf(
		"EventBasedBillingSegmentError[Segments=%v, AdditionalProperties=%v]",
		e.Segments, e.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for EventBasedBillingSegmentError.
// It customizes the JSON marshaling process for EventBasedBillingSegmentError objects.
func (e EventBasedBillingSegmentError) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(e.AdditionalProperties,
		"segments"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(e.toMap())
}

// toMap converts the EventBasedBillingSegmentError object to a map representation for JSON marshaling.
func (e EventBasedBillingSegmentError) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, e.AdditionalProperties)
	structMap["segments"] = e.Segments
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EventBasedBillingSegmentError.
// It customizes the JSON unmarshaling process for EventBasedBillingSegmentError objects.
func (e *EventBasedBillingSegmentError) UnmarshalJSON(input []byte) error {
	var temp tempEventBasedBillingSegmentError
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "segments")
	if err != nil {
		return err
	}
	e.AdditionalProperties = additionalProperties

	e.Segments = *temp.Segments
	return nil
}

// tempEventBasedBillingSegmentError is a temporary struct used for validating the fields of EventBasedBillingSegmentError.
type tempEventBasedBillingSegmentError struct {
	Segments *map[string]interface{} `json:"segments"`
}

func (e *tempEventBasedBillingSegmentError) validate() error {
	var errs []string
	if e.Segments == nil {
		errs = append(errs, "required field `segments` is missing for type `Event Based Billing Segment Error`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
