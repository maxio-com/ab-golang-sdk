package models

import (
    "encoding/json"
)

// SubscriptionProductMigrationRequest represents a SubscriptionProductMigrationRequest struct.
type SubscriptionProductMigrationRequest struct {
    Migration SubscriptionProductMigration `json:"migration"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionProductMigrationRequest.
// It customizes the JSON marshaling process for SubscriptionProductMigrationRequest objects.
func (s *SubscriptionProductMigrationRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionProductMigrationRequest object to a map representation for JSON marshaling.
func (s *SubscriptionProductMigrationRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["migration"] = s.Migration
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionProductMigrationRequest.
// It customizes the JSON unmarshaling process for SubscriptionProductMigrationRequest objects.
func (s *SubscriptionProductMigrationRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Migration SubscriptionProductMigration `json:"migration"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    s.Migration = temp.Migration
    return nil
}
