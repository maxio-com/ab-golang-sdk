package models

import (
	"encoding/json"
)

// PortalManagementLink represents a PortalManagementLink struct.
type PortalManagementLink struct {
	Url                *string          `json:"url,omitempty"`
	FetchCount         *int             `json:"fetch_count,omitempty"`
	CreatedAt          *string          `json:"created_at,omitempty"`
	NewLinkAvailableAt *string          `json:"new_link_available_at,omitempty"`
	ExpiresAt          *string          `json:"expires_at,omitempty"`
	LastInviteSentAt   Optional[string] `json:"last_invite_sent_at"`
}

// MarshalJSON implements the json.Marshaler interface for PortalManagementLink.
// It customizes the JSON marshaling process for PortalManagementLink objects.
func (p *PortalManagementLink) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(p.toMap())
}

// toMap converts the PortalManagementLink object to a map representation for JSON marshaling.
func (p *PortalManagementLink) toMap() map[string]any {
	structMap := make(map[string]any)
	if p.Url != nil {
		structMap["url"] = p.Url
	}
	if p.FetchCount != nil {
		structMap["fetch_count"] = p.FetchCount
	}
	if p.CreatedAt != nil {
		structMap["created_at"] = p.CreatedAt
	}
	if p.NewLinkAvailableAt != nil {
		structMap["new_link_available_at"] = p.NewLinkAvailableAt
	}
	if p.ExpiresAt != nil {
		structMap["expires_at"] = p.ExpiresAt
	}
	if p.LastInviteSentAt.IsValueSet() {
		structMap["last_invite_sent_at"] = p.LastInviteSentAt.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PortalManagementLink.
// It customizes the JSON unmarshaling process for PortalManagementLink objects.
func (p *PortalManagementLink) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Url                *string          `json:"url,omitempty"`
		FetchCount         *int             `json:"fetch_count,omitempty"`
		CreatedAt          *string          `json:"created_at,omitempty"`
		NewLinkAvailableAt *string          `json:"new_link_available_at,omitempty"`
		ExpiresAt          *string          `json:"expires_at,omitempty"`
		LastInviteSentAt   Optional[string] `json:"last_invite_sent_at"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	p.Url = temp.Url
	p.FetchCount = temp.FetchCount
	p.CreatedAt = temp.CreatedAt
	p.NewLinkAvailableAt = temp.NewLinkAvailableAt
	p.ExpiresAt = temp.ExpiresAt
	p.LastInviteSentAt = temp.LastInviteSentAt
	return nil
}
