package models

import (
    "encoding/json"
)

// UpdateComponent represents a UpdateComponent struct.
type UpdateComponent struct {
    Handle               *string                `json:"handle,omitempty"`
    // The name of the Component, suitable for display on statements. i.e. Text Messages.
    Name                 *string                `json:"name,omitempty"`
    // The description of the component.
    Description          Optional[string]       `json:"description"`
    AccountingCode       Optional[string]       `json:"accounting_code"`
    // Boolean flag describing whether a component is taxable or not.
    Taxable              *bool                  `json:"taxable,omitempty"`
    // A string representing the tax code related to the component type. This is especially important when using the Avalara service to tax based on locale. This attribute has a max length of 10 characters.
    TaxCode              Optional[string]       `json:"tax_code"`
    // One of the following: Business Software, Consumer Software, Digital Services, Physical Goods, Other
    ItemCategory         Optional[ItemCategory] `json:"item_category"`
    DisplayOnHostedPage  *bool                  `json:"display_on_hosted_page,omitempty"`
    // The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
    // Available values: `full`, `prorated`, `none`.
    UpgradeCharge        Optional[CreditType]   `json:"upgrade_charge"`
    AdditionalProperties map[string]any         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateComponent.
// It customizes the JSON marshaling process for UpdateComponent objects.
func (u UpdateComponent) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateComponent object to a map representation for JSON marshaling.
func (u UpdateComponent) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
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
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateComponent.
// It customizes the JSON unmarshaling process for UpdateComponent objects.
func (u *UpdateComponent) UnmarshalJSON(input []byte) error {
    var temp updateComponent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "handle", "name", "description", "accounting_code", "taxable", "tax_code", "item_category", "display_on_hosted_page", "upgrade_charge")
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
    return nil
}

// updateComponent is a temporary struct used for validating the fields of UpdateComponent.
type updateComponent  struct {
    Handle              *string                `json:"handle,omitempty"`
    Name                *string                `json:"name,omitempty"`
    Description         Optional[string]       `json:"description"`
    AccountingCode      Optional[string]       `json:"accounting_code"`
    Taxable             *bool                  `json:"taxable,omitempty"`
    TaxCode             Optional[string]       `json:"tax_code"`
    ItemCategory        Optional[ItemCategory] `json:"item_category"`
    DisplayOnHostedPage *bool                  `json:"display_on_hosted_page,omitempty"`
    UpgradeCharge       Optional[CreditType]   `json:"upgrade_charge"`
}
