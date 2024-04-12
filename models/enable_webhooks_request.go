package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// EnableWebhooksRequest represents a EnableWebhooksRequest struct.
type EnableWebhooksRequest struct {
    WebhooksEnabled      bool           `json:"webhooks_enabled"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for EnableWebhooksRequest.
// It customizes the JSON marshaling process for EnableWebhooksRequest objects.
func (e EnableWebhooksRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(e.toMap())
}

// toMap converts the EnableWebhooksRequest object to a map representation for JSON marshaling.
func (e EnableWebhooksRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, e.AdditionalProperties)
    structMap["webhooks_enabled"] = e.WebhooksEnabled
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EnableWebhooksRequest.
// It customizes the JSON unmarshaling process for EnableWebhooksRequest objects.
func (e *EnableWebhooksRequest) UnmarshalJSON(input []byte) error {
    var temp enableWebhooksRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "webhooks_enabled")
    if err != nil {
    	return err
    }
    
    e.AdditionalProperties = additionalProperties
    e.WebhooksEnabled = *temp.WebhooksEnabled
    return nil
}

// TODO
type enableWebhooksRequest  struct {
    WebhooksEnabled *bool `json:"webhooks_enabled"`
}

func (e *enableWebhooksRequest) validate() error {
    var errs []string
    if e.WebhooksEnabled == nil {
        errs = append(errs, "required field `webhooks_enabled` is missing for type `Enable Webhooks Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
