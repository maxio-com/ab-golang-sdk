/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// SubscriptionNote represents a SubscriptionNote struct.
type SubscriptionNote struct {
    Id                   *int                   `json:"id,omitempty"`
    Body                 *string                `json:"body,omitempty"`
    SubscriptionId       *int                   `json:"subscription_id,omitempty"`
    CreatedAt            *time.Time             `json:"created_at,omitempty"`
    UpdatedAt            *time.Time             `json:"updated_at,omitempty"`
    Sticky               *bool                  `json:"sticky,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SubscriptionNote,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionNote) String() string {
    return fmt.Sprintf(
    	"SubscriptionNote[Id=%v, Body=%v, SubscriptionId=%v, CreatedAt=%v, UpdatedAt=%v, Sticky=%v, AdditionalProperties=%v]",
    	s.Id, s.Body, s.SubscriptionId, s.CreatedAt, s.UpdatedAt, s.Sticky, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionNote.
// It customizes the JSON marshaling process for SubscriptionNote objects.
func (s SubscriptionNote) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "id", "body", "subscription_id", "created_at", "updated_at", "sticky"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionNote object to a map representation for JSON marshaling.
func (s SubscriptionNote) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
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
    var temp tempSubscriptionNote
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "body", "subscription_id", "created_at", "updated_at", "sticky")
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

// tempSubscriptionNote is a temporary struct used for validating the fields of SubscriptionNote.
type tempSubscriptionNote  struct {
    Id             *int    `json:"id,omitempty"`
    Body           *string `json:"body,omitempty"`
    SubscriptionId *int    `json:"subscription_id,omitempty"`
    CreatedAt      *string `json:"created_at,omitempty"`
    UpdatedAt      *string `json:"updated_at,omitempty"`
    Sticky         *bool   `json:"sticky,omitempty"`
}
