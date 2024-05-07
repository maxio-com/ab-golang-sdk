package models

import (
    "encoding/json"
    "log"
    "time"
)

// SubscriptionMigrationPreviewOptions represents a SubscriptionMigrationPreviewOptions struct.
type SubscriptionMigrationPreviewOptions struct {
    // The ID of the target Product. Either a product_id or product_handle must be present. A Subscription can be migrated to another product for both the current Product Family and another Product Family. Note: Going to another Product Family, components will not be migrated as well.
    ProductId               *int           `json:"product_id,omitempty"`
    // The ID of the specified product's price point. This can be passed to migrate to a non-default price point.
    ProductPricePointId     *int           `json:"product_price_point_id,omitempty"`
    // Whether to include the trial period configured for the product price point when starting a new billing period. Note that if preserve_period is set, then include_trial will be ignored.
    IncludeTrial            *bool          `json:"include_trial,omitempty"`
    // If `true` is sent initial charges will be assessed.
    IncludeInitialCharge    *bool          `json:"include_initial_charge,omitempty"`
    // If `true` is sent, any coupons associated with the subscription will be applied to the migration. If `false` is sent, coupons will not be applied. Note: When migrating to a new product family, the coupon cannot migrate.
    IncludeCoupons          *bool          `json:"include_coupons,omitempty"`
    // If `false` is sent, the subscription's billing period will be reset to today and the full price of the new product will be charged. If `true` is sent, the billing period will not change and a prorated charge will be issued for the new product.
    PreservePeriod          *bool          `json:"preserve_period,omitempty"`
    // The handle of the target Product. Either a product_id or product_handle must be present. A Subscription can be migrated to another product for both the current Product Family and another Product Family. Note: Going to another Product Family, components will not be migrated as well.
    ProductHandle           *string        `json:"product_handle,omitempty"`
    // The ID or handle of the specified product's price point. This can be passed to migrate to a non-default price point.
    ProductPricePointHandle *string        `json:"product_price_point_handle,omitempty"`
    Proration               *Proration     `json:"proration,omitempty"`
    // The date that the proration is calculated from for the preview
    ProrationDate           *time.Time     `json:"proration_date,omitempty"`
    AdditionalProperties    map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionMigrationPreviewOptions.
// It customizes the JSON marshaling process for SubscriptionMigrationPreviewOptions objects.
func (s SubscriptionMigrationPreviewOptions) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionMigrationPreviewOptions object to a map representation for JSON marshaling.
func (s SubscriptionMigrationPreviewOptions) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.ProductId != nil {
        structMap["product_id"] = s.ProductId
    }
    if s.ProductPricePointId != nil {
        structMap["product_price_point_id"] = s.ProductPricePointId
    }
    if s.IncludeTrial != nil {
        structMap["include_trial"] = s.IncludeTrial
    }
    if s.IncludeInitialCharge != nil {
        structMap["include_initial_charge"] = s.IncludeInitialCharge
    }
    if s.IncludeCoupons != nil {
        structMap["include_coupons"] = s.IncludeCoupons
    }
    if s.PreservePeriod != nil {
        structMap["preserve_period"] = s.PreservePeriod
    }
    if s.ProductHandle != nil {
        structMap["product_handle"] = s.ProductHandle
    }
    if s.ProductPricePointHandle != nil {
        structMap["product_price_point_handle"] = s.ProductPricePointHandle
    }
    if s.Proration != nil {
        structMap["proration"] = s.Proration.toMap()
    }
    if s.ProrationDate != nil {
        structMap["proration_date"] = s.ProrationDate.Format(time.RFC3339)
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionMigrationPreviewOptions.
// It customizes the JSON unmarshaling process for SubscriptionMigrationPreviewOptions objects.
func (s *SubscriptionMigrationPreviewOptions) UnmarshalJSON(input []byte) error {
    var temp subscriptionMigrationPreviewOptions
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "product_id", "product_price_point_id", "include_trial", "include_initial_charge", "include_coupons", "preserve_period", "product_handle", "product_price_point_handle", "proration", "proration_date")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.ProductId = temp.ProductId
    s.ProductPricePointId = temp.ProductPricePointId
    s.IncludeTrial = temp.IncludeTrial
    s.IncludeInitialCharge = temp.IncludeInitialCharge
    s.IncludeCoupons = temp.IncludeCoupons
    s.PreservePeriod = temp.PreservePeriod
    s.ProductHandle = temp.ProductHandle
    s.ProductPricePointHandle = temp.ProductPricePointHandle
    s.Proration = temp.Proration
    if temp.ProrationDate != nil {
        ProrationDateVal, err := time.Parse(time.RFC3339, *temp.ProrationDate)
        if err != nil {
            log.Fatalf("Cannot Parse proration_date as % s format.", time.RFC3339)
        }
        s.ProrationDate = &ProrationDateVal
    }
    return nil
}

// subscriptionMigrationPreviewOptions is a temporary struct used for validating the fields of SubscriptionMigrationPreviewOptions.
type subscriptionMigrationPreviewOptions  struct {
    ProductId               *int       `json:"product_id,omitempty"`
    ProductPricePointId     *int       `json:"product_price_point_id,omitempty"`
    IncludeTrial            *bool      `json:"include_trial,omitempty"`
    IncludeInitialCharge    *bool      `json:"include_initial_charge,omitempty"`
    IncludeCoupons          *bool      `json:"include_coupons,omitempty"`
    PreservePeriod          *bool      `json:"preserve_period,omitempty"`
    ProductHandle           *string    `json:"product_handle,omitempty"`
    ProductPricePointHandle *string    `json:"product_price_point_handle,omitempty"`
    Proration               *Proration `json:"proration,omitempty"`
    ProrationDate           *string    `json:"proration_date,omitempty"`
}
