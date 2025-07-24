// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CreateSubscriptionRequest represents a CreateSubscriptionRequest struct.
type CreateSubscriptionRequest struct {
    Subscription         CreateSubscription     `json:"subscription"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CreateSubscriptionRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateSubscriptionRequest) String() string {
    return fmt.Sprintf(
    	"CreateSubscriptionRequest[Subscription=%v, AdditionalProperties=%v]",
    	c.Subscription, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreateSubscriptionRequest.
// It customizes the JSON marshaling process for CreateSubscriptionRequest objects.
func (c CreateSubscriptionRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "subscription"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateSubscriptionRequest object to a map representation for JSON marshaling.
func (c CreateSubscriptionRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["subscription"] = c.Subscription.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateSubscriptionRequest.
// It customizes the JSON unmarshaling process for CreateSubscriptionRequest objects.
func (c *CreateSubscriptionRequest) UnmarshalJSON(input []byte) error {
    var temp tempCreateSubscriptionRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "subscription")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Subscription = *temp.Subscription
    return nil
}

// tempCreateSubscriptionRequest is a temporary struct used for validating the fields of CreateSubscriptionRequest.
type tempCreateSubscriptionRequest  struct {
    Subscription *CreateSubscription `json:"subscription"`
}

func (c *tempCreateSubscriptionRequest) validate() error {
    var errs []string
    if c.Subscription == nil {
        errs = append(errs, "required field `subscription` is missing for type `Create Subscription Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
