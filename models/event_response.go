/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// EventResponse represents a EventResponse struct.
type EventResponse struct {
    Event                Event                  `json:"event"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for EventResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (e EventResponse) String() string {
    return fmt.Sprintf(
    	"EventResponse[Event=%v, AdditionalProperties=%v]",
    	e.Event, e.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for EventResponse.
// It customizes the JSON marshaling process for EventResponse objects.
func (e EventResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(e.AdditionalProperties,
        "event"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(e.toMap())
}

// toMap converts the EventResponse object to a map representation for JSON marshaling.
func (e EventResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, e.AdditionalProperties)
    structMap["event"] = e.Event.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EventResponse.
// It customizes the JSON unmarshaling process for EventResponse objects.
func (e *EventResponse) UnmarshalJSON(input []byte) error {
    var temp tempEventResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "event")
    if err != nil {
    	return err
    }
    e.AdditionalProperties = additionalProperties
    
    e.Event = *temp.Event
    return nil
}

// tempEventResponse is a temporary struct used for validating the fields of EventResponse.
type tempEventResponse  struct {
    Event *Event `json:"event"`
}

func (e *tempEventResponse) validate() error {
    var errs []string
    if e.Event == nil {
        errs = append(errs, "required field `event` is missing for type `Event Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
