package models

import (
	"encoding/json"
)

// CreateInvoice represents a CreateInvoice struct.
type CreateInvoice struct {
	LineItems []CreateInvoiceItem `json:"line_items,omitempty"`
	IssueDate *string             `json:"issue_date,omitempty"`
	// By default, invoices will be created with a due date matching the date of invoice creation. If a different due date is desired, the net_terms parameter can be sent indicating the number of days in advance the due date should be.
	NetTerms            *int    `json:"net_terms,omitempty"`
	PaymentInstructions *string `json:"payment_instructions,omitempty"`
	// A custom memo can be sent to override the site's default.
	Memo *string `json:"memo,omitempty"`
	// Overrides the defaults for the site
	SellerAddress *CreateInvoiceAddress `json:"seller_address,omitempty"`
	// Overrides the default for the customer
	BillingAddress *CreateInvoiceAddress `json:"billing_address,omitempty"`
	// Overrides the default for the customer
	ShippingAddress *CreateInvoiceAddress `json:"shipping_address,omitempty"`
	Coupons         []CreateInvoiceCoupon `json:"coupons,omitempty"`
	Status          *Status1Enum          `json:"status,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateInvoice.
// It customizes the JSON marshaling process for CreateInvoice objects.
func (c *CreateInvoice) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateInvoice object to a map representation for JSON marshaling.
func (c *CreateInvoice) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.LineItems != nil {
		structMap["line_items"] = c.LineItems
	}
	if c.IssueDate != nil {
		structMap["issue_date"] = c.IssueDate
	}
	if c.NetTerms != nil {
		structMap["net_terms"] = c.NetTerms
	}
	if c.PaymentInstructions != nil {
		structMap["payment_instructions"] = c.PaymentInstructions
	}
	if c.Memo != nil {
		structMap["memo"] = c.Memo
	}
	if c.SellerAddress != nil {
		structMap["seller_address"] = c.SellerAddress
	}
	if c.BillingAddress != nil {
		structMap["billing_address"] = c.BillingAddress
	}
	if c.ShippingAddress != nil {
		structMap["shipping_address"] = c.ShippingAddress
	}
	if c.Coupons != nil {
		structMap["coupons"] = c.Coupons
	}
	if c.Status != nil {
		structMap["status"] = c.Status
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateInvoice.
// It customizes the JSON unmarshaling process for CreateInvoice objects.
func (c *CreateInvoice) UnmarshalJSON(input []byte) error {
	temp := &struct {
		LineItems           []CreateInvoiceItem   `json:"line_items,omitempty"`
		IssueDate           *string               `json:"issue_date,omitempty"`
		NetTerms            *int                  `json:"net_terms,omitempty"`
		PaymentInstructions *string               `json:"payment_instructions,omitempty"`
		Memo                *string               `json:"memo,omitempty"`
		SellerAddress       *CreateInvoiceAddress `json:"seller_address,omitempty"`
		BillingAddress      *CreateInvoiceAddress `json:"billing_address,omitempty"`
		ShippingAddress     *CreateInvoiceAddress `json:"shipping_address,omitempty"`
		Coupons             []CreateInvoiceCoupon `json:"coupons,omitempty"`
		Status              *Status1Enum          `json:"status,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.LineItems = temp.LineItems
	c.IssueDate = temp.IssueDate
	c.NetTerms = temp.NetTerms
	c.PaymentInstructions = temp.PaymentInstructions
	c.Memo = temp.Memo
	c.SellerAddress = temp.SellerAddress
	c.BillingAddress = temp.BillingAddress
	c.ShippingAddress = temp.ShippingAddress
	c.Coupons = temp.Coupons
	c.Status = temp.Status
	return nil
}
