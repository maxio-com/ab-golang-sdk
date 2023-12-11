package models

import (
	"encoding/json"
)

// AllocateComponents represents a AllocateComponents struct.
type AllocateComponents struct {
	ProrationUpgradeScheme   *string                   `json:"proration_upgrade_scheme,omitempty"`   // Deprecated
	ProrationDowngradeScheme *string                   `json:"proration_downgrade_scheme,omitempty"` // Deprecated
	Allocations              []CreateAllocationRequest `json:"allocations,omitempty"`
	AccrueCharge             *bool                     `json:"accrue_charge,omitempty"`
	// The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
	// Available values: `full`, `prorated`, `none`.
	UpgradeCharge Optional[CreditTypeEnum] `json:"upgrade_charge"`
	// The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
	// Available values: `full`, `prorated`, `none`.
	DowngradeCredit Optional[CreditTypeEnum] `json:"downgrade_credit"`
	// (Optional) If not passed, the allocation(s) will use the payment collection method on the subscription
	PaymentCollectionMethod *PaymentCollectionMethod1Enum `json:"payment_collection_method,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for AllocateComponents.
// It customizes the JSON marshaling process for AllocateComponents objects.
func (a *AllocateComponents) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(a.toMap())
}

// toMap converts the AllocateComponents object to a map representation for JSON marshaling.
func (a *AllocateComponents) toMap() map[string]any {
	structMap := make(map[string]any)
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
		structMap["upgrade_charge"] = a.UpgradeCharge.Value()
	}
	if a.DowngradeCredit.IsValueSet() {
		structMap["downgrade_credit"] = a.DowngradeCredit.Value()
	}
	if a.PaymentCollectionMethod != nil {
		structMap["payment_collection_method"] = a.PaymentCollectionMethod
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AllocateComponents.
// It customizes the JSON unmarshaling process for AllocateComponents objects.
func (a *AllocateComponents) UnmarshalJSON(input []byte) error {
	temp := &struct {
		ProrationUpgradeScheme   *string                       `json:"proration_upgrade_scheme,omitempty"`
		ProrationDowngradeScheme *string                       `json:"proration_downgrade_scheme,omitempty"`
		Allocations              []CreateAllocationRequest     `json:"allocations,omitempty"`
		AccrueCharge             *bool                         `json:"accrue_charge,omitempty"`
		UpgradeCharge            Optional[CreditTypeEnum]      `json:"upgrade_charge"`
		DowngradeCredit          Optional[CreditTypeEnum]      `json:"downgrade_credit"`
		PaymentCollectionMethod  *PaymentCollectionMethod1Enum `json:"payment_collection_method,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	a.ProrationUpgradeScheme = temp.ProrationUpgradeScheme
	a.ProrationDowngradeScheme = temp.ProrationDowngradeScheme
	a.Allocations = temp.Allocations
	a.AccrueCharge = temp.AccrueCharge
	a.UpgradeCharge = temp.UpgradeCharge
	a.DowngradeCredit = temp.DowngradeCredit
	a.PaymentCollectionMethod = temp.PaymentCollectionMethod
	return nil
}
