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

// EnableWebhooksRequest represents a EnableWebhooksRequest struct.
type EnableWebhooksRequest struct {
    WebhooksEnabled      bool                   `json:"webhooks_enabled"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for EnableWebhooksRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (e EnableWebhooksRequest) String() string {
    return fmt.Sprintf(
    	"EnableWebhooksRequest[WebhooksEnabled=%v, AdditionalProperties=%v]",
    	e.WebhooksEnabled, e.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for EnableWebhooksRequest.
// It customizes the JSON marshaling process for EnableWebhooksRequest objects.
func (e EnableWebhooksRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(e.AdditionalProperties,
        "webhooks_enabled"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(e.toMap())
}

// toMap converts the EnableWebhooksRequest object to a map representation for JSON marshaling.
func (e EnableWebhooksRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, e.AdditionalProperties)
    structMap["webhooks_enabled"] = e.WebhooksEnabled
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EnableWebhooksRequest.
// It customizes the JSON unmarshaling process for EnableWebhooksRequest objects.
func (e *EnableWebhooksRequest) UnmarshalJSON(input []byte) error {
    var temp tempEnableWebhooksRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "webhooks_enabled")
    if err != nil {
    	return err
    }
    e.AdditionalProperties = additionalProperties
    
    e.WebhooksEnabled = *temp.WebhooksEnabled
    return nil
}

// tempEnableWebhooksRequest is a temporary struct used for validating the fields of EnableWebhooksRequest.
type tempEnableWebhooksRequest  struct {
    WebhooksEnabled *bool `json:"webhooks_enabled"`
}

func (e *tempEnableWebhooksRequest) validate() error {
    var errs []string
    if e.WebhooksEnabled == nil {
        errs = append(errs, "required field `webhooks_enabled` is missing for type `Enable Webhooks Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
