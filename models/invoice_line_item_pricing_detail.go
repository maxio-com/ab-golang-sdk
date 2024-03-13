package models

import (
	"encoding/json"
)

// InvoiceLineItemPricingDetail represents a InvoiceLineItemPricingDetail struct.
type InvoiceLineItemPricingDetail struct {
	Label  *string `json:"label,omitempty"`
	Amount *string `json:"amount,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for InvoiceLineItemPricingDetail.
// It customizes the JSON marshaling process for InvoiceLineItemPricingDetail objects.
func (i *InvoiceLineItemPricingDetail) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(i.toMap())
}

// toMap converts the InvoiceLineItemPricingDetail object to a map representation for JSON marshaling.
func (i *InvoiceLineItemPricingDetail) toMap() map[string]any {
	structMap := make(map[string]any)
	if i.Label != nil {
		structMap["label"] = i.Label
	}
	if i.Amount != nil {
		structMap["amount"] = i.Amount
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceLineItemPricingDetail.
// It customizes the JSON unmarshaling process for InvoiceLineItemPricingDetail objects.
func (i *InvoiceLineItemPricingDetail) UnmarshalJSON(input []byte) error {
	var temp invoiceLineItemPricingDetail
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	i.Label = temp.Label
	i.Amount = temp.Amount
	return nil
}

// TODO
type invoiceLineItemPricingDetail struct {
	Label  *string `json:"label,omitempty"`
	Amount *string `json:"amount,omitempty"`
}
