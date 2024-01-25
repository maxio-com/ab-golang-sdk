package models

import (
	"encoding/json"
	"log"
	"time"
)

// PreviewAllocationsRequest represents a PreviewAllocationsRequest struct.
type PreviewAllocationsRequest struct {
	Allocations []CreateAllocation `json:"allocations"`
	// To calculate proration amounts for a future time. Only within a current subscription period. Only ISO8601 format is supported.
	EffectiveProrationDate *time.Time `json:"effective_proration_date,omitempty"`
	// The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
	// Available values: `full`, `prorated`, `none`.
	UpgradeCharge Optional[CreditType] `json:"upgrade_charge"`
	// The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
	// Available values: `full`, `prorated`, `none`.
	DowngradeCredit Optional[CreditType] `json:"downgrade_credit"`
}

// MarshalJSON implements the json.Marshaler interface for PreviewAllocationsRequest.
// It customizes the JSON marshaling process for PreviewAllocationsRequest objects.
func (p *PreviewAllocationsRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(p.toMap())
}

// toMap converts the PreviewAllocationsRequest object to a map representation for JSON marshaling.
func (p *PreviewAllocationsRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["allocations"] = p.Allocations
	if p.EffectiveProrationDate != nil {
		structMap["effective_proration_date"] = p.EffectiveProrationDate.Format(DEFAULT_DATE)
	}
	if p.UpgradeCharge.IsValueSet() {
		structMap["upgrade_charge"] = p.UpgradeCharge.Value()
	}
	if p.DowngradeCredit.IsValueSet() {
		structMap["downgrade_credit"] = p.DowngradeCredit.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PreviewAllocationsRequest.
// It customizes the JSON unmarshaling process for PreviewAllocationsRequest objects.
func (p *PreviewAllocationsRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Allocations            []CreateAllocation   `json:"allocations"`
		EffectiveProrationDate *string              `json:"effective_proration_date,omitempty"`
		UpgradeCharge          Optional[CreditType] `json:"upgrade_charge"`
		DowngradeCredit        Optional[CreditType] `json:"downgrade_credit"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	p.Allocations = temp.Allocations
	if temp.EffectiveProrationDate != nil {
		EffectiveProrationDateVal, err := time.Parse(DEFAULT_DATE, *temp.EffectiveProrationDate)
		if err != nil {
			log.Fatalf("Cannot Parse effective_proration_date as % s format.", DEFAULT_DATE)
		}
		p.EffectiveProrationDate = &EffectiveProrationDateVal
	}
	p.UpgradeCharge = temp.UpgradeCharge
	p.DowngradeCredit = temp.DowngradeCredit
	return nil
}
