package models

import (
	"encoding/json"
)

// UpdateComponent represents a UpdateComponent struct.
type UpdateComponent struct {
	Handle *string `json:"handle,omitempty"`
	// The name of the Component, suitable for display on statements. i.e. Text Messages.
	Name *string `json:"name,omitempty"`
	// The description of the component.
	Description    Optional[string] `json:"description"`
	AccountingCode Optional[string] `json:"accounting_code"`
	// Boolean flag describing whether a component is taxable or not.
	Taxable *bool `json:"taxable,omitempty"`
	// A string representing the tax code related to the component type. This is especially important when using the Avalara service to tax based on locale. This attribute has a max length of 10 characters.
	TaxCode Optional[string] `json:"tax_code"`
	// One of the following: Business Software, Consumer Software, Digital Services, Physical Goods, Other
	ItemCategory        Optional[ItemCategoryEnum] `json:"item_category"`
	DisplayOnHostedPage *bool                      `json:"display_on_hosted_page,omitempty"`
	// The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
	// Available values: `full`, `prorated`, `none`.
	UpgradeCharge Optional[CreditTypeEnum] `json:"upgrade_charge"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateComponent.
// It customizes the JSON marshaling process for UpdateComponent objects.
func (u *UpdateComponent) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateComponent object to a map representation for JSON marshaling.
func (u *UpdateComponent) toMap() map[string]any {
	structMap := make(map[string]any)
	if u.Handle != nil {
		structMap["handle"] = u.Handle
	}
	if u.Name != nil {
		structMap["name"] = u.Name
	}
	if u.Description.IsValueSet() {
		structMap["description"] = u.Description.Value()
	}
	if u.AccountingCode.IsValueSet() {
		structMap["accounting_code"] = u.AccountingCode.Value()
	}
	if u.Taxable != nil {
		structMap["taxable"] = u.Taxable
	}
	if u.TaxCode.IsValueSet() {
		structMap["tax_code"] = u.TaxCode.Value()
	}
	if u.ItemCategory.IsValueSet() {
		structMap["item_category"] = u.ItemCategory.Value()
	}
	if u.DisplayOnHostedPage != nil {
		structMap["display_on_hosted_page"] = u.DisplayOnHostedPage
	}
	if u.UpgradeCharge.IsValueSet() {
		structMap["upgrade_charge"] = u.UpgradeCharge.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateComponent.
// It customizes the JSON unmarshaling process for UpdateComponent objects.
func (u *UpdateComponent) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Handle              *string                    `json:"handle,omitempty"`
		Name                *string                    `json:"name,omitempty"`
		Description         Optional[string]           `json:"description"`
		AccountingCode      Optional[string]           `json:"accounting_code"`
		Taxable             *bool                      `json:"taxable,omitempty"`
		TaxCode             Optional[string]           `json:"tax_code"`
		ItemCategory        Optional[ItemCategoryEnum] `json:"item_category"`
		DisplayOnHostedPage *bool                      `json:"display_on_hosted_page,omitempty"`
		UpgradeCharge       Optional[CreditTypeEnum]   `json:"upgrade_charge"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

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
