// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CreateMultiInvoicePaymentRequest represents a CreateMultiInvoicePaymentRequest struct.
type CreateMultiInvoicePaymentRequest struct {
    Payment              CreateMultiInvoicePayment `json:"payment"`
    AdditionalProperties map[string]interface{}    `json:"_"`
}

// String implements the fmt.Stringer interface for CreateMultiInvoicePaymentRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateMultiInvoicePaymentRequest) String() string {
    return fmt.Sprintf(
    	"CreateMultiInvoicePaymentRequest[Payment=%v, AdditionalProperties=%v]",
    	c.Payment, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreateMultiInvoicePaymentRequest.
// It customizes the JSON marshaling process for CreateMultiInvoicePaymentRequest objects.
func (c CreateMultiInvoicePaymentRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "payment"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateMultiInvoicePaymentRequest object to a map representation for JSON marshaling.
func (c CreateMultiInvoicePaymentRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["payment"] = c.Payment.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateMultiInvoicePaymentRequest.
// It customizes the JSON unmarshaling process for CreateMultiInvoicePaymentRequest objects.
func (c *CreateMultiInvoicePaymentRequest) UnmarshalJSON(input []byte) error {
    var temp tempCreateMultiInvoicePaymentRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "payment")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Payment = *temp.Payment
    return nil
}

// tempCreateMultiInvoicePaymentRequest is a temporary struct used for validating the fields of CreateMultiInvoicePaymentRequest.
type tempCreateMultiInvoicePaymentRequest  struct {
    Payment *CreateMultiInvoicePayment `json:"payment"`
}

func (c *tempCreateMultiInvoicePaymentRequest) validate() error {
    var errs []string
    if c.Payment == nil {
        errs = append(errs, "required field `payment` is missing for type `Create Multi Invoice Payment Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
