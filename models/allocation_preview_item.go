package models

import (
	"encoding/json"
)

// AllocationPreviewItem represents a AllocationPreviewItem struct.
type AllocationPreviewItem struct {
	ComponentId              *int             `json:"component_id,omitempty"`
	SubscriptionId           *int             `json:"subscription_id,omitempty"`
	Quantity                 *float64         `json:"quantity,omitempty"`
	PreviousQuantity         *int             `json:"previous_quantity,omitempty"`
	Memo                     *string          `json:"memo,omitempty"`
	Timestamp                Optional[string] `json:"timestamp"`
	ProrationUpgradeScheme   *string          `json:"proration_upgrade_scheme,omitempty"`
	ProrationDowngradeScheme *string          `json:"proration_downgrade_scheme,omitempty"`
	AccrueCharge             *bool            `json:"accrue_charge,omitempty"`
	// The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
	// Available values: `full`, `prorated`, `none`.
	UpgradeCharge Optional[CreditTypeEnum] `json:"upgrade_charge"`
	// The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
	// Available values: `full`, `prorated`, `none`.
	DowngradeCredit      Optional[CreditTypeEnum] `json:"downgrade_credit"`
	PricePointId         *int                     `json:"price_point_id,omitempty"`
	PreviousPricePointId *int                     `json:"previous_price_point_id,omitempty"`
	ComponentHandle      *string                  `json:"component_handle,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for AllocationPreviewItem.
// It customizes the JSON marshaling process for AllocationPreviewItem objects.
func (a *AllocationPreviewItem) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(a.toMap())
}

// toMap converts the AllocationPreviewItem object to a map representation for JSON marshaling.
func (a *AllocationPreviewItem) toMap() map[string]any {
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
	if a.Memo != nil {
		structMap["memo"] = a.Memo
	}
	if a.Timestamp.IsValueSet() {
		structMap["timestamp"] = a.Timestamp.Value()
	}
	if a.ProrationUpgradeScheme != nil {
		structMap["proration_upgrade_scheme"] = a.ProrationUpgradeScheme
	}
	if a.ProrationDowngradeScheme != nil {
		structMap["proration_downgrade_scheme"] = a.ProrationDowngradeScheme
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
	if a.PricePointId != nil {
		structMap["price_point_id"] = a.PricePointId
	}
	if a.PreviousPricePointId != nil {
		structMap["previous_price_point_id"] = a.PreviousPricePointId
	}
	if a.ComponentHandle != nil {
		structMap["component_handle"] = a.ComponentHandle
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AllocationPreviewItem.
// It customizes the JSON unmarshaling process for AllocationPreviewItem objects.
func (a *AllocationPreviewItem) UnmarshalJSON(input []byte) error {
	temp := &struct {
		ComponentId              *int                     `json:"component_id,omitempty"`
		SubscriptionId           *int                     `json:"subscription_id,omitempty"`
		Quantity                 *float64                 `json:"quantity,omitempty"`
		PreviousQuantity         *int                     `json:"previous_quantity,omitempty"`
		Memo                     *string                  `json:"memo,omitempty"`
		Timestamp                Optional[string]         `json:"timestamp"`
		ProrationUpgradeScheme   *string                  `json:"proration_upgrade_scheme,omitempty"`
		ProrationDowngradeScheme *string                  `json:"proration_downgrade_scheme,omitempty"`
		AccrueCharge             *bool                    `json:"accrue_charge,omitempty"`
		UpgradeCharge            Optional[CreditTypeEnum] `json:"upgrade_charge"`
		DowngradeCredit          Optional[CreditTypeEnum] `json:"downgrade_credit"`
		PricePointId             *int                     `json:"price_point_id,omitempty"`
		PreviousPricePointId     *int                     `json:"previous_price_point_id,omitempty"`
		ComponentHandle          *string                  `json:"component_handle,omitempty"`
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
	a.AccrueCharge = temp.AccrueCharge
	a.UpgradeCharge = temp.UpgradeCharge
	a.DowngradeCredit = temp.DowngradeCredit
	a.PricePointId = temp.PricePointId
	a.PreviousPricePointId = temp.PreviousPricePointId
	a.ComponentHandle = temp.ComponentHandle
	return nil
}
