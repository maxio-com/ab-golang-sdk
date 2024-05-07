package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// EventBasedBillingSegmentError represents a EventBasedBillingSegmentError struct.
type EventBasedBillingSegmentError struct {
    // The key of the object would be a number (an index in the request array) where the error occurred. In the value object, the key represents the field and the value is an array with error messages. In most cases, this object would contain just one key.
    Segments             map[string]interface{} `json:"segments"`
    AdditionalProperties map[string]any         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for EventBasedBillingSegmentError.
// It customizes the JSON marshaling process for EventBasedBillingSegmentError objects.
func (e EventBasedBillingSegmentError) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(e.toMap())
}

// toMap converts the EventBasedBillingSegmentError object to a map representation for JSON marshaling.
func (e EventBasedBillingSegmentError) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, e.AdditionalProperties)
    structMap["segments"] = e.Segments
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EventBasedBillingSegmentError.
// It customizes the JSON unmarshaling process for EventBasedBillingSegmentError objects.
func (e *EventBasedBillingSegmentError) UnmarshalJSON(input []byte) error {
    var temp eventBasedBillingSegmentError
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "segments")
    if err != nil {
    	return err
    }
    
    e.AdditionalProperties = additionalProperties
    e.Segments = *temp.Segments
    return nil
}

// eventBasedBillingSegmentError is a temporary struct used for validating the fields of EventBasedBillingSegmentError.
type eventBasedBillingSegmentError  struct {
    Segments *map[string]interface{} `json:"segments"`
}

func (e *eventBasedBillingSegmentError) validate() error {
    var errs []string
    if e.Segments == nil {
        errs = append(errs, "required field `segments` is missing for type `Event Based Billing Segment Error`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
