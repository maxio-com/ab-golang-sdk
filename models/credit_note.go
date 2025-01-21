/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// CreditNote represents a CreditNote struct.
type CreditNote struct {
    // Unique identifier for the credit note. It is generated automatically by Chargify and has the prefix "cn_" followed by alphanumeric characters.
    Uid                  *string                 `json:"uid,omitempty"`
    // ID of the site to which the credit note belongs.
    SiteId               *int                    `json:"site_id,omitempty"`
    // ID of the customer to which the credit note belongs.
    CustomerId           *int                    `json:"customer_id,omitempty"`
    // ID of the subscription that generated the credit note.
    SubscriptionId       *int                    `json:"subscription_id,omitempty"`
    // A unique, identifying string that appears on the credit note and in places it is referenced.
    // While the UID is long and not appropriate to show to customers, the number is usually shorter and consumable by the customer and the merchant alike.
    Number               *string                 `json:"number,omitempty"`
    // A monotonically increasing number assigned to credit notes as they are created.  This number is unique within a site and can be used to sort and order credit notes.
    SequenceNumber       *int                    `json:"sequence_number,omitempty"`
    // Date the credit note was issued to the customer.  This is the date that the credit was made available for application, and may come before it is fully applied.
    // The format is `"YYYY-MM-DD"`.
    IssueDate            *time.Time              `json:"issue_date,omitempty"`
    // Credit notes are applied to invoices to offset invoiced amounts - they reduce the amount due. This field is the date the credit note became fully applied to invoices.
    // If the credit note has been partially applied, this field will not have a value until it has been fully applied.
    // The format is `"YYYY-MM-DD"`.
    AppliedDate          *time.Time              `json:"applied_date,omitempty"`
    // Current status of the credit note.
    Status               *CreditNoteStatus       `json:"status,omitempty"`
    // The ISO 4217 currency code (3 character string) representing the currency of the credit note amount fields.
    Currency             *string                 `json:"currency,omitempty"`
    // The memo printed on credit note, which is a description of the reason for the credit.
    Memo                 *string                 `json:"memo,omitempty"`
    // Information about the seller (merchant) listed on the masthead of the credit note.
    Seller               *InvoiceSeller          `json:"seller,omitempty"`
    // Information about the customer who is owner or recipient the credited subscription.
    Customer             *InvoiceCustomer        `json:"customer,omitempty"`
    // The billing address of the credit subscription.
    BillingAddress       *InvoiceAddress         `json:"billing_address,omitempty"`
    // The shipping address of the credited subscription.
    ShippingAddress      *InvoiceAddress         `json:"shipping_address,omitempty"`
    // Subtotal of the credit note, which is the sum of all line items before discounts or taxes. Note that this is a positive amount representing the credit back to the customer.
    SubtotalAmount       *string                 `json:"subtotal_amount,omitempty"`
    // Total discount applied to the credit note. Note that this is a positive amount representing the discount amount being credited back to the customer (i.e. a credit on an earlier discount). For example, if the original purchase was $1.00 and the original discount was $0.10, a credit of $0.50 of the original purchase (half) would have a discount credit of $0.05 (also half).
    DiscountAmount       *string                 `json:"discount_amount,omitempty"`
    // Total tax of the credit note. Note that this is a positive amount representing a previously taxex amount being credited back to the customer (i.e. a credit of an earlier tax). For example, if the original purchase was $1.00 and the original tax was $0.10, a credit of $0.50 of the original purchase (half) would also have a tax credit of $0.05 (also half).
    TaxAmount            *string                 `json:"tax_amount,omitempty"`
    // The credit note total, which is `subtotal_amount - discount_amount + tax_amount`.'
    TotalAmount          *string                 `json:"total_amount,omitempty"`
    // The amount of the credit note that has already been applied to invoices.
    AppliedAmount        *string                 `json:"applied_amount,omitempty"`
    // The amount of the credit note remaining to be applied to invoices, which is `total_amount - applied_amount`.
    RemainingAmount      *string                 `json:"remaining_amount,omitempty"`
    // Line items on the credit note.
    LineItems            []CreditNoteLineItem    `json:"line_items,omitempty"`
    Discounts            []InvoiceDiscount       `json:"discounts,omitempty"`
    Taxes                []InvoiceTax            `json:"taxes,omitempty"`
    Applications         []CreditNoteApplication `json:"applications,omitempty"`
    Refunds              []InvoiceRefund         `json:"refunds,omitempty"`
    // An array of origin invoices for the credit note. Learn more about [Origin Invoice from our docs](https://maxio.zendesk.com/hc/en-us/articles/24252261284749-Credit-Notes-Proration#origin-invoices)
    OriginInvoices       []OriginInvoice         `json:"origin_invoices,omitempty"`
    AdditionalProperties map[string]interface{}  `json:"_"`
}

// String implements the fmt.Stringer interface for CreditNote,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreditNote) String() string {
    return fmt.Sprintf(
    	"CreditNote[Uid=%v, SiteId=%v, CustomerId=%v, SubscriptionId=%v, Number=%v, SequenceNumber=%v, IssueDate=%v, AppliedDate=%v, Status=%v, Currency=%v, Memo=%v, Seller=%v, Customer=%v, BillingAddress=%v, ShippingAddress=%v, SubtotalAmount=%v, DiscountAmount=%v, TaxAmount=%v, TotalAmount=%v, AppliedAmount=%v, RemainingAmount=%v, LineItems=%v, Discounts=%v, Taxes=%v, Applications=%v, Refunds=%v, OriginInvoices=%v, AdditionalProperties=%v]",
    	c.Uid, c.SiteId, c.CustomerId, c.SubscriptionId, c.Number, c.SequenceNumber, c.IssueDate, c.AppliedDate, c.Status, c.Currency, c.Memo, c.Seller, c.Customer, c.BillingAddress, c.ShippingAddress, c.SubtotalAmount, c.DiscountAmount, c.TaxAmount, c.TotalAmount, c.AppliedAmount, c.RemainingAmount, c.LineItems, c.Discounts, c.Taxes, c.Applications, c.Refunds, c.OriginInvoices, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreditNote.
// It customizes the JSON marshaling process for CreditNote objects.
func (c CreditNote) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "uid", "site_id", "customer_id", "subscription_id", "number", "sequence_number", "issue_date", "applied_date", "status", "currency", "memo", "seller", "customer", "billing_address", "shipping_address", "subtotal_amount", "discount_amount", "tax_amount", "total_amount", "applied_amount", "remaining_amount", "line_items", "discounts", "taxes", "applications", "refunds", "origin_invoices"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreditNote object to a map representation for JSON marshaling.
func (c CreditNote) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Uid != nil {
        structMap["uid"] = c.Uid
    }
    if c.SiteId != nil {
        structMap["site_id"] = c.SiteId
    }
    if c.CustomerId != nil {
        structMap["customer_id"] = c.CustomerId
    }
    if c.SubscriptionId != nil {
        structMap["subscription_id"] = c.SubscriptionId
    }
    if c.Number != nil {
        structMap["number"] = c.Number
    }
    if c.SequenceNumber != nil {
        structMap["sequence_number"] = c.SequenceNumber
    }
    if c.IssueDate != nil {
        structMap["issue_date"] = c.IssueDate.Format(DEFAULT_DATE)
    }
    if c.AppliedDate != nil {
        structMap["applied_date"] = c.AppliedDate.Format(DEFAULT_DATE)
    }
    if c.Status != nil {
        structMap["status"] = c.Status
    }
    if c.Currency != nil {
        structMap["currency"] = c.Currency
    }
    if c.Memo != nil {
        structMap["memo"] = c.Memo
    }
    if c.Seller != nil {
        structMap["seller"] = c.Seller.toMap()
    }
    if c.Customer != nil {
        structMap["customer"] = c.Customer.toMap()
    }
    if c.BillingAddress != nil {
        structMap["billing_address"] = c.BillingAddress.toMap()
    }
    if c.ShippingAddress != nil {
        structMap["shipping_address"] = c.ShippingAddress.toMap()
    }
    if c.SubtotalAmount != nil {
        structMap["subtotal_amount"] = c.SubtotalAmount
    }
    if c.DiscountAmount != nil {
        structMap["discount_amount"] = c.DiscountAmount
    }
    if c.TaxAmount != nil {
        structMap["tax_amount"] = c.TaxAmount
    }
    if c.TotalAmount != nil {
        structMap["total_amount"] = c.TotalAmount
    }
    if c.AppliedAmount != nil {
        structMap["applied_amount"] = c.AppliedAmount
    }
    if c.RemainingAmount != nil {
        structMap["remaining_amount"] = c.RemainingAmount
    }
    if c.LineItems != nil {
        structMap["line_items"] = c.LineItems
    }
    if c.Discounts != nil {
        structMap["discounts"] = c.Discounts
    }
    if c.Taxes != nil {
        structMap["taxes"] = c.Taxes
    }
    if c.Applications != nil {
        structMap["applications"] = c.Applications
    }
    if c.Refunds != nil {
        structMap["refunds"] = c.Refunds
    }
    if c.OriginInvoices != nil {
        structMap["origin_invoices"] = c.OriginInvoices
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreditNote.
// It customizes the JSON unmarshaling process for CreditNote objects.
func (c *CreditNote) UnmarshalJSON(input []byte) error {
    var temp tempCreditNote
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "uid", "site_id", "customer_id", "subscription_id", "number", "sequence_number", "issue_date", "applied_date", "status", "currency", "memo", "seller", "customer", "billing_address", "shipping_address", "subtotal_amount", "discount_amount", "tax_amount", "total_amount", "applied_amount", "remaining_amount", "line_items", "discounts", "taxes", "applications", "refunds", "origin_invoices")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Uid = temp.Uid
    c.SiteId = temp.SiteId
    c.CustomerId = temp.CustomerId
    c.SubscriptionId = temp.SubscriptionId
    c.Number = temp.Number
    c.SequenceNumber = temp.SequenceNumber
    if temp.IssueDate != nil {
        IssueDateVal, err := time.Parse(DEFAULT_DATE, *temp.IssueDate)
        if err != nil {
            log.Fatalf("Cannot Parse issue_date as % s format.", DEFAULT_DATE)
        }
        c.IssueDate = &IssueDateVal
    }
    if temp.AppliedDate != nil {
        AppliedDateVal, err := time.Parse(DEFAULT_DATE, *temp.AppliedDate)
        if err != nil {
            log.Fatalf("Cannot Parse applied_date as % s format.", DEFAULT_DATE)
        }
        c.AppliedDate = &AppliedDateVal
    }
    c.Status = temp.Status
    c.Currency = temp.Currency
    c.Memo = temp.Memo
    c.Seller = temp.Seller
    c.Customer = temp.Customer
    c.BillingAddress = temp.BillingAddress
    c.ShippingAddress = temp.ShippingAddress
    c.SubtotalAmount = temp.SubtotalAmount
    c.DiscountAmount = temp.DiscountAmount
    c.TaxAmount = temp.TaxAmount
    c.TotalAmount = temp.TotalAmount
    c.AppliedAmount = temp.AppliedAmount
    c.RemainingAmount = temp.RemainingAmount
    c.LineItems = temp.LineItems
    c.Discounts = temp.Discounts
    c.Taxes = temp.Taxes
    c.Applications = temp.Applications
    c.Refunds = temp.Refunds
    c.OriginInvoices = temp.OriginInvoices
    return nil
}

// tempCreditNote is a temporary struct used for validating the fields of CreditNote.
type tempCreditNote  struct {
    Uid             *string                 `json:"uid,omitempty"`
    SiteId          *int                    `json:"site_id,omitempty"`
    CustomerId      *int                    `json:"customer_id,omitempty"`
    SubscriptionId  *int                    `json:"subscription_id,omitempty"`
    Number          *string                 `json:"number,omitempty"`
    SequenceNumber  *int                    `json:"sequence_number,omitempty"`
    IssueDate       *string                 `json:"issue_date,omitempty"`
    AppliedDate     *string                 `json:"applied_date,omitempty"`
    Status          *CreditNoteStatus       `json:"status,omitempty"`
    Currency        *string                 `json:"currency,omitempty"`
    Memo            *string                 `json:"memo,omitempty"`
    Seller          *InvoiceSeller          `json:"seller,omitempty"`
    Customer        *InvoiceCustomer        `json:"customer,omitempty"`
    BillingAddress  *InvoiceAddress         `json:"billing_address,omitempty"`
    ShippingAddress *InvoiceAddress         `json:"shipping_address,omitempty"`
    SubtotalAmount  *string                 `json:"subtotal_amount,omitempty"`
    DiscountAmount  *string                 `json:"discount_amount,omitempty"`
    TaxAmount       *string                 `json:"tax_amount,omitempty"`
    TotalAmount     *string                 `json:"total_amount,omitempty"`
    AppliedAmount   *string                 `json:"applied_amount,omitempty"`
    RemainingAmount *string                 `json:"remaining_amount,omitempty"`
    LineItems       []CreditNoteLineItem    `json:"line_items,omitempty"`
    Discounts       []InvoiceDiscount       `json:"discounts,omitempty"`
    Taxes           []InvoiceTax            `json:"taxes,omitempty"`
    Applications    []CreditNoteApplication `json:"applications,omitempty"`
    Refunds         []InvoiceRefund         `json:"refunds,omitempty"`
    OriginInvoices  []OriginInvoice         `json:"origin_invoices,omitempty"`
}
