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

// CreateSubscriptionGroupRequest represents a CreateSubscriptionGroupRequest struct.
type CreateSubscriptionGroupRequest struct {
    SubscriptionGroup    CreateSubscriptionGroup `json:"subscription_group"`
    AdditionalProperties map[string]interface{}  `json:"_"`
}

// String implements the fmt.Stringer interface for CreateSubscriptionGroupRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateSubscriptionGroupRequest) String() string {
    return fmt.Sprintf(
    	"CreateSubscriptionGroupRequest[SubscriptionGroup=%v, AdditionalProperties=%v]",
    	c.SubscriptionGroup, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreateSubscriptionGroupRequest.
// It customizes the JSON marshaling process for CreateSubscriptionGroupRequest objects.
func (c CreateSubscriptionGroupRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "subscription_group"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateSubscriptionGroupRequest object to a map representation for JSON marshaling.
func (c CreateSubscriptionGroupRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["subscription_group"] = c.SubscriptionGroup.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateSubscriptionGroupRequest.
// It customizes the JSON unmarshaling process for CreateSubscriptionGroupRequest objects.
func (c *CreateSubscriptionGroupRequest) UnmarshalJSON(input []byte) error {
    var temp tempCreateSubscriptionGroupRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "subscription_group")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.SubscriptionGroup = *temp.SubscriptionGroup
    return nil
}

// tempCreateSubscriptionGroupRequest is a temporary struct used for validating the fields of CreateSubscriptionGroupRequest.
type tempCreateSubscriptionGroupRequest  struct {
    SubscriptionGroup *CreateSubscriptionGroup `json:"subscription_group"`
}

func (c *tempCreateSubscriptionGroupRequest) validate() error {
    var errs []string
    if c.SubscriptionGroup == nil {
        errs = append(errs, "required field `subscription_group` is missing for type `Create Subscription Group Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
