/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CreateInvoicePaymentApplication represents a CreateInvoicePaymentApplication struct.
type CreateInvoicePaymentApplication struct {
    // Unique identifier for the invoice. It has the prefix "inv_" followed by alphanumeric characters.
    InvoiceUid           string                 `json:"invoice_uid"`
    // Dollar amount of the invoice payment (eg. "10.50" => $10.50).
    Amount               string                 `json:"amount"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CreateInvoicePaymentApplication,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateInvoicePaymentApplication) String() string {
    return fmt.Sprintf(
    	"CreateInvoicePaymentApplication[InvoiceUid=%v, Amount=%v, AdditionalProperties=%v]",
    	c.InvoiceUid, c.Amount, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreateInvoicePaymentApplication.
// It customizes the JSON marshaling process for CreateInvoicePaymentApplication objects.
func (c CreateInvoicePaymentApplication) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "invoice_uid", "amount"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateInvoicePaymentApplication object to a map representation for JSON marshaling.
func (c CreateInvoicePaymentApplication) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["invoice_uid"] = c.InvoiceUid
    structMap["amount"] = c.Amount
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateInvoicePaymentApplication.
// It customizes the JSON unmarshaling process for CreateInvoicePaymentApplication objects.
func (c *CreateInvoicePaymentApplication) UnmarshalJSON(input []byte) error {
    var temp tempCreateInvoicePaymentApplication
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "invoice_uid", "amount")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.InvoiceUid = *temp.InvoiceUid
    c.Amount = *temp.Amount
    return nil
}

// tempCreateInvoicePaymentApplication is a temporary struct used for validating the fields of CreateInvoicePaymentApplication.
type tempCreateInvoicePaymentApplication  struct {
    InvoiceUid *string `json:"invoice_uid"`
    Amount     *string `json:"amount"`
}

func (c *tempCreateInvoicePaymentApplication) validate() error {
    var errs []string
    if c.InvoiceUid == nil {
        errs = append(errs, "required field `invoice_uid` is missing for type `Create Invoice Payment Application`")
    }
    if c.Amount == nil {
        errs = append(errs, "required field `amount` is missing for type `Create Invoice Payment Application`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
