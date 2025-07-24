// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// MRR represents a MRR struct.
type MRR struct {
    AmountInCents        *int64                 `json:"amount_in_cents,omitempty"`
    AmountFormatted      *string                `json:"amount_formatted,omitempty"`
    Currency             *string                `json:"currency,omitempty"`
    CurrencySymbol       *string                `json:"currency_symbol,omitempty"`
    Breakouts            *Breakouts             `json:"breakouts,omitempty"`
    // ISO8601 timestamp
    AtTime               *time.Time             `json:"at_time,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MRR,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MRR) String() string {
    return fmt.Sprintf(
    	"MRR[AmountInCents=%v, AmountFormatted=%v, Currency=%v, CurrencySymbol=%v, Breakouts=%v, AtTime=%v, AdditionalProperties=%v]",
    	m.AmountInCents, m.AmountFormatted, m.Currency, m.CurrencySymbol, m.Breakouts, m.AtTime, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MRR.
// It customizes the JSON marshaling process for MRR objects.
func (m MRR) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "amount_in_cents", "amount_formatted", "currency", "currency_symbol", "breakouts", "at_time"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MRR object to a map representation for JSON marshaling.
func (m MRR) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    if m.AmountInCents != nil {
        structMap["amount_in_cents"] = m.AmountInCents
    }
    if m.AmountFormatted != nil {
        structMap["amount_formatted"] = m.AmountFormatted
    }
    if m.Currency != nil {
        structMap["currency"] = m.Currency
    }
    if m.CurrencySymbol != nil {
        structMap["currency_symbol"] = m.CurrencySymbol
    }
    if m.Breakouts != nil {
        structMap["breakouts"] = m.Breakouts.toMap()
    }
    if m.AtTime != nil {
        structMap["at_time"] = m.AtTime.Format(time.RFC3339)
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MRR.
// It customizes the JSON unmarshaling process for MRR objects.
func (m *MRR) UnmarshalJSON(input []byte) error {
    var temp tempMRR
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "amount_in_cents", "amount_formatted", "currency", "currency_symbol", "breakouts", "at_time")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.AmountInCents = temp.AmountInCents
    m.AmountFormatted = temp.AmountFormatted
    m.Currency = temp.Currency
    m.CurrencySymbol = temp.CurrencySymbol
    m.Breakouts = temp.Breakouts
    if temp.AtTime != nil {
        AtTimeVal, err := time.Parse(time.RFC3339, *temp.AtTime)
        if err != nil {
            log.Fatalf("Cannot Parse at_time as % s format.", time.RFC3339)
        }
        m.AtTime = &AtTimeVal
    }
    return nil
}

// tempMRR is a temporary struct used for validating the fields of MRR.
type tempMRR  struct {
    AmountInCents   *int64     `json:"amount_in_cents,omitempty"`
    AmountFormatted *string    `json:"amount_formatted,omitempty"`
    Currency        *string    `json:"currency,omitempty"`
    CurrencySymbol  *string    `json:"currency_symbol,omitempty"`
    Breakouts       *Breakouts `json:"breakouts,omitempty"`
    AtTime          *string    `json:"at_time,omitempty"`
}
