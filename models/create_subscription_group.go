/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CreateSubscriptionGroup represents a CreateSubscriptionGroup struct.
type CreateSubscriptionGroup struct {
    SubscriptionId       int                    `json:"subscription_id"`
    MemberIds            []int                  `json:"member_ids,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CreateSubscriptionGroup,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateSubscriptionGroup) String() string {
    return fmt.Sprintf(
    	"CreateSubscriptionGroup[SubscriptionId=%v, MemberIds=%v, AdditionalProperties=%v]",
    	c.SubscriptionId, c.MemberIds, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreateSubscriptionGroup.
// It customizes the JSON marshaling process for CreateSubscriptionGroup objects.
func (c CreateSubscriptionGroup) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "subscription_id", "member_ids"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateSubscriptionGroup object to a map representation for JSON marshaling.
func (c CreateSubscriptionGroup) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["subscription_id"] = c.SubscriptionId
    if c.MemberIds != nil {
        structMap["member_ids"] = c.MemberIds
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateSubscriptionGroup.
// It customizes the JSON unmarshaling process for CreateSubscriptionGroup objects.
func (c *CreateSubscriptionGroup) UnmarshalJSON(input []byte) error {
    var temp tempCreateSubscriptionGroup
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "subscription_id", "member_ids")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.SubscriptionId = *temp.SubscriptionId
    c.MemberIds = temp.MemberIds
    return nil
}

// tempCreateSubscriptionGroup is a temporary struct used for validating the fields of CreateSubscriptionGroup.
type tempCreateSubscriptionGroup  struct {
    SubscriptionId *int  `json:"subscription_id"`
    MemberIds      []int `json:"member_ids,omitempty"`
}

func (c *tempCreateSubscriptionGroup) validate() error {
    var errs []string
    if c.SubscriptionId == nil {
        errs = append(errs, "required field `subscription_id` is missing for type `Create Subscription Group`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
