// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// UpdateInvoice represents a UpdateInvoice struct.
// Attributes of a draft ad hoc invoice which can be updated. Only the submitted attributes are changed.
type UpdateInvoice struct {
    // Line item changes to apply. Line items without a `uid` are added, line items with a `uid` are updated, and line items with a `uid` and `_destroy` set to `true` are removed. Existing line items not referenced in the array remain unchanged.
    LineItems            []UpdateInvoiceItem    `json:"line_items,omitempty"`
    // New issue date for the invoice (format YYYY-MM-DD). This date is interpreted and validated in your site's time zone. It must be today or a date in the past — future dates are not accepted. The due date is recalculated from the issue date and net terms.
    IssueDate            *time.Time             `json:"issue_date,omitempty"`
    // Number of days after the issue date on which the invoice is due. The due date is recalculated when net terms or the issue date change.
    NetTerms             *int                   `json:"net_terms,omitempty"`
    // Custom payment instructions displayed on the invoice.
    PaymentInstructions  *string                `json:"payment_instructions,omitempty"`
    // A custom memo displayed on the invoice.
    Memo                 *string                `json:"memo,omitempty"`
    // Replaces the seller address on the invoice
    SellerAddress        *CreateInvoiceAddress  `json:"seller_address,omitempty"`
    // Replaces the billing address on the invoice
    BillingAddress       *CreateInvoiceAddress  `json:"billing_address,omitempty"`
    // Replaces the shipping address on the invoice
    ShippingAddress      *CreateInvoiceAddress  `json:"shipping_address,omitempty"`
    // When present, replaces all discounts currently applied to the invoice. Send an empty array to remove all discounts.
    Coupons              []CreateInvoiceCoupon  `json:"coupons,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UpdateInvoice,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpdateInvoice) String() string {
    return fmt.Sprintf(
    	"UpdateInvoice[LineItems=%v, IssueDate=%v, NetTerms=%v, PaymentInstructions=%v, Memo=%v, SellerAddress=%v, BillingAddress=%v, ShippingAddress=%v, Coupons=%v, AdditionalProperties=%v]",
    	u.LineItems, u.IssueDate, u.NetTerms, u.PaymentInstructions, u.Memo, u.SellerAddress, u.BillingAddress, u.ShippingAddress, u.Coupons, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UpdateInvoice.
// It customizes the JSON marshaling process for UpdateInvoice objects.
func (u UpdateInvoice) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "line_items", "issue_date", "net_terms", "payment_instructions", "memo", "seller_address", "billing_address", "shipping_address", "coupons"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateInvoice object to a map representation for JSON marshaling.
func (u UpdateInvoice) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    if u.LineItems != nil {
        structMap["line_items"] = u.LineItems
    }
    if u.IssueDate != nil {
        structMap["issue_date"] = u.IssueDate.Format(DEFAULT_DATE)
    }
    if u.NetTerms != nil {
        structMap["net_terms"] = u.NetTerms
    }
    if u.PaymentInstructions != nil {
        structMap["payment_instructions"] = u.PaymentInstructions
    }
    if u.Memo != nil {
        structMap["memo"] = u.Memo
    }
    if u.SellerAddress != nil {
        structMap["seller_address"] = u.SellerAddress.toMap()
    }
    if u.BillingAddress != nil {
        structMap["billing_address"] = u.BillingAddress.toMap()
    }
    if u.ShippingAddress != nil {
        structMap["shipping_address"] = u.ShippingAddress.toMap()
    }
    if u.Coupons != nil {
        structMap["coupons"] = u.Coupons
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateInvoice.
// It customizes the JSON unmarshaling process for UpdateInvoice objects.
func (u *UpdateInvoice) UnmarshalJSON(input []byte) error {
    var temp tempUpdateInvoice
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "line_items", "issue_date", "net_terms", "payment_instructions", "memo", "seller_address", "billing_address", "shipping_address", "coupons")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.LineItems = temp.LineItems
    if temp.IssueDate != nil {
        IssueDateVal, err := time.Parse(DEFAULT_DATE, *temp.IssueDate)
        if err != nil {
            log.Fatalf("Cannot Parse issue_date as % s format.", DEFAULT_DATE)
        }
        u.IssueDate = &IssueDateVal
    }
    u.NetTerms = temp.NetTerms
    u.PaymentInstructions = temp.PaymentInstructions
    u.Memo = temp.Memo
    u.SellerAddress = temp.SellerAddress
    u.BillingAddress = temp.BillingAddress
    u.ShippingAddress = temp.ShippingAddress
    u.Coupons = temp.Coupons
    return nil
}

// tempUpdateInvoice is a temporary struct used for validating the fields of UpdateInvoice.
type tempUpdateInvoice  struct {
    LineItems           []UpdateInvoiceItem   `json:"line_items,omitempty"`
    IssueDate           *string               `json:"issue_date,omitempty"`
    NetTerms            *int                  `json:"net_terms,omitempty"`
    PaymentInstructions *string               `json:"payment_instructions,omitempty"`
    Memo                *string               `json:"memo,omitempty"`
    SellerAddress       *CreateInvoiceAddress `json:"seller_address,omitempty"`
    BillingAddress      *CreateInvoiceAddress `json:"billing_address,omitempty"`
    ShippingAddress     *CreateInvoiceAddress `json:"shipping_address,omitempty"`
    Coupons             []CreateInvoiceCoupon `json:"coupons,omitempty"`
}
