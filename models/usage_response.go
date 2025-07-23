// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// UsageResponse represents a UsageResponse struct.
type UsageResponse struct {
    Usage                Usage                  `json:"usage"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UsageResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UsageResponse) String() string {
    return fmt.Sprintf(
    	"UsageResponse[Usage=%v, AdditionalProperties=%v]",
    	u.Usage, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UsageResponse.
// It customizes the JSON marshaling process for UsageResponse objects.
func (u UsageResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "usage"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UsageResponse object to a map representation for JSON marshaling.
func (u UsageResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    structMap["usage"] = u.Usage.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UsageResponse.
// It customizes the JSON unmarshaling process for UsageResponse objects.
func (u *UsageResponse) UnmarshalJSON(input []byte) error {
    var temp tempUsageResponse
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
    u.AdditionalProperties = additionalProperties
    
    u.Usage = *temp.Usage
    return nil
}

// tempUsageResponse is a temporary struct used for validating the fields of UsageResponse.
type tempUsageResponse  struct {
    Usage *Usage `json:"usage"`
}

func (u *tempUsageResponse) validate() error {
    var errs []string
    if u.Usage == nil {
        errs = append(errs, "required field `usage` is missing for type `Usage Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
