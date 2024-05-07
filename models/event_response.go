package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// EventResponse represents a EventResponse struct.
type EventResponse struct {
    Event                Event          `json:"event"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for EventResponse.
// It customizes the JSON marshaling process for EventResponse objects.
func (e EventResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(e.toMap())
}

// toMap converts the EventResponse object to a map representation for JSON marshaling.
func (e EventResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, e.AdditionalProperties)
    structMap["event"] = e.Event.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EventResponse.
// It customizes the JSON unmarshaling process for EventResponse objects.
func (e *EventResponse) UnmarshalJSON(input []byte) error {
    var temp eventResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "event")
    if err != nil {
    	return err
    }
    
    e.AdditionalProperties = additionalProperties
    e.Event = *temp.Event
    return nil
}

// eventResponse is a temporary struct used for validating the fields of EventResponse.
type eventResponse  struct {
    Event *Event `json:"event"`
}

func (e *eventResponse) validate() error {
    var errs []string
    if e.Event == nil {
        errs = append(errs, "required field `event` is missing for type `Event Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
