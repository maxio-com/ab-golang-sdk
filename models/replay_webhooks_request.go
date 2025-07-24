// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ReplayWebhooksRequest represents a ReplayWebhooksRequest struct.
type ReplayWebhooksRequest struct {
    Ids                  []int64                `json:"ids"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ReplayWebhooksRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ReplayWebhooksRequest) String() string {
    return fmt.Sprintf(
    	"ReplayWebhooksRequest[Ids=%v, AdditionalProperties=%v]",
    	r.Ids, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ReplayWebhooksRequest.
// It customizes the JSON marshaling process for ReplayWebhooksRequest objects.
func (r ReplayWebhooksRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "ids"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ReplayWebhooksRequest object to a map representation for JSON marshaling.
func (r ReplayWebhooksRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["ids"] = r.Ids
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ReplayWebhooksRequest.
// It customizes the JSON unmarshaling process for ReplayWebhooksRequest objects.
func (r *ReplayWebhooksRequest) UnmarshalJSON(input []byte) error {
    var temp tempReplayWebhooksRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ids")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Ids = *temp.Ids
    return nil
}

// tempReplayWebhooksRequest is a temporary struct used for validating the fields of ReplayWebhooksRequest.
type tempReplayWebhooksRequest  struct {
    Ids *[]int64 `json:"ids"`
}

func (r *tempReplayWebhooksRequest) validate() error {
    var errs []string
    if r.Ids == nil {
        errs = append(errs, "required field `ids` is missing for type `Replay Webhooks Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
