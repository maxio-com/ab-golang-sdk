package models

import (
	"encoding/json"
)

// Allocation represents a Allocation struct.
type Allocation struct {
	// The integer component ID for the allocation. This references a component that you have created in your Product setup
	ComponentId *int `json:"component_id,omitempty"`
	// The integer subscription ID for the allocation. This references a unique subscription in your Site
	SubscriptionId *int `json:"subscription_id,omitempty"`
	// The allocated quantity set in to effect by the allocation
	Quantity *int `json:"quantity,omitempty"`
	// The allocated quantity that was in effect before this allocation was created
	PreviousQuantity *int `json:"previous_quantity,omitempty"`
	// The memo passed when the allocation was created
	Memo Optional[string] `json:"memo"`
	// The time that the allocation was recorded, in  format and UTC timezone, i.e. 2012-11-20T22:00:37Z
	Timestamp *string `json:"timestamp,omitempty"`
	// The scheme used if the proration was an upgrade. This is only present when the allocation was created mid-period.
	ProrationUpgradeScheme *string `json:"proration_upgrade_scheme,omitempty"`
	// The scheme used if the proration was a downgrade. This is only present when the allocation was created mid-period.
	ProrationDowngradeScheme *string `json:"proration_downgrade_scheme,omitempty"`
	PricePointId             *int    `json:"price_point_id,omitempty"`
	PricePointName           *string `json:"price_point_name,omitempty"`
	PricePointHandle         *string `json:"price_point_handle,omitempty"`
	PreviousPricePointId     *int    `json:"previous_price_point_id,omitempty"`
	// If the change in cost is an upgrade, this determines if the charge should accrue to the next renewal or if capture should be attempted immediately.
	AccrueCharge *bool `json:"accrue_charge,omitempty"`
	// The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
	// Available values: `full`, `prorated`, `none`.
	UpgradeCharge Optional[CreditTypeEnum] `json:"upgrade_charge"`
	// The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
	// Available values: `full`, `prorated`, `none`.
	DowngradeCredit Optional[CreditTypeEnum]    `json:"downgrade_credit"`
	Payment         Optional[AllocationPayment] `json:"payment"`
}

// MarshalJSON implements the json.Marshaler interface for Allocation.
// It customizes the JSON marshaling process for Allocation objects.
func (a *Allocation) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(a.toMap())
}

// toMap converts the Allocation object to a map representation for JSON marshaling.
func (a *Allocation) toMap() map[string]any {
	structMap := make(map[string]any)
	if a.ComponentId != nil {
		structMap["component_id"] = a.ComponentId
	}
	if a.SubscriptionId != nil {
		structMap["subscription_id"] = a.SubscriptionId
	}
	if a.Quantity != nil {
		structMap["quantity"] = a.Quantity
	}
	if a.PreviousQuantity != nil {
		structMap["previous_quantity"] = a.PreviousQuantity
	}
	if a.Memo.IsValueSet() {
		structMap["memo"] = a.Memo.Value()
	}
	if a.Timestamp != nil {
		structMap["timestamp"] = a.Timestamp
	}
	if a.ProrationUpgradeScheme != nil {
		structMap["proration_upgrade_scheme"] = a.ProrationUpgradeScheme
	}
	if a.ProrationDowngradeScheme != nil {
		structMap["proration_downgrade_scheme"] = a.ProrationDowngradeScheme
	}
	if a.PricePointId != nil {
		structMap["price_point_id"] = a.PricePointId
	}
	if a.PricePointName != nil {
		structMap["price_point_name"] = a.PricePointName
	}
	if a.PricePointHandle != nil {
		structMap["price_point_handle"] = a.PricePointHandle
	}
	if a.PreviousPricePointId != nil {
		structMap["previous_price_point_id"] = a.PreviousPricePointId
	}
	if a.AccrueCharge != nil {
		structMap["accrue_charge"] = a.AccrueCharge
	}
	if a.UpgradeCharge.IsValueSet() {
		structMap["upgrade_charge"] = a.UpgradeCharge.Value()
	}
	if a.DowngradeCredit.IsValueSet() {
		structMap["downgrade_credit"] = a.DowngradeCredit.Value()
	}
	if a.Payment.IsValueSet() {
		structMap["payment"] = a.Payment.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Allocation.
// It customizes the JSON unmarshaling process for Allocation objects.
func (a *Allocation) UnmarshalJSON(input []byte) error {
	temp := &struct {
		ComponentId              *int                        `json:"component_id,omitempty"`
		SubscriptionId           *int                        `json:"subscription_id,omitempty"`
		Quantity                 *int                        `json:"quantity,omitempty"`
		PreviousQuantity         *int                        `json:"previous_quantity,omitempty"`
		Memo                     Optional[string]            `json:"memo"`
		Timestamp                *string                     `json:"timestamp,omitempty"`
		ProrationUpgradeScheme   *string                     `json:"proration_upgrade_scheme,omitempty"`
		ProrationDowngradeScheme *string                     `json:"proration_downgrade_scheme,omitempty"`
		PricePointId             *int                        `json:"price_point_id,omitempty"`
		PricePointName           *string                     `json:"price_point_name,omitempty"`
		PricePointHandle         *string                     `json:"price_point_handle,omitempty"`
		PreviousPricePointId     *int                        `json:"previous_price_point_id,omitempty"`
		AccrueCharge             *bool                       `json:"accrue_charge,omitempty"`
		UpgradeCharge            Optional[CreditTypeEnum]    `json:"upgrade_charge"`
		DowngradeCredit          Optional[CreditTypeEnum]    `json:"downgrade_credit"`
		Payment                  Optional[AllocationPayment] `json:"payment"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	a.ComponentId = temp.ComponentId
	a.SubscriptionId = temp.SubscriptionId
	a.Quantity = temp.Quantity
	a.PreviousQuantity = temp.PreviousQuantity
	a.Memo = temp.Memo
	a.Timestamp = temp.Timestamp
	a.ProrationUpgradeScheme = temp.ProrationUpgradeScheme
	a.ProrationDowngradeScheme = temp.ProrationDowngradeScheme
	a.PricePointId = temp.PricePointId
	a.PricePointName = temp.PricePointName
	a.PricePointHandle = temp.PricePointHandle
	a.PreviousPricePointId = temp.PreviousPricePointId
	a.AccrueCharge = temp.AccrueCharge
	a.UpgradeCharge = temp.UpgradeCharge
	a.DowngradeCredit = temp.DowngradeCredit
	a.Payment = temp.Payment
	return nil
}
