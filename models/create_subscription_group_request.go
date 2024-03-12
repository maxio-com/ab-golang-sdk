package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// CreateSubscriptionGroupRequest represents a CreateSubscriptionGroupRequest struct.
type CreateSubscriptionGroupRequest struct {
	SubscriptionGroup CreateSubscriptionGroup `json:"subscription_group"`
}

// MarshalJSON implements the json.Marshaler interface for CreateSubscriptionGroupRequest.
// It customizes the JSON marshaling process for CreateSubscriptionGroupRequest objects.
func (c *CreateSubscriptionGroupRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateSubscriptionGroupRequest object to a map representation for JSON marshaling.
func (c *CreateSubscriptionGroupRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["subscription_group"] = c.SubscriptionGroup.toMap()
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateSubscriptionGroupRequest.
// It customizes the JSON unmarshaling process for CreateSubscriptionGroupRequest objects.
func (c *CreateSubscriptionGroupRequest) UnmarshalJSON(input []byte) error {
	var temp createSubscriptionGroupRequest
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}

	c.SubscriptionGroup = *temp.SubscriptionGroup
	return nil
}

// TODO
type createSubscriptionGroupRequest struct {
	SubscriptionGroup *CreateSubscriptionGroup `json:"subscription_group"`
}

func (c *createSubscriptionGroupRequest) validate() error {
	var errs []string
	if c.SubscriptionGroup == nil {
		errs = append(errs, "required field `subscription_group` is missing for type `Create Subscription Group Request`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
