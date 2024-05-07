package models

import (
    "encoding/json"
)

// InvoiceLineItemPricingDetail represents a InvoiceLineItemPricingDetail struct.
type InvoiceLineItemPricingDetail struct {
    Label                *string        `json:"label,omitempty"`
    Amount               *string        `json:"amount,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for InvoiceLineItemPricingDetail.
// It customizes the JSON marshaling process for InvoiceLineItemPricingDetail objects.
func (i InvoiceLineItemPricingDetail) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(i.toMap())
}

// toMap converts the InvoiceLineItemPricingDetail object to a map representation for JSON marshaling.
func (i InvoiceLineItemPricingDetail) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, i.AdditionalProperties)
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "label", "amount")
    if err != nil {
    	return err
    }
    
    i.AdditionalProperties = additionalProperties
    i.Label = temp.Label
    i.Amount = temp.Amount
    return nil
}

// invoiceLineItemPricingDetail is a temporary struct used for validating the fields of InvoiceLineItemPricingDetail.
type invoiceLineItemPricingDetail  struct {
    Label  *string `json:"label,omitempty"`
    Amount *string `json:"amount,omitempty"`
}
