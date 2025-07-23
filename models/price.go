// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// Price represents a Price struct.
type Price struct {
	StartingQuantity PriceStartingQuantity         `json:"starting_quantity"`
	EndingQuantity   Optional[PriceEndingQuantity] `json:"ending_quantity"`
	// The price can contain up to 8 decimal places. i.e. 1.00 or 0.0012 or 0.00000065
	UnitPrice            PriceUnitPrice         `json:"unit_price"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Price,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p Price) String() string {
	return fmt.Sprintf(
		"Price[StartingQuantity=%v, EndingQuantity=%v, UnitPrice=%v, AdditionalProperties=%v]",
		p.StartingQuantity, p.EndingQuantity, p.UnitPrice, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Price.
// It customizes the JSON marshaling process for Price objects.
func (p Price) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(p.AdditionalProperties,
		"starting_quantity", "ending_quantity", "unit_price"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(p.toMap())
}

// toMap converts the Price object to a map representation for JSON marshaling.
func (p Price) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, p.AdditionalProperties)
	structMap["starting_quantity"] = p.StartingQuantity.toMap()
	if p.EndingQuantity.IsValueSet() {
		if p.EndingQuantity.Value() != nil {
			structMap["ending_quantity"] = p.EndingQuantity.Value().toMap()
		} else {
			structMap["ending_quantity"] = nil
		}
	}
	structMap["unit_price"] = p.UnitPrice.toMap()
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Price.
// It customizes the JSON unmarshaling process for Price objects.
func (p *Price) UnmarshalJSON(input []byte) error {
	var temp tempPrice
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "starting_quantity", "ending_quantity", "unit_price")
	if err != nil {
		return err
	}
	p.AdditionalProperties = additionalProperties

	p.StartingQuantity = *temp.StartingQuantity
	p.EndingQuantity = temp.EndingQuantity
	p.UnitPrice = *temp.UnitPrice
	return nil
}

// tempPrice is a temporary struct used for validating the fields of Price.
type tempPrice struct {
	StartingQuantity *PriceStartingQuantity        `json:"starting_quantity"`
	EndingQuantity   Optional[PriceEndingQuantity] `json:"ending_quantity"`
	UnitPrice        *PriceUnitPrice               `json:"unit_price"`
}

func (p *tempPrice) validate() error {
	var errs []string
	if p.StartingQuantity == nil {
		errs = append(errs, "required field `starting_quantity` is missing for type `Price`")
	}
	if p.UnitPrice == nil {
		errs = append(errs, "required field `unit_price` is missing for type `Price`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
