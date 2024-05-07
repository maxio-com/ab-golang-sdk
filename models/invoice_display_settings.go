package models

import (
    "encoding/json"
)

// InvoiceDisplaySettings represents a InvoiceDisplaySettings struct.
type InvoiceDisplaySettings struct {
    HideZeroSubtotalLines   *bool          `json:"hide_zero_subtotal_lines,omitempty"`
    IncludeDiscountsOnLines *bool          `json:"include_discounts_on_lines,omitempty"`
    AdditionalProperties    map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for InvoiceDisplaySettings.
// It customizes the JSON marshaling process for InvoiceDisplaySettings objects.
func (i InvoiceDisplaySettings) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(i.toMap())
}

// toMap converts the InvoiceDisplaySettings object to a map representation for JSON marshaling.
func (i InvoiceDisplaySettings) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, i.AdditionalProperties)
    if i.HideZeroSubtotalLines != nil {
        structMap["hide_zero_subtotal_lines"] = i.HideZeroSubtotalLines
    }
    if i.IncludeDiscountsOnLines != nil {
        structMap["include_discounts_on_lines"] = i.IncludeDiscountsOnLines
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceDisplaySettings.
// It customizes the JSON unmarshaling process for InvoiceDisplaySettings objects.
func (i *InvoiceDisplaySettings) UnmarshalJSON(input []byte) error {
    var temp invoiceDisplaySettings
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "hide_zero_subtotal_lines", "include_discounts_on_lines")
    if err != nil {
    	return err
    }
    
    i.AdditionalProperties = additionalProperties
    i.HideZeroSubtotalLines = temp.HideZeroSubtotalLines
    i.IncludeDiscountsOnLines = temp.IncludeDiscountsOnLines
    return nil
}

// invoiceDisplaySettings is a temporary struct used for validating the fields of InvoiceDisplaySettings.
type invoiceDisplaySettings  struct {
    HideZeroSubtotalLines   *bool `json:"hide_zero_subtotal_lines,omitempty"`
    IncludeDiscountsOnLines *bool `json:"include_discounts_on_lines,omitempty"`
}
