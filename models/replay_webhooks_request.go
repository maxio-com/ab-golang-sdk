package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ReplayWebhooksRequest represents a ReplayWebhooksRequest struct.
type ReplayWebhooksRequest struct {
    Ids                  []int64        `json:"ids"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ReplayWebhooksRequest.
// It customizes the JSON marshaling process for ReplayWebhooksRequest objects.
func (r ReplayWebhooksRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ReplayWebhooksRequest object to a map representation for JSON marshaling.
func (r ReplayWebhooksRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["ids"] = r.Ids
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ReplayWebhooksRequest.
// It customizes the JSON unmarshaling process for ReplayWebhooksRequest objects.
func (r *ReplayWebhooksRequest) UnmarshalJSON(input []byte) error {
    var temp replayWebhooksRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ids")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Ids = *temp.Ids
    return nil
}

// replayWebhooksRequest is a temporary struct used for validating the fields of ReplayWebhooksRequest.
type replayWebhooksRequest  struct {
    Ids *[]int64 `json:"ids"`
}

func (r *replayWebhooksRequest) validate() error {
    var errs []string
    if r.Ids == nil {
        errs = append(errs, "required field `ids` is missing for type `Replay Webhooks Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
