package models

import (
    "encoding/json"
)

// EventBasedBillingSegmentError represents a EventBasedBillingSegmentError struct.
type EventBasedBillingSegmentError struct {
    // The key of the object would be a number (an index in the request array) where the error occurred. In the value object, the key represents the field and the value is an array with error messages. In most cases, this object would contain just one key.
    Segments map[string]interface{} `json:"segments"`
}

// MarshalJSON implements the json.Marshaler interface for EventBasedBillingSegmentError.
// It customizes the JSON marshaling process for EventBasedBillingSegmentError objects.
func (e *EventBasedBillingSegmentError) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(e.toMap())
}

// toMap converts the EventBasedBillingSegmentError object to a map representation for JSON marshaling.
func (e *EventBasedBillingSegmentError) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["segments"] = e.Segments
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EventBasedBillingSegmentError.
// It customizes the JSON unmarshaling process for EventBasedBillingSegmentError objects.
func (e *EventBasedBillingSegmentError) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Segments map[string]interface{} `json:"segments"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    e.Segments = temp.Segments
    return nil
}
