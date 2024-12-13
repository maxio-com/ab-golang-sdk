/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SubscriptionMigrationPreviewResponse represents a SubscriptionMigrationPreviewResponse struct.
type SubscriptionMigrationPreviewResponse struct {
    Migration            SubscriptionMigrationPreview `json:"migration"`
    AdditionalProperties map[string]interface{}       `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionMigrationPreviewResponse.
// It customizes the JSON marshaling process for SubscriptionMigrationPreviewResponse objects.
func (s SubscriptionMigrationPreviewResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "migration"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionMigrationPreviewResponse object to a map representation for JSON marshaling.
func (s SubscriptionMigrationPreviewResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["migration"] = s.Migration.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionMigrationPreviewResponse.
// It customizes the JSON unmarshaling process for SubscriptionMigrationPreviewResponse objects.
func (s *SubscriptionMigrationPreviewResponse) UnmarshalJSON(input []byte) error {
    var temp tempSubscriptionMigrationPreviewResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "migration")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Migration = *temp.Migration
    return nil
}

// tempSubscriptionMigrationPreviewResponse is a temporary struct used for validating the fields of SubscriptionMigrationPreviewResponse.
type tempSubscriptionMigrationPreviewResponse  struct {
    Migration *SubscriptionMigrationPreview `json:"migration"`
}

func (s *tempSubscriptionMigrationPreviewResponse) validate() error {
    var errs []string
    if s.Migration == nil {
        errs = append(errs, "required field `migration` is missing for type `Subscription Migration Preview Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
