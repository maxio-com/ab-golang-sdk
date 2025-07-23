// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// RevokedInvitation represents a RevokedInvitation struct.
type RevokedInvitation struct {
    LastSentAt           *string                `json:"last_sent_at,omitempty"`
    LastAcceptedAt       *string                `json:"last_accepted_at,omitempty"`
    UninvitedCount       *int                   `json:"uninvited_count,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for RevokedInvitation,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RevokedInvitation) String() string {
    return fmt.Sprintf(
    	"RevokedInvitation[LastSentAt=%v, LastAcceptedAt=%v, UninvitedCount=%v, AdditionalProperties=%v]",
    	r.LastSentAt, r.LastAcceptedAt, r.UninvitedCount, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for RevokedInvitation.
// It customizes the JSON marshaling process for RevokedInvitation objects.
func (r RevokedInvitation) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "last_sent_at", "last_accepted_at", "uninvited_count"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the RevokedInvitation object to a map representation for JSON marshaling.
func (r RevokedInvitation) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.LastSentAt != nil {
        structMap["last_sent_at"] = r.LastSentAt
    }
    if r.LastAcceptedAt != nil {
        structMap["last_accepted_at"] = r.LastAcceptedAt
    }
    if r.UninvitedCount != nil {
        structMap["uninvited_count"] = r.UninvitedCount
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RevokedInvitation.
// It customizes the JSON unmarshaling process for RevokedInvitation objects.
func (r *RevokedInvitation) UnmarshalJSON(input []byte) error {
    var temp tempRevokedInvitation
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "last_sent_at", "last_accepted_at", "uninvited_count")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.LastSentAt = temp.LastSentAt
    r.LastAcceptedAt = temp.LastAcceptedAt
    r.UninvitedCount = temp.UninvitedCount
    return nil
}

// tempRevokedInvitation is a temporary struct used for validating the fields of RevokedInvitation.
type tempRevokedInvitation  struct {
    LastSentAt     *string `json:"last_sent_at,omitempty"`
    LastAcceptedAt *string `json:"last_accepted_at,omitempty"`
    UninvitedCount *int    `json:"uninvited_count,omitempty"`
}
