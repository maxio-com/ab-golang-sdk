// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// UpdateFeatureCatalogItemRequest represents a UpdateFeatureCatalogItemRequest struct.
type UpdateFeatureCatalogItemRequest struct {
    Feature              Feature3               `json:"feature"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UpdateFeatureCatalogItemRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpdateFeatureCatalogItemRequest) String() string {
    return fmt.Sprintf(
    	"UpdateFeatureCatalogItemRequest[Feature=%v, AdditionalProperties=%v]",
    	u.Feature, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UpdateFeatureCatalogItemRequest.
// It customizes the JSON marshaling process for UpdateFeatureCatalogItemRequest objects.
func (u UpdateFeatureCatalogItemRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "feature"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateFeatureCatalogItemRequest object to a map representation for JSON marshaling.
func (u UpdateFeatureCatalogItemRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    structMap["feature"] = u.Feature.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateFeatureCatalogItemRequest.
// It customizes the JSON unmarshaling process for UpdateFeatureCatalogItemRequest objects.
func (u *UpdateFeatureCatalogItemRequest) UnmarshalJSON(input []byte) error {
    var temp tempUpdateFeatureCatalogItemRequest
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

// tempUpdateFeatureCatalogItemRequest is a temporary struct used for validating the fields of UpdateFeatureCatalogItemRequest.
type tempUpdateFeatureCatalogItemRequest  struct {
    Feature *Feature3 `json:"feature"`
}

func (u *tempUpdateFeatureCatalogItemRequest) validate() error {
    var errs []string
    if u.Feature == nil {
        errs = append(errs, "required field `feature` is missing for type `Update Feature Catalog Item Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
