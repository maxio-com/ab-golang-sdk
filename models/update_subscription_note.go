package models

import (
	"encoding/json"
)

// UpdateSubscriptionNote represents a UpdateSubscriptionNote struct.
// Updatable fields for Subscription Note
type UpdateSubscriptionNote struct {
	Body   string `json:"body"`
	Sticky bool   `json:"sticky"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateSubscriptionNote.
// It customizes the JSON marshaling process for UpdateSubscriptionNote objects.
func (u *UpdateSubscriptionNote) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateSubscriptionNote object to a map representation for JSON marshaling.
func (u *UpdateSubscriptionNote) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["body"] = u.Body
	structMap["sticky"] = u.Sticky
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateSubscriptionNote.
// It customizes the JSON unmarshaling process for UpdateSubscriptionNote objects.
func (u *UpdateSubscriptionNote) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Body   string `json:"body"`
		Sticky bool   `json:"sticky"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.Body = temp.Body
	u.Sticky = temp.Sticky
	return nil
}
