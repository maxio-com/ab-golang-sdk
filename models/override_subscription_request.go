package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// OverrideSubscriptionRequest represents a OverrideSubscriptionRequest struct.
type OverrideSubscriptionRequest struct {
    Subscription         OverrideSubscription `json:"subscription"`
    AdditionalProperties map[string]any       `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OverrideSubscriptionRequest.
// It customizes the JSON marshaling process for OverrideSubscriptionRequest objects.
func (o OverrideSubscriptionRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OverrideSubscriptionRequest object to a map representation for JSON marshaling.
func (o OverrideSubscriptionRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, o.AdditionalProperties)
    structMap["subscription"] = o.Subscription.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OverrideSubscriptionRequest.
// It customizes the JSON unmarshaling process for OverrideSubscriptionRequest objects.
func (o *OverrideSubscriptionRequest) UnmarshalJSON(input []byte) error {
    var temp overrideSubscriptionRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "subscription")
    if err != nil {
    	return err
    }
    
    o.AdditionalProperties = additionalProperties
    o.Subscription = *temp.Subscription
    return nil
}

// TODO
type overrideSubscriptionRequest  struct {
    Subscription *OverrideSubscription `json:"subscription"`
}

func (o *overrideSubscriptionRequest) validate() error {
    var errs []string
    if o.Subscription == nil {
        errs = append(errs, "required field `subscription` is missing for type `Override Subscription Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
