package models

import (
	"encoding/json"
)

// UpsertPrepaidConfiguration represents a UpsertPrepaidConfiguration struct.
type UpsertPrepaidConfiguration struct {
	InitialFundingAmountInCents     *int64 `json:"initial_funding_amount_in_cents,omitempty"`
	ReplenishToAmountInCents        *int64 `json:"replenish_to_amount_in_cents,omitempty"`
	AutoReplenish                   *bool  `json:"auto_replenish,omitempty"`
	ReplenishThresholdAmountInCents *int64 `json:"replenish_threshold_amount_in_cents,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for UpsertPrepaidConfiguration.
// It customizes the JSON marshaling process for UpsertPrepaidConfiguration objects.
func (u *UpsertPrepaidConfiguration) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpsertPrepaidConfiguration object to a map representation for JSON marshaling.
func (u *UpsertPrepaidConfiguration) toMap() map[string]any {
	structMap := make(map[string]any)
	if u.InitialFundingAmountInCents != nil {
		structMap["initial_funding_amount_in_cents"] = u.InitialFundingAmountInCents
	}
	if u.ReplenishToAmountInCents != nil {
		structMap["replenish_to_amount_in_cents"] = u.ReplenishToAmountInCents
	}
	if u.AutoReplenish != nil {
		structMap["auto_replenish"] = u.AutoReplenish
	}
	if u.ReplenishThresholdAmountInCents != nil {
		structMap["replenish_threshold_amount_in_cents"] = u.ReplenishThresholdAmountInCents
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpsertPrepaidConfiguration.
// It customizes the JSON unmarshaling process for UpsertPrepaidConfiguration objects.
func (u *UpsertPrepaidConfiguration) UnmarshalJSON(input []byte) error {
	var temp upsertPrepaidConfiguration
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	u.InitialFundingAmountInCents = temp.InitialFundingAmountInCents
	u.ReplenishToAmountInCents = temp.ReplenishToAmountInCents
	u.AutoReplenish = temp.AutoReplenish
	u.ReplenishThresholdAmountInCents = temp.ReplenishThresholdAmountInCents
	return nil
}

// TODO
type upsertPrepaidConfiguration struct {
	InitialFundingAmountInCents     *int64 `json:"initial_funding_amount_in_cents,omitempty"`
	ReplenishToAmountInCents        *int64 `json:"replenish_to_amount_in_cents,omitempty"`
	AutoReplenish                   *bool  `json:"auto_replenish,omitempty"`
	ReplenishThresholdAmountInCents *int64 `json:"replenish_threshold_amount_in_cents,omitempty"`
}
