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

// UpdateSubscriptionRequest represents a UpdateSubscriptionRequest struct.
type UpdateSubscriptionRequest struct {
    Subscription         UpdateSubscription     `json:"subscription"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UpdateSubscriptionRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpdateSubscriptionRequest) String() string {
    return fmt.Sprintf(
    	"UpdateSubscriptionRequest[Subscription=%v, AdditionalProperties=%v]",
    	u.Subscription, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UpdateSubscriptionRequest.
// It customizes the JSON marshaling process for UpdateSubscriptionRequest objects.
func (u UpdateSubscriptionRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "subscription"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateSubscriptionRequest object to a map representation for JSON marshaling.
func (u UpdateSubscriptionRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    structMap["subscription"] = u.Subscription.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateSubscriptionRequest.
// It customizes the JSON unmarshaling process for UpdateSubscriptionRequest objects.
func (u *UpdateSubscriptionRequest) UnmarshalJSON(input []byte) error {
    var temp tempUpdateSubscriptionRequest
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
    u.AdditionalProperties = additionalProperties
    
    u.Subscription = *temp.Subscription
    return nil
}

// tempUpdateSubscriptionRequest is a temporary struct used for validating the fields of UpdateSubscriptionRequest.
type tempUpdateSubscriptionRequest  struct {
    Subscription *UpdateSubscription `json:"subscription"`
}

func (u *tempUpdateSubscriptionRequest) validate() error {
    var errs []string
    if u.Subscription == nil {
        errs = append(errs, "required field `subscription` is missing for type `Update Subscription Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
