// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// FeatureCatalogItemResponse represents a FeatureCatalogItemResponse struct.
type FeatureCatalogItemResponse struct {
    // A feature template attached to a specific product or component (or one of their price points), with a concrete value. When a subscriber signs up for or is assigned this product/component, the feature catalog item is provisioned as an entitlement on their subscription.
    Feature              FeatureCatalogItem     `json:"feature"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for FeatureCatalogItemResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (f FeatureCatalogItemResponse) String() string {
    return fmt.Sprintf(
    	"FeatureCatalogItemResponse[Feature=%v, AdditionalProperties=%v]",
    	f.Feature, f.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for FeatureCatalogItemResponse.
// It customizes the JSON marshaling process for FeatureCatalogItemResponse objects.
func (f FeatureCatalogItemResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(f.AdditionalProperties,
        "feature"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(f.toMap())
}

// toMap converts the FeatureCatalogItemResponse object to a map representation for JSON marshaling.
func (f FeatureCatalogItemResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, f.AdditionalProperties)
    structMap["feature"] = f.Feature.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for FeatureCatalogItemResponse.
// It customizes the JSON unmarshaling process for FeatureCatalogItemResponse objects.
func (f *FeatureCatalogItemResponse) UnmarshalJSON(input []byte) error {
    var temp tempFeatureCatalogItemResponse
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

// tempFeatureCatalogItemResponse is a temporary struct used for validating the fields of FeatureCatalogItemResponse.
type tempFeatureCatalogItemResponse  struct {
    Feature *FeatureCatalogItem `json:"feature"`
}

func (f *tempFeatureCatalogItemResponse) validate() error {
    var errs []string
    if f.Feature == nil {
        errs = append(errs, "required field `feature` is missing for type `Feature Catalog Item Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
