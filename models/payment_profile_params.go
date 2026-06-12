// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// PaymentProfileParams represents a PaymentProfileParams struct.
// PCI-safe cardholder fields only. Full card numbers, CVV, and billing address are never included.
type PaymentProfileParams struct {
    FirstName            *string                `json:"first_name,omitempty"`
    LastName             *string                `json:"last_name,omitempty"`
    CardType             *string                `json:"card_type,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for PaymentProfileParams,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PaymentProfileParams) String() string {
    return fmt.Sprintf(
    	"PaymentProfileParams[FirstName=%v, LastName=%v, CardType=%v, AdditionalProperties=%v]",
    	p.FirstName, p.LastName, p.CardType, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for PaymentProfileParams.
// It customizes the JSON marshaling process for PaymentProfileParams objects.
func (p PaymentProfileParams) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(p.AdditionalProperties,
        "first_name", "last_name", "card_type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(p.toMap())
}

// toMap converts the PaymentProfileParams object to a map representation for JSON marshaling.
func (p PaymentProfileParams) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, p.AdditionalProperties)
    if p.FirstName != nil {
        structMap["first_name"] = p.FirstName
    }
    if p.LastName != nil {
        structMap["last_name"] = p.LastName
    }
    if p.CardType != nil {
        structMap["card_type"] = p.CardType
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PaymentProfileParams.
// It customizes the JSON unmarshaling process for PaymentProfileParams objects.
func (p *PaymentProfileParams) UnmarshalJSON(input []byte) error {
    var temp tempPaymentProfileParams
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "first_name", "last_name", "card_type")
    if err != nil {
    	return err
    }
    p.AdditionalProperties = additionalProperties
    
    p.FirstName = temp.FirstName
    p.LastName = temp.LastName
    p.CardType = temp.CardType
    return nil
}

// tempPaymentProfileParams is a temporary struct used for validating the fields of PaymentProfileParams.
type tempPaymentProfileParams  struct {
    FirstName *string `json:"first_name,omitempty"`
    LastName  *string `json:"last_name,omitempty"`
    CardType  *string `json:"card_type,omitempty"`
}
