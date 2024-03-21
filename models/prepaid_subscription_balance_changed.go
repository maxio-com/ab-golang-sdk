package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// PrepaidSubscriptionBalanceChanged represents a PrepaidSubscriptionBalanceChanged struct.
type PrepaidSubscriptionBalanceChanged struct {
	Reason                          string `json:"reason"`
	CurrentAccountBalanceInCents    int64  `json:"current_account_balance_in_cents"`
	PrepaymentAccountBalanceInCents int64  `json:"prepayment_account_balance_in_cents"`
	CurrentUsageAmountInCents       int64  `json:"current_usage_amount_in_cents"`
}

// MarshalJSON implements the json.Marshaler interface for PrepaidSubscriptionBalanceChanged.
// It customizes the JSON marshaling process for PrepaidSubscriptionBalanceChanged objects.
func (p *PrepaidSubscriptionBalanceChanged) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(p.toMap())
}

// toMap converts the PrepaidSubscriptionBalanceChanged object to a map representation for JSON marshaling.
func (p *PrepaidSubscriptionBalanceChanged) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["reason"] = p.Reason
	structMap["current_account_balance_in_cents"] = p.CurrentAccountBalanceInCents
	structMap["prepayment_account_balance_in_cents"] = p.PrepaymentAccountBalanceInCents
	structMap["current_usage_amount_in_cents"] = p.CurrentUsageAmountInCents
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PrepaidSubscriptionBalanceChanged.
// It customizes the JSON unmarshaling process for PrepaidSubscriptionBalanceChanged objects.
func (p *PrepaidSubscriptionBalanceChanged) UnmarshalJSON(input []byte) error {
	var temp prepaidSubscriptionBalanceChanged
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}

	p.Reason = *temp.Reason
	p.CurrentAccountBalanceInCents = *temp.CurrentAccountBalanceInCents
	p.PrepaymentAccountBalanceInCents = *temp.PrepaymentAccountBalanceInCents
	p.CurrentUsageAmountInCents = *temp.CurrentUsageAmountInCents
	return nil
}

// TODO
type prepaidSubscriptionBalanceChanged struct {
	Reason                          *string `json:"reason"`
	CurrentAccountBalanceInCents    *int64  `json:"current_account_balance_in_cents"`
	PrepaymentAccountBalanceInCents *int64  `json:"prepayment_account_balance_in_cents"`
	CurrentUsageAmountInCents       *int64  `json:"current_usage_amount_in_cents"`
}

func (p *prepaidSubscriptionBalanceChanged) validate() error {
	var errs []string
	if p.Reason == nil {
		errs = append(errs, "required field `reason` is missing for type `Prepaid Subscription Balance Changed`")
	}
	if p.CurrentAccountBalanceInCents == nil {
		errs = append(errs, "required field `current_account_balance_in_cents` is missing for type `Prepaid Subscription Balance Changed`")
	}
	if p.PrepaymentAccountBalanceInCents == nil {
		errs = append(errs, "required field `prepayment_account_balance_in_cents` is missing for type `Prepaid Subscription Balance Changed`")
	}
	if p.CurrentUsageAmountInCents == nil {
		errs = append(errs, "required field `current_usage_amount_in_cents` is missing for type `Prepaid Subscription Balance Changed`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
