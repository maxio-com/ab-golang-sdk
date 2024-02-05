package models

import (
    "encoding/json"
)

// OfferSignupPage represents a OfferSignupPage struct.
type OfferSignupPage struct {
    Id           *int    `json:"id,omitempty"`
    Nickname     *string `json:"nickname,omitempty"`
    Enabled      *bool   `json:"enabled,omitempty"`
    ReturnUrl    *string `json:"return_url,omitempty"`
    ReturnParams *string `json:"return_params,omitempty"`
    Url          *string `json:"url,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for OfferSignupPage.
// It customizes the JSON marshaling process for OfferSignupPage objects.
func (o *OfferSignupPage) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OfferSignupPage object to a map representation for JSON marshaling.
func (o *OfferSignupPage) toMap() map[string]any {
    structMap := make(map[string]any)
    if o.Id != nil {
        structMap["id"] = o.Id
    }
    if o.Nickname != nil {
        structMap["nickname"] = o.Nickname
    }
    if o.Enabled != nil {
        structMap["enabled"] = o.Enabled
    }
    if o.ReturnUrl != nil {
        structMap["return_url"] = o.ReturnUrl
    }
    if o.ReturnParams != nil {
        structMap["return_params"] = o.ReturnParams
    }
    if o.Url != nil {
        structMap["url"] = o.Url
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OfferSignupPage.
// It customizes the JSON unmarshaling process for OfferSignupPage objects.
func (o *OfferSignupPage) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id           *int    `json:"id,omitempty"`
        Nickname     *string `json:"nickname,omitempty"`
        Enabled      *bool   `json:"enabled,omitempty"`
        ReturnUrl    *string `json:"return_url,omitempty"`
        ReturnParams *string `json:"return_params,omitempty"`
        Url          *string `json:"url,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    o.Id = temp.Id
    o.Nickname = temp.Nickname
    o.Enabled = temp.Enabled
    o.ReturnUrl = temp.ReturnUrl
    o.ReturnParams = temp.ReturnParams
    o.Url = temp.Url
    return nil
}
