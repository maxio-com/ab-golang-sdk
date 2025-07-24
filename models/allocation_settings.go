// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// AllocationSettings represents a AllocationSettings struct.
type AllocationSettings struct {
    // The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
    // Available values: `full`, `prorated`, `none`.
    UpgradeCharge        Optional[CreditType]   `json:"upgrade_charge"`
    // The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
    // Available values: `full`, `prorated`, `none`.
    DowngradeCredit      Optional[CreditType]   `json:"downgrade_credit"`
    // Either "true" or "false".
    AccrueCharge         *string                `json:"accrue_charge,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AllocationSettings,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AllocationSettings) String() string {
    return fmt.Sprintf(
    	"AllocationSettings[UpgradeCharge=%v, DowngradeCredit=%v, AccrueCharge=%v, AdditionalProperties=%v]",
    	a.UpgradeCharge, a.DowngradeCredit, a.AccrueCharge, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AllocationSettings.
// It customizes the JSON marshaling process for AllocationSettings objects.
func (a AllocationSettings) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "upgrade_charge", "downgrade_credit", "accrue_charge"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AllocationSettings object to a map representation for JSON marshaling.
func (a AllocationSettings) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
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
    if a.AccrueCharge != nil {
        structMap["accrue_charge"] = a.AccrueCharge
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AllocationSettings.
// It customizes the JSON unmarshaling process for AllocationSettings objects.
func (a *AllocationSettings) UnmarshalJSON(input []byte) error {
    var temp tempAllocationSettings
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "upgrade_charge", "downgrade_credit", "accrue_charge")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.UpgradeCharge = temp.UpgradeCharge
    a.DowngradeCredit = temp.DowngradeCredit
    a.AccrueCharge = temp.AccrueCharge
    return nil
}

// tempAllocationSettings is a temporary struct used for validating the fields of AllocationSettings.
type tempAllocationSettings  struct {
    UpgradeCharge   Optional[CreditType] `json:"upgrade_charge"`
    DowngradeCredit Optional[CreditType] `json:"downgrade_credit"`
    AccrueCharge    *string              `json:"accrue_charge,omitempty"`
}
