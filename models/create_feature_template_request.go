// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CreateFeatureTemplateRequest represents a CreateFeatureTemplateRequest struct.
type CreateFeatureTemplateRequest struct {
    Feature              Feature                `json:"feature"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CreateFeatureTemplateRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateFeatureTemplateRequest) String() string {
    return fmt.Sprintf(
    	"CreateFeatureTemplateRequest[Feature=%v, AdditionalProperties=%v]",
    	c.Feature, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreateFeatureTemplateRequest.
// It customizes the JSON marshaling process for CreateFeatureTemplateRequest objects.
func (c CreateFeatureTemplateRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "feature"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateFeatureTemplateRequest object to a map representation for JSON marshaling.
func (c CreateFeatureTemplateRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["feature"] = c.Feature.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateFeatureTemplateRequest.
// It customizes the JSON unmarshaling process for CreateFeatureTemplateRequest objects.
func (c *CreateFeatureTemplateRequest) UnmarshalJSON(input []byte) error {
    var temp tempCreateFeatureTemplateRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "feature")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Feature = *temp.Feature
    return nil
}

// tempCreateFeatureTemplateRequest is a temporary struct used for validating the fields of CreateFeatureTemplateRequest.
type tempCreateFeatureTemplateRequest  struct {
    Feature *Feature `json:"feature"`
}

func (c *tempCreateFeatureTemplateRequest) validate() error {
    var errs []string
    if c.Feature == nil {
        errs = append(errs, "required field `feature` is missing for type `Create Feature Template Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
