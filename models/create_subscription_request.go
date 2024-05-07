package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CreateSubscriptionRequest represents a CreateSubscriptionRequest struct.
type CreateSubscriptionRequest struct {
    Subscription         CreateSubscription `json:"subscription"`
    AdditionalProperties map[string]any     `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreateSubscriptionRequest.
// It customizes the JSON marshaling process for CreateSubscriptionRequest objects.
func (c CreateSubscriptionRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateSubscriptionRequest object to a map representation for JSON marshaling.
func (c CreateSubscriptionRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["subscription"] = c.Subscription.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateSubscriptionRequest.
// It customizes the JSON unmarshaling process for CreateSubscriptionRequest objects.
func (c *CreateSubscriptionRequest) UnmarshalJSON(input []byte) error {
    var temp createSubscriptionRequest
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
    
    c.AdditionalProperties = additionalProperties
    c.Subscription = *temp.Subscription
    return nil
}

// createSubscriptionRequest is a temporary struct used for validating the fields of CreateSubscriptionRequest.
type createSubscriptionRequest  struct {
    Subscription *CreateSubscription `json:"subscription"`
}

func (c *createSubscriptionRequest) validate() error {
    var errs []string
    if c.Subscription == nil {
        errs = append(errs, "required field `subscription` is missing for type `Create Subscription Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
