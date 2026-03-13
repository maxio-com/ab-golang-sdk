// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ScheduledRenewalItemRequestBodyProduct represents a ScheduledRenewalItemRequestBodyProduct struct.
type ScheduledRenewalItemRequestBodyProduct struct {
    // Item type to add. Either Product or Component.
    ItemType             string                             `json:"item_type"`
    // Product or component identifier.
    ItemId               int                                `json:"item_id"`
    // Price point identifier.
    PricePointId         *int                               `json:"price_point_id,omitempty"`
    // Optional quantity for the item.
    Quantity             *int                               `json:"quantity,omitempty"`
    // Custom pricing for a product within a scheduled renewal.
    CustomPrice          *ScheduledRenewalProductPricePoint `json:"custom_price,omitempty"`
    AdditionalProperties map[string]interface{}             `json:"_"`
}

// String implements the fmt.Stringer interface for ScheduledRenewalItemRequestBodyProduct,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s ScheduledRenewalItemRequestBodyProduct) String() string {
    return fmt.Sprintf(
    	"ScheduledRenewalItemRequestBodyProduct[ItemType=%v, ItemId=%v, PricePointId=%v, Quantity=%v, CustomPrice=%v, AdditionalProperties=%v]",
    	s.ItemType, s.ItemId, s.PricePointId, s.Quantity, s.CustomPrice, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ScheduledRenewalItemRequestBodyProduct.
// It customizes the JSON marshaling process for ScheduledRenewalItemRequestBodyProduct objects.
func (s ScheduledRenewalItemRequestBodyProduct) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "item_type", "item_id", "price_point_id", "quantity", "custom_price"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the ScheduledRenewalItemRequestBodyProduct object to a map representation for JSON marshaling.
func (s ScheduledRenewalItemRequestBodyProduct) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["item_type"] = s.ItemType
    structMap["item_id"] = s.ItemId
    if s.PricePointId != nil {
        structMap["price_point_id"] = s.PricePointId
    }
    if s.Quantity != nil {
        structMap["quantity"] = s.Quantity
    }
    if s.CustomPrice != nil {
        structMap["custom_price"] = s.CustomPrice.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ScheduledRenewalItemRequestBodyProduct.
// It customizes the JSON unmarshaling process for ScheduledRenewalItemRequestBodyProduct objects.
func (s *ScheduledRenewalItemRequestBodyProduct) UnmarshalJSON(input []byte) error {
    var temp tempScheduledRenewalItemRequestBodyProduct
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "item_type", "item_id", "price_point_id", "quantity", "custom_price")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.ItemType = *temp.ItemType
    s.ItemId = *temp.ItemId
    s.PricePointId = temp.PricePointId
    s.Quantity = temp.Quantity
    s.CustomPrice = temp.CustomPrice
    return nil
}

// tempScheduledRenewalItemRequestBodyProduct is a temporary struct used for validating the fields of ScheduledRenewalItemRequestBodyProduct.
type tempScheduledRenewalItemRequestBodyProduct  struct {
    ItemType     *string                            `json:"item_type"`
    ItemId       *int                               `json:"item_id"`
    PricePointId *int                               `json:"price_point_id,omitempty"`
    Quantity     *int                               `json:"quantity,omitempty"`
    CustomPrice  *ScheduledRenewalProductPricePoint `json:"custom_price,omitempty"`
}

func (s *tempScheduledRenewalItemRequestBodyProduct) validate() error {
    var errs []string
    if s.ItemType == nil {
        errs = append(errs, "required field `item_type` is missing for type `Scheduled Renewal Item Request Body Product`")
    }
    if s.ItemId == nil {
        errs = append(errs, "required field `item_id` is missing for type `Scheduled Renewal Item Request Body Product`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
