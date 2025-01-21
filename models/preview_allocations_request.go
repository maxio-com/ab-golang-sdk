/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "log"
    "strings"
    "time"
)

// PreviewAllocationsRequest represents a PreviewAllocationsRequest struct.
type PreviewAllocationsRequest struct {
    Allocations            []CreateAllocation     `json:"allocations"`
    // To calculate proration amounts for a future time. Only within a current subscription period. Only ISO8601 format is supported.
    EffectiveProrationDate *time.Time             `json:"effective_proration_date,omitempty"`
    // The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
    // Available values: `full`, `prorated`, `none`.
    UpgradeCharge          Optional[CreditType]   `json:"upgrade_charge"`
    // The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
    // Available values: `full`, `prorated`, `none`.
    DowngradeCredit        Optional[CreditType]   `json:"downgrade_credit"`
    AdditionalProperties   map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for PreviewAllocationsRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PreviewAllocationsRequest) String() string {
    return fmt.Sprintf(
    	"PreviewAllocationsRequest[Allocations=%v, EffectiveProrationDate=%v, UpgradeCharge=%v, DowngradeCredit=%v, AdditionalProperties=%v]",
    	p.Allocations, p.EffectiveProrationDate, p.UpgradeCharge, p.DowngradeCredit, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for PreviewAllocationsRequest.
// It customizes the JSON marshaling process for PreviewAllocationsRequest objects.
func (p PreviewAllocationsRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(p.AdditionalProperties,
        "allocations", "effective_proration_date", "upgrade_charge", "downgrade_credit"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(p.toMap())
}

// toMap converts the PreviewAllocationsRequest object to a map representation for JSON marshaling.
func (p PreviewAllocationsRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, p.AdditionalProperties)
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
    var temp tempPreviewAllocationsRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "allocations", "effective_proration_date", "upgrade_charge", "downgrade_credit")
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

// tempPreviewAllocationsRequest is a temporary struct used for validating the fields of PreviewAllocationsRequest.
type tempPreviewAllocationsRequest  struct {
    Allocations            *[]CreateAllocation  `json:"allocations"`
    EffectiveProrationDate *string              `json:"effective_proration_date,omitempty"`
    UpgradeCharge          Optional[CreditType] `json:"upgrade_charge"`
    DowngradeCredit        Optional[CreditType] `json:"downgrade_credit"`
}

func (p *tempPreviewAllocationsRequest) validate() error {
    var errs []string
    if p.Allocations == nil {
        errs = append(errs, "required field `allocations` is missing for type `Preview Allocations Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
