// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// UpdateFeatureTemplateRequest represents a UpdateFeatureTemplateRequest struct.
type UpdateFeatureTemplateRequest struct {
    // `key` cannot be changed once set. `kind` cannot be changed once any feature catalog item has been created from this template.
    Feature              Feature1               `json:"feature"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UpdateFeatureTemplateRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpdateFeatureTemplateRequest) String() string {
    return fmt.Sprintf(
    	"UpdateFeatureTemplateRequest[Feature=%v, AdditionalProperties=%v]",
    	u.Feature, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UpdateFeatureTemplateRequest.
// It customizes the JSON marshaling process for UpdateFeatureTemplateRequest objects.
func (u UpdateFeatureTemplateRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "feature"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateFeatureTemplateRequest object to a map representation for JSON marshaling.
func (u UpdateFeatureTemplateRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    structMap["feature"] = u.Feature.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateFeatureTemplateRequest.
// It customizes the JSON unmarshaling process for UpdateFeatureTemplateRequest objects.
func (u *UpdateFeatureTemplateRequest) UnmarshalJSON(input []byte) error {
    var temp tempUpdateFeatureTemplateRequest
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
    u.AdditionalProperties = additionalProperties
    
    u.Feature = *temp.Feature
    return nil
}

// tempUpdateFeatureTemplateRequest is a temporary struct used for validating the fields of UpdateFeatureTemplateRequest.
type tempUpdateFeatureTemplateRequest  struct {
    Feature *Feature1 `json:"feature"`
}

func (u *tempUpdateFeatureTemplateRequest) validate() error {
    var errs []string
    if u.Feature == nil {
        errs = append(errs, "required field `feature` is missing for type `Update Feature Template Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
