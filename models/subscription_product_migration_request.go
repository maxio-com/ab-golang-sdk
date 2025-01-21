/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// SubscriptionProductMigrationRequest represents a SubscriptionProductMigrationRequest struct.
type SubscriptionProductMigrationRequest struct {
    Migration            SubscriptionProductMigration `json:"migration"`
    AdditionalProperties map[string]interface{}       `json:"_"`
}

// String implements the fmt.Stringer interface for SubscriptionProductMigrationRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionProductMigrationRequest) String() string {
    return fmt.Sprintf(
    	"SubscriptionProductMigrationRequest[Migration=%v, AdditionalProperties=%v]",
    	s.Migration, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionProductMigrationRequest.
// It customizes the JSON marshaling process for SubscriptionProductMigrationRequest objects.
func (s SubscriptionProductMigrationRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "migration"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionProductMigrationRequest object to a map representation for JSON marshaling.
func (s SubscriptionProductMigrationRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["migration"] = s.Migration.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionProductMigrationRequest.
// It customizes the JSON unmarshaling process for SubscriptionProductMigrationRequest objects.
func (s *SubscriptionProductMigrationRequest) UnmarshalJSON(input []byte) error {
    var temp tempSubscriptionProductMigrationRequest
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

// tempSubscriptionProductMigrationRequest is a temporary struct used for validating the fields of SubscriptionProductMigrationRequest.
type tempSubscriptionProductMigrationRequest  struct {
    Migration *SubscriptionProductMigration `json:"migration"`
}

func (s *tempSubscriptionProductMigrationRequest) validate() error {
    var errs []string
    if s.Migration == nil {
        errs = append(errs, "required field `migration` is missing for type `Subscription Product Migration Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
