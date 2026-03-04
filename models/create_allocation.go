// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CreateAllocation represents a CreateAllocation struct.
type CreateAllocation struct {
    // The allocated quantity to which to set the line-items allocated quantity. By default, this is an integer. If decimal allocations are enabled for the component, it will be a decimal number. For On/Off components, use 1for on and 0 for off.
    Quantity                 float64                                `json:"quantity"`
    // Decimal representation of the allocated quantity. Only valid when decimal
    // allocations are enabled for the component.
    DecimalQuantity          *string                                `json:"decimal_quantity,omitempty"`
    // The quantity that was in effect before this allocation. Responses always
    // include this value; it may be supplied on preview requests to ensure the
    // expected change is evaluated.
    PreviousQuantity         *float64                               `json:"previous_quantity,omitempty"`
    // Decimal representation of `previous_quantity`. Only valid when decimal
    // allocations are enabled for the component.
    DecimalPreviousQuantity  *string                                `json:"decimal_previous_quantity,omitempty"`
    // (required for the multiple allocations endpoint) The id associated with the component for which the allocation is being made.
    ComponentId              *int                                   `json:"component_id,omitempty"`
    // A memo to record along with the allocation.
    Memo                     *string                                `json:"memo,omitempty"`
    // The scheme used if the proration is a downgrade. Defaults to the site setting if one is not provided.
    ProrationDowngradeScheme *string                                `json:"proration_downgrade_scheme,omitempty"` // Deprecated
    // The scheme used if the proration is an upgrade. Defaults to the site setting if one is not provided.
    ProrationUpgradeScheme   *string                                `json:"proration_upgrade_scheme,omitempty"`   // Deprecated
    // The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided. Values are:
    // `full` -  A full price credit is added for the amount owed.
    // `prorated` - A prorated credit is added for the amount owed.
    // `none` - No charge is added.
    DowngradeCredit          Optional[DowngradeCreditCreditType]    `json:"downgrade_credit"`
    // The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided. Values are:
    // `full` - A charge is added for the full price of the component.
    // `prorated` - A charge is added for the prorated price of the component change.
    // `none` - No charge is added.
    UpgradeCharge            Optional[UpgradeChargeCreditType]      `json:"upgrade_charge"`
    // "If the change in cost is an upgrade, this determines if the charge should accrue to the next renewal or if capture should be attempted immediately.
    // `true` - Attempt to charge the customer at the next renewal.
    // `false` - Attempt to charge the customer right away. If it fails, the charge will be accrued until the next renewal.
    // Defaults to the site setting if unspecified in the request.
    AccrueCharge             *bool                                  `json:"accrue_charge,omitempty"`
    // If set to true, if the immediate component payment fails, initiate dunning for the subscription.
    // Otherwise, leave the charges on the subscription to pay for at renewal. Defaults to false.
    InitiateDunning          *bool                                  `json:"initiate_dunning,omitempty"`
    // Price point that the allocation should be charged at. Accepts either the price point's id (integer) or handle (string). When not specified, the default price point will be used.
    PricePointId             Optional[CreateAllocationPricePointId] `json:"price_point_id"`
    // This attribute is particularly useful when you need to align billing events for different components on distinct schedules within a subscription. This only works for site with Multifrequency enabled.
    BillingSchedule          *BillingSchedule                       `json:"billing_schedule,omitempty"`
    // Create or update custom pricing unique to the subscription. Used in place of `price_point_id`.
    CustomPrice              *ComponentCustomPrice                  `json:"custom_price,omitempty"`
    AdditionalProperties     map[string]interface{}                 `json:"_"`
}

// String implements the fmt.Stringer interface for CreateAllocation,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateAllocation) String() string {
    return fmt.Sprintf(
    	"CreateAllocation[Quantity=%v, DecimalQuantity=%v, PreviousQuantity=%v, DecimalPreviousQuantity=%v, ComponentId=%v, Memo=%v, ProrationDowngradeScheme=%v, ProrationUpgradeScheme=%v, DowngradeCredit=%v, UpgradeCharge=%v, AccrueCharge=%v, InitiateDunning=%v, PricePointId=%v, BillingSchedule=%v, CustomPrice=%v, AdditionalProperties=%v]",
    	c.Quantity, c.DecimalQuantity, c.PreviousQuantity, c.DecimalPreviousQuantity, c.ComponentId, c.Memo, c.ProrationDowngradeScheme, c.ProrationUpgradeScheme, c.DowngradeCredit, c.UpgradeCharge, c.AccrueCharge, c.InitiateDunning, c.PricePointId, c.BillingSchedule, c.CustomPrice, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreateAllocation.
// It customizes the JSON marshaling process for CreateAllocation objects.
func (c CreateAllocation) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "quantity", "decimal_quantity", "previous_quantity", "decimal_previous_quantity", "component_id", "memo", "proration_downgrade_scheme", "proration_upgrade_scheme", "downgrade_credit", "upgrade_charge", "accrue_charge", "initiate_dunning", "price_point_id", "billing_schedule", "custom_price"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateAllocation object to a map representation for JSON marshaling.
func (c CreateAllocation) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["quantity"] = c.Quantity
    if c.DecimalQuantity != nil {
        structMap["decimal_quantity"] = c.DecimalQuantity
    }
    if c.PreviousQuantity != nil {
        structMap["previous_quantity"] = c.PreviousQuantity
    }
    if c.DecimalPreviousQuantity != nil {
        structMap["decimal_previous_quantity"] = c.DecimalPreviousQuantity
    }
    if c.ComponentId != nil {
        structMap["component_id"] = c.ComponentId
    }
    if c.Memo != nil {
        structMap["memo"] = c.Memo
    }
    if c.ProrationDowngradeScheme != nil {
        structMap["proration_downgrade_scheme"] = c.ProrationDowngradeScheme
    }
    if c.ProrationUpgradeScheme != nil {
        structMap["proration_upgrade_scheme"] = c.ProrationUpgradeScheme
    }
    if c.DowngradeCredit.IsValueSet() {
        if c.DowngradeCredit.Value() != nil {
            structMap["downgrade_credit"] = c.DowngradeCredit.Value()
        } else {
            structMap["downgrade_credit"] = nil
        }
    }
    if c.UpgradeCharge.IsValueSet() {
        if c.UpgradeCharge.Value() != nil {
            structMap["upgrade_charge"] = c.UpgradeCharge.Value()
        } else {
            structMap["upgrade_charge"] = nil
        }
    }
    if c.AccrueCharge != nil {
        structMap["accrue_charge"] = c.AccrueCharge
    }
    if c.InitiateDunning != nil {
        structMap["initiate_dunning"] = c.InitiateDunning
    }
    if c.PricePointId.IsValueSet() {
        if c.PricePointId.Value() != nil {
            structMap["price_point_id"] = c.PricePointId.Value().toMap()
        } else {
            structMap["price_point_id"] = nil
        }
    }
    if c.BillingSchedule != nil {
        structMap["billing_schedule"] = c.BillingSchedule.toMap()
    }
    if c.CustomPrice != nil {
        structMap["custom_price"] = c.CustomPrice.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateAllocation.
// It customizes the JSON unmarshaling process for CreateAllocation objects.
func (c *CreateAllocation) UnmarshalJSON(input []byte) error {
    var temp tempCreateAllocation
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "quantity", "decimal_quantity", "previous_quantity", "decimal_previous_quantity", "component_id", "memo", "proration_downgrade_scheme", "proration_upgrade_scheme", "downgrade_credit", "upgrade_charge", "accrue_charge", "initiate_dunning", "price_point_id", "billing_schedule", "custom_price")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Quantity = *temp.Quantity
    c.DecimalQuantity = temp.DecimalQuantity
    c.PreviousQuantity = temp.PreviousQuantity
    c.DecimalPreviousQuantity = temp.DecimalPreviousQuantity
    c.ComponentId = temp.ComponentId
    c.Memo = temp.Memo
    c.ProrationDowngradeScheme = temp.ProrationDowngradeScheme
    c.ProrationUpgradeScheme = temp.ProrationUpgradeScheme
    c.DowngradeCredit = temp.DowngradeCredit
    c.UpgradeCharge = temp.UpgradeCharge
    c.AccrueCharge = temp.AccrueCharge
    c.InitiateDunning = temp.InitiateDunning
    c.PricePointId = temp.PricePointId
    c.BillingSchedule = temp.BillingSchedule
    c.CustomPrice = temp.CustomPrice
    return nil
}

// tempCreateAllocation is a temporary struct used for validating the fields of CreateAllocation.
type tempCreateAllocation  struct {
    Quantity                 *float64                               `json:"quantity"`
    DecimalQuantity          *string                                `json:"decimal_quantity,omitempty"`
    PreviousQuantity         *float64                               `json:"previous_quantity,omitempty"`
    DecimalPreviousQuantity  *string                                `json:"decimal_previous_quantity,omitempty"`
    ComponentId              *int                                   `json:"component_id,omitempty"`
    Memo                     *string                                `json:"memo,omitempty"`
    ProrationDowngradeScheme *string                                `json:"proration_downgrade_scheme,omitempty"`
    ProrationUpgradeScheme   *string                                `json:"proration_upgrade_scheme,omitempty"`
    DowngradeCredit          Optional[DowngradeCreditCreditType]    `json:"downgrade_credit"`
    UpgradeCharge            Optional[UpgradeChargeCreditType]      `json:"upgrade_charge"`
    AccrueCharge             *bool                                  `json:"accrue_charge,omitempty"`
    InitiateDunning          *bool                                  `json:"initiate_dunning,omitempty"`
    PricePointId             Optional[CreateAllocationPricePointId] `json:"price_point_id"`
    BillingSchedule          *BillingSchedule                       `json:"billing_schedule,omitempty"`
    CustomPrice              *ComponentCustomPrice                  `json:"custom_price,omitempty"`
}

func (c *tempCreateAllocation) validate() error {
    var errs []string
    if c.Quantity == nil {
        errs = append(errs, "required field `quantity` is missing for type `Create Allocation`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
