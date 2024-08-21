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

// SubscriptionMigrationPreviewRequest represents a SubscriptionMigrationPreviewRequest struct.
type SubscriptionMigrationPreviewRequest struct {
    Migration            SubscriptionMigrationPreviewOptions `json:"migration"`
    AdditionalProperties map[string]any                      `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionMigrationPreviewRequest.
// It customizes the JSON marshaling process for SubscriptionMigrationPreviewRequest objects.
func (s SubscriptionMigrationPreviewRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionMigrationPreviewRequest object to a map representation for JSON marshaling.
func (s SubscriptionMigrationPreviewRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["migration"] = s.Migration.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionMigrationPreviewRequest.
// It customizes the JSON unmarshaling process for SubscriptionMigrationPreviewRequest objects.
func (s *SubscriptionMigrationPreviewRequest) UnmarshalJSON(input []byte) error {
    var temp tempSubscriptionMigrationPreviewRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "migration")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Migration = *temp.Migration
    return nil
}

// tempSubscriptionMigrationPreviewRequest is a temporary struct used for validating the fields of SubscriptionMigrationPreviewRequest.
type tempSubscriptionMigrationPreviewRequest  struct {
    Migration *SubscriptionMigrationPreviewOptions `json:"migration"`
}

func (s *tempSubscriptionMigrationPreviewRequest) validate() error {
    var errs []string
    if s.Migration == nil {
        errs = append(errs, "required field `migration` is missing for type `Subscription Migration Preview Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
