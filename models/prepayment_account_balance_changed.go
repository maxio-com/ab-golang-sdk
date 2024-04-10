package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// PrepaymentAccountBalanceChanged represents a PrepaymentAccountBalanceChanged struct.
type PrepaymentAccountBalanceChanged struct {
    Reason                          string         `json:"reason"`
    PrepaymentAccountBalanceInCents int64          `json:"prepayment_account_balance_in_cents"`
    PrepaymentBalanceChangeInCents  int64          `json:"prepayment_balance_change_in_cents"`
    CurrencyCode                    string         `json:"currency_code"`
    AdditionalProperties            map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for PrepaymentAccountBalanceChanged.
// It customizes the JSON marshaling process for PrepaymentAccountBalanceChanged objects.
func (p PrepaymentAccountBalanceChanged) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PrepaymentAccountBalanceChanged object to a map representation for JSON marshaling.
func (p PrepaymentAccountBalanceChanged) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
    structMap["reason"] = p.Reason
    structMap["prepayment_account_balance_in_cents"] = p.PrepaymentAccountBalanceInCents
    structMap["prepayment_balance_change_in_cents"] = p.PrepaymentBalanceChangeInCents
    structMap["currency_code"] = p.CurrencyCode
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PrepaymentAccountBalanceChanged.
// It customizes the JSON unmarshaling process for PrepaymentAccountBalanceChanged objects.
func (p *PrepaymentAccountBalanceChanged) UnmarshalJSON(input []byte) error {
    var temp prepaymentAccountBalanceChanged
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "reason", "prepayment_account_balance_in_cents", "prepayment_balance_change_in_cents", "currency_code")
    if err != nil {
    	return err
    }
    
    p.AdditionalProperties = additionalProperties
    p.Reason = *temp.Reason
    p.PrepaymentAccountBalanceInCents = *temp.PrepaymentAccountBalanceInCents
    p.PrepaymentBalanceChangeInCents = *temp.PrepaymentBalanceChangeInCents
    p.CurrencyCode = *temp.CurrencyCode
    return nil
}

// TODO
type prepaymentAccountBalanceChanged  struct {
    Reason                          *string `json:"reason"`
    PrepaymentAccountBalanceInCents *int64  `json:"prepayment_account_balance_in_cents"`
    PrepaymentBalanceChangeInCents  *int64  `json:"prepayment_balance_change_in_cents"`
    CurrencyCode                    *string `json:"currency_code"`
}

func (p *prepaymentAccountBalanceChanged) validate() error {
    var errs []string
    if p.Reason == nil {
        errs = append(errs, "required field `reason` is missing for type `Prepayment Account Balance Changed`")
    }
    if p.PrepaymentAccountBalanceInCents == nil {
        errs = append(errs, "required field `prepayment_account_balance_in_cents` is missing for type `Prepayment Account Balance Changed`")
    }
    if p.PrepaymentBalanceChangeInCents == nil {
        errs = append(errs, "required field `prepayment_balance_change_in_cents` is missing for type `Prepayment Account Balance Changed`")
    }
    if p.CurrencyCode == nil {
        errs = append(errs, "required field `currency_code` is missing for type `Prepayment Account Balance Changed`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
