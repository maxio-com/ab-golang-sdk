// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CreateFeatureCatalogItemRequest represents a CreateFeatureCatalogItemRequest struct.
// The owning product or component is taken from the URL and must not be included in the request body.
type CreateFeatureCatalogItemRequest struct {
    Feature              Feature2               `json:"feature"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CreateFeatureCatalogItemRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateFeatureCatalogItemRequest) String() string {
    return fmt.Sprintf(
    	"CreateFeatureCatalogItemRequest[Feature=%v, AdditionalProperties=%v]",
    	c.Feature, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreateFeatureCatalogItemRequest.
// It customizes the JSON marshaling process for CreateFeatureCatalogItemRequest objects.
func (c CreateFeatureCatalogItemRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "feature"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateFeatureCatalogItemRequest object to a map representation for JSON marshaling.
func (c CreateFeatureCatalogItemRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["feature"] = c.Feature.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateFeatureCatalogItemRequest.
// It customizes the JSON unmarshaling process for CreateFeatureCatalogItemRequest objects.
func (c *CreateFeatureCatalogItemRequest) UnmarshalJSON(input []byte) error {
    var temp tempCreateFeatureCatalogItemRequest
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

// tempCreateFeatureCatalogItemRequest is a temporary struct used for validating the fields of CreateFeatureCatalogItemRequest.
type tempCreateFeatureCatalogItemRequest  struct {
    Feature *Feature2 `json:"feature"`
}

func (c *tempCreateFeatureCatalogItemRequest) validate() error {
    var errs []string
    if c.Feature == nil {
        errs = append(errs, "required field `feature` is missing for type `Create Feature Catalog Item Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
