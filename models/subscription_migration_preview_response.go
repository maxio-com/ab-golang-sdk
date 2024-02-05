package models

import (
    "encoding/json"
)

// SubscriptionMigrationPreviewResponse represents a SubscriptionMigrationPreviewResponse struct.
type SubscriptionMigrationPreviewResponse struct {
    Migration SubscriptionMigrationPreview `json:"migration"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionMigrationPreviewResponse.
// It customizes the JSON marshaling process for SubscriptionMigrationPreviewResponse objects.
func (s *SubscriptionMigrationPreviewResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionMigrationPreviewResponse object to a map representation for JSON marshaling.
func (s *SubscriptionMigrationPreviewResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["migration"] = s.Migration.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionMigrationPreviewResponse.
// It customizes the JSON unmarshaling process for SubscriptionMigrationPreviewResponse objects.
func (s *SubscriptionMigrationPreviewResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Migration SubscriptionMigrationPreview `json:"migration"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    s.Migration = temp.Migration
    return nil
}
