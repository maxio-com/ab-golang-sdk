// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// PublicKey represents a PublicKey struct.
type PublicKey struct {
    PublicKey             *string                `json:"public_key,omitempty"`
    RequiresSecurityToken *bool                  `json:"requires_security_token,omitempty"`
    CreatedAt             *time.Time             `json:"created_at,omitempty"`
    AdditionalProperties  map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for PublicKey,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PublicKey) String() string {
    return fmt.Sprintf(
    	"PublicKey[PublicKey=%v, RequiresSecurityToken=%v, CreatedAt=%v, AdditionalProperties=%v]",
    	p.PublicKey, p.RequiresSecurityToken, p.CreatedAt, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for PublicKey.
// It customizes the JSON marshaling process for PublicKey objects.
func (p PublicKey) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(p.AdditionalProperties,
        "public_key", "requires_security_token", "created_at"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(p.toMap())
}

// toMap converts the PublicKey object to a map representation for JSON marshaling.
func (p PublicKey) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, p.AdditionalProperties)
    if p.PublicKey != nil {
        structMap["public_key"] = p.PublicKey
    }
    if p.RequiresSecurityToken != nil {
        structMap["requires_security_token"] = p.RequiresSecurityToken
    }
    if p.CreatedAt != nil {
        structMap["created_at"] = p.CreatedAt.Format(time.RFC3339)
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PublicKey.
// It customizes the JSON unmarshaling process for PublicKey objects.
func (p *PublicKey) UnmarshalJSON(input []byte) error {
    var temp tempPublicKey
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "public_key", "requires_security_token", "created_at")
    if err != nil {
    	return err
    }
    p.AdditionalProperties = additionalProperties
    
    p.PublicKey = temp.PublicKey
    p.RequiresSecurityToken = temp.RequiresSecurityToken
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        p.CreatedAt = &CreatedAtVal
    }
    return nil
}

// tempPublicKey is a temporary struct used for validating the fields of PublicKey.
type tempPublicKey  struct {
    PublicKey             *string `json:"public_key,omitempty"`
    RequiresSecurityToken *bool   `json:"requires_security_token,omitempty"`
    CreatedAt             *string `json:"created_at,omitempty"`
}
