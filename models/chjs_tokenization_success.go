// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ChjsTokenizationSuccess represents a ChjsTokenizationSuccess struct.
type ChjsTokenizationSuccess struct {
    PaymentProfile       TokenizedPaymentProfile `json:"payment_profile"`
    GatewayCustomerId    Optional[int]           `json:"gateway_customer_id"`
    AdditionalProperties map[string]interface{}  `json:"_"`
}

// String implements the fmt.Stringer interface for ChjsTokenizationSuccess,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ChjsTokenizationSuccess) String() string {
    return fmt.Sprintf(
    	"ChjsTokenizationSuccess[PaymentProfile=%v, GatewayCustomerId=%v, AdditionalProperties=%v]",
    	c.PaymentProfile, c.GatewayCustomerId, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ChjsTokenizationSuccess.
// It customizes the JSON marshaling process for ChjsTokenizationSuccess objects.
func (c ChjsTokenizationSuccess) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "payment_profile", "gateway_customer_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ChjsTokenizationSuccess object to a map representation for JSON marshaling.
func (c ChjsTokenizationSuccess) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["payment_profile"] = c.PaymentProfile.toMap()
    if c.GatewayCustomerId.IsValueSet() {
        if c.GatewayCustomerId.Value() != nil {
            structMap["gateway_customer_id"] = c.GatewayCustomerId.Value()
        } else {
            structMap["gateway_customer_id"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ChjsTokenizationSuccess.
// It customizes the JSON unmarshaling process for ChjsTokenizationSuccess objects.
func (c *ChjsTokenizationSuccess) UnmarshalJSON(input []byte) error {
    var temp tempChjsTokenizationSuccess
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "payment_profile", "gateway_customer_id")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.PaymentProfile = *temp.PaymentProfile
    c.GatewayCustomerId = temp.GatewayCustomerId
    return nil
}

// tempChjsTokenizationSuccess is a temporary struct used for validating the fields of ChjsTokenizationSuccess.
type tempChjsTokenizationSuccess  struct {
    PaymentProfile    *TokenizedPaymentProfile `json:"payment_profile"`
    GatewayCustomerId Optional[int]            `json:"gateway_customer_id"`
}

func (c *tempChjsTokenizationSuccess) validate() error {
    var errs []string
    if c.PaymentProfile == nil {
        errs = append(errs, "required field `payment_profile` is missing for type `Chjs Tokenization Success`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
