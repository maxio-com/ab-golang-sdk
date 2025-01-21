/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// AllocationPreviewItem represents a AllocationPreviewItem struct.
type AllocationPreviewItem struct {
    ComponentId              *int                                   `json:"component_id,omitempty"`
    SubscriptionId           *int                                   `json:"subscription_id,omitempty"`
    Quantity                 *AllocationPreviewItemQuantity         `json:"quantity,omitempty"`
    PreviousQuantity         *AllocationPreviewItemPreviousQuantity `json:"previous_quantity,omitempty"`
    Memo                     Optional[string]                       `json:"memo"`
    Timestamp                Optional[string]                       `json:"timestamp"`
    ProrationUpgradeScheme   *string                                `json:"proration_upgrade_scheme,omitempty"`   // Deprecated
    ProrationDowngradeScheme *string                                `json:"proration_downgrade_scheme,omitempty"` // Deprecated
    AccrueCharge             *bool                                  `json:"accrue_charge,omitempty"`
    // The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
    // Available values: `full`, `prorated`, `none`.
    UpgradeCharge            Optional[CreditType]                   `json:"upgrade_charge"`
    // The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
    // Available values: `full`, `prorated`, `none`.
    DowngradeCredit          Optional[CreditType]                   `json:"downgrade_credit"`
    PricePointId             *int                                   `json:"price_point_id,omitempty"`
    // The numerical interval. i.e. an interval of ‘30’ coupled with an interval_unit of day would mean this component price point would renew every 30 days. This property is only available for sites with Multifrequency enabled.
    Interval                 *int                                   `json:"interval,omitempty"`
    // A string representing the interval unit for this component price point, either month or day. This property is only available for sites with Multifrequency enabled.
    IntervalUnit             Optional[IntervalUnit]                 `json:"interval_unit"`
    PreviousPricePointId     *int                                   `json:"previous_price_point_id,omitempty"`
    PricePointHandle         *string                                `json:"price_point_handle,omitempty"`
    PricePointName           *string                                `json:"price_point_name,omitempty"`
    ComponentHandle          Optional[string]                       `json:"component_handle"`
    AdditionalProperties     map[string]interface{}                 `json:"_"`
}

// String implements the fmt.Stringer interface for AllocationPreviewItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AllocationPreviewItem) String() string {
    return fmt.Sprintf(
    	"AllocationPreviewItem[ComponentId=%v, SubscriptionId=%v, Quantity=%v, PreviousQuantity=%v, Memo=%v, Timestamp=%v, ProrationUpgradeScheme=%v, ProrationDowngradeScheme=%v, AccrueCharge=%v, UpgradeCharge=%v, DowngradeCredit=%v, PricePointId=%v, Interval=%v, IntervalUnit=%v, PreviousPricePointId=%v, PricePointHandle=%v, PricePointName=%v, ComponentHandle=%v, AdditionalProperties=%v]",
    	a.ComponentId, a.SubscriptionId, a.Quantity, a.PreviousQuantity, a.Memo, a.Timestamp, a.ProrationUpgradeScheme, a.ProrationDowngradeScheme, a.AccrueCharge, a.UpgradeCharge, a.DowngradeCredit, a.PricePointId, a.Interval, a.IntervalUnit, a.PreviousPricePointId, a.PricePointHandle, a.PricePointName, a.ComponentHandle, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AllocationPreviewItem.
// It customizes the JSON marshaling process for AllocationPreviewItem objects.
func (a AllocationPreviewItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "component_id", "subscription_id", "quantity", "previous_quantity", "memo", "timestamp", "proration_upgrade_scheme", "proration_downgrade_scheme", "accrue_charge", "upgrade_charge", "downgrade_credit", "price_point_id", "interval", "interval_unit", "previous_price_point_id", "price_point_handle", "price_point_name", "component_handle"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AllocationPreviewItem object to a map representation for JSON marshaling.
func (a AllocationPreviewItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.ComponentId != nil {
        structMap["component_id"] = a.ComponentId
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
    if a.Timestamp.IsValueSet() {
        if a.Timestamp.Value() != nil {
            structMap["timestamp"] = a.Timestamp.Value()
        } else {
            structMap["timestamp"] = nil
        }
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
    if a.PricePointId != nil {
        structMap["price_point_id"] = a.PricePointId
    }
    if a.Interval != nil {
        structMap["interval"] = a.Interval
    }
    if a.IntervalUnit.IsValueSet() {
        if a.IntervalUnit.Value() != nil {
            structMap["interval_unit"] = a.IntervalUnit.Value()
        } else {
            structMap["interval_unit"] = nil
        }
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
        if a.ComponentHandle.Value() != nil {
            structMap["component_handle"] = a.ComponentHandle.Value()
        } else {
            structMap["component_handle"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AllocationPreviewItem.
// It customizes the JSON unmarshaling process for AllocationPreviewItem objects.
func (a *AllocationPreviewItem) UnmarshalJSON(input []byte) error {
    var temp tempAllocationPreviewItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "component_id", "subscription_id", "quantity", "previous_quantity", "memo", "timestamp", "proration_upgrade_scheme", "proration_downgrade_scheme", "accrue_charge", "upgrade_charge", "downgrade_credit", "price_point_id", "interval", "interval_unit", "previous_price_point_id", "price_point_handle", "price_point_name", "component_handle")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
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

// tempAllocationPreviewItem is a temporary struct used for validating the fields of AllocationPreviewItem.
type tempAllocationPreviewItem  struct {
    ComponentId              *int                                   `json:"component_id,omitempty"`
    SubscriptionId           *int                                   `json:"subscription_id,omitempty"`
    Quantity                 *AllocationPreviewItemQuantity         `json:"quantity,omitempty"`
    PreviousQuantity         *AllocationPreviewItemPreviousQuantity `json:"previous_quantity,omitempty"`
    Memo                     Optional[string]                       `json:"memo"`
    Timestamp                Optional[string]                       `json:"timestamp"`
    ProrationUpgradeScheme   *string                                `json:"proration_upgrade_scheme,omitempty"`
    ProrationDowngradeScheme *string                                `json:"proration_downgrade_scheme,omitempty"`
    AccrueCharge             *bool                                  `json:"accrue_charge,omitempty"`
    UpgradeCharge            Optional[CreditType]                   `json:"upgrade_charge"`
    DowngradeCredit          Optional[CreditType]                   `json:"downgrade_credit"`
    PricePointId             *int                                   `json:"price_point_id,omitempty"`
    Interval                 *int                                   `json:"interval,omitempty"`
    IntervalUnit             Optional[IntervalUnit]                 `json:"interval_unit"`
    PreviousPricePointId     *int                                   `json:"previous_price_point_id,omitempty"`
    PricePointHandle         *string                                `json:"price_point_handle,omitempty"`
    PricePointName           *string                                `json:"price_point_name,omitempty"`
    ComponentHandle          Optional[string]                       `json:"component_handle"`
}
