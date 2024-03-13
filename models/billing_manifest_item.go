package models

import (
	"encoding/json"
)

// BillingManifestItem represents a BillingManifestItem struct.
type BillingManifestItem struct {
	// A handle for the line item transaction type
	TransactionType *LineItemTransactionType `json:"transaction_type,omitempty"`
	// A handle for the billing manifest line item kind
	Kind                  *BillingManifestLineItemKind `json:"kind,omitempty"`
	AmountInCents         *int64                       `json:"amount_in_cents,omitempty"`
	Memo                  *string                      `json:"memo,omitempty"`
	DiscountAmountInCents *int64                       `json:"discount_amount_in_cents,omitempty"`
	TaxableAmountInCents  *int64                       `json:"taxable_amount_in_cents,omitempty"`
	ComponentId           *int                         `json:"component_id,omitempty"`
	ComponentHandle       *string                      `json:"component_handle,omitempty"`
	ComponentName         *string                      `json:"component_name,omitempty"`
	ProductId             *int                         `json:"product_id,omitempty"`
	ProductHandle         *string                      `json:"product_handle,omitempty"`
	ProductName           *string                      `json:"product_name,omitempty"`
	PeriodRangeStart      *string                      `json:"period_range_start,omitempty"`
	PeriodRangeEnd        *string                      `json:"period_range_end,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for BillingManifestItem.
// It customizes the JSON marshaling process for BillingManifestItem objects.
func (b *BillingManifestItem) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(b.toMap())
}

// toMap converts the BillingManifestItem object to a map representation for JSON marshaling.
func (b *BillingManifestItem) toMap() map[string]any {
	structMap := make(map[string]any)
	if b.TransactionType != nil {
		structMap["transaction_type"] = b.TransactionType
	}
	if b.Kind != nil {
		structMap["kind"] = b.Kind
	}
	if b.AmountInCents != nil {
		structMap["amount_in_cents"] = b.AmountInCents
	}
	if b.Memo != nil {
		structMap["memo"] = b.Memo
	}
	if b.DiscountAmountInCents != nil {
		structMap["discount_amount_in_cents"] = b.DiscountAmountInCents
	}
	if b.TaxableAmountInCents != nil {
		structMap["taxable_amount_in_cents"] = b.TaxableAmountInCents
	}
	if b.ComponentId != nil {
		structMap["component_id"] = b.ComponentId
	}
	if b.ComponentHandle != nil {
		structMap["component_handle"] = b.ComponentHandle
	}
	if b.ComponentName != nil {
		structMap["component_name"] = b.ComponentName
	}
	if b.ProductId != nil {
		structMap["product_id"] = b.ProductId
	}
	if b.ProductHandle != nil {
		structMap["product_handle"] = b.ProductHandle
	}
	if b.ProductName != nil {
		structMap["product_name"] = b.ProductName
	}
	if b.PeriodRangeStart != nil {
		structMap["period_range_start"] = b.PeriodRangeStart
	}
	if b.PeriodRangeEnd != nil {
		structMap["period_range_end"] = b.PeriodRangeEnd
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BillingManifestItem.
// It customizes the JSON unmarshaling process for BillingManifestItem objects.
func (b *BillingManifestItem) UnmarshalJSON(input []byte) error {
	var temp billingManifestItem
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	b.TransactionType = temp.TransactionType
	b.Kind = temp.Kind
	b.AmountInCents = temp.AmountInCents
	b.Memo = temp.Memo
	b.DiscountAmountInCents = temp.DiscountAmountInCents
	b.TaxableAmountInCents = temp.TaxableAmountInCents
	b.ComponentId = temp.ComponentId
	b.ComponentHandle = temp.ComponentHandle
	b.ComponentName = temp.ComponentName
	b.ProductId = temp.ProductId
	b.ProductHandle = temp.ProductHandle
	b.ProductName = temp.ProductName
	b.PeriodRangeStart = temp.PeriodRangeStart
	b.PeriodRangeEnd = temp.PeriodRangeEnd
	return nil
}

// TODO
type billingManifestItem struct {
	TransactionType       *LineItemTransactionType     `json:"transaction_type,omitempty"`
	Kind                  *BillingManifestLineItemKind `json:"kind,omitempty"`
	AmountInCents         *int64                       `json:"amount_in_cents,omitempty"`
	Memo                  *string                      `json:"memo,omitempty"`
	DiscountAmountInCents *int64                       `json:"discount_amount_in_cents,omitempty"`
	TaxableAmountInCents  *int64                       `json:"taxable_amount_in_cents,omitempty"`
	ComponentId           *int                         `json:"component_id,omitempty"`
	ComponentHandle       *string                      `json:"component_handle,omitempty"`
	ComponentName         *string                      `json:"component_name,omitempty"`
	ProductId             *int                         `json:"product_id,omitempty"`
	ProductHandle         *string                      `json:"product_handle,omitempty"`
	ProductName           *string                      `json:"product_name,omitempty"`
	PeriodRangeStart      *string                      `json:"period_range_start,omitempty"`
	PeriodRangeEnd        *string                      `json:"period_range_end,omitempty"`
}
