// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// WalletProduct represents a WalletProduct struct.
type WalletProduct struct {
    Name                 string                              `json:"name"`
    // The name of the unit the Wallet Product's usage is measured in, for example `tokens`.
    UnitName             string                              `json:"unit_name"`
    Handle               *string                             `json:"handle,omitempty"`
    Description          Optional[string]                    `json:"description"`
    Taxable              *bool                               `json:"taxable,omitempty"`
    TaxCode              Optional[string]                    `json:"tax_code"`
    AccountingCode       Optional[string]                    `json:"accounting_code"`
    // One of the following: Business Software, Consumer Software, Digital Services, Physical Goods, Other
    ItemCategory         Optional[ItemCategory]              `json:"item_category"`
    // Creates one price point for the Wallet Product in the same call. Ignored if `price_points` is also sent.
    PricePoint           *WalletProductPricePointAttributes  `json:"price_point,omitempty"`
    // Creates several price points for the Wallet Product in the same call, winning over `price_point` if both are sent. The first price point created becomes the Wallet Product's default. The whole call is one transaction: if any price point in the array fails validation, nothing is created, including the Wallet Product itself.
    PricePoints          []WalletProductPricePointAttributes `json:"price_points,omitempty"`
    AdditionalProperties map[string]interface{}              `json:"_"`
}

// String implements the fmt.Stringer interface for WalletProduct,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WalletProduct) String() string {
    return fmt.Sprintf(
    	"WalletProduct[Name=%v, UnitName=%v, Handle=%v, Description=%v, Taxable=%v, TaxCode=%v, AccountingCode=%v, ItemCategory=%v, PricePoint=%v, PricePoints=%v, AdditionalProperties=%v]",
    	w.Name, w.UnitName, w.Handle, w.Description, w.Taxable, w.TaxCode, w.AccountingCode, w.ItemCategory, w.PricePoint, w.PricePoints, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WalletProduct.
// It customizes the JSON marshaling process for WalletProduct objects.
func (w WalletProduct) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "name", "unit_name", "handle", "description", "taxable", "tax_code", "accounting_code", "item_category", "price_point", "price_points"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WalletProduct object to a map representation for JSON marshaling.
func (w WalletProduct) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["name"] = w.Name
    structMap["unit_name"] = w.UnitName
    if w.Handle != nil {
        structMap["handle"] = w.Handle
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
    if w.PricePoint != nil {
        structMap["price_point"] = w.PricePoint.toMap()
    }
    if w.PricePoints != nil {
        structMap["price_points"] = w.PricePoints
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WalletProduct.
// It customizes the JSON unmarshaling process for WalletProduct objects.
func (w *WalletProduct) UnmarshalJSON(input []byte) error {
    var temp tempWalletProduct
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "name", "unit_name", "handle", "description", "taxable", "tax_code", "accounting_code", "item_category", "price_point", "price_points")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.Name = *temp.Name
    w.UnitName = *temp.UnitName
    w.Handle = temp.Handle
    w.Description = temp.Description
    w.Taxable = temp.Taxable
    w.TaxCode = temp.TaxCode
    w.AccountingCode = temp.AccountingCode
    w.ItemCategory = temp.ItemCategory
    w.PricePoint = temp.PricePoint
    w.PricePoints = temp.PricePoints
    return nil
}

// tempWalletProduct is a temporary struct used for validating the fields of WalletProduct.
type tempWalletProduct  struct {
    Name           *string                             `json:"name"`
    UnitName       *string                             `json:"unit_name"`
    Handle         *string                             `json:"handle,omitempty"`
    Description    Optional[string]                    `json:"description"`
    Taxable        *bool                               `json:"taxable,omitempty"`
    TaxCode        Optional[string]                    `json:"tax_code"`
    AccountingCode Optional[string]                    `json:"accounting_code"`
    ItemCategory   Optional[ItemCategory]              `json:"item_category"`
    PricePoint     *WalletProductPricePointAttributes  `json:"price_point,omitempty"`
    PricePoints    []WalletProductPricePointAttributes `json:"price_points,omitempty"`
}

func (w *tempWalletProduct) validate() error {
    var errs []string
    if w.Name == nil {
        errs = append(errs, "required field `name` is missing for type `WalletProduct`")
    }
    if w.UnitName == nil {
        errs = append(errs, "required field `unit_name` is missing for type `WalletProduct`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
