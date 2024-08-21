/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// AllocateComponents represents a AllocateComponents struct.
type AllocateComponents struct {
    ProrationUpgradeScheme   *string              `json:"proration_upgrade_scheme,omitempty"`   // Deprecated
    ProrationDowngradeScheme *string              `json:"proration_downgrade_scheme,omitempty"` // Deprecated
    Allocations              []CreateAllocation   `json:"allocations,omitempty"`
    AccrueCharge             *bool                `json:"accrue_charge,omitempty"`
    // The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
    // Available values: `full`, `prorated`, `none`.
    UpgradeCharge            Optional[CreditType] `json:"upgrade_charge"`
    // The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
    // Available values: `full`, `prorated`, `none`.
    DowngradeCredit          Optional[CreditType] `json:"downgrade_credit"`
    // (Optional) If not passed, the allocation(s) will use the payment collection method on the subscription
    PaymentCollectionMethod  *CollectionMethod    `json:"payment_collection_method,omitempty"`
    // If true, if the immediate component payment fails, initiate dunning for the subscription. 
    // Otherwise, leave the charges on the subscription to pay for at renewal.
    InitiateDunning          *bool                `json:"initiate_dunning,omitempty"`
    AdditionalProperties     map[string]any       `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for AllocateComponents.
// It customizes the JSON marshaling process for AllocateComponents objects.
func (a AllocateComponents) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the AllocateComponents object to a map representation for JSON marshaling.
func (a AllocateComponents) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.ProrationUpgradeScheme != nil {
        structMap["proration_upgrade_scheme"] = a.ProrationUpgradeScheme
    }
    if a.ProrationDowngradeScheme != nil {
        structMap["proration_downgrade_scheme"] = a.ProrationDowngradeScheme
    }
    if a.Allocations != nil {
        structMap["allocations"] = a.Allocations
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
    if a.PaymentCollectionMethod != nil {
        structMap["payment_collection_method"] = a.PaymentCollectionMethod
    }
    if a.InitiateDunning != nil {
        structMap["initiate_dunning"] = a.InitiateDunning
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AllocateComponents.
// It customizes the JSON unmarshaling process for AllocateComponents objects.
func (a *AllocateComponents) UnmarshalJSON(input []byte) error {
    var temp tempAllocateComponents
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "proration_upgrade_scheme", "proration_downgrade_scheme", "allocations", "accrue_charge", "upgrade_charge", "downgrade_credit", "payment_collection_method", "initiate_dunning")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.ProrationUpgradeScheme = temp.ProrationUpgradeScheme
    a.ProrationDowngradeScheme = temp.ProrationDowngradeScheme
    a.Allocations = temp.Allocations
    a.AccrueCharge = temp.AccrueCharge
    a.UpgradeCharge = temp.UpgradeCharge
    a.DowngradeCredit = temp.DowngradeCredit
    a.PaymentCollectionMethod = temp.PaymentCollectionMethod
    a.InitiateDunning = temp.InitiateDunning
    return nil
}

// tempAllocateComponents is a temporary struct used for validating the fields of AllocateComponents.
type tempAllocateComponents  struct {
    ProrationUpgradeScheme   *string              `json:"proration_upgrade_scheme,omitempty"`
    ProrationDowngradeScheme *string              `json:"proration_downgrade_scheme,omitempty"`
    Allocations              []CreateAllocation   `json:"allocations,omitempty"`
    AccrueCharge             *bool                `json:"accrue_charge,omitempty"`
    UpgradeCharge            Optional[CreditType] `json:"upgrade_charge"`
    DowngradeCredit          Optional[CreditType] `json:"downgrade_credit"`
    PaymentCollectionMethod  *CollectionMethod    `json:"payment_collection_method,omitempty"`
    InitiateDunning          *bool                `json:"initiate_dunning,omitempty"`
}
