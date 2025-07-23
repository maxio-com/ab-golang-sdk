// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"fmt"
)

// ListProductsFilter represents a ListProductsFilter struct.
type ListProductsFilter struct {
	// Allows fetching products with matching id based on provided values. Use in query `filter[ids]=1,2,3`.
	Ids []int `json:"ids,omitempty"`
	// Allows fetching products only if a prepaid product price point is present or not. To use this filter you also have to include the following param in the request `include=prepaid_product_price_point`. Use in query `filter[prepaid_product_price_point][product_price_point_id]=not_null`.
	PrepaidProductPricePoint *PrepaidProductPricePointFilter `json:"prepaid_product_price_point,omitempty"`
	// Allows fetching products with matching use_site_exchange_rate based on provided value (refers to default price point). Use in query `filter[use_site_exchange_rate]=true`.
	UseSiteExchangeRate  *bool                  `json:"use_site_exchange_rate,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ListProductsFilter,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (l ListProductsFilter) String() string {
	return fmt.Sprintf(
		"ListProductsFilter[Ids=%v, PrepaidProductPricePoint=%v, UseSiteExchangeRate=%v, AdditionalProperties=%v]",
		l.Ids, l.PrepaidProductPricePoint, l.UseSiteExchangeRate, l.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ListProductsFilter.
// It customizes the JSON marshaling process for ListProductsFilter objects.
func (l ListProductsFilter) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(l.AdditionalProperties,
		"ids", "prepaid_product_price_point", "use_site_exchange_rate"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(l.toMap())
}

// toMap converts the ListProductsFilter object to a map representation for JSON marshaling.
func (l ListProductsFilter) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, l.AdditionalProperties)
	if l.Ids != nil {
		structMap["ids"] = l.Ids
	}
	if l.PrepaidProductPricePoint != nil {
		structMap["prepaid_product_price_point"] = l.PrepaidProductPricePoint.toMap()
	}
	if l.UseSiteExchangeRate != nil {
		structMap["use_site_exchange_rate"] = l.UseSiteExchangeRate
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListProductsFilter.
// It customizes the JSON unmarshaling process for ListProductsFilter objects.
func (l *ListProductsFilter) UnmarshalJSON(input []byte) error {
	var temp tempListProductsFilter
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ids", "prepaid_product_price_point", "use_site_exchange_rate")
	if err != nil {
		return err
	}
	l.AdditionalProperties = additionalProperties

	l.Ids = temp.Ids
	l.PrepaidProductPricePoint = temp.PrepaidProductPricePoint
	l.UseSiteExchangeRate = temp.UseSiteExchangeRate
	return nil
}

// tempListProductsFilter is a temporary struct used for validating the fields of ListProductsFilter.
type tempListProductsFilter struct {
	Ids                      []int                           `json:"ids,omitempty"`
	PrepaidProductPricePoint *PrepaidProductPricePointFilter `json:"prepaid_product_price_point,omitempty"`
	UseSiteExchangeRate      *bool                           `json:"use_site_exchange_rate,omitempty"`
}
