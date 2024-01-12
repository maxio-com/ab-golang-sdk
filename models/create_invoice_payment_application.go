package models

import (
    "encoding/json"
)

// CreateInvoicePaymentApplication represents a CreateInvoicePaymentApplication struct.
type CreateInvoicePaymentApplication struct {
    // Unique identifier for the invoice. It has the prefix "inv_" followed by alphanumeric characters.
    InvoiceUid string `json:"invoice_uid"`
    // Dollar amount of the invoice payment (eg. "10.50" => $10.50).
    Amount     string `json:"amount"`
}

// MarshalJSON implements the json.Marshaler interface for CreateInvoicePaymentApplication.
// It customizes the JSON marshaling process for CreateInvoicePaymentApplication objects.
func (c *CreateInvoicePaymentApplication) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateInvoicePaymentApplication object to a map representation for JSON marshaling.
func (c *CreateInvoicePaymentApplication) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["invoice_uid"] = c.InvoiceUid
    structMap["amount"] = c.Amount
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateInvoicePaymentApplication.
// It customizes the JSON unmarshaling process for CreateInvoicePaymentApplication objects.
func (c *CreateInvoicePaymentApplication) UnmarshalJSON(input []byte) error {
    temp := &struct {
        InvoiceUid string `json:"invoice_uid"`
        Amount     string `json:"amount"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.InvoiceUid = temp.InvoiceUid
    c.Amount = temp.Amount
    return nil
}
