/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// InvoiceDisplaySettings represents a InvoiceDisplaySettings struct.
type InvoiceDisplaySettings struct {
    HideZeroSubtotalLines   *bool                  `json:"hide_zero_subtotal_lines,omitempty"`
    IncludeDiscountsOnLines *bool                  `json:"include_discounts_on_lines,omitempty"`
    AdditionalProperties    map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for InvoiceDisplaySettings,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i InvoiceDisplaySettings) String() string {
    return fmt.Sprintf(
    	"InvoiceDisplaySettings[HideZeroSubtotalLines=%v, IncludeDiscountsOnLines=%v, AdditionalProperties=%v]",
    	i.HideZeroSubtotalLines, i.IncludeDiscountsOnLines, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for InvoiceDisplaySettings.
// It customizes the JSON marshaling process for InvoiceDisplaySettings objects.
func (i InvoiceDisplaySettings) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(i.AdditionalProperties,
        "hide_zero_subtotal_lines", "include_discounts_on_lines"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(i.toMap())
}

// toMap converts the InvoiceDisplaySettings object to a map representation for JSON marshaling.
func (i InvoiceDisplaySettings) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, i.AdditionalProperties)
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
    var temp tempInvoiceDisplaySettings
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "hide_zero_subtotal_lines", "include_discounts_on_lines")
    if err != nil {
    	return err
    }
    i.AdditionalProperties = additionalProperties
    
    i.HideZeroSubtotalLines = temp.HideZeroSubtotalLines
    i.IncludeDiscountsOnLines = temp.IncludeDiscountsOnLines
    return nil
}

// tempInvoiceDisplaySettings is a temporary struct used for validating the fields of InvoiceDisplaySettings.
type tempInvoiceDisplaySettings  struct {
    HideZeroSubtotalLines   *bool `json:"hide_zero_subtotal_lines,omitempty"`
    IncludeDiscountsOnLines *bool `json:"include_discounts_on_lines,omitempty"`
}
