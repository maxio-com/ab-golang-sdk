package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SubscriptionProductMigrationRequest represents a SubscriptionProductMigrationRequest struct.
type SubscriptionProductMigrationRequest struct {
    Migration            SubscriptionProductMigration `json:"migration"`
    AdditionalProperties map[string]any               `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionProductMigrationRequest.
// It customizes the JSON marshaling process for SubscriptionProductMigrationRequest objects.
func (s SubscriptionProductMigrationRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionProductMigrationRequest object to a map representation for JSON marshaling.
func (s SubscriptionProductMigrationRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["migration"] = s.Migration.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionProductMigrationRequest.
// It customizes the JSON unmarshaling process for SubscriptionProductMigrationRequest objects.
func (s *SubscriptionProductMigrationRequest) UnmarshalJSON(input []byte) error {
    var temp subscriptionProductMigrationRequest
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

// TODO
type subscriptionProductMigrationRequest  struct {
    Migration *SubscriptionProductMigration `json:"migration"`
}

func (s *subscriptionProductMigrationRequest) validate() error {
    var errs []string
    if s.Migration == nil {
        errs = append(errs, "required field `migration` is missing for type `Subscription Product Migration Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
