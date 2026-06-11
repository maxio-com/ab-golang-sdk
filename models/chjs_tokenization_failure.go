// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ChjsTokenizationFailure represents a ChjsTokenizationFailure struct.
type ChjsTokenizationFailure struct {
    Errors               string                 `json:"errors"`
    // PCI-safe cardholder fields only. Full card numbers, CVV, and billing address are never included.
    PaymentProfileParams *PaymentProfileParams  `json:"payment_profile_params,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ChjsTokenizationFailure,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ChjsTokenizationFailure) String() string {
    return fmt.Sprintf(
    	"ChjsTokenizationFailure[Errors=%v, PaymentProfileParams=%v, AdditionalProperties=%v]",
    	c.Errors, c.PaymentProfileParams, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ChjsTokenizationFailure.
// It customizes the JSON marshaling process for ChjsTokenizationFailure objects.
func (c ChjsTokenizationFailure) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "errors", "payment_profile_params"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ChjsTokenizationFailure object to a map representation for JSON marshaling.
func (c ChjsTokenizationFailure) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["errors"] = c.Errors
    if c.PaymentProfileParams != nil {
        structMap["payment_profile_params"] = c.PaymentProfileParams.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ChjsTokenizationFailure.
// It customizes the JSON unmarshaling process for ChjsTokenizationFailure objects.
func (c *ChjsTokenizationFailure) UnmarshalJSON(input []byte) error {
    var temp tempChjsTokenizationFailure
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "errors", "payment_profile_params")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Errors = *temp.Errors
    c.PaymentProfileParams = temp.PaymentProfileParams
    return nil
}

// tempChjsTokenizationFailure is a temporary struct used for validating the fields of ChjsTokenizationFailure.
type tempChjsTokenizationFailure  struct {
    Errors               *string               `json:"errors"`
    PaymentProfileParams *PaymentProfileParams `json:"payment_profile_params,omitempty"`
}

func (c *tempChjsTokenizationFailure) validate() error {
    var errs []string
    if c.Errors == nil {
        errs = append(errs, "required field `errors` is missing for type `Chjs Tokenization Failure`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
