/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// AllocationPreviewLineItem represents a AllocationPreviewLineItem struct.
type AllocationPreviewLineItem struct {
    // A handle for the line item transaction type
    TransactionType       *LineItemTransactionType       `json:"transaction_type,omitempty"`
    // A handle for the line item kind for allocation preview
    Kind                  *AllocationPreviewLineItemKind `json:"kind,omitempty"`
    AmountInCents         *int64                         `json:"amount_in_cents,omitempty"`
    Memo                  *string                        `json:"memo,omitempty"`
    DiscountAmountInCents *int64                         `json:"discount_amount_in_cents,omitempty"`
    TaxableAmountInCents  *int64                         `json:"taxable_amount_in_cents,omitempty"`
    ComponentId           *int                           `json:"component_id,omitempty"`
    ComponentHandle       *string                        `json:"component_handle,omitempty"`
    // Visible when using Fine-grained Component Control
    Direction             *AllocationPreviewDirection    `json:"direction,omitempty"`
    AdditionalProperties  map[string]interface{}         `json:"_"`
}

// String implements the fmt.Stringer interface for AllocationPreviewLineItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AllocationPreviewLineItem) String() string {
    return fmt.Sprintf(
    	"AllocationPreviewLineItem[TransactionType=%v, Kind=%v, AmountInCents=%v, Memo=%v, DiscountAmountInCents=%v, TaxableAmountInCents=%v, ComponentId=%v, ComponentHandle=%v, Direction=%v, AdditionalProperties=%v]",
    	a.TransactionType, a.Kind, a.AmountInCents, a.Memo, a.DiscountAmountInCents, a.TaxableAmountInCents, a.ComponentId, a.ComponentHandle, a.Direction, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AllocationPreviewLineItem.
// It customizes the JSON marshaling process for AllocationPreviewLineItem objects.
func (a AllocationPreviewLineItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "transaction_type", "kind", "amount_in_cents", "memo", "discount_amount_in_cents", "taxable_amount_in_cents", "component_id", "component_handle", "direction"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AllocationPreviewLineItem object to a map representation for JSON marshaling.
func (a AllocationPreviewLineItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.TransactionType != nil {
        structMap["transaction_type"] = a.TransactionType
    }
    if a.Kind != nil {
        structMap["kind"] = a.Kind
    }
    if a.AmountInCents != nil {
        structMap["amount_in_cents"] = a.AmountInCents
    }
    if a.Memo != nil {
        structMap["memo"] = a.Memo
    }
    if a.DiscountAmountInCents != nil {
        structMap["discount_amount_in_cents"] = a.DiscountAmountInCents
    }
    if a.TaxableAmountInCents != nil {
        structMap["taxable_amount_in_cents"] = a.TaxableAmountInCents
    }
    if a.ComponentId != nil {
        structMap["component_id"] = a.ComponentId
    }
    if a.ComponentHandle != nil {
        structMap["component_handle"] = a.ComponentHandle
    }
    if a.Direction != nil {
        structMap["direction"] = a.Direction
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AllocationPreviewLineItem.
// It customizes the JSON unmarshaling process for AllocationPreviewLineItem objects.
func (a *AllocationPreviewLineItem) UnmarshalJSON(input []byte) error {
    var temp tempAllocationPreviewLineItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "transaction_type", "kind", "amount_in_cents", "memo", "discount_amount_in_cents", "taxable_amount_in_cents", "component_id", "component_handle", "direction")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.TransactionType = temp.TransactionType
    a.Kind = temp.Kind
    a.AmountInCents = temp.AmountInCents
    a.Memo = temp.Memo
    a.DiscountAmountInCents = temp.DiscountAmountInCents
    a.TaxableAmountInCents = temp.TaxableAmountInCents
    a.ComponentId = temp.ComponentId
    a.ComponentHandle = temp.ComponentHandle
    a.Direction = temp.Direction
    return nil
}

// tempAllocationPreviewLineItem is a temporary struct used for validating the fields of AllocationPreviewLineItem.
type tempAllocationPreviewLineItem  struct {
    TransactionType       *LineItemTransactionType       `json:"transaction_type,omitempty"`
    Kind                  *AllocationPreviewLineItemKind `json:"kind,omitempty"`
    AmountInCents         *int64                         `json:"amount_in_cents,omitempty"`
    Memo                  *string                        `json:"memo,omitempty"`
    DiscountAmountInCents *int64                         `json:"discount_amount_in_cents,omitempty"`
    TaxableAmountInCents  *int64                         `json:"taxable_amount_in_cents,omitempty"`
    ComponentId           *int                           `json:"component_id,omitempty"`
    ComponentHandle       *string                        `json:"component_handle,omitempty"`
    Direction             *AllocationPreviewDirection    `json:"direction,omitempty"`
}
