// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// DebitNote represents a DebitNote struct.
type DebitNote struct {
	// Unique identifier for the debit note. It is generated automatically by Chargify and has the prefix "db_" followed by alphanumeric characters.
	Uid *string `json:"uid,omitempty"`
	// ID of the site to which the debit note belongs.
	SiteId *int `json:"site_id,omitempty"`
	// ID of the customer to which the debit note belongs.
	CustomerId *int `json:"customer_id,omitempty"`
	// ID of the subscription that generated the debit note.
	SubscriptionId *int `json:"subscription_id,omitempty"`
	// A unique, identifier that appears on the debit note and in places it is referenced.
	Number *int `json:"number,omitempty"`
	// A monotonically increasing number assigned to debit notes as they are created.
	SequenceNumber *int `json:"sequence_number,omitempty"`
	// Unique identifier for the connected credit note. It is generated automatically by Chargify and has the prefix "cn_" followed by alphanumeric characters.
	// While the UID is long and not appropriate to show to customers, the number is usually shorter and consumable by the customer and the merchant alike.
	OriginCreditNoteUid *string `json:"origin_credit_note_uid,omitempty"`
	// A unique, identifying string of the connected credit note.
	OriginCreditNoteNumber *string `json:"origin_credit_note_number,omitempty"`
	// Date the document was issued to the customer. This is the date that the document was made available for payment.
	// The format is "YYYY-MM-DD".
	IssueDate *time.Time `json:"issue_date,omitempty"`
	// Debit notes are applied to invoices to offset invoiced amounts - they adjust the amount due. This field is the date the debit note document became fully applied to the invoice.
	// The format is "YYYY-MM-DD".
	AppliedDate *time.Time `json:"applied_date,omitempty"`
	// Date the document is due for payment. The format is "YYYY-MM-DD".
	DueDate *time.Time `json:"due_date,omitempty"`
	// Current status of the debit note.
	Status *DebitNoteStatus `json:"status,omitempty"`
	// The memo printed on debit note, which is a description of the reason for the debit.
	Memo *string `json:"memo,omitempty"`
	// The role of the debit note.
	Role *DebitNoteRole `json:"role,omitempty"`
	// The ISO 4217 currency code (3 character string) representing the currency of the credit note amount fields.
	Currency *string `json:"currency,omitempty"`
	// Information about the seller (merchant) listed on the masthead of the debit note.
	Seller *InvoiceSeller `json:"seller,omitempty"`
	// Information about the customer who is owner or recipient the debited subscription.
	Customer *InvoiceCustomer `json:"customer,omitempty"`
	// The billing address of the debited subscription.
	BillingAddress *InvoiceAddress `json:"billing_address,omitempty"`
	// The shipping address of the debited subscription.
	ShippingAddress *InvoiceAddress `json:"shipping_address,omitempty"`
	// Line items on the debit note.
	LineItems            []CreditNoteLineItem   `json:"line_items,omitempty"`
	Discounts            []InvoiceDiscount      `json:"discounts,omitempty"`
	Taxes                []InvoiceTax           `json:"taxes,omitempty"`
	Refunds              []InvoiceRefund        `json:"refunds,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for DebitNote,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (d DebitNote) String() string {
	return fmt.Sprintf(
		"DebitNote[Uid=%v, SiteId=%v, CustomerId=%v, SubscriptionId=%v, Number=%v, SequenceNumber=%v, OriginCreditNoteUid=%v, OriginCreditNoteNumber=%v, IssueDate=%v, AppliedDate=%v, DueDate=%v, Status=%v, Memo=%v, Role=%v, Currency=%v, Seller=%v, Customer=%v, BillingAddress=%v, ShippingAddress=%v, LineItems=%v, Discounts=%v, Taxes=%v, Refunds=%v, AdditionalProperties=%v]",
		d.Uid, d.SiteId, d.CustomerId, d.SubscriptionId, d.Number, d.SequenceNumber, d.OriginCreditNoteUid, d.OriginCreditNoteNumber, d.IssueDate, d.AppliedDate, d.DueDate, d.Status, d.Memo, d.Role, d.Currency, d.Seller, d.Customer, d.BillingAddress, d.ShippingAddress, d.LineItems, d.Discounts, d.Taxes, d.Refunds, d.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for DebitNote.
// It customizes the JSON marshaling process for DebitNote objects.
func (d DebitNote) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(d.AdditionalProperties,
		"uid", "site_id", "customer_id", "subscription_id", "number", "sequence_number", "origin_credit_note_uid", "origin_credit_note_number", "issue_date", "applied_date", "due_date", "status", "memo", "role", "currency", "seller", "customer", "billing_address", "shipping_address", "line_items", "discounts", "taxes", "refunds"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(d.toMap())
}

// toMap converts the DebitNote object to a map representation for JSON marshaling.
func (d DebitNote) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, d.AdditionalProperties)
	if d.Uid != nil {
		structMap["uid"] = d.Uid
	}
	if d.SiteId != nil {
		structMap["site_id"] = d.SiteId
	}
	if d.CustomerId != nil {
		structMap["customer_id"] = d.CustomerId
	}
	if d.SubscriptionId != nil {
		structMap["subscription_id"] = d.SubscriptionId
	}
	if d.Number != nil {
		structMap["number"] = d.Number
	}
	if d.SequenceNumber != nil {
		structMap["sequence_number"] = d.SequenceNumber
	}
	if d.OriginCreditNoteUid != nil {
		structMap["origin_credit_note_uid"] = d.OriginCreditNoteUid
	}
	if d.OriginCreditNoteNumber != nil {
		structMap["origin_credit_note_number"] = d.OriginCreditNoteNumber
	}
	if d.IssueDate != nil {
		structMap["issue_date"] = d.IssueDate.Format(DEFAULT_DATE)
	}
	if d.AppliedDate != nil {
		structMap["applied_date"] = d.AppliedDate.Format(DEFAULT_DATE)
	}
	if d.DueDate != nil {
		structMap["due_date"] = d.DueDate.Format(DEFAULT_DATE)
	}
	if d.Status != nil {
		structMap["status"] = d.Status
	}
	if d.Memo != nil {
		structMap["memo"] = d.Memo
	}
	if d.Role != nil {
		structMap["role"] = d.Role
	}
	if d.Currency != nil {
		structMap["currency"] = d.Currency
	}
	if d.Seller != nil {
		structMap["seller"] = d.Seller.toMap()
	}
	if d.Customer != nil {
		structMap["customer"] = d.Customer.toMap()
	}
	if d.BillingAddress != nil {
		structMap["billing_address"] = d.BillingAddress.toMap()
	}
	if d.ShippingAddress != nil {
		structMap["shipping_address"] = d.ShippingAddress.toMap()
	}
	if d.LineItems != nil {
		structMap["line_items"] = d.LineItems
	}
	if d.Discounts != nil {
		structMap["discounts"] = d.Discounts
	}
	if d.Taxes != nil {
		structMap["taxes"] = d.Taxes
	}
	if d.Refunds != nil {
		structMap["refunds"] = d.Refunds
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DebitNote.
// It customizes the JSON unmarshaling process for DebitNote objects.
func (d *DebitNote) UnmarshalJSON(input []byte) error {
	var temp tempDebitNote
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "uid", "site_id", "customer_id", "subscription_id", "number", "sequence_number", "origin_credit_note_uid", "origin_credit_note_number", "issue_date", "applied_date", "due_date", "status", "memo", "role", "currency", "seller", "customer", "billing_address", "shipping_address", "line_items", "discounts", "taxes", "refunds")
	if err != nil {
		return err
	}
	d.AdditionalProperties = additionalProperties

	d.Uid = temp.Uid
	d.SiteId = temp.SiteId
	d.CustomerId = temp.CustomerId
	d.SubscriptionId = temp.SubscriptionId
	d.Number = temp.Number
	d.SequenceNumber = temp.SequenceNumber
	d.OriginCreditNoteUid = temp.OriginCreditNoteUid
	d.OriginCreditNoteNumber = temp.OriginCreditNoteNumber
	if temp.IssueDate != nil {
		IssueDateVal, err := time.Parse(DEFAULT_DATE, *temp.IssueDate)
		if err != nil {
			log.Fatalf("Cannot Parse issue_date as % s format.", DEFAULT_DATE)
		}
		d.IssueDate = &IssueDateVal
	}
	if temp.AppliedDate != nil {
		AppliedDateVal, err := time.Parse(DEFAULT_DATE, *temp.AppliedDate)
		if err != nil {
			log.Fatalf("Cannot Parse applied_date as % s format.", DEFAULT_DATE)
		}
		d.AppliedDate = &AppliedDateVal
	}
	if temp.DueDate != nil {
		DueDateVal, err := time.Parse(DEFAULT_DATE, *temp.DueDate)
		if err != nil {
			log.Fatalf("Cannot Parse due_date as % s format.", DEFAULT_DATE)
		}
		d.DueDate = &DueDateVal
	}
	d.Status = temp.Status
	d.Memo = temp.Memo
	d.Role = temp.Role
	d.Currency = temp.Currency
	d.Seller = temp.Seller
	d.Customer = temp.Customer
	d.BillingAddress = temp.BillingAddress
	d.ShippingAddress = temp.ShippingAddress
	d.LineItems = temp.LineItems
	d.Discounts = temp.Discounts
	d.Taxes = temp.Taxes
	d.Refunds = temp.Refunds
	return nil
}

// tempDebitNote is a temporary struct used for validating the fields of DebitNote.
type tempDebitNote struct {
	Uid                    *string              `json:"uid,omitempty"`
	SiteId                 *int                 `json:"site_id,omitempty"`
	CustomerId             *int                 `json:"customer_id,omitempty"`
	SubscriptionId         *int                 `json:"subscription_id,omitempty"`
	Number                 *int                 `json:"number,omitempty"`
	SequenceNumber         *int                 `json:"sequence_number,omitempty"`
	OriginCreditNoteUid    *string              `json:"origin_credit_note_uid,omitempty"`
	OriginCreditNoteNumber *string              `json:"origin_credit_note_number,omitempty"`
	IssueDate              *string              `json:"issue_date,omitempty"`
	AppliedDate            *string              `json:"applied_date,omitempty"`
	DueDate                *string              `json:"due_date,omitempty"`
	Status                 *DebitNoteStatus     `json:"status,omitempty"`
	Memo                   *string              `json:"memo,omitempty"`
	Role                   *DebitNoteRole       `json:"role,omitempty"`
	Currency               *string              `json:"currency,omitempty"`
	Seller                 *InvoiceSeller       `json:"seller,omitempty"`
	Customer               *InvoiceCustomer     `json:"customer,omitempty"`
	BillingAddress         *InvoiceAddress      `json:"billing_address,omitempty"`
	ShippingAddress        *InvoiceAddress      `json:"shipping_address,omitempty"`
	LineItems              []CreditNoteLineItem `json:"line_items,omitempty"`
	Discounts              []InvoiceDiscount    `json:"discounts,omitempty"`
	Taxes                  []InvoiceTax         `json:"taxes,omitempty"`
	Refunds                []InvoiceRefund      `json:"refunds,omitempty"`
}
