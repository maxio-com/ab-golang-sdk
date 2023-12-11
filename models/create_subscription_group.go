package models

import (
	"encoding/json"
)

// CreateSubscriptionGroup represents a CreateSubscriptionGroup struct.
type CreateSubscriptionGroup struct {
	SubscriptionId interface{} `json:"subscription_id"`
	MemberIds      []int       `json:"member_ids,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateSubscriptionGroup.
// It customizes the JSON marshaling process for CreateSubscriptionGroup objects.
func (c *CreateSubscriptionGroup) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateSubscriptionGroup object to a map representation for JSON marshaling.
func (c *CreateSubscriptionGroup) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["subscription_id"] = c.SubscriptionId
	if c.MemberIds != nil {
		structMap["member_ids"] = c.MemberIds
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateSubscriptionGroup.
// It customizes the JSON unmarshaling process for CreateSubscriptionGroup objects.
func (c *CreateSubscriptionGroup) UnmarshalJSON(input []byte) error {
	temp := &struct {
		SubscriptionId interface{} `json:"subscription_id"`
		MemberIds      []int       `json:"member_ids,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.SubscriptionId = temp.SubscriptionId
	c.MemberIds = temp.MemberIds
	return nil
}
