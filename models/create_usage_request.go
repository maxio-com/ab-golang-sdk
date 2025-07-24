// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CreateUsageRequest represents a CreateUsageRequest struct.
type CreateUsageRequest struct {
    Usage                CreateUsage            `json:"usage"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CreateUsageRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateUsageRequest) String() string {
    return fmt.Sprintf(
    	"CreateUsageRequest[Usage=%v, AdditionalProperties=%v]",
    	c.Usage, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreateUsageRequest.
// It customizes the JSON marshaling process for CreateUsageRequest objects.
func (c CreateUsageRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "usage"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateUsageRequest object to a map representation for JSON marshaling.
func (c CreateUsageRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["usage"] = c.Usage.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateUsageRequest.
// It customizes the JSON unmarshaling process for CreateUsageRequest objects.
func (c *CreateUsageRequest) UnmarshalJSON(input []byte) error {
    var temp tempCreateUsageRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "usage")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Usage = *temp.Usage
    return nil
}

// tempCreateUsageRequest is a temporary struct used for validating the fields of CreateUsageRequest.
type tempCreateUsageRequest  struct {
    Usage *CreateUsage `json:"usage"`
}

func (c *tempCreateUsageRequest) validate() error {
    var errs []string
    if c.Usage == nil {
        errs = append(errs, "required field `usage` is missing for type `Create Usage Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
