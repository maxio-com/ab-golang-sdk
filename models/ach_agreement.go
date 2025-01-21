/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// ACHAgreement represents a ACHAgreement struct.
// (Optional) If passed, the proof of the authorized ACH agreement terms will be persisted.
type ACHAgreement struct {
    // (Required when providing ACH agreement params) The ACH authorization agreement terms.
    AgreementTerms       *string                `json:"agreement_terms,omitempty"`
    // (Required when providing ACH agreement params) The first name of the person authorizing the ACH agreement.
    AuthorizerFirstName  *string                `json:"authorizer_first_name,omitempty"`
    // (Required when providing ACH agreement params) The last name of the person authorizing the ACH agreement.
    AuthorizerLastName   *string                `json:"authorizer_last_name,omitempty"`
    // (Required when providing ACH agreement params) The IP address of the person authorizing the ACH agreement.
    IpAddress            *string                `json:"ip_address,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ACHAgreement,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a ACHAgreement) String() string {
    return fmt.Sprintf(
    	"ACHAgreement[AgreementTerms=%v, AuthorizerFirstName=%v, AuthorizerLastName=%v, IpAddress=%v, AdditionalProperties=%v]",
    	a.AgreementTerms, a.AuthorizerFirstName, a.AuthorizerLastName, a.IpAddress, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ACHAgreement.
// It customizes the JSON marshaling process for ACHAgreement objects.
func (a ACHAgreement) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "agreement_terms", "authorizer_first_name", "authorizer_last_name", "ip_address"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the ACHAgreement object to a map representation for JSON marshaling.
func (a ACHAgreement) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.AgreementTerms != nil {
        structMap["agreement_terms"] = a.AgreementTerms
    }
    if a.AuthorizerFirstName != nil {
        structMap["authorizer_first_name"] = a.AuthorizerFirstName
    }
    if a.AuthorizerLastName != nil {
        structMap["authorizer_last_name"] = a.AuthorizerLastName
    }
    if a.IpAddress != nil {
        structMap["ip_address"] = a.IpAddress
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ACHAgreement.
// It customizes the JSON unmarshaling process for ACHAgreement objects.
func (a *ACHAgreement) UnmarshalJSON(input []byte) error {
    var temp tempACHAgreement
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "agreement_terms", "authorizer_first_name", "authorizer_last_name", "ip_address")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.AgreementTerms = temp.AgreementTerms
    a.AuthorizerFirstName = temp.AuthorizerFirstName
    a.AuthorizerLastName = temp.AuthorizerLastName
    a.IpAddress = temp.IpAddress
    return nil
}

// tempACHAgreement is a temporary struct used for validating the fields of ACHAgreement.
type tempACHAgreement  struct {
    AgreementTerms      *string `json:"agreement_terms,omitempty"`
    AuthorizerFirstName *string `json:"authorizer_first_name,omitempty"`
    AuthorizerLastName  *string `json:"authorizer_last_name,omitempty"`
    IpAddress           *string `json:"ip_address,omitempty"`
}
