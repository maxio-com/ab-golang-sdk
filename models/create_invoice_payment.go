/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// CreateInvoicePayment represents a CreateInvoicePayment struct.
type CreateInvoicePayment struct {
    // A string of the dollar amount to be refunded (eg. "10.50" => $10.50)
    Amount               *CreateInvoicePaymentAmount `json:"amount,omitempty"`
    // A description to be attached to the payment.
    Memo                 *string                     `json:"memo,omitempty"`
    // The type of payment method used. Defaults to other.
    Method               *InvoicePaymentMethodType   `json:"method,omitempty"`
    // Additional information related to the payment method (eg. Check #)
    Details              *string                     `json:"details,omitempty"`
    // The ID of the payment profile to be used for the payment.
    PaymentProfileId     *int                        `json:"payment_profile_id,omitempty"`
    AdditionalProperties map[string]any              `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreateInvoicePayment.
// It customizes the JSON marshaling process for CreateInvoicePayment objects.
func (c CreateInvoicePayment) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateInvoicePayment object to a map representation for JSON marshaling.
func (c CreateInvoicePayment) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Amount != nil {
        structMap["amount"] = c.Amount.toMap()
    }
    if c.Memo != nil {
        structMap["memo"] = c.Memo
    }
    if c.Method != nil {
        structMap["method"] = c.Method
    }
    if c.Details != nil {
        structMap["details"] = c.Details
    }
    if c.PaymentProfileId != nil {
        structMap["payment_profile_id"] = c.PaymentProfileId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateInvoicePayment.
// It customizes the JSON unmarshaling process for CreateInvoicePayment objects.
func (c *CreateInvoicePayment) UnmarshalJSON(input []byte) error {
    var temp tempCreateInvoicePayment
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "amount", "memo", "method", "details", "payment_profile_id")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Amount = temp.Amount
    c.Memo = temp.Memo
    c.Method = temp.Method
    c.Details = temp.Details
    c.PaymentProfileId = temp.PaymentProfileId
    return nil
}

// tempCreateInvoicePayment is a temporary struct used for validating the fields of CreateInvoicePayment.
type tempCreateInvoicePayment  struct {
    Amount           *CreateInvoicePaymentAmount `json:"amount,omitempty"`
    Memo             *string                     `json:"memo,omitempty"`
    Method           *InvoicePaymentMethodType   `json:"method,omitempty"`
    Details          *string                     `json:"details,omitempty"`
    PaymentProfileId *int                        `json:"payment_profile_id,omitempty"`
}
