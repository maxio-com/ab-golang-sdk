// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CustomerChangesPreviewResponse represents a CustomerChangesPreviewResponse struct.
type CustomerChangesPreviewResponse struct {
    Changes              CustomerChange         `json:"changes"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CustomerChangesPreviewResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CustomerChangesPreviewResponse) String() string {
    return fmt.Sprintf(
    	"CustomerChangesPreviewResponse[Changes=%v, AdditionalProperties=%v]",
    	c.Changes, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CustomerChangesPreviewResponse.
// It customizes the JSON marshaling process for CustomerChangesPreviewResponse objects.
func (c CustomerChangesPreviewResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "changes"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CustomerChangesPreviewResponse object to a map representation for JSON marshaling.
func (c CustomerChangesPreviewResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["changes"] = c.Changes.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CustomerChangesPreviewResponse.
// It customizes the JSON unmarshaling process for CustomerChangesPreviewResponse objects.
func (c *CustomerChangesPreviewResponse) UnmarshalJSON(input []byte) error {
    var temp tempCustomerChangesPreviewResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "changes")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Changes = *temp.Changes
    return nil
}

// tempCustomerChangesPreviewResponse is a temporary struct used for validating the fields of CustomerChangesPreviewResponse.
type tempCustomerChangesPreviewResponse  struct {
    Changes *CustomerChange `json:"changes"`
}

func (c *tempCustomerChangesPreviewResponse) validate() error {
    var errs []string
    if c.Changes == nil {
        errs = append(errs, "required field `changes` is missing for type `Customer Changes Preview Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
