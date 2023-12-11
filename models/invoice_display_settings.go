package models

import (
	"encoding/json"
)

// InvoiceDisplaySettings represents a InvoiceDisplaySettings struct.
type InvoiceDisplaySettings struct {
	HideZeroSubtotalLines   *bool `json:"hide_zero_subtotal_lines,omitempty"`
	IncludeDiscountsOnLines *bool `json:"include_discounts_on_lines,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for InvoiceDisplaySettings.
// It customizes the JSON marshaling process for InvoiceDisplaySettings objects.
func (i *InvoiceDisplaySettings) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(i.toMap())
}

// toMap converts the InvoiceDisplaySettings object to a map representation for JSON marshaling.
func (i *InvoiceDisplaySettings) toMap() map[string]any {
	structMap := make(map[string]any)
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
	temp := &struct {
		HideZeroSubtotalLines   *bool `json:"hide_zero_subtotal_lines,omitempty"`
		IncludeDiscountsOnLines *bool `json:"include_discounts_on_lines,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	i.HideZeroSubtotalLines = temp.HideZeroSubtotalLines
	i.IncludeDiscountsOnLines = temp.IncludeDiscountsOnLines
	return nil
}
