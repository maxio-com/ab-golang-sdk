// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// Purchase represents a Purchase struct.
type Purchase struct {
    // The number of wallet units to purchase, as a decimal string with at most 8 fractional digits. Send fractional values as strings: a JSON number with a decimal point is rejected with `422`, because a double-precision float may already have lost precision before the request arrives.
    Quantity             PurchaseQuantity       `json:"quantity"`
    // The price point to purchase at. Must belong to the wallet's billing product. Defaults to the `price_point_id` set on the wallet.
    PricePointId         *int                   `json:"price_point_id,omitempty"`
    // (Optional) A note recorded on the resulting grant's ledger row.
    Memo                 Optional[string]       `json:"memo"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Purchase,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p Purchase) String() string {
    return fmt.Sprintf(
    	"Purchase[Quantity=%v, PricePointId=%v, Memo=%v, AdditionalProperties=%v]",
    	p.Quantity, p.PricePointId, p.Memo, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Purchase.
// It customizes the JSON marshaling process for Purchase objects.
func (p Purchase) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(p.AdditionalProperties,
        "quantity", "price_point_id", "memo"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(p.toMap())
}

// toMap converts the Purchase object to a map representation for JSON marshaling.
func (p Purchase) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, p.AdditionalProperties)
    structMap["quantity"] = p.Quantity.toMap()
    if p.PricePointId != nil {
        structMap["price_point_id"] = p.PricePointId
    }
    if p.Memo.IsValueSet() {
        if p.Memo.Value() != nil {
            structMap["memo"] = p.Memo.Value()
        } else {
            structMap["memo"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Purchase.
// It customizes the JSON unmarshaling process for Purchase objects.
func (p *Purchase) UnmarshalJSON(input []byte) error {
    var temp tempPurchase
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "quantity", "price_point_id", "memo")
    if err != nil {
    	return err
    }
    p.AdditionalProperties = additionalProperties
    
    p.Quantity = *temp.Quantity
    p.PricePointId = temp.PricePointId
    p.Memo = temp.Memo
    return nil
}

// tempPurchase is a temporary struct used for validating the fields of Purchase.
type tempPurchase  struct {
    Quantity     *PurchaseQuantity `json:"quantity"`
    PricePointId *int              `json:"price_point_id,omitempty"`
    Memo         Optional[string]  `json:"memo"`
}

func (p *tempPurchase) validate() error {
    var errs []string
    if p.Quantity == nil {
        errs = append(errs, "required field `quantity` is missing for type `Purchase`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
