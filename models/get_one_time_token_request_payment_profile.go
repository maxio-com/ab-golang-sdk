// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// GetOneTimeTokenRequestPaymentProfile represents a GetOneTimeTokenRequestPaymentProfile struct.
// This is a container for any-of cases.
type GetOneTimeTokenRequestPaymentProfile struct {
    value                                      any
    isGetOneTimeTokenPaymentProfile            bool
    isGetOneTimeTokenBankAccountPaymentProfile bool
}

// String implements the fmt.Stringer interface for GetOneTimeTokenRequestPaymentProfile,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GetOneTimeTokenRequestPaymentProfile) String() string {
    return fmt.Sprintf("%v", g.value)
}

// MarshalJSON implements the json.Marshaler interface for GetOneTimeTokenRequestPaymentProfile.
// It customizes the JSON marshaling process for GetOneTimeTokenRequestPaymentProfile objects.
func (g GetOneTimeTokenRequestPaymentProfile) MarshalJSON() (
    []byte,
    error) {
    if g.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.GetOneTimeTokenRequestPaymentProfileContainer.From*` functions to initialize the GetOneTimeTokenRequestPaymentProfile object.")
    }
    return json.Marshal(g.toMap())
}

// toMap converts the GetOneTimeTokenRequestPaymentProfile object to a map representation for JSON marshaling.
func (g *GetOneTimeTokenRequestPaymentProfile) toMap() any {
    switch obj := g.value.(type) {
    case *GetOneTimeTokenPaymentProfile:
        return obj.toMap()
    case *GetOneTimeTokenBankAccountPaymentProfile:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetOneTimeTokenRequestPaymentProfile.
// It customizes the JSON unmarshaling process for GetOneTimeTokenRequestPaymentProfile objects.
func (g *GetOneTimeTokenRequestPaymentProfile) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(&GetOneTimeTokenPaymentProfile{}, false, &g.isGetOneTimeTokenPaymentProfile),
        NewTypeHolder(&GetOneTimeTokenBankAccountPaymentProfile{}, false, &g.isGetOneTimeTokenBankAccountPaymentProfile),
    )
    
    g.value = result
    return err
}

func (g *GetOneTimeTokenRequestPaymentProfile) AsGetOneTimeTokenPaymentProfile() (
    *GetOneTimeTokenPaymentProfile,
    bool) {
    if !g.isGetOneTimeTokenPaymentProfile {
        return nil, false
    }
    return g.value.(*GetOneTimeTokenPaymentProfile), true
}

func (g *GetOneTimeTokenRequestPaymentProfile) AsGetOneTimeTokenBankAccountPaymentProfile() (
    *GetOneTimeTokenBankAccountPaymentProfile,
    bool) {
    if !g.isGetOneTimeTokenBankAccountPaymentProfile {
        return nil, false
    }
    return g.value.(*GetOneTimeTokenBankAccountPaymentProfile), true
}

// internalGetOneTimeTokenRequestPaymentProfile represents a getOneTimeTokenRequestPaymentProfile struct.
// This is a container for any-of cases.
type internalGetOneTimeTokenRequestPaymentProfile struct {}

var GetOneTimeTokenRequestPaymentProfileContainer internalGetOneTimeTokenRequestPaymentProfile

// The internalGetOneTimeTokenRequestPaymentProfile instance, wrapping the provided GetOneTimeTokenPaymentProfile value.
func (g *internalGetOneTimeTokenRequestPaymentProfile) FromGetOneTimeTokenPaymentProfile(val GetOneTimeTokenPaymentProfile) GetOneTimeTokenRequestPaymentProfile {
    return GetOneTimeTokenRequestPaymentProfile{value: &val}
}

// The internalGetOneTimeTokenRequestPaymentProfile instance, wrapping the provided GetOneTimeTokenBankAccountPaymentProfile value.
func (g *internalGetOneTimeTokenRequestPaymentProfile) FromGetOneTimeTokenBankAccountPaymentProfile(val GetOneTimeTokenBankAccountPaymentProfile) GetOneTimeTokenRequestPaymentProfile {
    return GetOneTimeTokenRequestPaymentProfile{value: &val}
}
