// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// FeatureCatalogItemsListResponse represents a FeatureCatalogItemsListResponse struct.
type FeatureCatalogItemsListResponse struct {
    Features             []FeatureCatalogItem   `json:"features"`
    // The number of subscriptions on this product/component that would be affected if a feature catalog item change were propagated with `propagate_to_subscriptions=true`.
    SubscriptionsCount   int                    `json:"subscriptions_count"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for FeatureCatalogItemsListResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (f FeatureCatalogItemsListResponse) String() string {
    return fmt.Sprintf(
    	"FeatureCatalogItemsListResponse[Features=%v, SubscriptionsCount=%v, AdditionalProperties=%v]",
    	f.Features, f.SubscriptionsCount, f.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for FeatureCatalogItemsListResponse.
// It customizes the JSON marshaling process for FeatureCatalogItemsListResponse objects.
func (f FeatureCatalogItemsListResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(f.AdditionalProperties,
        "features", "subscriptions_count"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(f.toMap())
}

// toMap converts the FeatureCatalogItemsListResponse object to a map representation for JSON marshaling.
func (f FeatureCatalogItemsListResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, f.AdditionalProperties)
    structMap["features"] = f.Features
    structMap["subscriptions_count"] = f.SubscriptionsCount
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for FeatureCatalogItemsListResponse.
// It customizes the JSON unmarshaling process for FeatureCatalogItemsListResponse objects.
func (f *FeatureCatalogItemsListResponse) UnmarshalJSON(input []byte) error {
    var temp tempFeatureCatalogItemsListResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "features", "subscriptions_count")
    if err != nil {
    	return err
    }
    f.AdditionalProperties = additionalProperties
    
    f.Features = *temp.Features
    f.SubscriptionsCount = *temp.SubscriptionsCount
    return nil
}

// tempFeatureCatalogItemsListResponse is a temporary struct used for validating the fields of FeatureCatalogItemsListResponse.
type tempFeatureCatalogItemsListResponse  struct {
    Features           *[]FeatureCatalogItem `json:"features"`
    SubscriptionsCount *int                  `json:"subscriptions_count"`
}

func (f *tempFeatureCatalogItemsListResponse) validate() error {
    var errs []string
    if f.Features == nil {
        errs = append(errs, "required field `features` is missing for type `Feature Catalog Items List Response`")
    }
    if f.SubscriptionsCount == nil {
        errs = append(errs, "required field `subscriptions_count` is missing for type `Feature Catalog Items List Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
