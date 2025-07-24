// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CreateInvoicePaymentRequest represents a CreateInvoicePaymentRequest struct.
type CreateInvoicePaymentRequest struct {
    Payment              CreateInvoicePayment   `json:"payment"`
    // The type of payment to be applied to an Invoice. Defaults to external.
    Type                 *InvoicePaymentType    `json:"type,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CreateInvoicePaymentRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateInvoicePaymentRequest) String() string {
    return fmt.Sprintf(
    	"CreateInvoicePaymentRequest[Payment=%v, Type=%v, AdditionalProperties=%v]",
    	c.Payment, c.Type, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreateInvoicePaymentRequest.
// It customizes the JSON marshaling process for CreateInvoicePaymentRequest objects.
func (c CreateInvoicePaymentRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "payment", "type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateInvoicePaymentRequest object to a map representation for JSON marshaling.
func (c CreateInvoicePaymentRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["payment"] = c.Payment.toMap()
    if c.Type != nil {
        structMap["type"] = c.Type
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateInvoicePaymentRequest.
// It customizes the JSON unmarshaling process for CreateInvoicePaymentRequest objects.
func (c *CreateInvoicePaymentRequest) UnmarshalJSON(input []byte) error {
    var temp tempCreateInvoicePaymentRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "payment", "type")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Payment = *temp.Payment
    c.Type = temp.Type
    return nil
}

// tempCreateInvoicePaymentRequest is a temporary struct used for validating the fields of CreateInvoicePaymentRequest.
type tempCreateInvoicePaymentRequest  struct {
    Payment *CreateInvoicePayment `json:"payment"`
    Type    *InvoicePaymentType   `json:"type,omitempty"`
}

func (c *tempCreateInvoicePaymentRequest) validate() error {
    var errs []string
    if c.Payment == nil {
        errs = append(errs, "required field `payment` is missing for type `Create Invoice Payment Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
