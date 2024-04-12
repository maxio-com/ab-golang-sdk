package models

import (
    "encoding/json"
)

// InvoiceDiscount represents a InvoiceDiscount struct.
type InvoiceDiscount struct {
    Uid                  *string                    `json:"uid,omitempty"`
    Title                *string                    `json:"title,omitempty"`
    Description          Optional[string]           `json:"description"`
    Code                 *string                    `json:"code,omitempty"`
    SourceType           *InvoiceDiscountSourceType `json:"source_type,omitempty"`
    SourceId             *int                       `json:"source_id,omitempty"`
    DiscountType         *InvoiceDiscountType       `json:"discount_type,omitempty"`
    Percentage           *string                    `json:"percentage,omitempty"`
    EligibleAmount       *string                    `json:"eligible_amount,omitempty"`
    DiscountAmount       *string                    `json:"discount_amount,omitempty"`
    TransactionId        *int                       `json:"transaction_id,omitempty"`
    LineItemBreakouts    []InvoiceDiscountBreakout  `json:"line_item_breakouts,omitempty"`
    AdditionalProperties map[string]any             `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for InvoiceDiscount.
// It customizes the JSON marshaling process for InvoiceDiscount objects.
func (i InvoiceDiscount) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(i.toMap())
}

// toMap converts the InvoiceDiscount object to a map representation for JSON marshaling.
func (i InvoiceDiscount) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, i.AdditionalProperties)
    if i.Uid != nil {
        structMap["uid"] = i.Uid
    }
    if i.Title != nil {
        structMap["title"] = i.Title
    }
    if i.Description.IsValueSet() {
        if i.Description.Value() != nil {
            structMap["description"] = i.Description.Value()
        } else {
            structMap["description"] = nil
        }
    }
    if i.Code != nil {
        structMap["code"] = i.Code
    }
    if i.SourceType != nil {
        structMap["source_type"] = i.SourceType
    }
    if i.SourceId != nil {
        structMap["source_id"] = i.SourceId
    }
    if i.DiscountType != nil {
        structMap["discount_type"] = i.DiscountType
    }
    if i.Percentage != nil {
        structMap["percentage"] = i.Percentage
    }
    if i.EligibleAmount != nil {
        structMap["eligible_amount"] = i.EligibleAmount
    }
    if i.DiscountAmount != nil {
        structMap["discount_amount"] = i.DiscountAmount
    }
    if i.TransactionId != nil {
        structMap["transaction_id"] = i.TransactionId
    }
    if i.LineItemBreakouts != nil {
        structMap["line_item_breakouts"] = i.LineItemBreakouts
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceDiscount.
// It customizes the JSON unmarshaling process for InvoiceDiscount objects.
func (i *InvoiceDiscount) UnmarshalJSON(input []byte) error {
    var temp invoiceDiscount
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "uid", "title", "description", "code", "source_type", "source_id", "discount_type", "percentage", "eligible_amount", "discount_amount", "transaction_id", "line_item_breakouts")
    if err != nil {
    	return err
    }
    
    i.AdditionalProperties = additionalProperties
    i.Uid = temp.Uid
    i.Title = temp.Title
    i.Description = temp.Description
    i.Code = temp.Code
    i.SourceType = temp.SourceType
    i.SourceId = temp.SourceId
    i.DiscountType = temp.DiscountType
    i.Percentage = temp.Percentage
    i.EligibleAmount = temp.EligibleAmount
    i.DiscountAmount = temp.DiscountAmount
    i.TransactionId = temp.TransactionId
    i.LineItemBreakouts = temp.LineItemBreakouts
    return nil
}

// TODO
type invoiceDiscount  struct {
    Uid               *string                    `json:"uid,omitempty"`
    Title             *string                    `json:"title,omitempty"`
    Description       Optional[string]           `json:"description"`
    Code              *string                    `json:"code,omitempty"`
    SourceType        *InvoiceDiscountSourceType `json:"source_type,omitempty"`
    SourceId          *int                       `json:"source_id,omitempty"`
    DiscountType      *InvoiceDiscountType       `json:"discount_type,omitempty"`
    Percentage        *string                    `json:"percentage,omitempty"`
    EligibleAmount    *string                    `json:"eligible_amount,omitempty"`
    DiscountAmount    *string                    `json:"discount_amount,omitempty"`
    TransactionId     *int                       `json:"transaction_id,omitempty"`
    LineItemBreakouts []InvoiceDiscountBreakout  `json:"line_item_breakouts,omitempty"`
}
