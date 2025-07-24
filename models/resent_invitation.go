// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// ResentInvitation represents a ResentInvitation struct.
type ResentInvitation struct {
    LastSentAt           *string                `json:"last_sent_at,omitempty"`            // Deprecated
    LastAcceptedAt       *string                `json:"last_accepted_at,omitempty"`        // Deprecated
    SendInviteLinkText   *string                `json:"send_invite_link_text,omitempty"`
    UninvitedCount       *int                   `json:"uninvited_count,omitempty"`
    LastInviteSentAt     *time.Time             `json:"last_invite_sent_at,omitempty"`
    LastInviteAcceptedAt *time.Time             `json:"last_invite_accepted_at,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResentInvitation,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResentInvitation) String() string {
    return fmt.Sprintf(
    	"ResentInvitation[LastSentAt=%v, LastAcceptedAt=%v, SendInviteLinkText=%v, UninvitedCount=%v, LastInviteSentAt=%v, LastInviteAcceptedAt=%v, AdditionalProperties=%v]",
    	r.LastSentAt, r.LastAcceptedAt, r.SendInviteLinkText, r.UninvitedCount, r.LastInviteSentAt, r.LastInviteAcceptedAt, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResentInvitation.
// It customizes the JSON marshaling process for ResentInvitation objects.
func (r ResentInvitation) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "last_sent_at", "last_accepted_at", "send_invite_link_text", "uninvited_count", "last_invite_sent_at", "last_invite_accepted_at"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResentInvitation object to a map representation for JSON marshaling.
func (r ResentInvitation) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.LastSentAt != nil {
        structMap["last_sent_at"] = r.LastSentAt
    }
    if r.LastAcceptedAt != nil {
        structMap["last_accepted_at"] = r.LastAcceptedAt
    }
    if r.SendInviteLinkText != nil {
        structMap["send_invite_link_text"] = r.SendInviteLinkText
    }
    if r.UninvitedCount != nil {
        structMap["uninvited_count"] = r.UninvitedCount
    }
    if r.LastInviteSentAt != nil {
        structMap["last_invite_sent_at"] = r.LastInviteSentAt.Format(time.RFC3339)
    }
    if r.LastInviteAcceptedAt != nil {
        structMap["last_invite_accepted_at"] = r.LastInviteAcceptedAt.Format(time.RFC3339)
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResentInvitation.
// It customizes the JSON unmarshaling process for ResentInvitation objects.
func (r *ResentInvitation) UnmarshalJSON(input []byte) error {
    var temp tempResentInvitation
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "last_sent_at", "last_accepted_at", "send_invite_link_text", "uninvited_count", "last_invite_sent_at", "last_invite_accepted_at")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.LastSentAt = temp.LastSentAt
    r.LastAcceptedAt = temp.LastAcceptedAt
    r.SendInviteLinkText = temp.SendInviteLinkText
    r.UninvitedCount = temp.UninvitedCount
    if temp.LastInviteSentAt != nil {
        LastInviteSentAtVal, err := time.Parse(time.RFC3339, *temp.LastInviteSentAt)
        if err != nil {
            log.Fatalf("Cannot Parse last_invite_sent_at as % s format.", time.RFC3339)
        }
        r.LastInviteSentAt = &LastInviteSentAtVal
    }
    if temp.LastInviteAcceptedAt != nil {
        LastInviteAcceptedAtVal, err := time.Parse(time.RFC3339, *temp.LastInviteAcceptedAt)
        if err != nil {
            log.Fatalf("Cannot Parse last_invite_accepted_at as % s format.", time.RFC3339)
        }
        r.LastInviteAcceptedAt = &LastInviteAcceptedAtVal
    }
    return nil
}

// tempResentInvitation is a temporary struct used for validating the fields of ResentInvitation.
type tempResentInvitation  struct {
    LastSentAt           *string `json:"last_sent_at,omitempty"`
    LastAcceptedAt       *string `json:"last_accepted_at,omitempty"`
    SendInviteLinkText   *string `json:"send_invite_link_text,omitempty"`
    UninvitedCount       *int    `json:"uninvited_count,omitempty"`
    LastInviteSentAt     *string `json:"last_invite_sent_at,omitempty"`
    LastInviteAcceptedAt *string `json:"last_invite_accepted_at,omitempty"`
}
