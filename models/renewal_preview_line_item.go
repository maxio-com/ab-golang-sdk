// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"fmt"
)

// RenewalPreviewLineItem represents a RenewalPreviewLineItem struct.
type RenewalPreviewLineItem struct {
	// A handle for the line item transaction type
	TransactionType *LineItemTransactionType `json:"transaction_type,omitempty"`
	// A handle for the line item kind
	Kind                  *LineItemKind          `json:"kind,omitempty"`
	AmountInCents         *int64                 `json:"amount_in_cents,omitempty"`
	Memo                  *string                `json:"memo,omitempty"`
	DiscountAmountInCents *int64                 `json:"discount_amount_in_cents,omitempty"`
	TaxableAmountInCents  *int64                 `json:"taxable_amount_in_cents,omitempty"`
	ProductId             *int                   `json:"product_id,omitempty"`
	ProductName           *string                `json:"product_name,omitempty"`
	ComponentId           *int                   `json:"component_id,omitempty"`
	ComponentHandle       *string                `json:"component_handle,omitempty"`
	ComponentName         *string                `json:"component_name,omitempty"`
	ProductHandle         *string                `json:"product_handle,omitempty"`
	PeriodRangeStart      *string                `json:"period_range_start,omitempty"`
	PeriodRangeEnd        *string                `json:"period_range_end,omitempty"`
	AdditionalProperties  map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for RenewalPreviewLineItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RenewalPreviewLineItem) String() string {
	return fmt.Sprintf(
		"RenewalPreviewLineItem[TransactionType=%v, Kind=%v, AmountInCents=%v, Memo=%v, DiscountAmountInCents=%v, TaxableAmountInCents=%v, ProductId=%v, ProductName=%v, ComponentId=%v, ComponentHandle=%v, ComponentName=%v, ProductHandle=%v, PeriodRangeStart=%v, PeriodRangeEnd=%v, AdditionalProperties=%v]",
		r.TransactionType, r.Kind, r.AmountInCents, r.Memo, r.DiscountAmountInCents, r.TaxableAmountInCents, r.ProductId, r.ProductName, r.ComponentId, r.ComponentHandle, r.ComponentName, r.ProductHandle, r.PeriodRangeStart, r.PeriodRangeEnd, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for RenewalPreviewLineItem.
// It customizes the JSON marshaling process for RenewalPreviewLineItem objects.
func (r RenewalPreviewLineItem) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"transaction_type", "kind", "amount_in_cents", "memo", "discount_amount_in_cents", "taxable_amount_in_cents", "product_id", "product_name", "component_id", "component_handle", "component_name", "product_handle", "period_range_start", "period_range_end"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the RenewalPreviewLineItem object to a map representation for JSON marshaling.
func (r RenewalPreviewLineItem) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	if r.TransactionType != nil {
		structMap["transaction_type"] = r.TransactionType
	}
	if r.Kind != nil {
		structMap["kind"] = r.Kind
	}
	if r.AmountInCents != nil {
		structMap["amount_in_cents"] = r.AmountInCents
	}
	if r.Memo != nil {
		structMap["memo"] = r.Memo
	}
	if r.DiscountAmountInCents != nil {
		structMap["discount_amount_in_cents"] = r.DiscountAmountInCents
	}
	if r.TaxableAmountInCents != nil {
		structMap["taxable_amount_in_cents"] = r.TaxableAmountInCents
	}
	if r.ProductId != nil {
		structMap["product_id"] = r.ProductId
	}
	if r.ProductName != nil {
		structMap["product_name"] = r.ProductName
	}
	if r.ComponentId != nil {
		structMap["component_id"] = r.ComponentId
	}
	if r.ComponentHandle != nil {
		structMap["component_handle"] = r.ComponentHandle
	}
	if r.ComponentName != nil {
		structMap["component_name"] = r.ComponentName
	}
	if r.ProductHandle != nil {
		structMap["product_handle"] = r.ProductHandle
	}
	if r.PeriodRangeStart != nil {
		structMap["period_range_start"] = r.PeriodRangeStart
	}
	if r.PeriodRangeEnd != nil {
		structMap["period_range_end"] = r.PeriodRangeEnd
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RenewalPreviewLineItem.
// It customizes the JSON unmarshaling process for RenewalPreviewLineItem objects.
func (r *RenewalPreviewLineItem) UnmarshalJSON(input []byte) error {
	var temp tempRenewalPreviewLineItem
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "transaction_type", "kind", "amount_in_cents", "memo", "discount_amount_in_cents", "taxable_amount_in_cents", "product_id", "product_name", "component_id", "component_handle", "component_name", "product_handle", "period_range_start", "period_range_end")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.TransactionType = temp.TransactionType
	r.Kind = temp.Kind
	r.AmountInCents = temp.AmountInCents
	r.Memo = temp.Memo
	r.DiscountAmountInCents = temp.DiscountAmountInCents
	r.TaxableAmountInCents = temp.TaxableAmountInCents
	r.ProductId = temp.ProductId
	r.ProductName = temp.ProductName
	r.ComponentId = temp.ComponentId
	r.ComponentHandle = temp.ComponentHandle
	r.ComponentName = temp.ComponentName
	r.ProductHandle = temp.ProductHandle
	r.PeriodRangeStart = temp.PeriodRangeStart
	r.PeriodRangeEnd = temp.PeriodRangeEnd
	return nil
}

// tempRenewalPreviewLineItem is a temporary struct used for validating the fields of RenewalPreviewLineItem.
type tempRenewalPreviewLineItem struct {
	TransactionType       *LineItemTransactionType `json:"transaction_type,omitempty"`
	Kind                  *LineItemKind            `json:"kind,omitempty"`
	AmountInCents         *int64                   `json:"amount_in_cents,omitempty"`
	Memo                  *string                  `json:"memo,omitempty"`
	DiscountAmountInCents *int64                   `json:"discount_amount_in_cents,omitempty"`
	TaxableAmountInCents  *int64                   `json:"taxable_amount_in_cents,omitempty"`
	ProductId             *int                     `json:"product_id,omitempty"`
	ProductName           *string                  `json:"product_name,omitempty"`
	ComponentId           *int                     `json:"component_id,omitempty"`
	ComponentHandle       *string                  `json:"component_handle,omitempty"`
	ComponentName         *string                  `json:"component_name,omitempty"`
	ProductHandle         *string                  `json:"product_handle,omitempty"`
	PeriodRangeStart      *string                  `json:"period_range_start,omitempty"`
	PeriodRangeEnd        *string                  `json:"period_range_end,omitempty"`
}
