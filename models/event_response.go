package models

import (
    "encoding/json"
)

// EventResponse represents a EventResponse struct.
type EventResponse struct {
    Event Event `json:"event"`
}

// MarshalJSON implements the json.Marshaler interface for EventResponse.
// It customizes the JSON marshaling process for EventResponse objects.
func (e *EventResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(e.toMap())
}

// toMap converts the EventResponse object to a map representation for JSON marshaling.
func (e *EventResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["event"] = e.Event
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EventResponse.
// It customizes the JSON unmarshaling process for EventResponse objects.
func (e *EventResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Event Event `json:"event"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    e.Event = temp.Event
    return nil
}
