// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// CreateInvoicePayment represents a CreateInvoicePayment struct.
type CreateInvoicePayment struct {
    // A string of the dollar amount to be refunded (eg. "10.50" => $10.50)
    Amount               *CreateInvoicePaymentAmount `json:"amount,omitempty"`
    // A description to be attached to the payment. Applicable only to `external` payments.
    Memo                 *string                     `json:"memo,omitempty"`
    // The type of payment method used. Defaults to other.
    Method               *InvoicePaymentMethodType   `json:"method,omitempty"`
    // Additional information related to the payment method (eg. Check #). Applicable only to `external` payments.
    Details              *string                     `json:"details,omitempty"`
    // The ID of the payment profile to be used for the payment.
    PaymentProfileId     *int                        `json:"payment_profile_id,omitempty"`
    // Date reflecting when the payment was received from a customer. Must be in the past. Applicable only to
    // `external` payments.
    ReceivedOn           *time.Time                  `json:"received_on,omitempty"`
    AdditionalProperties map[string]interface{}      `json:"_"`
}

// String implements the fmt.Stringer interface for CreateInvoicePayment,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateInvoicePayment) String() string {
    return fmt.Sprintf(
    	"CreateInvoicePayment[Amount=%v, Memo=%v, Method=%v, Details=%v, PaymentProfileId=%v, ReceivedOn=%v, AdditionalProperties=%v]",
    	c.Amount, c.Memo, c.Method, c.Details, c.PaymentProfileId, c.ReceivedOn, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreateInvoicePayment.
// It customizes the JSON marshaling process for CreateInvoicePayment objects.
func (c CreateInvoicePayment) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "amount", "memo", "method", "details", "payment_profile_id", "received_on"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateInvoicePayment object to a map representation for JSON marshaling.
func (c CreateInvoicePayment) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
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
    if c.ReceivedOn != nil {
        structMap["received_on"] = c.ReceivedOn.Format(DEFAULT_DATE)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "amount", "memo", "method", "details", "payment_profile_id", "received_on")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Amount = temp.Amount
    c.Memo = temp.Memo
    c.Method = temp.Method
    c.Details = temp.Details
    c.PaymentProfileId = temp.PaymentProfileId
    if temp.ReceivedOn != nil {
        ReceivedOnVal, err := time.Parse(DEFAULT_DATE, *temp.ReceivedOn)
        if err != nil {
            log.Fatalf("Cannot Parse received_on as % s format.", DEFAULT_DATE)
        }
        c.ReceivedOn = &ReceivedOnVal
    }
    return nil
}

// tempCreateInvoicePayment is a temporary struct used for validating the fields of CreateInvoicePayment.
type tempCreateInvoicePayment  struct {
    Amount           *CreateInvoicePaymentAmount `json:"amount,omitempty"`
    Memo             *string                     `json:"memo,omitempty"`
    Method           *InvoicePaymentMethodType   `json:"method,omitempty"`
    Details          *string                     `json:"details,omitempty"`
    PaymentProfileId *int                        `json:"payment_profile_id,omitempty"`
    ReceivedOn       *string                     `json:"received_on,omitempty"`
}
