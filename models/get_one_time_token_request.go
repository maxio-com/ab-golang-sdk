// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// GetOneTimeTokenRequest represents a GetOneTimeTokenRequest struct.
type GetOneTimeTokenRequest struct {
    PaymentProfile       GetOneTimeTokenPaymentProfile `json:"payment_profile"`
    AdditionalProperties map[string]interface{}        `json:"_"`
}

// String implements the fmt.Stringer interface for GetOneTimeTokenRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GetOneTimeTokenRequest) String() string {
    return fmt.Sprintf(
    	"GetOneTimeTokenRequest[PaymentProfile=%v, AdditionalProperties=%v]",
    	g.PaymentProfile, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for GetOneTimeTokenRequest.
// It customizes the JSON marshaling process for GetOneTimeTokenRequest objects.
func (g GetOneTimeTokenRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(g.AdditionalProperties,
        "payment_profile"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(g.toMap())
}

// toMap converts the GetOneTimeTokenRequest object to a map representation for JSON marshaling.
func (g GetOneTimeTokenRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, g.AdditionalProperties)
    structMap["payment_profile"] = g.PaymentProfile.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetOneTimeTokenRequest.
// It customizes the JSON unmarshaling process for GetOneTimeTokenRequest objects.
func (g *GetOneTimeTokenRequest) UnmarshalJSON(input []byte) error {
    var temp tempGetOneTimeTokenRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "payment_profile")
    if err != nil {
    	return err
    }
    g.AdditionalProperties = additionalProperties
    
    g.PaymentProfile = *temp.PaymentProfile
    return nil
}

// tempGetOneTimeTokenRequest is a temporary struct used for validating the fields of GetOneTimeTokenRequest.
type tempGetOneTimeTokenRequest  struct {
    PaymentProfile *GetOneTimeTokenPaymentProfile `json:"payment_profile"`
}

func (g *tempGetOneTimeTokenRequest) validate() error {
    var errs []string
    if g.PaymentProfile == nil {
        errs = append(errs, "required field `payment_profile` is missing for type `Get One Time Token Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
