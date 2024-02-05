package models

import (
    "encoding/json"
)

// PublicKey represents a PublicKey struct.
type PublicKey struct {
    PublicKey             *string `json:"public_key,omitempty"`
    RequiresSecurityToken *bool   `json:"requires_security_token,omitempty"`
    CreatedAt             *string `json:"created_at,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for PublicKey.
// It customizes the JSON marshaling process for PublicKey objects.
func (p *PublicKey) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PublicKey object to a map representation for JSON marshaling.
func (p *PublicKey) toMap() map[string]any {
    structMap := make(map[string]any)
    if p.PublicKey != nil {
        structMap["public_key"] = p.PublicKey
    }
    if p.RequiresSecurityToken != nil {
        structMap["requires_security_token"] = p.RequiresSecurityToken
    }
    if p.CreatedAt != nil {
        structMap["created_at"] = p.CreatedAt
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PublicKey.
// It customizes the JSON unmarshaling process for PublicKey objects.
func (p *PublicKey) UnmarshalJSON(input []byte) error {
    temp := &struct {
        PublicKey             *string `json:"public_key,omitempty"`
        RequiresSecurityToken *bool   `json:"requires_security_token,omitempty"`
        CreatedAt             *string `json:"created_at,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    p.PublicKey = temp.PublicKey
    p.RequiresSecurityToken = temp.RequiresSecurityToken
    p.CreatedAt = temp.CreatedAt
    return nil
}
