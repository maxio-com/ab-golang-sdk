package models

import (
	"encoding/json"
)

// RenewalPreviewLineItem represents a RenewalPreviewLineItem struct.
type RenewalPreviewLineItem struct {
	TransactionType       *string `json:"transaction_type,omitempty"`
	Kind                  *string `json:"kind,omitempty"`
	AmountInCents         *int64  `json:"amount_in_cents,omitempty"`
	Memo                  *string `json:"memo,omitempty"`
	DiscountAmountInCents *int64  `json:"discount_amount_in_cents,omitempty"`
	TaxableAmountInCents  *int64  `json:"taxable_amount_in_cents,omitempty"`
	ProductId             *int    `json:"product_id,omitempty"`
	ProductName           *string `json:"product_name,omitempty"`
	ComponentId           *int    `json:"component_id,omitempty"`
	ComponentHandle       *string `json:"component_handle,omitempty"`
	ComponentName         *string `json:"component_name,omitempty"`
	ProductHandle         *string `json:"product_handle,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for RenewalPreviewLineItem.
// It customizes the JSON marshaling process for RenewalPreviewLineItem objects.
func (r *RenewalPreviewLineItem) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RenewalPreviewLineItem object to a map representation for JSON marshaling.
func (r *RenewalPreviewLineItem) toMap() map[string]any {
	structMap := make(map[string]any)
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
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RenewalPreviewLineItem.
// It customizes the JSON unmarshaling process for RenewalPreviewLineItem objects.
func (r *RenewalPreviewLineItem) UnmarshalJSON(input []byte) error {
	temp := &struct {
		TransactionType       *string `json:"transaction_type,omitempty"`
		Kind                  *string `json:"kind,omitempty"`
		AmountInCents         *int64  `json:"amount_in_cents,omitempty"`
		Memo                  *string `json:"memo,omitempty"`
		DiscountAmountInCents *int64  `json:"discount_amount_in_cents,omitempty"`
		TaxableAmountInCents  *int64  `json:"taxable_amount_in_cents,omitempty"`
		ProductId             *int    `json:"product_id,omitempty"`
		ProductName           *string `json:"product_name,omitempty"`
		ComponentId           *int    `json:"component_id,omitempty"`
		ComponentHandle       *string `json:"component_handle,omitempty"`
		ComponentName         *string `json:"component_name,omitempty"`
		ProductHandle         *string `json:"product_handle,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

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
	return nil
}
