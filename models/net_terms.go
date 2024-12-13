/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// NetTerms represents a NetTerms struct.
type NetTerms struct {
    DefaultNetTerms                    *int                   `json:"default_net_terms,omitempty"`
    AutomaticNetTerms                  *int                   `json:"automatic_net_terms,omitempty"`
    RemittanceNetTerms                 *int                   `json:"remittance_net_terms,omitempty"`
    NetTermsOnRemittanceSignupsEnabled *bool                  `json:"net_terms_on_remittance_signups_enabled,omitempty"`
    CustomNetTermsEnabled              *bool                  `json:"custom_net_terms_enabled,omitempty"`
    AdditionalProperties               map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for NetTerms.
// It customizes the JSON marshaling process for NetTerms objects.
func (n NetTerms) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(n.AdditionalProperties,
        "default_net_terms", "automatic_net_terms", "remittance_net_terms", "net_terms_on_remittance_signups_enabled", "custom_net_terms_enabled"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(n.toMap())
}

// toMap converts the NetTerms object to a map representation for JSON marshaling.
func (n NetTerms) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, n.AdditionalProperties)
    if n.DefaultNetTerms != nil {
        structMap["default_net_terms"] = n.DefaultNetTerms
    }
    if n.AutomaticNetTerms != nil {
        structMap["automatic_net_terms"] = n.AutomaticNetTerms
    }
    if n.RemittanceNetTerms != nil {
        structMap["remittance_net_terms"] = n.RemittanceNetTerms
    }
    if n.NetTermsOnRemittanceSignupsEnabled != nil {
        structMap["net_terms_on_remittance_signups_enabled"] = n.NetTermsOnRemittanceSignupsEnabled
    }
    if n.CustomNetTermsEnabled != nil {
        structMap["custom_net_terms_enabled"] = n.CustomNetTermsEnabled
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NetTerms.
// It customizes the JSON unmarshaling process for NetTerms objects.
func (n *NetTerms) UnmarshalJSON(input []byte) error {
    var temp tempNetTerms
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "default_net_terms", "automatic_net_terms", "remittance_net_terms", "net_terms_on_remittance_signups_enabled", "custom_net_terms_enabled")
    if err != nil {
    	return err
    }
    n.AdditionalProperties = additionalProperties
    
    n.DefaultNetTerms = temp.DefaultNetTerms
    n.AutomaticNetTerms = temp.AutomaticNetTerms
    n.RemittanceNetTerms = temp.RemittanceNetTerms
    n.NetTermsOnRemittanceSignupsEnabled = temp.NetTermsOnRemittanceSignupsEnabled
    n.CustomNetTermsEnabled = temp.CustomNetTermsEnabled
    return nil
}

// tempNetTerms is a temporary struct used for validating the fields of NetTerms.
type tempNetTerms  struct {
    DefaultNetTerms                    *int  `json:"default_net_terms,omitempty"`
    AutomaticNetTerms                  *int  `json:"automatic_net_terms,omitempty"`
    RemittanceNetTerms                 *int  `json:"remittance_net_terms,omitempty"`
    NetTermsOnRemittanceSignupsEnabled *bool `json:"net_terms_on_remittance_signups_enabled,omitempty"`
    CustomNetTermsEnabled              *bool `json:"custom_net_terms_enabled,omitempty"`
}
