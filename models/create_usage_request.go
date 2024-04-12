package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CreateUsageRequest represents a CreateUsageRequest struct.
type CreateUsageRequest struct {
    Usage                CreateUsage    `json:"usage"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreateUsageRequest.
// It customizes the JSON marshaling process for CreateUsageRequest objects.
func (c CreateUsageRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateUsageRequest object to a map representation for JSON marshaling.
func (c CreateUsageRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["usage"] = c.Usage.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateUsageRequest.
// It customizes the JSON unmarshaling process for CreateUsageRequest objects.
func (c *CreateUsageRequest) UnmarshalJSON(input []byte) error {
    var temp createUsageRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "usage")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Usage = *temp.Usage
    return nil
}

// TODO
type createUsageRequest  struct {
    Usage *CreateUsage `json:"usage"`
}

func (c *createUsageRequest) validate() error {
    var errs []string
    if c.Usage == nil {
        errs = append(errs, "required field `usage` is missing for type `Create Usage Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
