package models

import (
    "encoding/json"
)

// ListProductsFilter represents a ListProductsFilter struct.
type ListProductsFilter struct {
    // Allows fetching products only if a prepaid product price point is present or not. To use this filter you also have to include the following param in the request `include=prepaid_product_price_point`. Use in query `filter[prepaid_product_price_point][product_price_point_id]=not_null`.
    PrepaidProductPricePoint *PrepaidProductPricePointFilter `json:"prepaid_product_price_point,omitempty"`
    // Allows fetching products with matching use_site_exchange_rate based on provided value (refers to default price point). Use in query `filter[use_site_exchange_rate]=true`.
    UseSiteExchangeRate      *bool                           `json:"use_site_exchange_rate,omitempty"`
    AdditionalProperties     map[string]any                  `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ListProductsFilter.
// It customizes the JSON marshaling process for ListProductsFilter objects.
func (l ListProductsFilter) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the ListProductsFilter object to a map representation for JSON marshaling.
func (l ListProductsFilter) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, l.AdditionalProperties)
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
    var temp listProductsFilter
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "prepaid_product_price_point", "use_site_exchange_rate")
    if err != nil {
    	return err
    }
    
    l.AdditionalProperties = additionalProperties
    l.PrepaidProductPricePoint = temp.PrepaidProductPricePoint
    l.UseSiteExchangeRate = temp.UseSiteExchangeRate
    return nil
}

// listProductsFilter is a temporary struct used for validating the fields of ListProductsFilter.
type listProductsFilter  struct {
    PrepaidProductPricePoint *PrepaidProductPricePointFilter `json:"prepaid_product_price_point,omitempty"`
    UseSiteExchangeRate      *bool                           `json:"use_site_exchange_rate,omitempty"`
}
