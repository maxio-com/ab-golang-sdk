// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// OverrideSubscriptionRequest represents a OverrideSubscriptionRequest struct.
type OverrideSubscriptionRequest struct {
    Subscription         OverrideSubscription   `json:"subscription"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OverrideSubscriptionRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OverrideSubscriptionRequest) String() string {
    return fmt.Sprintf(
    	"OverrideSubscriptionRequest[Subscription=%v, AdditionalProperties=%v]",
    	o.Subscription, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OverrideSubscriptionRequest.
// It customizes the JSON marshaling process for OverrideSubscriptionRequest objects.
func (o OverrideSubscriptionRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "subscription"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OverrideSubscriptionRequest object to a map representation for JSON marshaling.
func (o OverrideSubscriptionRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    structMap["subscription"] = o.Subscription.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OverrideSubscriptionRequest.
// It customizes the JSON unmarshaling process for OverrideSubscriptionRequest objects.
func (o *OverrideSubscriptionRequest) UnmarshalJSON(input []byte) error {
    var temp tempOverrideSubscriptionRequest
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
    o.AdditionalProperties = additionalProperties
    
    o.Subscription = *temp.Subscription
    return nil
}

// tempOverrideSubscriptionRequest is a temporary struct used for validating the fields of OverrideSubscriptionRequest.
type tempOverrideSubscriptionRequest  struct {
    Subscription *OverrideSubscription `json:"subscription"`
}

func (o *tempOverrideSubscriptionRequest) validate() error {
    var errs []string
    if o.Subscription == nil {
        errs = append(errs, "required field `subscription` is missing for type `Override Subscription Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
