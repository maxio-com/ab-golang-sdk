package models

import (
	"encoding/json"
	"log"
	"time"
)

// Invoice represents a Invoice struct.
type Invoice struct {
	Id *int64 `json:"id,omitempty"`
	// Unique identifier for the invoice. It is generated automatically by Chargify and has the prefix "inv_" followed by alphanumeric characters.
	Uid *string `json:"uid,omitempty"`
	// ID of the site to which the invoice belongs.
	SiteId *int `json:"site_id,omitempty"`
	// ID of the customer to which the invoice belongs.
	CustomerId *int `json:"customer_id,omitempty"`
	// ID of the subscription that generated the invoice.
	SubscriptionId *int `json:"subscription_id,omitempty"`
	// A unique, identifying string that appears on the invoice and in places the invoice is referenced.
	// While the UID is long and not appropriate to show to customers, the number is usually shorter and consumable by the customer and the merchant alike.
	Number *string `json:"number,omitempty"`
	// A monotonically increasing number assigned to invoices as they are created.  This number is unique within a site and can be used to sort and order invoices.
	SequenceNumber  *int       `json:"sequence_number,omitempty"`
	TransactionTime *time.Time `json:"transaction_time,omitempty"`
	CreatedAt       *time.Time `json:"created_at,omitempty"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty"`
	// Date the invoice was issued to the customer.  This is the date that the invoice was made available for payment.
	// The format is `"YYYY-MM-DD"`.
	IssueDate *time.Time `json:"issue_date,omitempty"`
	// Date the invoice is due.
	// The format is `"YYYY-MM-DD"`.
	DueDate *time.Time `json:"due_date,omitempty"`
	// Date the invoice became fully paid.
	// If partial payments are applied to the invoice, this date will not be present until payment has been made in full.
	// The format is `"YYYY-MM-DD"`.
	PaidDate Optional[time.Time] `json:"paid_date"`
	// The current status of the invoice. See [Invoice Statuses](https://chargify.zendesk.com/hc/en-us/articles/4407737494171#line-item-breakdowns) for more.
	Status          *InvoiceStatus `json:"status,omitempty"`
	Role            *InvoiceRole   `json:"role,omitempty"`
	ParentInvoiceId Optional[int]  `json:"parent_invoice_id"`
	// The type of payment collection to be used in the subscription. For legacy Statements Architecture valid options are - `invoice`, `automatic`. For current Relationship Invoicing Architecture valid options are - `remittance`, `automatic`, `prepaid`.
	CollectionMethod *CollectionMethod `json:"collection_method,omitempty"`
	// A message that is printed on the invoice when it is marked for remittance collection. It is intended to describe to the customer how they may make payment, and is configured by the merchant.
	PaymentInstructions *string `json:"payment_instructions,omitempty"`
	// The ISO 4217 currency code (3 character string) representing the currency of invoice transaction.
	Currency *string `json:"currency,omitempty"`
	// Consolidation level of the invoice, which is applicable to invoice consolidation.  It will hold one of the following values:
	// * "none": A normal invoice with no consolidation.
	// * "child": An invoice segment which has been combined into a consolidated invoice.
	// * "parent": A consolidated invoice, whose contents are composed of invoice segments.
	// "Parent" invoices do not have lines of their own, but they have subtotals and totals which aggregate the member invoice segments.
	// See also the [invoice consolidation documentation](https://chargify.zendesk.com/hc/en-us/articles/4407746391835).
	ConsolidationLevel *InvoiceConsolidationLevel `json:"consolidation_level,omitempty"`
	// For invoices with `consolidation_level` of `child`, this specifies the UID of the parent (consolidated) invoice.
	ParentInvoiceUid    Optional[string] `json:"parent_invoice_uid"`
	SubscriptionGroupId Optional[int]    `json:"subscription_group_id"`
	// For invoices with `consolidation_level` of `child`, this specifies the number of the parent (consolidated) invoice.
	ParentInvoiceNumber Optional[int] `json:"parent_invoice_number"`
	// For invoices with `consolidation_level` of `parent`, this specifies the ID of the subscription which was the primary subscription of the subscription group that generated the invoice.
	GroupPrimarySubscriptionId Optional[int] `json:"group_primary_subscription_id"`
	// The name of the product subscribed when the invoice was generated.
	ProductName *string `json:"product_name,omitempty"`
	// The name of the product family subscribed when the invoice was generated.
	ProductFamilyName *string `json:"product_family_name,omitempty"`
	// Information about the seller (merchant) listed on the masthead of the invoice.
	Seller *InvoiceSeller `json:"seller,omitempty"`
	// Information about the customer who is owner or recipient the invoiced subscription.
	Customer        *InvoiceCustomer `json:"customer,omitempty"`
	Payer           *InvoicePayer    `json:"payer,omitempty"`
	RecipientEmails []string         `json:"recipient_emails,omitempty"`
	NetTerms        *int             `json:"net_terms,omitempty"`
	// The memo printed on invoices of any collection type.  This message is in control of the merchant.
	Memo *string `json:"memo,omitempty"`
	// The invoice billing address.
	BillingAddress *InvoiceAddress `json:"billing_address,omitempty"`
	// The invoice shipping address.
	ShippingAddress *InvoiceAddress `json:"shipping_address,omitempty"`
	// Subtotal of the invoice, which is the sum of all line items before discounts or taxes.
	SubtotalAmount *string `json:"subtotal_amount,omitempty"`
	// Total discount applied to the invoice.
	DiscountAmount *string `json:"discount_amount,omitempty"`
	// Total tax on the invoice.
	TaxAmount *string `json:"tax_amount,omitempty"`
	// The invoice total, which is `subtotal_amount - discount_amount + tax_amount`.'
	TotalAmount *string `json:"total_amount,omitempty"`
	// The amount of credit (from credit notes) applied to this invoice.
	// Credits offset the amount due from the customer.
	CreditAmount *string `json:"credit_amount,omitempty"`
	RefundAmount *string `json:"refund_amount,omitempty"`
	// The amount paid on the invoice by the customer.
	PaidAmount *string `json:"paid_amount,omitempty"`
	// Amount due on the invoice, which is `total_amount - credit_amount - paid_amount`.
	DueAmount *string `json:"due_amount,omitempty"`
	// Line items on the invoice.
	LineItems       []InvoiceLineItem       `json:"line_items,omitempty"`
	Discounts       []InvoiceDiscount       `json:"discounts,omitempty"`
	Taxes           []InvoiceTax            `json:"taxes,omitempty"`
	Credits         []InvoiceCredit         `json:"credits,omitempty"`
	Refunds         []InvoiceRefund         `json:"refunds,omitempty"`
	Payments        []InvoicePayment        `json:"payments,omitempty"`
	CustomFields    []InvoiceCustomField    `json:"custom_fields,omitempty"`
	DisplaySettings *InvoiceDisplaySettings `json:"display_settings,omitempty"`
	// The public URL of the invoice
	PublicUrl           *string                 `json:"public_url,omitempty"`
	PreviousBalanceData *InvoicePreviousBalance `json:"previous_balance_data,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for Invoice.
// It customizes the JSON marshaling process for Invoice objects.
func (i *Invoice) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(i.toMap())
}

// toMap converts the Invoice object to a map representation for JSON marshaling.
func (i *Invoice) toMap() map[string]any {
	structMap := make(map[string]any)
	if i.Id != nil {
		structMap["id"] = i.Id
	}
	if i.Uid != nil {
		structMap["uid"] = i.Uid
	}
	if i.SiteId != nil {
		structMap["site_id"] = i.SiteId
	}
	if i.CustomerId != nil {
		structMap["customer_id"] = i.CustomerId
	}
	if i.SubscriptionId != nil {
		structMap["subscription_id"] = i.SubscriptionId
	}
	if i.Number != nil {
		structMap["number"] = i.Number
	}
	if i.SequenceNumber != nil {
		structMap["sequence_number"] = i.SequenceNumber
	}
	if i.TransactionTime != nil {
		structMap["transaction_time"] = i.TransactionTime.Format(time.RFC3339)
	}
	if i.CreatedAt != nil {
		structMap["created_at"] = i.CreatedAt.Format(time.RFC3339)
	}
	if i.UpdatedAt != nil {
		structMap["updated_at"] = i.UpdatedAt.Format(time.RFC3339)
	}
	if i.IssueDate != nil {
		structMap["issue_date"] = i.IssueDate.Format(DEFAULT_DATE)
	}
	if i.DueDate != nil {
		structMap["due_date"] = i.DueDate.Format(DEFAULT_DATE)
	}
	if i.PaidDate.IsValueSet() {
		var PaidDateVal *string = nil
		if i.PaidDate.Value() != nil {
			val := i.PaidDate.Value().Format(DEFAULT_DATE)
			PaidDateVal = &val
		}
		if i.PaidDate.Value() != nil {
			structMap["paid_date"] = PaidDateVal
		} else {
			structMap["paid_date"] = nil
		}
	}
	if i.Status != nil {
		structMap["status"] = i.Status
	}
	if i.Role != nil {
		structMap["role"] = i.Role
	}
	if i.ParentInvoiceId.IsValueSet() {
		if i.ParentInvoiceId.Value() != nil {
			structMap["parent_invoice_id"] = i.ParentInvoiceId.Value()
		} else {
			structMap["parent_invoice_id"] = nil
		}
	}
	if i.CollectionMethod != nil {
		structMap["collection_method"] = i.CollectionMethod
	}
	if i.PaymentInstructions != nil {
		structMap["payment_instructions"] = i.PaymentInstructions
	}
	if i.Currency != nil {
		structMap["currency"] = i.Currency
	}
	if i.ConsolidationLevel != nil {
		structMap["consolidation_level"] = i.ConsolidationLevel
	}
	if i.ParentInvoiceUid.IsValueSet() {
		if i.ParentInvoiceUid.Value() != nil {
			structMap["parent_invoice_uid"] = i.ParentInvoiceUid.Value()
		} else {
			structMap["parent_invoice_uid"] = nil
		}
	}
	if i.SubscriptionGroupId.IsValueSet() {
		if i.SubscriptionGroupId.Value() != nil {
			structMap["subscription_group_id"] = i.SubscriptionGroupId.Value()
		} else {
			structMap["subscription_group_id"] = nil
		}
	}
	if i.ParentInvoiceNumber.IsValueSet() {
		if i.ParentInvoiceNumber.Value() != nil {
			structMap["parent_invoice_number"] = i.ParentInvoiceNumber.Value()
		} else {
			structMap["parent_invoice_number"] = nil
		}
	}
	if i.GroupPrimarySubscriptionId.IsValueSet() {
		if i.GroupPrimarySubscriptionId.Value() != nil {
			structMap["group_primary_subscription_id"] = i.GroupPrimarySubscriptionId.Value()
		} else {
			structMap["group_primary_subscription_id"] = nil
		}
	}
	if i.ProductName != nil {
		structMap["product_name"] = i.ProductName
	}
	if i.ProductFamilyName != nil {
		structMap["product_family_name"] = i.ProductFamilyName
	}
	if i.Seller != nil {
		structMap["seller"] = i.Seller.toMap()
	}
	if i.Customer != nil {
		structMap["customer"] = i.Customer.toMap()
	}
	if i.Payer != nil {
		structMap["payer"] = i.Payer.toMap()
	}
	if i.RecipientEmails != nil {
		structMap["recipient_emails"] = i.RecipientEmails
	}
	if i.NetTerms != nil {
		structMap["net_terms"] = i.NetTerms
	}
	if i.Memo != nil {
		structMap["memo"] = i.Memo
	}
	if i.BillingAddress != nil {
		structMap["billing_address"] = i.BillingAddress.toMap()
	}
	if i.ShippingAddress != nil {
		structMap["shipping_address"] = i.ShippingAddress.toMap()
	}
	if i.SubtotalAmount != nil {
		structMap["subtotal_amount"] = i.SubtotalAmount
	}
	if i.DiscountAmount != nil {
		structMap["discount_amount"] = i.DiscountAmount
	}
	if i.TaxAmount != nil {
		structMap["tax_amount"] = i.TaxAmount
	}
	if i.TotalAmount != nil {
		structMap["total_amount"] = i.TotalAmount
	}
	if i.CreditAmount != nil {
		structMap["credit_amount"] = i.CreditAmount
	}
	if i.RefundAmount != nil {
		structMap["refund_amount"] = i.RefundAmount
	}
	if i.PaidAmount != nil {
		structMap["paid_amount"] = i.PaidAmount
	}
	if i.DueAmount != nil {
		structMap["due_amount"] = i.DueAmount
	}
	if i.LineItems != nil {
		structMap["line_items"] = i.LineItems
	}
	if i.Discounts != nil {
		structMap["discounts"] = i.Discounts
	}
	if i.Taxes != nil {
		structMap["taxes"] = i.Taxes
	}
	if i.Credits != nil {
		structMap["credits"] = i.Credits
	}
	if i.Refunds != nil {
		structMap["refunds"] = i.Refunds
	}
	if i.Payments != nil {
		structMap["payments"] = i.Payments
	}
	if i.CustomFields != nil {
		structMap["custom_fields"] = i.CustomFields
	}
	if i.DisplaySettings != nil {
		structMap["display_settings"] = i.DisplaySettings.toMap()
	}
	if i.PublicUrl != nil {
		structMap["public_url"] = i.PublicUrl
	}
	if i.PreviousBalanceData != nil {
		structMap["previous_balance_data"] = i.PreviousBalanceData.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Invoice.
// It customizes the JSON unmarshaling process for Invoice objects.
func (i *Invoice) UnmarshalJSON(input []byte) error {
	var temp invoice
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	i.Id = temp.Id
	i.Uid = temp.Uid
	i.SiteId = temp.SiteId
	i.CustomerId = temp.CustomerId
	i.SubscriptionId = temp.SubscriptionId
	i.Number = temp.Number
	i.SequenceNumber = temp.SequenceNumber
	if temp.TransactionTime != nil {
		TransactionTimeVal, err := time.Parse(time.RFC3339, *temp.TransactionTime)
		if err != nil {
			log.Fatalf("Cannot Parse transaction_time as % s format.", time.RFC3339)
		}
		i.TransactionTime = &TransactionTimeVal
	}
	if temp.CreatedAt != nil {
		CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
		if err != nil {
			log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
		}
		i.CreatedAt = &CreatedAtVal
	}
	if temp.UpdatedAt != nil {
		UpdatedAtVal, err := time.Parse(time.RFC3339, *temp.UpdatedAt)
		if err != nil {
			log.Fatalf("Cannot Parse updated_at as % s format.", time.RFC3339)
		}
		i.UpdatedAt = &UpdatedAtVal
	}
	if temp.IssueDate != nil {
		IssueDateVal, err := time.Parse(DEFAULT_DATE, *temp.IssueDate)
		if err != nil {
			log.Fatalf("Cannot Parse issue_date as % s format.", DEFAULT_DATE)
		}
		i.IssueDate = &IssueDateVal
	}
	if temp.DueDate != nil {
		DueDateVal, err := time.Parse(DEFAULT_DATE, *temp.DueDate)
		if err != nil {
			log.Fatalf("Cannot Parse due_date as % s format.", DEFAULT_DATE)
		}
		i.DueDate = &DueDateVal
	}
	i.PaidDate.ShouldSetValue(temp.PaidDate.IsValueSet())
	if temp.PaidDate.Value() != nil {
		PaidDateVal, err := time.Parse(DEFAULT_DATE, (*temp.PaidDate.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse paid_date as % s format.", DEFAULT_DATE)
		}
		i.PaidDate.SetValue(&PaidDateVal)
	}
	i.Status = temp.Status
	i.Role = temp.Role
	i.ParentInvoiceId = temp.ParentInvoiceId
	i.CollectionMethod = temp.CollectionMethod
	i.PaymentInstructions = temp.PaymentInstructions
	i.Currency = temp.Currency
	i.ConsolidationLevel = temp.ConsolidationLevel
	i.ParentInvoiceUid = temp.ParentInvoiceUid
	i.SubscriptionGroupId = temp.SubscriptionGroupId
	i.ParentInvoiceNumber = temp.ParentInvoiceNumber
	i.GroupPrimarySubscriptionId = temp.GroupPrimarySubscriptionId
	i.ProductName = temp.ProductName
	i.ProductFamilyName = temp.ProductFamilyName
	i.Seller = temp.Seller
	i.Customer = temp.Customer
	i.Payer = temp.Payer
	i.RecipientEmails = temp.RecipientEmails
	i.NetTerms = temp.NetTerms
	i.Memo = temp.Memo
	i.BillingAddress = temp.BillingAddress
	i.ShippingAddress = temp.ShippingAddress
	i.SubtotalAmount = temp.SubtotalAmount
	i.DiscountAmount = temp.DiscountAmount
	i.TaxAmount = temp.TaxAmount
	i.TotalAmount = temp.TotalAmount
	i.CreditAmount = temp.CreditAmount
	i.RefundAmount = temp.RefundAmount
	i.PaidAmount = temp.PaidAmount
	i.DueAmount = temp.DueAmount
	i.LineItems = temp.LineItems
	i.Discounts = temp.Discounts
	i.Taxes = temp.Taxes
	i.Credits = temp.Credits
	i.Refunds = temp.Refunds
	i.Payments = temp.Payments
	i.CustomFields = temp.CustomFields
	i.DisplaySettings = temp.DisplaySettings
	i.PublicUrl = temp.PublicUrl
	i.PreviousBalanceData = temp.PreviousBalanceData
	return nil
}

// TODO
type invoice struct {
	Id                         *int64                     `json:"id,omitempty"`
	Uid                        *string                    `json:"uid,omitempty"`
	SiteId                     *int                       `json:"site_id,omitempty"`
	CustomerId                 *int                       `json:"customer_id,omitempty"`
	SubscriptionId             *int                       `json:"subscription_id,omitempty"`
	Number                     *string                    `json:"number,omitempty"`
	SequenceNumber             *int                       `json:"sequence_number,omitempty"`
	TransactionTime            *string                    `json:"transaction_time,omitempty"`
	CreatedAt                  *string                    `json:"created_at,omitempty"`
	UpdatedAt                  *string                    `json:"updated_at,omitempty"`
	IssueDate                  *string                    `json:"issue_date,omitempty"`
	DueDate                    *string                    `json:"due_date,omitempty"`
	PaidDate                   Optional[string]           `json:"paid_date"`
	Status                     *InvoiceStatus             `json:"status,omitempty"`
	Role                       *InvoiceRole               `json:"role,omitempty"`
	ParentInvoiceId            Optional[int]              `json:"parent_invoice_id"`
	CollectionMethod           *CollectionMethod          `json:"collection_method,omitempty"`
	PaymentInstructions        *string                    `json:"payment_instructions,omitempty"`
	Currency                   *string                    `json:"currency,omitempty"`
	ConsolidationLevel         *InvoiceConsolidationLevel `json:"consolidation_level,omitempty"`
	ParentInvoiceUid           Optional[string]           `json:"parent_invoice_uid"`
	SubscriptionGroupId        Optional[int]              `json:"subscription_group_id"`
	ParentInvoiceNumber        Optional[int]              `json:"parent_invoice_number"`
	GroupPrimarySubscriptionId Optional[int]              `json:"group_primary_subscription_id"`
	ProductName                *string                    `json:"product_name,omitempty"`
	ProductFamilyName          *string                    `json:"product_family_name,omitempty"`
	Seller                     *InvoiceSeller             `json:"seller,omitempty"`
	Customer                   *InvoiceCustomer           `json:"customer,omitempty"`
	Payer                      *InvoicePayer              `json:"payer,omitempty"`
	RecipientEmails            []string                   `json:"recipient_emails,omitempty"`
	NetTerms                   *int                       `json:"net_terms,omitempty"`
	Memo                       *string                    `json:"memo,omitempty"`
	BillingAddress             *InvoiceAddress            `json:"billing_address,omitempty"`
	ShippingAddress            *InvoiceAddress            `json:"shipping_address,omitempty"`
	SubtotalAmount             *string                    `json:"subtotal_amount,omitempty"`
	DiscountAmount             *string                    `json:"discount_amount,omitempty"`
	TaxAmount                  *string                    `json:"tax_amount,omitempty"`
	TotalAmount                *string                    `json:"total_amount,omitempty"`
	CreditAmount               *string                    `json:"credit_amount,omitempty"`
	RefundAmount               *string                    `json:"refund_amount,omitempty"`
	PaidAmount                 *string                    `json:"paid_amount,omitempty"`
	DueAmount                  *string                    `json:"due_amount,omitempty"`
	LineItems                  []InvoiceLineItem          `json:"line_items,omitempty"`
	Discounts                  []InvoiceDiscount          `json:"discounts,omitempty"`
	Taxes                      []InvoiceTax               `json:"taxes,omitempty"`
	Credits                    []InvoiceCredit            `json:"credits,omitempty"`
	Refunds                    []InvoiceRefund            `json:"refunds,omitempty"`
	Payments                   []InvoicePayment           `json:"payments,omitempty"`
	CustomFields               []InvoiceCustomField       `json:"custom_fields,omitempty"`
	DisplaySettings            *InvoiceDisplaySettings    `json:"display_settings,omitempty"`
	PublicUrl                  *string                    `json:"public_url,omitempty"`
	PreviousBalanceData        *InvoicePreviousBalance    `json:"previous_balance_data,omitempty"`
}
