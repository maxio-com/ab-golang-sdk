package models

import (
	"encoding/json"
)

// ProformaInvoice represents a ProformaInvoice struct.
type ProformaInvoice struct {
	Uid                 *string `json:"uid,omitempty"`
	SiteId              *int    `json:"site_id,omitempty"`
	CustomerId          *int    `json:"customer_id,omitempty"`
	SubscriptionId      *int    `json:"subscription_id,omitempty"`
	Number              *int    `json:"number,omitempty"`
	SequenceNumber      *int    `json:"sequence_number,omitempty"`
	CreatedAt           *string `json:"created_at,omitempty"`
	DeliveryDate        *string `json:"delivery_date,omitempty"`
	Status              *string `json:"status,omitempty"`
	CollectionMethod    *string `json:"collection_method,omitempty"`
	PaymentInstructions *string `json:"payment_instructions,omitempty"`
	Currency            *string `json:"currency,omitempty"`
	ConsolidationLevel  *string `json:"consolidation_level,omitempty"`
	ProductName         *string `json:"product_name,omitempty"`
	ProductFamilyName   *string `json:"product_family_name,omitempty"`
	Role                *string `json:"role,omitempty"`
	// Information about the seller (merchant) listed on the masthead of the invoice.
	Seller *InvoiceSeller `json:"seller,omitempty"`
	// Information about the customer who is owner or recipient the invoiced subscription.
	Customer        *InvoiceCustomer          `json:"customer,omitempty"`
	Memo            *string                   `json:"memo,omitempty"`
	BillingAddress  *InvoiceAddress           `json:"billing_address,omitempty"`
	ShippingAddress *InvoiceAddress           `json:"shipping_address,omitempty"`
	SubtotalAmount  *string                   `json:"subtotal_amount,omitempty"`
	DiscountAmount  *string                   `json:"discount_amount,omitempty"`
	TaxAmount       *string                   `json:"tax_amount,omitempty"`
	TotalAmount     *string                   `json:"total_amount,omitempty"`
	CreditAmount    *string                   `json:"credit_amount,omitempty"`
	PaidAmount      *string                   `json:"paid_amount,omitempty"`
	RefundAmount    *string                   `json:"refund_amount,omitempty"`
	DueAmount       *string                   `json:"due_amount,omitempty"`
	LineItems       []InvoiceLineItem         `json:"line_items,omitempty"`
	Discounts       []ProformaInvoiceDiscount `json:"discounts,omitempty"`
	Taxes           []ProformaInvoiceTax      `json:"taxes,omitempty"`
	Credits         []ProformaInvoiceCredit   `json:"credits,omitempty"`
	Payments        []ProformaInvoicePayment  `json:"payments,omitempty"`
	CustomFields    []ProformaCustomField     `json:"custom_fields,omitempty"`
	PublicUrl       *string                   `json:"public_url,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ProformaInvoice.
// It customizes the JSON marshaling process for ProformaInvoice objects.
func (p *ProformaInvoice) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(p.toMap())
}

// toMap converts the ProformaInvoice object to a map representation for JSON marshaling.
func (p *ProformaInvoice) toMap() map[string]any {
	structMap := make(map[string]any)
	if p.Uid != nil {
		structMap["uid"] = p.Uid
	}
	if p.SiteId != nil {
		structMap["site_id"] = p.SiteId
	}
	if p.CustomerId != nil {
		structMap["customer_id"] = p.CustomerId
	}
	if p.SubscriptionId != nil {
		structMap["subscription_id"] = p.SubscriptionId
	}
	if p.Number != nil {
		structMap["number"] = p.Number
	}
	if p.SequenceNumber != nil {
		structMap["sequence_number"] = p.SequenceNumber
	}
	if p.CreatedAt != nil {
		structMap["created_at"] = p.CreatedAt
	}
	if p.DeliveryDate != nil {
		structMap["delivery_date"] = p.DeliveryDate
	}
	if p.Status != nil {
		structMap["status"] = p.Status
	}
	if p.CollectionMethod != nil {
		structMap["collection_method"] = p.CollectionMethod
	}
	if p.PaymentInstructions != nil {
		structMap["payment_instructions"] = p.PaymentInstructions
	}
	if p.Currency != nil {
		structMap["currency"] = p.Currency
	}
	if p.ConsolidationLevel != nil {
		structMap["consolidation_level"] = p.ConsolidationLevel
	}
	if p.ProductName != nil {
		structMap["product_name"] = p.ProductName
	}
	if p.ProductFamilyName != nil {
		structMap["product_family_name"] = p.ProductFamilyName
	}
	if p.Role != nil {
		structMap["role"] = p.Role
	}
	if p.Seller != nil {
		structMap["seller"] = p.Seller
	}
	if p.Customer != nil {
		structMap["customer"] = p.Customer
	}
	if p.Memo != nil {
		structMap["memo"] = p.Memo
	}
	if p.BillingAddress != nil {
		structMap["billing_address"] = p.BillingAddress
	}
	if p.ShippingAddress != nil {
		structMap["shipping_address"] = p.ShippingAddress
	}
	if p.SubtotalAmount != nil {
		structMap["subtotal_amount"] = p.SubtotalAmount
	}
	if p.DiscountAmount != nil {
		structMap["discount_amount"] = p.DiscountAmount
	}
	if p.TaxAmount != nil {
		structMap["tax_amount"] = p.TaxAmount
	}
	if p.TotalAmount != nil {
		structMap["total_amount"] = p.TotalAmount
	}
	if p.CreditAmount != nil {
		structMap["credit_amount"] = p.CreditAmount
	}
	if p.PaidAmount != nil {
		structMap["paid_amount"] = p.PaidAmount
	}
	if p.RefundAmount != nil {
		structMap["refund_amount"] = p.RefundAmount
	}
	if p.DueAmount != nil {
		structMap["due_amount"] = p.DueAmount
	}
	if p.LineItems != nil {
		structMap["line_items"] = p.LineItems
	}
	if p.Discounts != nil {
		structMap["discounts"] = p.Discounts
	}
	if p.Taxes != nil {
		structMap["taxes"] = p.Taxes
	}
	if p.Credits != nil {
		structMap["credits"] = p.Credits
	}
	if p.Payments != nil {
		structMap["payments"] = p.Payments
	}
	if p.CustomFields != nil {
		structMap["custom_fields"] = p.CustomFields
	}
	if p.PublicUrl != nil {
		structMap["public_url"] = p.PublicUrl
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ProformaInvoice.
// It customizes the JSON unmarshaling process for ProformaInvoice objects.
func (p *ProformaInvoice) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Uid                 *string                   `json:"uid,omitempty"`
		SiteId              *int                      `json:"site_id,omitempty"`
		CustomerId          *int                      `json:"customer_id,omitempty"`
		SubscriptionId      *int                      `json:"subscription_id,omitempty"`
		Number              *int                      `json:"number,omitempty"`
		SequenceNumber      *int                      `json:"sequence_number,omitempty"`
		CreatedAt           *string                   `json:"created_at,omitempty"`
		DeliveryDate        *string                   `json:"delivery_date,omitempty"`
		Status              *string                   `json:"status,omitempty"`
		CollectionMethod    *string                   `json:"collection_method,omitempty"`
		PaymentInstructions *string                   `json:"payment_instructions,omitempty"`
		Currency            *string                   `json:"currency,omitempty"`
		ConsolidationLevel  *string                   `json:"consolidation_level,omitempty"`
		ProductName         *string                   `json:"product_name,omitempty"`
		ProductFamilyName   *string                   `json:"product_family_name,omitempty"`
		Role                *string                   `json:"role,omitempty"`
		Seller              *InvoiceSeller            `json:"seller,omitempty"`
		Customer            *InvoiceCustomer          `json:"customer,omitempty"`
		Memo                *string                   `json:"memo,omitempty"`
		BillingAddress      *InvoiceAddress           `json:"billing_address,omitempty"`
		ShippingAddress     *InvoiceAddress           `json:"shipping_address,omitempty"`
		SubtotalAmount      *string                   `json:"subtotal_amount,omitempty"`
		DiscountAmount      *string                   `json:"discount_amount,omitempty"`
		TaxAmount           *string                   `json:"tax_amount,omitempty"`
		TotalAmount         *string                   `json:"total_amount,omitempty"`
		CreditAmount        *string                   `json:"credit_amount,omitempty"`
		PaidAmount          *string                   `json:"paid_amount,omitempty"`
		RefundAmount        *string                   `json:"refund_amount,omitempty"`
		DueAmount           *string                   `json:"due_amount,omitempty"`
		LineItems           []InvoiceLineItem         `json:"line_items,omitempty"`
		Discounts           []ProformaInvoiceDiscount `json:"discounts,omitempty"`
		Taxes               []ProformaInvoiceTax      `json:"taxes,omitempty"`
		Credits             []ProformaInvoiceCredit   `json:"credits,omitempty"`
		Payments            []ProformaInvoicePayment  `json:"payments,omitempty"`
		CustomFields        []ProformaCustomField     `json:"custom_fields,omitempty"`
		PublicUrl           *string                   `json:"public_url,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	p.Uid = temp.Uid
	p.SiteId = temp.SiteId
	p.CustomerId = temp.CustomerId
	p.SubscriptionId = temp.SubscriptionId
	p.Number = temp.Number
	p.SequenceNumber = temp.SequenceNumber
	p.CreatedAt = temp.CreatedAt
	p.DeliveryDate = temp.DeliveryDate
	p.Status = temp.Status
	p.CollectionMethod = temp.CollectionMethod
	p.PaymentInstructions = temp.PaymentInstructions
	p.Currency = temp.Currency
	p.ConsolidationLevel = temp.ConsolidationLevel
	p.ProductName = temp.ProductName
	p.ProductFamilyName = temp.ProductFamilyName
	p.Role = temp.Role
	p.Seller = temp.Seller
	p.Customer = temp.Customer
	p.Memo = temp.Memo
	p.BillingAddress = temp.BillingAddress
	p.ShippingAddress = temp.ShippingAddress
	p.SubtotalAmount = temp.SubtotalAmount
	p.DiscountAmount = temp.DiscountAmount
	p.TaxAmount = temp.TaxAmount
	p.TotalAmount = temp.TotalAmount
	p.CreditAmount = temp.CreditAmount
	p.PaidAmount = temp.PaidAmount
	p.RefundAmount = temp.RefundAmount
	p.DueAmount = temp.DueAmount
	p.LineItems = temp.LineItems
	p.Discounts = temp.Discounts
	p.Taxes = temp.Taxes
	p.Credits = temp.Credits
	p.Payments = temp.Payments
	p.CustomFields = temp.CustomFields
	p.PublicUrl = temp.PublicUrl
	return nil
}
