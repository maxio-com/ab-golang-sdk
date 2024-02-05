package models

import (
    "encoding/json"
)

// AllocationSettings represents a AllocationSettings struct.
type AllocationSettings struct {
    // The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
    // Available values: `full`, `prorated`, `none`.
    UpgradeCharge   Optional[CreditType] `json:"upgrade_charge"`
    // The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
    // Available values: `full`, `prorated`, `none`.
    DowngradeCredit Optional[CreditType] `json:"downgrade_credit"`
    // Either "true" or "false".
    AccrueCharge    *string              `json:"accrue_charge,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for AllocationSettings.
// It customizes the JSON marshaling process for AllocationSettings objects.
func (a *AllocationSettings) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the AllocationSettings object to a map representation for JSON marshaling.
func (a *AllocationSettings) toMap() map[string]any {
    structMap := make(map[string]any)
    if a.UpgradeCharge.IsValueSet() {
        structMap["upgrade_charge"] = a.UpgradeCharge.Value()
    }
    if a.DowngradeCredit.IsValueSet() {
        structMap["downgrade_credit"] = a.DowngradeCredit.Value()
    }
    if a.AccrueCharge != nil {
        structMap["accrue_charge"] = a.AccrueCharge
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AllocationSettings.
// It customizes the JSON unmarshaling process for AllocationSettings objects.
func (a *AllocationSettings) UnmarshalJSON(input []byte) error {
    temp := &struct {
        UpgradeCharge   Optional[CreditType] `json:"upgrade_charge"`
        DowngradeCredit Optional[CreditType] `json:"downgrade_credit"`
        AccrueCharge    *string              `json:"accrue_charge,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    a.UpgradeCharge = temp.UpgradeCharge
    a.DowngradeCredit = temp.DowngradeCredit
    a.AccrueCharge = temp.AccrueCharge
    return nil
}
