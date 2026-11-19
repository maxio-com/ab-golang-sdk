// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// FeatureTemplatesListResponse represents a FeatureTemplatesListResponse struct.
type FeatureTemplatesListResponse struct {
    Items                []FeatureTemplate      `json:"items"`
    // Total number of feature templates matching the filters, across all pages.
    TotalCount           int                    `json:"total_count"`
    // Number of archived feature templates matching the filters. Returned as `0` unless the active result set is empty or `status=archived` was requested.
    ArchivedCount        int                    `json:"archived_count"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for FeatureTemplatesListResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (f FeatureTemplatesListResponse) String() string {
    return fmt.Sprintf(
    	"FeatureTemplatesListResponse[Items=%v, TotalCount=%v, ArchivedCount=%v, AdditionalProperties=%v]",
    	f.Items, f.TotalCount, f.ArchivedCount, f.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for FeatureTemplatesListResponse.
// It customizes the JSON marshaling process for FeatureTemplatesListResponse objects.
func (f FeatureTemplatesListResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(f.AdditionalProperties,
        "items", "total_count", "archived_count"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(f.toMap())
}

// toMap converts the FeatureTemplatesListResponse object to a map representation for JSON marshaling.
func (f FeatureTemplatesListResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, f.AdditionalProperties)
    structMap["items"] = f.Items
    structMap["total_count"] = f.TotalCount
    structMap["archived_count"] = f.ArchivedCount
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for FeatureTemplatesListResponse.
// It customizes the JSON unmarshaling process for FeatureTemplatesListResponse objects.
func (f *FeatureTemplatesListResponse) UnmarshalJSON(input []byte) error {
    var temp tempFeatureTemplatesListResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "items", "total_count", "archived_count")
    if err != nil {
    	return err
    }
    f.AdditionalProperties = additionalProperties
    
    f.Items = *temp.Items
    f.TotalCount = *temp.TotalCount
    f.ArchivedCount = *temp.ArchivedCount
    return nil
}

// tempFeatureTemplatesListResponse is a temporary struct used for validating the fields of FeatureTemplatesListResponse.
type tempFeatureTemplatesListResponse  struct {
    Items         *[]FeatureTemplate `json:"items"`
    TotalCount    *int               `json:"total_count"`
    ArchivedCount *int               `json:"archived_count"`
}

func (f *tempFeatureTemplatesListResponse) validate() error {
    var errs []string
    if f.Items == nil {
        errs = append(errs, "required field `items` is missing for type `Feature Templates List Response`")
    }
    if f.TotalCount == nil {
        errs = append(errs, "required field `total_count` is missing for type `Feature Templates List Response`")
    }
    if f.ArchivedCount == nil {
        errs = append(errs, "required field `archived_count` is missing for type `Feature Templates List Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
