// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CreateOrUpdateEndpoint represents a CreateOrUpdateEndpoint struct.
// Used to Create or Update Endpoint
type CreateOrUpdateEndpoint struct {
    Url                  string                 `json:"url"`
    WebhookSubscriptions []WebhookSubscription  `json:"webhook_subscriptions"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CreateOrUpdateEndpoint,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateOrUpdateEndpoint) String() string {
    return fmt.Sprintf(
    	"CreateOrUpdateEndpoint[Url=%v, WebhookSubscriptions=%v, AdditionalProperties=%v]",
    	c.Url, c.WebhookSubscriptions, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreateOrUpdateEndpoint.
// It customizes the JSON marshaling process for CreateOrUpdateEndpoint objects.
func (c CreateOrUpdateEndpoint) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "url", "webhook_subscriptions"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateOrUpdateEndpoint object to a map representation for JSON marshaling.
func (c CreateOrUpdateEndpoint) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["url"] = c.Url
    structMap["webhook_subscriptions"] = c.WebhookSubscriptions
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateOrUpdateEndpoint.
// It customizes the JSON unmarshaling process for CreateOrUpdateEndpoint objects.
func (c *CreateOrUpdateEndpoint) UnmarshalJSON(input []byte) error {
    var temp tempCreateOrUpdateEndpoint
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "url", "webhook_subscriptions")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Url = *temp.Url
    c.WebhookSubscriptions = *temp.WebhookSubscriptions
    return nil
}

// tempCreateOrUpdateEndpoint is a temporary struct used for validating the fields of CreateOrUpdateEndpoint.
type tempCreateOrUpdateEndpoint  struct {
    Url                  *string                `json:"url"`
    WebhookSubscriptions *[]WebhookSubscription `json:"webhook_subscriptions"`
}

func (c *tempCreateOrUpdateEndpoint) validate() error {
    var errs []string
    if c.Url == nil {
        errs = append(errs, "required field `url` is missing for type `Create or Update Endpoint`")
    }
    if c.WebhookSubscriptions == nil {
        errs = append(errs, "required field `webhook_subscriptions` is missing for type `Create or Update Endpoint`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
