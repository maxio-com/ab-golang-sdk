// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// BulkUpdateSegmentsItem represents a BulkUpdateSegmentsItem struct.
type BulkUpdateSegmentsItem struct {
	// The ID of the segment you want to update.
	Id int `json:"id"`
	// The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes.
	PricingScheme        PricingScheme                `json:"pricing_scheme"`
	Prices               []CreateOrUpdateSegmentPrice `json:"prices"`
	AdditionalProperties map[string]interface{}       `json:"_"`
}

// String implements the fmt.Stringer interface for BulkUpdateSegmentsItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (b BulkUpdateSegmentsItem) String() string {
	return fmt.Sprintf(
		"BulkUpdateSegmentsItem[Id=%v, PricingScheme=%v, Prices=%v, AdditionalProperties=%v]",
		b.Id, b.PricingScheme, b.Prices, b.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for BulkUpdateSegmentsItem.
// It customizes the JSON marshaling process for BulkUpdateSegmentsItem objects.
func (b BulkUpdateSegmentsItem) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(b.AdditionalProperties,
		"id", "pricing_scheme", "prices"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(b.toMap())
}

// toMap converts the BulkUpdateSegmentsItem object to a map representation for JSON marshaling.
func (b BulkUpdateSegmentsItem) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, b.AdditionalProperties)
	structMap["id"] = b.Id
	structMap["pricing_scheme"] = b.PricingScheme
	structMap["prices"] = b.Prices
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BulkUpdateSegmentsItem.
// It customizes the JSON unmarshaling process for BulkUpdateSegmentsItem objects.
func (b *BulkUpdateSegmentsItem) UnmarshalJSON(input []byte) error {
	var temp tempBulkUpdateSegmentsItem
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "pricing_scheme", "prices")
	if err != nil {
		return err
	}
	b.AdditionalProperties = additionalProperties

	b.Id = *temp.Id
	b.PricingScheme = *temp.PricingScheme
	b.Prices = *temp.Prices
	return nil
}

// tempBulkUpdateSegmentsItem is a temporary struct used for validating the fields of BulkUpdateSegmentsItem.
type tempBulkUpdateSegmentsItem struct {
	Id            *int                          `json:"id"`
	PricingScheme *PricingScheme                `json:"pricing_scheme"`
	Prices        *[]CreateOrUpdateSegmentPrice `json:"prices"`
}

func (b *tempBulkUpdateSegmentsItem) validate() error {
	var errs []string
	if b.Id == nil {
		errs = append(errs, "required field `id` is missing for type `Bulk Update Segments Item`")
	}
	if b.PricingScheme == nil {
		errs = append(errs, "required field `pricing_scheme` is missing for type `Bulk Update Segments Item`")
	}
	if b.Prices == nil {
		errs = append(errs, "required field `prices` is missing for type `Bulk Update Segments Item`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
