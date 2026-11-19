// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// FeatureTemplateResponse represents a FeatureTemplateResponse struct.
type FeatureTemplateResponse struct {
    // A feature that can be granted to subscribers, defined once at the site level and then attached to products or components.
    Feature              FeatureTemplate        `json:"feature"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for FeatureTemplateResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (f FeatureTemplateResponse) String() string {
    return fmt.Sprintf(
    	"FeatureTemplateResponse[Feature=%v, AdditionalProperties=%v]",
    	f.Feature, f.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for FeatureTemplateResponse.
// It customizes the JSON marshaling process for FeatureTemplateResponse objects.
func (f FeatureTemplateResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(f.AdditionalProperties,
        "feature"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(f.toMap())
}

// toMap converts the FeatureTemplateResponse object to a map representation for JSON marshaling.
func (f FeatureTemplateResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, f.AdditionalProperties)
    structMap["feature"] = f.Feature.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for FeatureTemplateResponse.
// It customizes the JSON unmarshaling process for FeatureTemplateResponse objects.
func (f *FeatureTemplateResponse) UnmarshalJSON(input []byte) error {
    var temp tempFeatureTemplateResponse
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
    f.AdditionalProperties = additionalProperties
    
    f.Feature = *temp.Feature
    return nil
}

// tempFeatureTemplateResponse is a temporary struct used for validating the fields of FeatureTemplateResponse.
type tempFeatureTemplateResponse  struct {
    Feature *FeatureTemplate `json:"feature"`
}

func (f *tempFeatureTemplateResponse) validate() error {
    var errs []string
    if f.Feature == nil {
        errs = append(errs, "required field `feature` is missing for type `Feature Template Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
