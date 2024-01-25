package models

import (
	"encoding/json"
)

// UpdateSubscriptionGroup represents a UpdateSubscriptionGroup struct.
type UpdateSubscriptionGroup struct {
	MemberIds []int `json:"member_ids,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateSubscriptionGroup.
// It customizes the JSON marshaling process for UpdateSubscriptionGroup objects.
func (u *UpdateSubscriptionGroup) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateSubscriptionGroup object to a map representation for JSON marshaling.
func (u *UpdateSubscriptionGroup) toMap() map[string]any {
	structMap := make(map[string]any)
	if u.MemberIds != nil {
		structMap["member_ids"] = u.MemberIds
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateSubscriptionGroup.
// It customizes the JSON unmarshaling process for UpdateSubscriptionGroup objects.
func (u *UpdateSubscriptionGroup) UnmarshalJSON(input []byte) error {
	temp := &struct {
		MemberIds []int `json:"member_ids,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.MemberIds = temp.MemberIds
	return nil
}
