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

// PortalManagementLink represents a PortalManagementLink struct.
type PortalManagementLink struct {
    Url                  *string                `json:"url,omitempty"`
    FetchCount           *int                   `json:"fetch_count,omitempty"`
    CreatedAt            *time.Time             `json:"created_at,omitempty"`
    NewLinkAvailableAt   *time.Time             `json:"new_link_available_at,omitempty"`
    ExpiresAt            *time.Time             `json:"expires_at,omitempty"`
    LastInviteSentAt     Optional[time.Time]    `json:"last_invite_sent_at"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for PortalManagementLink,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PortalManagementLink) String() string {
    return fmt.Sprintf(
    	"PortalManagementLink[Url=%v, FetchCount=%v, CreatedAt=%v, NewLinkAvailableAt=%v, ExpiresAt=%v, LastInviteSentAt=%v, AdditionalProperties=%v]",
    	p.Url, p.FetchCount, p.CreatedAt, p.NewLinkAvailableAt, p.ExpiresAt, p.LastInviteSentAt, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for PortalManagementLink.
// It customizes the JSON marshaling process for PortalManagementLink objects.
func (p PortalManagementLink) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(p.AdditionalProperties,
        "url", "fetch_count", "created_at", "new_link_available_at", "expires_at", "last_invite_sent_at"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(p.toMap())
}

// toMap converts the PortalManagementLink object to a map representation for JSON marshaling.
func (p PortalManagementLink) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, p.AdditionalProperties)
    if p.Url != nil {
        structMap["url"] = p.Url
    }
    if p.FetchCount != nil {
        structMap["fetch_count"] = p.FetchCount
    }
    if p.CreatedAt != nil {
        structMap["created_at"] = p.CreatedAt.Format(time.RFC3339)
    }
    if p.NewLinkAvailableAt != nil {
        structMap["new_link_available_at"] = p.NewLinkAvailableAt.Format(time.RFC3339)
    }
    if p.ExpiresAt != nil {
        structMap["expires_at"] = p.ExpiresAt.Format(time.RFC3339)
    }
    if p.LastInviteSentAt.IsValueSet() {
        var LastInviteSentAtVal *string = nil
        if p.LastInviteSentAt.Value() != nil {
            val := p.LastInviteSentAt.Value().Format(time.RFC3339)
            LastInviteSentAtVal = &val
        }
        if p.LastInviteSentAt.Value() != nil {
            structMap["last_invite_sent_at"] = LastInviteSentAtVal
        } else {
            structMap["last_invite_sent_at"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PortalManagementLink.
// It customizes the JSON unmarshaling process for PortalManagementLink objects.
func (p *PortalManagementLink) UnmarshalJSON(input []byte) error {
    var temp tempPortalManagementLink
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "url", "fetch_count", "created_at", "new_link_available_at", "expires_at", "last_invite_sent_at")
    if err != nil {
    	return err
    }
    p.AdditionalProperties = additionalProperties
    
    p.Url = temp.Url
    p.FetchCount = temp.FetchCount
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        p.CreatedAt = &CreatedAtVal
    }
    if temp.NewLinkAvailableAt != nil {
        NewLinkAvailableAtVal, err := time.Parse(time.RFC3339, *temp.NewLinkAvailableAt)
        if err != nil {
            log.Fatalf("Cannot Parse new_link_available_at as % s format.", time.RFC3339)
        }
        p.NewLinkAvailableAt = &NewLinkAvailableAtVal
    }
    if temp.ExpiresAt != nil {
        ExpiresAtVal, err := time.Parse(time.RFC3339, *temp.ExpiresAt)
        if err != nil {
            log.Fatalf("Cannot Parse expires_at as % s format.", time.RFC3339)
        }
        p.ExpiresAt = &ExpiresAtVal
    }
    p.LastInviteSentAt.ShouldSetValue(temp.LastInviteSentAt.IsValueSet())
    if temp.LastInviteSentAt.Value() != nil {
        LastInviteSentAtVal, err := time.Parse(time.RFC3339, (*temp.LastInviteSentAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse last_invite_sent_at as % s format.", time.RFC3339)
        }
        p.LastInviteSentAt.SetValue(&LastInviteSentAtVal)
    }
    return nil
}

// tempPortalManagementLink is a temporary struct used for validating the fields of PortalManagementLink.
type tempPortalManagementLink  struct {
    Url                *string          `json:"url,omitempty"`
    FetchCount         *int             `json:"fetch_count,omitempty"`
    CreatedAt          *string          `json:"created_at,omitempty"`
    NewLinkAvailableAt *string          `json:"new_link_available_at,omitempty"`
    ExpiresAt          *string          `json:"expires_at,omitempty"`
    LastInviteSentAt   Optional[string] `json:"last_invite_sent_at"`
}
