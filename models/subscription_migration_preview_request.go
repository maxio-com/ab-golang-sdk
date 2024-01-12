package models

import (
    "encoding/json"
)

// SubscriptionMigrationPreviewRequest represents a SubscriptionMigrationPreviewRequest struct.
type SubscriptionMigrationPreviewRequest struct {
    Migration SubscriptionMigrationPreviewOptions `json:"migration"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionMigrationPreviewRequest.
// It customizes the JSON marshaling process for SubscriptionMigrationPreviewRequest objects.
func (s *SubscriptionMigrationPreviewRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionMigrationPreviewRequest object to a map representation for JSON marshaling.
func (s *SubscriptionMigrationPreviewRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["migration"] = s.Migration
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionMigrationPreviewRequest.
// It customizes the JSON unmarshaling process for SubscriptionMigrationPreviewRequest objects.
func (s *SubscriptionMigrationPreviewRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Migration SubscriptionMigrationPreviewOptions `json:"migration"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    s.Migration = temp.Migration
    return nil
}
