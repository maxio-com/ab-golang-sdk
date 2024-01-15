package models

import (
    "encoding/json"
    "log"
    "time"
)

// SubscriptionComponent represents a SubscriptionComponent struct.
type SubscriptionComponent struct {
    Id                        *int                               `json:"id,omitempty"`
    Name                      *string                            `json:"name,omitempty"`
    // A handle for the component type
    Kind                      *ComponentKind                     `json:"kind,omitempty"`
    UnitName                  *string                            `json:"unit_name,omitempty"`
    // (for on/off components) indicates if the component is enabled for the subscription
    Enabled                   *bool                              `json:"enabled,omitempty"`
    UnitBalance               *int                               `json:"unit_balance,omitempty"`
    Currency                  *string                            `json:"currency,omitempty"`
    // For Quantity-based components: The current allocation for the component on the given subscription. For On/Off components: Use 1 for on. Use 0 for off.
    AllocatedQuantity         *interface{}                       `json:"allocated_quantity,omitempty"`
    PricingScheme             Optional[PricingScheme]            `json:"pricing_scheme"`
    ComponentId               *int                               `json:"component_id,omitempty"`
    ComponentHandle           Optional[string]                   `json:"component_handle"`
    SubscriptionId            *int                               `json:"subscription_id,omitempty"`
    Recurring                 *bool                              `json:"recurring,omitempty"`
    // The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
    // Available values: `full`, `prorated`, `none`.
    UpgradeCharge             Optional[CreditType]               `json:"upgrade_charge"`
    // The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
    // Available values: `full`, `prorated`, `none`.
    DowngradeCredit           Optional[CreditType]               `json:"downgrade_credit"`
    ArchivedAt                Optional[string]                   `json:"archived_at"`
    PricePointId              Optional[int]                      `json:"price_point_id"`
    PricePointHandle          Optional[string]                   `json:"price_point_handle"`
    PricePointType            *interface{}                       `json:"price_point_type,omitempty"`
    PricePointName            Optional[string]                   `json:"price_point_name"`
    ProductFamilyId           *int                               `json:"product_family_id,omitempty"`
    ProductFamilyHandle       *string                            `json:"product_family_handle,omitempty"`
    CreatedAt                 *time.Time                         `json:"created_at,omitempty"`
    UpdatedAt                 *time.Time                         `json:"updated_at,omitempty"`
    UseSiteExchangeRate       Optional[bool]                     `json:"use_site_exchange_rate"`
    Description               Optional[string]                   `json:"description"`
    AllowFractionalQuantities *bool                              `json:"allow_fractional_quantities,omitempty"`
    // An optional object, will be returned if provided `include=subscription` query param.
    Subscription              *SubscriptionComponentSubscription `json:"subscription,omitempty"`
    DisplayOnHostedPage       *bool                              `json:"display_on_hosted_page,omitempty"`
    // The numerical interval. i.e. an interval of '30' coupled with an interval_unit of day would mean this component price point would renew every 30 days. This property is only available for sites with Multifrequency enabled.
    Interval                  *int                               `json:"interval,omitempty"`
    // A string representing the interval unit for this component price point, either month or day. This property is only available for sites with Multifrequency enabled.
    IntervalUnit              *IntervalUnit                      `json:"interval_unit,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionComponent.
// It customizes the JSON marshaling process for SubscriptionComponent objects.
func (s *SubscriptionComponent) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionComponent object to a map representation for JSON marshaling.
func (s *SubscriptionComponent) toMap() map[string]any {
    structMap := make(map[string]any)
    if s.Id != nil {
        structMap["id"] = s.Id
    }
    if s.Name != nil {
        structMap["name"] = s.Name
    }
    if s.Kind != nil {
        structMap["kind"] = s.Kind
    }
    if s.UnitName != nil {
        structMap["unit_name"] = s.UnitName
    }
    if s.Enabled != nil {
        structMap["enabled"] = s.Enabled
    }
    if s.UnitBalance != nil {
        structMap["unit_balance"] = s.UnitBalance
    }
    if s.Currency != nil {
        structMap["currency"] = s.Currency
    }
    if s.AllocatedQuantity != nil {
        structMap["allocated_quantity"] = s.AllocatedQuantity
    }
    if s.PricingScheme.IsValueSet() {
        structMap["pricing_scheme"] = s.PricingScheme.Value()
    }
    if s.ComponentId != nil {
        structMap["component_id"] = s.ComponentId
    }
    if s.ComponentHandle.IsValueSet() {
        structMap["component_handle"] = s.ComponentHandle.Value()
    }
    if s.SubscriptionId != nil {
        structMap["subscription_id"] = s.SubscriptionId
    }
    if s.Recurring != nil {
        structMap["recurring"] = s.Recurring
    }
    if s.UpgradeCharge.IsValueSet() {
        structMap["upgrade_charge"] = s.UpgradeCharge.Value()
    }
    if s.DowngradeCredit.IsValueSet() {
        structMap["downgrade_credit"] = s.DowngradeCredit.Value()
    }
    if s.ArchivedAt.IsValueSet() {
        structMap["archived_at"] = s.ArchivedAt.Value()
    }
    if s.PricePointId.IsValueSet() {
        structMap["price_point_id"] = s.PricePointId.Value()
    }
    if s.PricePointHandle.IsValueSet() {
        structMap["price_point_handle"] = s.PricePointHandle.Value()
    }
    if s.PricePointType != nil {
        structMap["price_point_type"] = s.PricePointType
    }
    if s.PricePointName.IsValueSet() {
        structMap["price_point_name"] = s.PricePointName.Value()
    }
    if s.ProductFamilyId != nil {
        structMap["product_family_id"] = s.ProductFamilyId
    }
    if s.ProductFamilyHandle != nil {
        structMap["product_family_handle"] = s.ProductFamilyHandle
    }
    if s.CreatedAt != nil {
        structMap["created_at"] = s.CreatedAt.Format(time.RFC3339)
    }
    if s.UpdatedAt != nil {
        structMap["updated_at"] = s.UpdatedAt.Format(time.RFC3339)
    }
    if s.UseSiteExchangeRate.IsValueSet() {
        structMap["use_site_exchange_rate"] = s.UseSiteExchangeRate.Value()
    }
    if s.Description.IsValueSet() {
        structMap["description"] = s.Description.Value()
    }
    if s.AllowFractionalQuantities != nil {
        structMap["allow_fractional_quantities"] = s.AllowFractionalQuantities
    }
    if s.Subscription != nil {
        structMap["subscription"] = s.Subscription
    }
    if s.DisplayOnHostedPage != nil {
        structMap["display_on_hosted_page"] = s.DisplayOnHostedPage
    }
    if s.Interval != nil {
        structMap["interval"] = s.Interval
    }
    if s.IntervalUnit != nil {
        structMap["interval_unit"] = s.IntervalUnit
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionComponent.
// It customizes the JSON unmarshaling process for SubscriptionComponent objects.
func (s *SubscriptionComponent) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id                        *int                               `json:"id,omitempty"`
        Name                      *string                            `json:"name,omitempty"`
        Kind                      *ComponentKind                     `json:"kind,omitempty"`
        UnitName                  *string                            `json:"unit_name,omitempty"`
        Enabled                   *bool                              `json:"enabled,omitempty"`
        UnitBalance               *int                               `json:"unit_balance,omitempty"`
        Currency                  *string                            `json:"currency,omitempty"`
        AllocatedQuantity         *interface{}                       `json:"allocated_quantity,omitempty"`
        PricingScheme             Optional[PricingScheme]            `json:"pricing_scheme"`
        ComponentId               *int                               `json:"component_id,omitempty"`
        ComponentHandle           Optional[string]                   `json:"component_handle"`
        SubscriptionId            *int                               `json:"subscription_id,omitempty"`
        Recurring                 *bool                              `json:"recurring,omitempty"`
        UpgradeCharge             Optional[CreditType]               `json:"upgrade_charge"`
        DowngradeCredit           Optional[CreditType]               `json:"downgrade_credit"`
        ArchivedAt                Optional[string]                   `json:"archived_at"`
        PricePointId              Optional[int]                      `json:"price_point_id"`
        PricePointHandle          Optional[string]                   `json:"price_point_handle"`
        PricePointType            *interface{}                       `json:"price_point_type,omitempty"`
        PricePointName            Optional[string]                   `json:"price_point_name"`
        ProductFamilyId           *int                               `json:"product_family_id,omitempty"`
        ProductFamilyHandle       *string                            `json:"product_family_handle,omitempty"`
        CreatedAt                 *string                            `json:"created_at,omitempty"`
        UpdatedAt                 *string                            `json:"updated_at,omitempty"`
        UseSiteExchangeRate       Optional[bool]                     `json:"use_site_exchange_rate"`
        Description               Optional[string]                   `json:"description"`
        AllowFractionalQuantities *bool                              `json:"allow_fractional_quantities,omitempty"`
        Subscription              *SubscriptionComponentSubscription `json:"subscription,omitempty"`
        DisplayOnHostedPage       *bool                              `json:"display_on_hosted_page,omitempty"`
        Interval                  *int                               `json:"interval,omitempty"`
        IntervalUnit              *IntervalUnit                      `json:"interval_unit,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    s.Id = temp.Id
    s.Name = temp.Name
    s.Kind = temp.Kind
    s.UnitName = temp.UnitName
    s.Enabled = temp.Enabled
    s.UnitBalance = temp.UnitBalance
    s.Currency = temp.Currency
    s.AllocatedQuantity = temp.AllocatedQuantity
    s.PricingScheme = temp.PricingScheme
    s.ComponentId = temp.ComponentId
    s.ComponentHandle = temp.ComponentHandle
    s.SubscriptionId = temp.SubscriptionId
    s.Recurring = temp.Recurring
    s.UpgradeCharge = temp.UpgradeCharge
    s.DowngradeCredit = temp.DowngradeCredit
    s.ArchivedAt = temp.ArchivedAt
    s.PricePointId = temp.PricePointId
    s.PricePointHandle = temp.PricePointHandle
    s.PricePointType = temp.PricePointType
    s.PricePointName = temp.PricePointName
    s.ProductFamilyId = temp.ProductFamilyId
    s.ProductFamilyHandle = temp.ProductFamilyHandle
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        s.CreatedAt = &CreatedAtVal
    }
    if temp.UpdatedAt != nil {
        UpdatedAtVal, err := time.Parse(time.RFC3339, *temp.UpdatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse updated_at as % s format.", time.RFC3339)
        }
        s.UpdatedAt = &UpdatedAtVal
    }
    s.UseSiteExchangeRate = temp.UseSiteExchangeRate
    s.Description = temp.Description
    s.AllowFractionalQuantities = temp.AllowFractionalQuantities
    s.Subscription = temp.Subscription
    s.DisplayOnHostedPage = temp.DisplayOnHostedPage
    s.Interval = temp.Interval
    s.IntervalUnit = temp.IntervalUnit
    return nil
}
