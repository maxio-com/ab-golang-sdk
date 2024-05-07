package models

import (
    "encoding/json"
    "log"
    "time"
)

// SubscriptionNote represents a SubscriptionNote struct.
type SubscriptionNote struct {
    Id                   *int           `json:"id,omitempty"`
    Body                 *string        `json:"body,omitempty"`
    SubscriptionId       *int           `json:"subscription_id,omitempty"`
    CreatedAt            *time.Time     `json:"created_at,omitempty"`
    UpdatedAt            *time.Time     `json:"updated_at,omitempty"`
    Sticky               *bool          `json:"sticky,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionNote.
// It customizes the JSON marshaling process for SubscriptionNote objects.
func (s SubscriptionNote) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionNote object to a map representation for JSON marshaling.
func (s SubscriptionNote) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Id != nil {
        structMap["id"] = s.Id
    }
    if s.Body != nil {
        structMap["body"] = s.Body
    }
    if s.SubscriptionId != nil {
        structMap["subscription_id"] = s.SubscriptionId
    }
    if s.CreatedAt != nil {
        structMap["created_at"] = s.CreatedAt.Format(time.RFC3339)
    }
    if s.UpdatedAt != nil {
        structMap["updated_at"] = s.UpdatedAt.Format(time.RFC3339)
    }
    if s.Sticky != nil {
        structMap["sticky"] = s.Sticky
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionNote.
// It customizes the JSON unmarshaling process for SubscriptionNote objects.
func (s *SubscriptionNote) UnmarshalJSON(input []byte) error {
    var temp subscriptionNote
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "id", "body", "subscription_id", "created_at", "updated_at", "sticky")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Id = temp.Id
    s.Body = temp.Body
    s.SubscriptionId = temp.SubscriptionId
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        s.CreatedAt = &CreatedAtVal
    }
    if temp.UpdatedAt != nil {
        UpdatedAtVal, err := time.Parse(time.RFC3339, *temp.UpdatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse updated_at as % s format.", time.RFC3339)
        }
        s.UpdatedAt = &UpdatedAtVal
    }
    s.Sticky = temp.Sticky
    return nil
}

// subscriptionNote is a temporary struct used for validating the fields of SubscriptionNote.
type subscriptionNote  struct {
    Id             *int    `json:"id,omitempty"`
    Body           *string `json:"body,omitempty"`
    SubscriptionId *int    `json:"subscription_id,omitempty"`
    CreatedAt      *string `json:"created_at,omitempty"`
    UpdatedAt      *string `json:"updated_at,omitempty"`
    Sticky         *bool   `json:"sticky,omitempty"`
}
