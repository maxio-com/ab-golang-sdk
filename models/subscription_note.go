package models

import (
	"encoding/json"
)

// SubscriptionNote represents a SubscriptionNote struct.
type SubscriptionNote struct {
	Id             *int    `json:"id,omitempty"`
	Body           *string `json:"body,omitempty"`
	SubscriptionId *int    `json:"subscription_id,omitempty"`
	CreatedAt      *string `json:"created_at,omitempty"`
	UpdatedAt      *string `json:"updated_at,omitempty"`
	Sticky         *bool   `json:"sticky,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionNote.
// It customizes the JSON marshaling process for SubscriptionNote objects.
func (s *SubscriptionNote) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionNote object to a map representation for JSON marshaling.
func (s *SubscriptionNote) toMap() map[string]any {
	structMap := make(map[string]any)
	if s.Id != nil {
		structMap["id"] = s.Id
	}
	if s.Body != nil {
		structMap["body"] = s.Body
	}
	if s.SubscriptionId != nil {
		structMap["subscription_id"] = s.SubscriptionId
	}
	if s.CreatedAt != nil {
		structMap["created_at"] = s.CreatedAt
	}
	if s.UpdatedAt != nil {
		structMap["updated_at"] = s.UpdatedAt
	}
	if s.Sticky != nil {
		structMap["sticky"] = s.Sticky
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionNote.
// It customizes the JSON unmarshaling process for SubscriptionNote objects.
func (s *SubscriptionNote) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id             *int    `json:"id,omitempty"`
		Body           *string `json:"body,omitempty"`
		SubscriptionId *int    `json:"subscription_id,omitempty"`
		CreatedAt      *string `json:"created_at,omitempty"`
		UpdatedAt      *string `json:"updated_at,omitempty"`
		Sticky         *bool   `json:"sticky,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	s.Id = temp.Id
	s.Body = temp.Body
	s.SubscriptionId = temp.SubscriptionId
	s.CreatedAt = temp.CreatedAt
	s.UpdatedAt = temp.UpdatedAt
	s.Sticky = temp.Sticky
	return nil
}
