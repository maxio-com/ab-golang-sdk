/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// PublicSignupPage represents a PublicSignupPage struct.
type PublicSignupPage struct {
    // The id of the signup page (public_signup_pages only)
    Id                   *int                   `json:"id,omitempty"`
    // The url to which a customer will be returned after a successful signup (public_signup_pages only)
    ReturnUrl            Optional[string]       `json:"return_url"`
    // The params to be appended to the return_url (public_signup_pages only)
    ReturnParams         Optional[string]       `json:"return_params"`
    // The url where the signup page can be viewed (public_signup_pages only)
    Url                  *string                `json:"url,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for PublicSignupPage.
// It customizes the JSON marshaling process for PublicSignupPage objects.
func (p PublicSignupPage) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(p.AdditionalProperties,
        "id", "return_url", "return_params", "url"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(p.toMap())
}

// toMap converts the PublicSignupPage object to a map representation for JSON marshaling.
func (p PublicSignupPage) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, p.AdditionalProperties)
    if p.Id != nil {
        structMap["id"] = p.Id
    }
    if p.ReturnUrl.IsValueSet() {
        if p.ReturnUrl.Value() != nil {
            structMap["return_url"] = p.ReturnUrl.Value()
        } else {
            structMap["return_url"] = nil
        }
    }
    if p.ReturnParams.IsValueSet() {
        if p.ReturnParams.Value() != nil {
            structMap["return_params"] = p.ReturnParams.Value()
        } else {
            structMap["return_params"] = nil
        }
    }
    if p.Url != nil {
        structMap["url"] = p.Url
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PublicSignupPage.
// It customizes the JSON unmarshaling process for PublicSignupPage objects.
func (p *PublicSignupPage) UnmarshalJSON(input []byte) error {
    var temp tempPublicSignupPage
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "return_url", "return_params", "url")
    if err != nil {
    	return err
    }
    p.AdditionalProperties = additionalProperties
    
    p.Id = temp.Id
    p.ReturnUrl = temp.ReturnUrl
    p.ReturnParams = temp.ReturnParams
    p.Url = temp.Url
    return nil
}

// tempPublicSignupPage is a temporary struct used for validating the fields of PublicSignupPage.
type tempPublicSignupPage  struct {
    Id           *int             `json:"id,omitempty"`
    ReturnUrl    Optional[string] `json:"return_url"`
    ReturnParams Optional[string] `json:"return_params"`
    Url          *string          `json:"url,omitempty"`
}
