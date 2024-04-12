package models

import (
    "encoding/json"
    "log"
    "time"
)

// Allocation represents a Allocation struct.
type Allocation struct {
    // The allocation unique id
    AllocationId             *int                           `json:"allocation_id,omitempty"`
    // The integer component ID for the allocation. This references a component that you have created in your Product setup
    ComponentId              *int                           `json:"component_id,omitempty"`
    // The handle of the component. This references a component that you have created in your Product setup
    ComponentHandle          Optional[string]               `json:"component_handle"`
    // The integer subscription ID for the allocation. This references a unique subscription in your Site
    SubscriptionId           *int                           `json:"subscription_id,omitempty"`
    // The allocated quantity set in to effect by the allocation. String for components supporting fractional quantities
    Quantity                 *AllocationQuantity            `json:"quantity,omitempty"`
    // The allocated quantity that was in effect before this allocation was created. String for components supporting fractional quantities
    PreviousQuantity         *AllocationPreviousQuantity    `json:"previous_quantity,omitempty"`
    // The memo passed when the allocation was created
    Memo                     Optional[string]               `json:"memo"`
    // The time that the allocation was recorded, in format and UTC timezone, i.e. 2012-11-20T22:00:37Z
    Timestamp                *time.Time                     `json:"timestamp,omitempty"`
    // Timestamp indicating when this allocation was created
    CreatedAt                *time.Time                     `json:"created_at,omitempty"`
    // The scheme used if the proration was an upgrade. This is only present when the allocation was created mid-period.
    ProrationUpgradeScheme   *string                        `json:"proration_upgrade_scheme,omitempty"`   // Deprecated
    // The scheme used if the proration was a downgrade. This is only present when the allocation was created mid-period.
    ProrationDowngradeScheme *string                        `json:"proration_downgrade_scheme,omitempty"` // Deprecated
    PricePointId             *int                           `json:"price_point_id,omitempty"`
    PricePointName           *string                        `json:"price_point_name,omitempty"`
    PricePointHandle         *string                        `json:"price_point_handle,omitempty"`
    // The numerical interval. i.e. an interval of ‘30’ coupled with an interval_unit of day would mean this component price point would renew every 30 days. This property is only available for sites with Multifrequency enabled.
    Interval                 *int                           `json:"interval,omitempty"`
    // A string representing the interval unit for this component price point, either month or day. This property is only available for sites with Multifrequency enabled.
    IntervalUnit             *IntervalUnit                  `json:"interval_unit,omitempty"`
    PreviousPricePointId     *int                           `json:"previous_price_point_id,omitempty"`
    // If the change in cost is an upgrade, this determines if the charge should accrue to the next renewal or if capture should be attempted immediately.
    AccrueCharge             *bool                          `json:"accrue_charge,omitempty"`
    // If true, if the immediate component payment fails, initiate dunning for the subscription. 
    // Otherwise, leave the charges on the subscription to pay for at renewal.
    InitiateDunning          *bool                          `json:"initiate_dunning,omitempty"`
    // The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
    // Available values: `full`, `prorated`, `none`.
    UpgradeCharge            Optional[CreditType]           `json:"upgrade_charge"`
    // The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
    // Available values: `full`, `prorated`, `none`.
    DowngradeCredit          Optional[CreditType]           `json:"downgrade_credit"`
    Payment                  Optional[PaymentForAllocation] `json:"payment"`
    ExpiresAt                *time.Time                     `json:"expires_at,omitempty"`
    UsedQuantity             *int64                         `json:"used_quantity,omitempty"`
    ChargeId                 *int64                         `json:"charge_id,omitempty"`
    AdditionalProperties     map[string]any                 `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Allocation.
// It customizes the JSON marshaling process for Allocation objects.
func (a Allocation) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the Allocation object to a map representation for JSON marshaling.
func (a Allocation) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.AllocationId != nil {
        structMap["allocation_id"] = a.AllocationId
    }
    if a.ComponentId != nil {
        structMap["component_id"] = a.ComponentId
    }
    if a.ComponentHandle.IsValueSet() {
        if a.ComponentHandle.Value() != nil {
            structMap["component_handle"] = a.ComponentHandle.Value()
        } else {
            structMap["component_handle"] = nil
        }
    }
    if a.SubscriptionId != nil {
        structMap["subscription_id"] = a.SubscriptionId
    }
    if a.Quantity != nil {
        structMap["quantity"] = a.Quantity.toMap()
    }
    if a.PreviousQuantity != nil {
        structMap["previous_quantity"] = a.PreviousQuantity.toMap()
    }
    if a.Memo.IsValueSet() {
        if a.Memo.Value() != nil {
            structMap["memo"] = a.Memo.Value()
        } else {
            structMap["memo"] = nil
        }
    }
    if a.Timestamp != nil {
        structMap["timestamp"] = a.Timestamp.Format(time.RFC3339)
    }
    if a.CreatedAt != nil {
        structMap["created_at"] = a.CreatedAt.Format(time.RFC3339)
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
    if a.Interval != nil {
        structMap["interval"] = a.Interval
    }
    if a.IntervalUnit != nil {
        structMap["interval_unit"] = a.IntervalUnit
    }
    if a.PreviousPricePointId != nil {
        structMap["previous_price_point_id"] = a.PreviousPricePointId
    }
    if a.AccrueCharge != nil {
        structMap["accrue_charge"] = a.AccrueCharge
    }
    if a.InitiateDunning != nil {
        structMap["initiate_dunning"] = a.InitiateDunning
    }
    if a.UpgradeCharge.IsValueSet() {
        if a.UpgradeCharge.Value() != nil {
            structMap["upgrade_charge"] = a.UpgradeCharge.Value()
        } else {
            structMap["upgrade_charge"] = nil
        }
    }
    if a.DowngradeCredit.IsValueSet() {
        if a.DowngradeCredit.Value() != nil {
            structMap["downgrade_credit"] = a.DowngradeCredit.Value()
        } else {
            structMap["downgrade_credit"] = nil
        }
    }
    if a.Payment.IsValueSet() {
        if a.Payment.Value() != nil {
            structMap["payment"] = a.Payment.Value().toMap()
        } else {
            structMap["payment"] = nil
        }
    }
    if a.ExpiresAt != nil {
        structMap["expires_at"] = a.ExpiresAt.Format(time.RFC3339)
    }
    if a.UsedQuantity != nil {
        structMap["used_quantity"] = a.UsedQuantity
    }
    if a.ChargeId != nil {
        structMap["charge_id"] = a.ChargeId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Allocation.
// It customizes the JSON unmarshaling process for Allocation objects.
func (a *Allocation) UnmarshalJSON(input []byte) error {
    var temp allocation
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "allocation_id", "component_id", "component_handle", "subscription_id", "quantity", "previous_quantity", "memo", "timestamp", "created_at", "proration_upgrade_scheme", "proration_downgrade_scheme", "price_point_id", "price_point_name", "price_point_handle", "interval", "interval_unit", "previous_price_point_id", "accrue_charge", "initiate_dunning", "upgrade_charge", "downgrade_credit", "payment", "expires_at", "used_quantity", "charge_id")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.AllocationId = temp.AllocationId
    a.ComponentId = temp.ComponentId
    a.ComponentHandle = temp.ComponentHandle
    a.SubscriptionId = temp.SubscriptionId
    a.Quantity = temp.Quantity
    a.PreviousQuantity = temp.PreviousQuantity
    a.Memo = temp.Memo
    if temp.Timestamp != nil {
        TimestampVal, err := time.Parse(time.RFC3339, *temp.Timestamp)
        if err != nil {
            log.Fatalf("Cannot Parse timestamp as % s format.", time.RFC3339)
        }
        a.Timestamp = &TimestampVal
    }
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        a.CreatedAt = &CreatedAtVal
    }
    a.ProrationUpgradeScheme = temp.ProrationUpgradeScheme
    a.ProrationDowngradeScheme = temp.ProrationDowngradeScheme
    a.PricePointId = temp.PricePointId
    a.PricePointName = temp.PricePointName
    a.PricePointHandle = temp.PricePointHandle
    a.Interval = temp.Interval
    a.IntervalUnit = temp.IntervalUnit
    a.PreviousPricePointId = temp.PreviousPricePointId
    a.AccrueCharge = temp.AccrueCharge
    a.InitiateDunning = temp.InitiateDunning
    a.UpgradeCharge = temp.UpgradeCharge
    a.DowngradeCredit = temp.DowngradeCredit
    a.Payment = temp.Payment
    if temp.ExpiresAt != nil {
        ExpiresAtVal, err := time.Parse(time.RFC3339, *temp.ExpiresAt)
        if err != nil {
            log.Fatalf("Cannot Parse expires_at as % s format.", time.RFC3339)
        }
        a.ExpiresAt = &ExpiresAtVal
    }
    a.UsedQuantity = temp.UsedQuantity
    a.ChargeId = temp.ChargeId
    return nil
}

// TODO
type allocation  struct {
    AllocationId             *int                           `json:"allocation_id,omitempty"`
    ComponentId              *int                           `json:"component_id,omitempty"`
    ComponentHandle          Optional[string]               `json:"component_handle"`
    SubscriptionId           *int                           `json:"subscription_id,omitempty"`
    Quantity                 *AllocationQuantity            `json:"quantity,omitempty"`
    PreviousQuantity         *AllocationPreviousQuantity    `json:"previous_quantity,omitempty"`
    Memo                     Optional[string]               `json:"memo"`
    Timestamp                *string                        `json:"timestamp,omitempty"`
    CreatedAt                *string                        `json:"created_at,omitempty"`
    ProrationUpgradeScheme   *string                        `json:"proration_upgrade_scheme,omitempty"`
    ProrationDowngradeScheme *string                        `json:"proration_downgrade_scheme,omitempty"`
    PricePointId             *int                           `json:"price_point_id,omitempty"`
    PricePointName           *string                        `json:"price_point_name,omitempty"`
    PricePointHandle         *string                        `json:"price_point_handle,omitempty"`
    Interval                 *int                           `json:"interval,omitempty"`
    IntervalUnit             *IntervalUnit                  `json:"interval_unit,omitempty"`
    PreviousPricePointId     *int                           `json:"previous_price_point_id,omitempty"`
    AccrueCharge             *bool                          `json:"accrue_charge,omitempty"`
    InitiateDunning          *bool                          `json:"initiate_dunning,omitempty"`
    UpgradeCharge            Optional[CreditType]           `json:"upgrade_charge"`
    DowngradeCredit          Optional[CreditType]           `json:"downgrade_credit"`
    Payment                  Optional[PaymentForAllocation] `json:"payment"`
    ExpiresAt                *string                        `json:"expires_at,omitempty"`
    UsedQuantity             *int64                         `json:"used_quantity,omitempty"`
    ChargeId                 *int64                         `json:"charge_id,omitempty"`
}
