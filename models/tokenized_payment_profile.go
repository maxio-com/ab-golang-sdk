// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// TokenizedPaymentProfile represents a TokenizedPaymentProfile struct.
type TokenizedPaymentProfile struct {
    Id                   int                    `json:"id"`
    VaultToken           *string                `json:"vault_token,omitempty"`
    GatewayHandle        Optional[string]       `json:"gateway_handle"`
    CustomerVaultToken   Optional[string]       `json:"customer_vault_token"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for TokenizedPaymentProfile,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (t TokenizedPaymentProfile) String() string {
    return fmt.Sprintf(
    	"TokenizedPaymentProfile[Id=%v, VaultToken=%v, GatewayHandle=%v, CustomerVaultToken=%v, AdditionalProperties=%v]",
    	t.Id, t.VaultToken, t.GatewayHandle, t.CustomerVaultToken, t.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for TokenizedPaymentProfile.
// It customizes the JSON marshaling process for TokenizedPaymentProfile objects.
func (t TokenizedPaymentProfile) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(t.AdditionalProperties,
        "id", "vault_token", "gateway_handle", "customer_vault_token"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(t.toMap())
}

// toMap converts the TokenizedPaymentProfile object to a map representation for JSON marshaling.
func (t TokenizedPaymentProfile) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, t.AdditionalProperties)
    structMap["id"] = t.Id
    if t.VaultToken != nil {
        structMap["vault_token"] = t.VaultToken
    }
    if t.GatewayHandle.IsValueSet() {
        if t.GatewayHandle.Value() != nil {
            structMap["gateway_handle"] = t.GatewayHandle.Value()
        } else {
            structMap["gateway_handle"] = nil
        }
    }
    if t.CustomerVaultToken.IsValueSet() {
        if t.CustomerVaultToken.Value() != nil {
            structMap["customer_vault_token"] = t.CustomerVaultToken.Value()
        } else {
            structMap["customer_vault_token"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TokenizedPaymentProfile.
// It customizes the JSON unmarshaling process for TokenizedPaymentProfile objects.
func (t *TokenizedPaymentProfile) UnmarshalJSON(input []byte) error {
    var temp tempTokenizedPaymentProfile
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "vault_token", "gateway_handle", "customer_vault_token")
    if err != nil {
    	return err
    }
    t.AdditionalProperties = additionalProperties
    
    t.Id = *temp.Id
    t.VaultToken = temp.VaultToken
    t.GatewayHandle = temp.GatewayHandle
    t.CustomerVaultToken = temp.CustomerVaultToken
    return nil
}

// tempTokenizedPaymentProfile is a temporary struct used for validating the fields of TokenizedPaymentProfile.
type tempTokenizedPaymentProfile  struct {
    Id                 *int             `json:"id"`
    VaultToken         *string          `json:"vault_token,omitempty"`
    GatewayHandle      Optional[string] `json:"gateway_handle"`
    CustomerVaultToken Optional[string] `json:"customer_vault_token"`
}

func (t *tempTokenizedPaymentProfile) validate() error {
    var errs []string
    if t.Id == nil {
        errs = append(errs, "required field `id` is missing for type `Tokenized Payment Profile`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
