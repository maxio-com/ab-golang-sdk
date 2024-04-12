package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CreateSubscriptionGroup represents a CreateSubscriptionGroup struct.
type CreateSubscriptionGroup struct {
    SubscriptionId       int            `json:"subscription_id"`
    MemberIds            []int          `json:"member_ids,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreateSubscriptionGroup.
// It customizes the JSON marshaling process for CreateSubscriptionGroup objects.
func (c CreateSubscriptionGroup) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateSubscriptionGroup object to a map representation for JSON marshaling.
func (c CreateSubscriptionGroup) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["subscription_id"] = c.SubscriptionId
    if c.MemberIds != nil {
        structMap["member_ids"] = c.MemberIds
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateSubscriptionGroup.
// It customizes the JSON unmarshaling process for CreateSubscriptionGroup objects.
func (c *CreateSubscriptionGroup) UnmarshalJSON(input []byte) error {
    var temp createSubscriptionGroup
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "subscription_id", "member_ids")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.SubscriptionId = *temp.SubscriptionId
    c.MemberIds = temp.MemberIds
    return nil
}

// TODO
type createSubscriptionGroup  struct {
    SubscriptionId *int  `json:"subscription_id"`
    MemberIds      []int `json:"member_ids,omitempty"`
}

func (c *createSubscriptionGroup) validate() error {
    var errs []string
    if c.SubscriptionId == nil {
        errs = append(errs, "required field `subscription_id` is missing for type `Create Subscription Group`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
