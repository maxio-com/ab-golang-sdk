package models

import (
    "encoding/json"
)

// MRR represents a MRR struct.
type MRR struct {
    AmountInCents   *int64     `json:"amount_in_cents,omitempty"`
    AmountFormatted *string    `json:"amount_formatted,omitempty"`
    Currency        *string    `json:"currency,omitempty"`
    CurrencySymbol  *string    `json:"currency_symbol,omitempty"`
    Breakouts       *Breakouts `json:"breakouts,omitempty"`
    // ISO8601 timestamp
    AtTime          *string    `json:"at_time,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for MRR.
// It customizes the JSON marshaling process for MRR objects.
func (m *MRR) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MRR object to a map representation for JSON marshaling.
func (m *MRR) toMap() map[string]any {
    structMap := make(map[string]any)
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
        structMap["breakouts"] = m.Breakouts
    }
    if m.AtTime != nil {
        structMap["at_time"] = m.AtTime
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MRR.
// It customizes the JSON unmarshaling process for MRR objects.
func (m *MRR) UnmarshalJSON(input []byte) error {
    temp := &struct {
        AmountInCents   *int64     `json:"amount_in_cents,omitempty"`
        AmountFormatted *string    `json:"amount_formatted,omitempty"`
        Currency        *string    `json:"currency,omitempty"`
        CurrencySymbol  *string    `json:"currency_symbol,omitempty"`
        Breakouts       *Breakouts `json:"breakouts,omitempty"`
        AtTime          *string    `json:"at_time,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    m.AmountInCents = temp.AmountInCents
    m.AmountFormatted = temp.AmountFormatted
    m.Currency = temp.Currency
    m.CurrencySymbol = temp.CurrencySymbol
    m.Breakouts = temp.Breakouts
    m.AtTime = temp.AtTime
    return nil
}
