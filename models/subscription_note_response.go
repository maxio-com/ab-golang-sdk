package models

import (
    "encoding/json"
)

// SubscriptionNoteResponse represents a SubscriptionNoteResponse struct.
type SubscriptionNoteResponse struct {
    Note SubscriptionNote `json:"note"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionNoteResponse.
// It customizes the JSON marshaling process for SubscriptionNoteResponse objects.
func (s *SubscriptionNoteResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionNoteResponse object to a map representation for JSON marshaling.
func (s *SubscriptionNoteResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["note"] = s.Note.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionNoteResponse.
// It customizes the JSON unmarshaling process for SubscriptionNoteResponse objects.
func (s *SubscriptionNoteResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Note SubscriptionNote `json:"note"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    s.Note = temp.Note
    return nil
}
