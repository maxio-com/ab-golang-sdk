// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// UpdateComponent represents a UpdateComponent struct.
type UpdateComponent struct {
    Handle                    *string                `json:"handle,omitempty"`
    // The name of the Component, suitable for display on statements. e.g., Text Messages.
    Name                      *string                `json:"name,omitempty"`
    // The description of the component.
    Description               Optional[string]       `json:"description"`
    AccountingCode            Optional[string]       `json:"accounting_code"`
    // Boolean flag describing whether a component is taxable or not.
    Taxable                   *bool                  `json:"taxable,omitempty"`
    // A string representing the tax code related to the component type. This is especially important when using AvaTax to tax based on locale. This attribute has a max length of 25 characters.
    TaxCode                   Optional[string]       `json:"tax_code"`
    // One of the following: Business Software, Consumer Software, Digital Services, Physical Goods, Other
    ItemCategory              Optional[ItemCategory] `json:"item_category"`
    DisplayOnHostedPage       *bool                  `json:"display_on_hosted_page,omitempty"`
    // The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
    UpgradeCharge             Optional[CreditType]   `json:"upgrade_charge"`
    // For components using legacy Events Based Billing, changes the attached site-owned metric. An unchanged metric ID does not change the source. Sites using Maxio Metering ignore this input. This ID does not identify a Maxio Metering meter.
    EventBasedBillingMetricId *int                   `json:"event_based_billing_metric_id,omitempty"`
    // For components using Maxio Metering, changes the attached merchant-owned meter. Do not supply together with a non-blank metering_formula. Ignored for components using legacy Events Based Billing; not applicable to other component kinds.
    MeteringMeterId           *int                   `json:"metering_meter_id,omitempty"`
    // For components using Maxio Metering, sets an Advanced Formula using meters available to the site. Supply a non-blank formula; do not supply together with metering_meter_id. Ignored for components using legacy Events Based Billing; not applicable to other component kinds.
    MeteringFormula           *string                `json:"metering_formula,omitempty"`
    // (Optional) Custom UNSPSC commodity code for Level 3/CEDP payment data. When set, this value is sent as the commodity code on invoice line items for this component instead of the default derived from item_category.
    UnspscCode                Optional[string]       `json:"unspsc_code"`
    AdditionalProperties      map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UpdateComponent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpdateComponent) String() string {
    return fmt.Sprintf(
    	"UpdateComponent[Handle=%v, Name=%v, Description=%v, AccountingCode=%v, Taxable=%v, TaxCode=%v, ItemCategory=%v, DisplayOnHostedPage=%v, UpgradeCharge=%v, EventBasedBillingMetricId=%v, MeteringMeterId=%v, MeteringFormula=%v, UnspscCode=%v, AdditionalProperties=%v]",
    	u.Handle, u.Name, u.Description, u.AccountingCode, u.Taxable, u.TaxCode, u.ItemCategory, u.DisplayOnHostedPage, u.UpgradeCharge, u.EventBasedBillingMetricId, u.MeteringMeterId, u.MeteringFormula, u.UnspscCode, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UpdateComponent.
// It customizes the JSON marshaling process for UpdateComponent objects.
func (u UpdateComponent) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "handle", "name", "description", "accounting_code", "taxable", "tax_code", "item_category", "display_on_hosted_page", "upgrade_charge", "event_based_billing_metric_id", "metering_meter_id", "metering_formula", "unspsc_code"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateComponent object to a map representation for JSON marshaling.
func (u UpdateComponent) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Handle != nil {
        structMap["handle"] = u.Handle
    }
    if u.Name != nil {
        structMap["name"] = u.Name
    }
    if u.Description.IsValueSet() {
        if u.Description.Value() != nil {
            structMap["description"] = u.Description.Value()
        } else {
            structMap["description"] = nil
        }
    }
    if u.AccountingCode.IsValueSet() {
        if u.AccountingCode.Value() != nil {
            structMap["accounting_code"] = u.AccountingCode.Value()
        } else {
            structMap["accounting_code"] = nil
        }
    }
    if u.Taxable != nil {
        structMap["taxable"] = u.Taxable
    }
    if u.TaxCode.IsValueSet() {
        if u.TaxCode.Value() != nil {
            structMap["tax_code"] = u.TaxCode.Value()
        } else {
            structMap["tax_code"] = nil
        }
    }
    if u.ItemCategory.IsValueSet() {
        if u.ItemCategory.Value() != nil {
            structMap["item_category"] = u.ItemCategory.Value()
        } else {
            structMap["item_category"] = nil
        }
    }
    if u.DisplayOnHostedPage != nil {
        structMap["display_on_hosted_page"] = u.DisplayOnHostedPage
    }
    if u.UpgradeCharge.IsValueSet() {
        if u.UpgradeCharge.Value() != nil {
            structMap["upgrade_charge"] = u.UpgradeCharge.Value()
        } else {
            structMap["upgrade_charge"] = nil
        }
    }
    if u.EventBasedBillingMetricId != nil {
        structMap["event_based_billing_metric_id"] = u.EventBasedBillingMetricId
    }
    if u.MeteringMeterId != nil {
        structMap["metering_meter_id"] = u.MeteringMeterId
    }
    if u.MeteringFormula != nil {
        structMap["metering_formula"] = u.MeteringFormula
    }
    if u.UnspscCode.IsValueSet() {
        if u.UnspscCode.Value() != nil {
            structMap["unspsc_code"] = u.UnspscCode.Value()
        } else {
            structMap["unspsc_code"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateComponent.
// It customizes the JSON unmarshaling process for UpdateComponent objects.
func (u *UpdateComponent) UnmarshalJSON(input []byte) error {
    var temp tempUpdateComponent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "handle", "name", "description", "accounting_code", "taxable", "tax_code", "item_category", "display_on_hosted_page", "upgrade_charge", "event_based_billing_metric_id", "metering_meter_id", "metering_formula", "unspsc_code")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.Handle = temp.Handle
    u.Name = temp.Name
    u.Description = temp.Description
    u.AccountingCode = temp.AccountingCode
    u.Taxable = temp.Taxable
    u.TaxCode = temp.TaxCode
    u.ItemCategory = temp.ItemCategory
    u.DisplayOnHostedPage = temp.DisplayOnHostedPage
    u.UpgradeCharge = temp.UpgradeCharge
    u.EventBasedBillingMetricId = temp.EventBasedBillingMetricId
    u.MeteringMeterId = temp.MeteringMeterId
    u.MeteringFormula = temp.MeteringFormula
    u.UnspscCode = temp.UnspscCode
    return nil
}

// tempUpdateComponent is a temporary struct used for validating the fields of UpdateComponent.
type tempUpdateComponent  struct {
    Handle                    *string                `json:"handle,omitempty"`
    Name                      *string                `json:"name,omitempty"`
    Description               Optional[string]       `json:"description"`
    AccountingCode            Optional[string]       `json:"accounting_code"`
    Taxable                   *bool                  `json:"taxable,omitempty"`
    TaxCode                   Optional[string]       `json:"tax_code"`
    ItemCategory              Optional[ItemCategory] `json:"item_category"`
    DisplayOnHostedPage       *bool                  `json:"display_on_hosted_page,omitempty"`
    UpgradeCharge             Optional[CreditType]   `json:"upgrade_charge"`
    EventBasedBillingMetricId *int                   `json:"event_based_billing_metric_id,omitempty"`
    MeteringMeterId           *int                   `json:"metering_meter_id,omitempty"`
    MeteringFormula           *string                `json:"metering_formula,omitempty"`
    UnspscCode                Optional[string]       `json:"unspsc_code"`
}
