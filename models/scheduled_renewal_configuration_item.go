// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// ScheduledRenewalConfigurationItem represents a ScheduledRenewalConfigurationItem struct.
type ScheduledRenewalConfigurationItem struct {
    Id                                 *int                   `json:"id,omitempty"`
    SubscriptionId                     *int                   `json:"subscription_id,omitempty"`
    SubscriptionRenewalConfigurationId *int                   `json:"subscription_renewal_configuration_id,omitempty"`
    ItemId                             *int                   `json:"item_id,omitempty"`
    ItemType                           *string                `json:"item_type,omitempty"`
    ItemSubclass                       *string                `json:"item_subclass,omitempty"`
    PricePointId                       *int                   `json:"price_point_id,omitempty"`
    PricePointType                     *string                `json:"price_point_type,omitempty"`
    Quantity                           *int                   `json:"quantity,omitempty"`
    DecimalQuantity                    *string                `json:"decimal_quantity,omitempty"`
    CreatedAt                          *time.Time             `json:"created_at,omitempty"`
    AdditionalProperties               map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ScheduledRenewalConfigurationItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s ScheduledRenewalConfigurationItem) String() string {
    return fmt.Sprintf(
    	"ScheduledRenewalConfigurationItem[Id=%v, SubscriptionId=%v, SubscriptionRenewalConfigurationId=%v, ItemId=%v, ItemType=%v, ItemSubclass=%v, PricePointId=%v, PricePointType=%v, Quantity=%v, DecimalQuantity=%v, CreatedAt=%v, AdditionalProperties=%v]",
    	s.Id, s.SubscriptionId, s.SubscriptionRenewalConfigurationId, s.ItemId, s.ItemType, s.ItemSubclass, s.PricePointId, s.PricePointType, s.Quantity, s.DecimalQuantity, s.CreatedAt, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ScheduledRenewalConfigurationItem.
// It customizes the JSON marshaling process for ScheduledRenewalConfigurationItem objects.
func (s ScheduledRenewalConfigurationItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "id", "subscription_id", "subscription_renewal_configuration_id", "item_id", "item_type", "item_subclass", "price_point_id", "price_point_type", "quantity", "decimal_quantity", "created_at"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the ScheduledRenewalConfigurationItem object to a map representation for JSON marshaling.
func (s ScheduledRenewalConfigurationItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Id != nil {
        structMap["id"] = s.Id
    }
    if s.SubscriptionId != nil {
        structMap["subscription_id"] = s.SubscriptionId
    }
    if s.SubscriptionRenewalConfigurationId != nil {
        structMap["subscription_renewal_configuration_id"] = s.SubscriptionRenewalConfigurationId
    }
    if s.ItemId != nil {
        structMap["item_id"] = s.ItemId
    }
    if s.ItemType != nil {
        structMap["item_type"] = s.ItemType
    }
    if s.ItemSubclass != nil {
        structMap["item_subclass"] = s.ItemSubclass
    }
    if s.PricePointId != nil {
        structMap["price_point_id"] = s.PricePointId
    }
    if s.PricePointType != nil {
        structMap["price_point_type"] = s.PricePointType
    }
    if s.Quantity != nil {
        structMap["quantity"] = s.Quantity
    }
    if s.DecimalQuantity != nil {
        structMap["decimal_quantity"] = s.DecimalQuantity
    }
    if s.CreatedAt != nil {
        structMap["created_at"] = s.CreatedAt.Format(time.RFC3339)
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ScheduledRenewalConfigurationItem.
// It customizes the JSON unmarshaling process for ScheduledRenewalConfigurationItem objects.
func (s *ScheduledRenewalConfigurationItem) UnmarshalJSON(input []byte) error {
    var temp tempScheduledRenewalConfigurationItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "subscription_id", "subscription_renewal_configuration_id", "item_id", "item_type", "item_subclass", "price_point_id", "price_point_type", "quantity", "decimal_quantity", "created_at")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Id = temp.Id
    s.SubscriptionId = temp.SubscriptionId
    s.SubscriptionRenewalConfigurationId = temp.SubscriptionRenewalConfigurationId
    s.ItemId = temp.ItemId
    s.ItemType = temp.ItemType
    s.ItemSubclass = temp.ItemSubclass
    s.PricePointId = temp.PricePointId
    s.PricePointType = temp.PricePointType
    s.Quantity = temp.Quantity
    s.DecimalQuantity = temp.DecimalQuantity
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        s.CreatedAt = &CreatedAtVal
    }
    return nil
}

// tempScheduledRenewalConfigurationItem is a temporary struct used for validating the fields of ScheduledRenewalConfigurationItem.
type tempScheduledRenewalConfigurationItem  struct {
    Id                                 *int    `json:"id,omitempty"`
    SubscriptionId                     *int    `json:"subscription_id,omitempty"`
    SubscriptionRenewalConfigurationId *int    `json:"subscription_renewal_configuration_id,omitempty"`
    ItemId                             *int    `json:"item_id,omitempty"`
    ItemType                           *string `json:"item_type,omitempty"`
    ItemSubclass                       *string `json:"item_subclass,omitempty"`
    PricePointId                       *int    `json:"price_point_id,omitempty"`
    PricePointType                     *string `json:"price_point_type,omitempty"`
    Quantity                           *int    `json:"quantity,omitempty"`
    DecimalQuantity                    *string `json:"decimal_quantity,omitempty"`
    CreatedAt                          *string `json:"created_at,omitempty"`
}
