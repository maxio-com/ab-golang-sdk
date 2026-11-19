// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// WalletProduct1 represents a WalletProduct1 struct.
type WalletProduct1 struct {
    Name                 *string                `json:"name,omitempty"`
    Description          Optional[string]       `json:"description"`
    Taxable              *bool                  `json:"taxable,omitempty"`
    TaxCode              Optional[string]       `json:"tax_code"`
    AccountingCode       Optional[string]       `json:"accounting_code"`
    // One of the following: Business Software, Consumer Software, Digital Services, Physical Goods, Other
    ItemCategory         Optional[ItemCategory] `json:"item_category"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WalletProduct1,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WalletProduct1) String() string {
    return fmt.Sprintf(
    	"WalletProduct1[Name=%v, Description=%v, Taxable=%v, TaxCode=%v, AccountingCode=%v, ItemCategory=%v, AdditionalProperties=%v]",
    	w.Name, w.Description, w.Taxable, w.TaxCode, w.AccountingCode, w.ItemCategory, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WalletProduct1.
// It customizes the JSON marshaling process for WalletProduct1 objects.
func (w WalletProduct1) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "name", "description", "taxable", "tax_code", "accounting_code", "item_category"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WalletProduct1 object to a map representation for JSON marshaling.
func (w WalletProduct1) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    if w.Name != nil {
        structMap["name"] = w.Name
    }
    if w.Description.IsValueSet() {
        if w.Description.Value() != nil {
            structMap["description"] = w.Description.Value()
        } else {
            structMap["description"] = nil
        }
    }
    if w.Taxable != nil {
        structMap["taxable"] = w.Taxable
    }
    if w.TaxCode.IsValueSet() {
        if w.TaxCode.Value() != nil {
            structMap["tax_code"] = w.TaxCode.Value()
        } else {
            structMap["tax_code"] = nil
        }
    }
    if w.AccountingCode.IsValueSet() {
        if w.AccountingCode.Value() != nil {
            structMap["accounting_code"] = w.AccountingCode.Value()
        } else {
            structMap["accounting_code"] = nil
        }
    }
    if w.ItemCategory.IsValueSet() {
        if w.ItemCategory.Value() != nil {
            structMap["item_category"] = w.ItemCategory.Value()
        } else {
            structMap["item_category"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WalletProduct1.
// It customizes the JSON unmarshaling process for WalletProduct1 objects.
func (w *WalletProduct1) UnmarshalJSON(input []byte) error {
    var temp tempWalletProduct1
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "name", "description", "taxable", "tax_code", "accounting_code", "item_category")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.Name = temp.Name
    w.Description = temp.Description
    w.Taxable = temp.Taxable
    w.TaxCode = temp.TaxCode
    w.AccountingCode = temp.AccountingCode
    w.ItemCategory = temp.ItemCategory
    return nil
}

// tempWalletProduct1 is a temporary struct used for validating the fields of WalletProduct1.
type tempWalletProduct1  struct {
    Name           *string                `json:"name,omitempty"`
    Description    Optional[string]       `json:"description"`
    Taxable        *bool                  `json:"taxable,omitempty"`
    TaxCode        Optional[string]       `json:"tax_code"`
    AccountingCode Optional[string]       `json:"accounting_code"`
    ItemCategory   Optional[ItemCategory] `json:"item_category"`
}
