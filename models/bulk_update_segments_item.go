package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// BulkUpdateSegmentsItem represents a BulkUpdateSegmentsItem struct.
type BulkUpdateSegmentsItem struct {
    // The ID of the segment you want to update.
    Id                   int                          `json:"id"`
    // The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes.
    PricingScheme        PricingScheme                `json:"pricing_scheme"`
    Prices               []CreateOrUpdateSegmentPrice `json:"prices"`
    AdditionalProperties map[string]any               `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for BulkUpdateSegmentsItem.
// It customizes the JSON marshaling process for BulkUpdateSegmentsItem objects.
func (b BulkUpdateSegmentsItem) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(b.toMap())
}

// toMap converts the BulkUpdateSegmentsItem object to a map representation for JSON marshaling.
func (b BulkUpdateSegmentsItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, b.AdditionalProperties)
    structMap["id"] = b.Id
    structMap["pricing_scheme"] = b.PricingScheme
    structMap["prices"] = b.Prices
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BulkUpdateSegmentsItem.
// It customizes the JSON unmarshaling process for BulkUpdateSegmentsItem objects.
func (b *BulkUpdateSegmentsItem) UnmarshalJSON(input []byte) error {
    var temp bulkUpdateSegmentsItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "id", "pricing_scheme", "prices")
    if err != nil {
    	return err
    }
    
    b.AdditionalProperties = additionalProperties
    b.Id = *temp.Id
    b.PricingScheme = *temp.PricingScheme
    b.Prices = *temp.Prices
    return nil
}

// TODO
type bulkUpdateSegmentsItem  struct {
    Id            *int                          `json:"id"`
    PricingScheme *PricingScheme                `json:"pricing_scheme"`
    Prices        *[]CreateOrUpdateSegmentPrice `json:"prices"`
}

func (b *bulkUpdateSegmentsItem) validate() error {
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
