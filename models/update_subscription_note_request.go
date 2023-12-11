package models

import (
	"encoding/json"
)

// UpdateSubscriptionNoteRequest represents a UpdateSubscriptionNoteRequest struct.
// Updatable fields for Subscription Note
type UpdateSubscriptionNoteRequest struct {
	// Updatable fields for Subscription Note
	Note UpdateSubscriptionNote `json:"note"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateSubscriptionNoteRequest.
// It customizes the JSON marshaling process for UpdateSubscriptionNoteRequest objects.
func (u *UpdateSubscriptionNoteRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateSubscriptionNoteRequest object to a map representation for JSON marshaling.
func (u *UpdateSubscriptionNoteRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["note"] = u.Note
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateSubscriptionNoteRequest.
// It customizes the JSON unmarshaling process for UpdateSubscriptionNoteRequest objects.
func (u *UpdateSubscriptionNoteRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Note UpdateSubscriptionNote `json:"note"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.Note = temp.Note
	return nil
}
