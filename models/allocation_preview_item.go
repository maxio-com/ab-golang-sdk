package models

import (
    "encoding/json"
)

// AllocationPreviewItem represents a AllocationPreviewItem struct.
type AllocationPreviewItem struct {
    ComponentId              *int                 `json:"component_id,omitempty"`
    SubscriptionId           *int                 `json:"subscription_id,omitempty"`
    Quantity                 *interface{}         `json:"quantity,omitempty"`
    PreviousQuantity         *interface{}         `json:"previous_quantity,omitempty"`
    Memo                     Optional[string]     `json:"memo"`
    Timestamp                Optional[string]     `json:"timestamp"`
    ProrationUpgradeScheme   *string              `json:"proration_upgrade_scheme,omitempty"`   // Deprecated
    ProrationDowngradeScheme *string              `json:"proration_downgrade_scheme,omitempty"` // Deprecated
    AccrueCharge             *bool                `json:"accrue_charge,omitempty"`
    // The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
    // Available values: `full`, `prorated`, `none`.
    UpgradeCharge            Optional[CreditType] `json:"upgrade_charge"`
    // The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
    // Available values: `full`, `prorated`, `none`.
    DowngradeCredit          Optional[CreditType] `json:"downgrade_credit"`
    PricePointId             *int                 `json:"price_point_id,omitempty"`
    // The numerical interval. i.e. an interval of ‘30’ coupled with an interval_unit of day would mean this component price point would renew every 30 days. This property is only available for sites with Multifrequency enabled.
    Interval                 *int                 `json:"interval,omitempty"`
    // A string representing the interval unit for this component price point, either month or day. This property is only available for sites with Multifrequency enabled.
    IntervalUnit             *IntervalUnit        `json:"interval_unit,omitempty"`
    PreviousPricePointId     *int                 `json:"previous_price_point_id,omitempty"`
    PricePointHandle         *string              `json:"price_point_handle,omitempty"`
    PricePointName           *string              `json:"price_point_name,omitempty"`
    ComponentHandle          Optional[string]     `json:"component_handle"`
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
    if a.Memo.IsValueSet() {
        structMap["memo"] = a.Memo.Value()
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
    if a.Interval != nil {
        structMap["interval"] = a.Interval
    }
    if a.IntervalUnit != nil {
        structMap["interval_unit"] = a.IntervalUnit
    }
    if a.PreviousPricePointId != nil {
        structMap["previous_price_point_id"] = a.PreviousPricePointId
    }
    if a.PricePointHandle != nil {
        structMap["price_point_handle"] = a.PricePointHandle
    }
    if a.PricePointName != nil {
        structMap["price_point_name"] = a.PricePointName
    }
    if a.ComponentHandle.IsValueSet() {
        structMap["component_handle"] = a.ComponentHandle.Value()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AllocationPreviewItem.
// It customizes the JSON unmarshaling process for AllocationPreviewItem objects.
func (a *AllocationPreviewItem) UnmarshalJSON(input []byte) error {
    temp := &struct {
        ComponentId              *int                 `json:"component_id,omitempty"`
        SubscriptionId           *int                 `json:"subscription_id,omitempty"`
        Quantity                 *interface{}         `json:"quantity,omitempty"`
        PreviousQuantity         *interface{}         `json:"previous_quantity,omitempty"`
        Memo                     Optional[string]     `json:"memo"`
        Timestamp                Optional[string]     `json:"timestamp"`
        ProrationUpgradeScheme   *string              `json:"proration_upgrade_scheme,omitempty"`
        ProrationDowngradeScheme *string              `json:"proration_downgrade_scheme,omitempty"`
        AccrueCharge             *bool                `json:"accrue_charge,omitempty"`
        UpgradeCharge            Optional[CreditType] `json:"upgrade_charge"`
        DowngradeCredit          Optional[CreditType] `json:"downgrade_credit"`
        PricePointId             *int                 `json:"price_point_id,omitempty"`
        Interval                 *int                 `json:"interval,omitempty"`
        IntervalUnit             *IntervalUnit        `json:"interval_unit,omitempty"`
        PreviousPricePointId     *int                 `json:"previous_price_point_id,omitempty"`
        PricePointHandle         *string              `json:"price_point_handle,omitempty"`
        PricePointName           *string              `json:"price_point_name,omitempty"`
        ComponentHandle          Optional[string]     `json:"component_handle"`
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
    a.Interval = temp.Interval
    a.IntervalUnit = temp.IntervalUnit
    a.PreviousPricePointId = temp.PreviousPricePointId
    a.PricePointHandle = temp.PricePointHandle
    a.PricePointName = temp.PricePointName
    a.ComponentHandle = temp.ComponentHandle
    return nil
}
