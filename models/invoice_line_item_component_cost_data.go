package models

import (
    "encoding/json"
)

// InvoiceLineItemComponentCostData represents a InvoiceLineItemComponentCostData struct.
type InvoiceLineItemComponentCostData struct {
    Rates                []ComponentCostData `json:"rates,omitempty"`
    AdditionalProperties map[string]any      `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for InvoiceLineItemComponentCostData.
// It customizes the JSON marshaling process for InvoiceLineItemComponentCostData objects.
func (i InvoiceLineItemComponentCostData) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(i.toMap())
}

// toMap converts the InvoiceLineItemComponentCostData object to a map representation for JSON marshaling.
func (i InvoiceLineItemComponentCostData) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, i.AdditionalProperties)
    if i.Rates != nil {
        structMap["rates"] = i.Rates
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceLineItemComponentCostData.
// It customizes the JSON unmarshaling process for InvoiceLineItemComponentCostData objects.
func (i *InvoiceLineItemComponentCostData) UnmarshalJSON(input []byte) error {
    var temp invoiceLineItemComponentCostData
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "rates")
    if err != nil {
    	return err
    }
    
    i.AdditionalProperties = additionalProperties
    i.Rates = temp.Rates
    return nil
}

// TODO
type invoiceLineItemComponentCostData  struct {
    Rates []ComponentCostData `json:"rates,omitempty"`
}
