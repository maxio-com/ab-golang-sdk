package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CreateInvoicePaymentRequest represents a CreateInvoicePaymentRequest struct.
type CreateInvoicePaymentRequest struct {
    Payment              CreateInvoicePayment `json:"payment"`
    // The type of payment to be applied to an Invoice. Defaults to external.
    Type                 *InvoicePaymentType  `json:"type,omitempty"`
    AdditionalProperties map[string]any       `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreateInvoicePaymentRequest.
// It customizes the JSON marshaling process for CreateInvoicePaymentRequest objects.
func (c CreateInvoicePaymentRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateInvoicePaymentRequest object to a map representation for JSON marshaling.
func (c CreateInvoicePaymentRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["payment"] = c.Payment.toMap()
    if c.Type != nil {
        structMap["type"] = c.Type
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateInvoicePaymentRequest.
// It customizes the JSON unmarshaling process for CreateInvoicePaymentRequest objects.
func (c *CreateInvoicePaymentRequest) UnmarshalJSON(input []byte) error {
    var temp createInvoicePaymentRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "payment", "type")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Payment = *temp.Payment
    c.Type = temp.Type
    return nil
}

// TODO
type createInvoicePaymentRequest  struct {
    Payment *CreateInvoicePayment `json:"payment"`
    Type    *InvoicePaymentType   `json:"type,omitempty"`
}

func (c *createInvoicePaymentRequest) validate() error {
    var errs []string
    if c.Payment == nil {
        errs = append(errs, "required field `payment` is missing for type `Create Invoice Payment Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
