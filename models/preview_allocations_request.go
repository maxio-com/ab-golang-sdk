package models

import (
    "encoding/json"
    "errors"
    "log"
    "strings"
    "time"
)

// PreviewAllocationsRequest represents a PreviewAllocationsRequest struct.
type PreviewAllocationsRequest struct {
    Allocations            []CreateAllocation   `json:"allocations"`
    // To calculate proration amounts for a future time. Only within a current subscription period. Only ISO8601 format is supported.
    EffectiveProrationDate *time.Time           `json:"effective_proration_date,omitempty"`
    // The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
    // Available values: `full`, `prorated`, `none`.
    UpgradeCharge          Optional[CreditType] `json:"upgrade_charge"`
    // The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
    // Available values: `full`, `prorated`, `none`.
    DowngradeCredit        Optional[CreditType] `json:"downgrade_credit"`
    AdditionalProperties   map[string]any       `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for PreviewAllocationsRequest.
// It customizes the JSON marshaling process for PreviewAllocationsRequest objects.
func (p PreviewAllocationsRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PreviewAllocationsRequest object to a map representation for JSON marshaling.
func (p PreviewAllocationsRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
    structMap["allocations"] = p.Allocations
    if p.EffectiveProrationDate != nil {
        structMap["effective_proration_date"] = p.EffectiveProrationDate.Format(DEFAULT_DATE)
    }
    if p.UpgradeCharge.IsValueSet() {
        if p.UpgradeCharge.Value() != nil {
            structMap["upgrade_charge"] = p.UpgradeCharge.Value()
        } else {
            structMap["upgrade_charge"] = nil
        }
    }
    if p.DowngradeCredit.IsValueSet() {
        if p.DowngradeCredit.Value() != nil {
            structMap["downgrade_credit"] = p.DowngradeCredit.Value()
        } else {
            structMap["downgrade_credit"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PreviewAllocationsRequest.
// It customizes the JSON unmarshaling process for PreviewAllocationsRequest objects.
func (p *PreviewAllocationsRequest) UnmarshalJSON(input []byte) error {
    var temp previewAllocationsRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "allocations", "effective_proration_date", "upgrade_charge", "downgrade_credit")
    if err != nil {
    	return err
    }
    
    p.AdditionalProperties = additionalProperties
    p.Allocations = *temp.Allocations
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

// previewAllocationsRequest is a temporary struct used for validating the fields of PreviewAllocationsRequest.
type previewAllocationsRequest  struct {
    Allocations            *[]CreateAllocation  `json:"allocations"`
    EffectiveProrationDate *string              `json:"effective_proration_date,omitempty"`
    UpgradeCharge          Optional[CreditType] `json:"upgrade_charge"`
    DowngradeCredit        Optional[CreditType] `json:"downgrade_credit"`
}

func (p *previewAllocationsRequest) validate() error {
    var errs []string
    if p.Allocations == nil {
        errs = append(errs, "required field `allocations` is missing for type `Preview Allocations Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
