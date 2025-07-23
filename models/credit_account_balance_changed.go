// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

// CreditAccountBalanceChanged represents a CreditAccountBalanceChanged struct.
type CreditAccountBalanceChanged struct {
	Reason                             string                 `json:"reason"`
	ServiceCreditAccountBalanceInCents int64                  `json:"service_credit_account_balance_in_cents"`
	ServiceCreditBalanceChangeInCents  int64                  `json:"service_credit_balance_change_in_cents"`
	CurrencyCode                       string                 `json:"currency_code"`
	AtTime                             time.Time              `json:"at_time"`
	AdditionalProperties               map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CreditAccountBalanceChanged,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreditAccountBalanceChanged) String() string {
	return fmt.Sprintf(
		"CreditAccountBalanceChanged[Reason=%v, ServiceCreditAccountBalanceInCents=%v, ServiceCreditBalanceChangeInCents=%v, CurrencyCode=%v, AtTime=%v, AdditionalProperties=%v]",
		c.Reason, c.ServiceCreditAccountBalanceInCents, c.ServiceCreditBalanceChangeInCents, c.CurrencyCode, c.AtTime, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreditAccountBalanceChanged.
// It customizes the JSON marshaling process for CreditAccountBalanceChanged objects.
func (c CreditAccountBalanceChanged) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(c.AdditionalProperties,
		"reason", "service_credit_account_balance_in_cents", "service_credit_balance_change_in_cents", "currency_code", "at_time"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(c.toMap())
}

// toMap converts the CreditAccountBalanceChanged object to a map representation for JSON marshaling.
func (c CreditAccountBalanceChanged) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, c.AdditionalProperties)
	structMap["reason"] = c.Reason
	structMap["service_credit_account_balance_in_cents"] = c.ServiceCreditAccountBalanceInCents
	structMap["service_credit_balance_change_in_cents"] = c.ServiceCreditBalanceChangeInCents
	structMap["currency_code"] = c.CurrencyCode
	structMap["at_time"] = c.AtTime.Format(time.RFC3339)
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreditAccountBalanceChanged.
// It customizes the JSON unmarshaling process for CreditAccountBalanceChanged objects.
func (c *CreditAccountBalanceChanged) UnmarshalJSON(input []byte) error {
	var temp tempCreditAccountBalanceChanged
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "reason", "service_credit_account_balance_in_cents", "service_credit_balance_change_in_cents", "currency_code", "at_time")
	if err != nil {
		return err
	}
	c.AdditionalProperties = additionalProperties

	c.Reason = *temp.Reason
	c.ServiceCreditAccountBalanceInCents = *temp.ServiceCreditAccountBalanceInCents
	c.ServiceCreditBalanceChangeInCents = *temp.ServiceCreditBalanceChangeInCents
	c.CurrencyCode = *temp.CurrencyCode
	AtTimeVal, err := time.Parse(time.RFC3339, *temp.AtTime)
	if err != nil {
		log.Fatalf("Cannot Parse at_time as % s format.", time.RFC3339)
	}
	c.AtTime = AtTimeVal
	return nil
}

// tempCreditAccountBalanceChanged is a temporary struct used for validating the fields of CreditAccountBalanceChanged.
type tempCreditAccountBalanceChanged struct {
	Reason                             *string `json:"reason"`
	ServiceCreditAccountBalanceInCents *int64  `json:"service_credit_account_balance_in_cents"`
	ServiceCreditBalanceChangeInCents  *int64  `json:"service_credit_balance_change_in_cents"`
	CurrencyCode                       *string `json:"currency_code"`
	AtTime                             *string `json:"at_time"`
}

func (c *tempCreditAccountBalanceChanged) validate() error {
	var errs []string
	if c.Reason == nil {
		errs = append(errs, "required field `reason` is missing for type `Credit Account Balance Changed`")
	}
	if c.ServiceCreditAccountBalanceInCents == nil {
		errs = append(errs, "required field `service_credit_account_balance_in_cents` is missing for type `Credit Account Balance Changed`")
	}
	if c.ServiceCreditBalanceChangeInCents == nil {
		errs = append(errs, "required field `service_credit_balance_change_in_cents` is missing for type `Credit Account Balance Changed`")
	}
	if c.CurrencyCode == nil {
		errs = append(errs, "required field `currency_code` is missing for type `Credit Account Balance Changed`")
	}
	if c.AtTime == nil {
		errs = append(errs, "required field `at_time` is missing for type `Credit Account Balance Changed`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
