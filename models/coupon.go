// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// Coupon represents a Coupon struct.
type Coupon struct {
    Id                            *int                          `json:"id,omitempty"`
    Name                          *string                       `json:"name,omitempty"`
    Code                          *string                       `json:"code,omitempty"`
    Description                   *string                       `json:"description,omitempty"`
    Amount                        Optional[float64]             `json:"amount"`
    AmountInCents                 Optional[int64]               `json:"amount_in_cents"`
    ProductFamilyId               *int                          `json:"product_family_id,omitempty"`
    ProductFamilyName             Optional[string]              `json:"product_family_name"`
    StartDate                     *time.Time                    `json:"start_date,omitempty"`
    // After the given time, this coupon code will be invalid for new signups. Recurring discounts started before this date will continue to recur even after this date.
    EndDate                       Optional[time.Time]           `json:"end_date"`
    Percentage                    Optional[string]              `json:"percentage"`
    Recurring                     *bool                         `json:"recurring,omitempty"`
    RecurringScheme               *RecurringScheme              `json:"recurring_scheme,omitempty"`
    DurationPeriodCount           Optional[int]                 `json:"duration_period_count"`
    DurationInterval              Optional[int]                 `json:"duration_interval"`
    DurationIntervalUnit          Optional[string]              `json:"duration_interval_unit"`
    DurationIntervalSpan          Optional[string]              `json:"duration_interval_span"`
    // If set to true, discount is not limited (credits will carry forward to next billing).
    AllowNegativeBalance          *bool                         `json:"allow_negative_balance,omitempty"`
    ArchivedAt                    Optional[time.Time]           `json:"archived_at"`
    ConversionLimit               Optional[string]              `json:"conversion_limit"`
    // A stackable coupon can be combined with other coupons on a Subscription.
    Stackable                     *bool                         `json:"stackable,omitempty"`
    // Applicable only to stackable coupons. For `compound`, Percentage-based discounts will be calculated against the remaining price, after prior discounts have been calculated. For `full-price`, Percentage-based discounts will always be calculated against the original item price, before other discounts are applied.
    CompoundingStrategy           Optional[CompoundingStrategy] `json:"compounding_strategy"`
    UseSiteExchangeRate           *bool                         `json:"use_site_exchange_rate,omitempty"`
    CreatedAt                     *time.Time                    `json:"created_at,omitempty"`
    UpdatedAt                     *time.Time                    `json:"updated_at,omitempty"`
    DiscountType                  *DiscountType                 `json:"discount_type,omitempty"`
    ExcludeMidPeriodAllocations   *bool                         `json:"exclude_mid_period_allocations,omitempty"`
    ApplyOnCancelAtEndOfPeriod    *bool                         `json:"apply_on_cancel_at_end_of_period,omitempty"`
    ApplyOnSubscriptionExpiration *bool                         `json:"apply_on_subscription_expiration,omitempty"`
    CouponRestrictions            []CouponRestriction           `json:"coupon_restrictions,omitempty"`
    // Returned in read, find, and list endpoints if the query parameter is provided.
    CurrencyPrices                []CouponCurrency              `json:"currency_prices,omitempty"`
    AdditionalProperties          map[string]interface{}        `json:"_"`
}

// String implements the fmt.Stringer interface for Coupon,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c Coupon) String() string {
    return fmt.Sprintf(
    	"Coupon[Id=%v, Name=%v, Code=%v, Description=%v, Amount=%v, AmountInCents=%v, ProductFamilyId=%v, ProductFamilyName=%v, StartDate=%v, EndDate=%v, Percentage=%v, Recurring=%v, RecurringScheme=%v, DurationPeriodCount=%v, DurationInterval=%v, DurationIntervalUnit=%v, DurationIntervalSpan=%v, AllowNegativeBalance=%v, ArchivedAt=%v, ConversionLimit=%v, Stackable=%v, CompoundingStrategy=%v, UseSiteExchangeRate=%v, CreatedAt=%v, UpdatedAt=%v, DiscountType=%v, ExcludeMidPeriodAllocations=%v, ApplyOnCancelAtEndOfPeriod=%v, ApplyOnSubscriptionExpiration=%v, CouponRestrictions=%v, CurrencyPrices=%v, AdditionalProperties=%v]",
    	c.Id, c.Name, c.Code, c.Description, c.Amount, c.AmountInCents, c.ProductFamilyId, c.ProductFamilyName, c.StartDate, c.EndDate, c.Percentage, c.Recurring, c.RecurringScheme, c.DurationPeriodCount, c.DurationInterval, c.DurationIntervalUnit, c.DurationIntervalSpan, c.AllowNegativeBalance, c.ArchivedAt, c.ConversionLimit, c.Stackable, c.CompoundingStrategy, c.UseSiteExchangeRate, c.CreatedAt, c.UpdatedAt, c.DiscountType, c.ExcludeMidPeriodAllocations, c.ApplyOnCancelAtEndOfPeriod, c.ApplyOnSubscriptionExpiration, c.CouponRestrictions, c.CurrencyPrices, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Coupon.
// It customizes the JSON marshaling process for Coupon objects.
func (c Coupon) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "id", "name", "code", "description", "amount", "amount_in_cents", "product_family_id", "product_family_name", "start_date", "end_date", "percentage", "recurring", "recurring_scheme", "duration_period_count", "duration_interval", "duration_interval_unit", "duration_interval_span", "allow_negative_balance", "archived_at", "conversion_limit", "stackable", "compounding_strategy", "use_site_exchange_rate", "created_at", "updated_at", "discount_type", "exclude_mid_period_allocations", "apply_on_cancel_at_end_of_period", "apply_on_subscription_expiration", "coupon_restrictions", "currency_prices"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the Coupon object to a map representation for JSON marshaling.
func (c Coupon) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Id != nil {
        structMap["id"] = c.Id
    }
    if c.Name != nil {
        structMap["name"] = c.Name
    }
    if c.Code != nil {
        structMap["code"] = c.Code
    }
    if c.Description != nil {
        structMap["description"] = c.Description
    }
    if c.Amount.IsValueSet() {
        if c.Amount.Value() != nil {
            structMap["amount"] = c.Amount.Value()
        } else {
            structMap["amount"] = nil
        }
    }
    if c.AmountInCents.IsValueSet() {
        if c.AmountInCents.Value() != nil {
            structMap["amount_in_cents"] = c.AmountInCents.Value()
        } else {
            structMap["amount_in_cents"] = nil
        }
    }
    if c.ProductFamilyId != nil {
        structMap["product_family_id"] = c.ProductFamilyId
    }
    if c.ProductFamilyName.IsValueSet() {
        if c.ProductFamilyName.Value() != nil {
            structMap["product_family_name"] = c.ProductFamilyName.Value()
        } else {
            structMap["product_family_name"] = nil
        }
    }
    if c.StartDate != nil {
        structMap["start_date"] = c.StartDate.Format(time.RFC3339)
    }
    if c.EndDate.IsValueSet() {
        var EndDateVal *string = nil
        if c.EndDate.Value() != nil {
            val := c.EndDate.Value().Format(time.RFC3339)
            EndDateVal = &val
        }
        if c.EndDate.Value() != nil {
            structMap["end_date"] = EndDateVal
        } else {
            structMap["end_date"] = nil
        }
    }
    if c.Percentage.IsValueSet() {
        if c.Percentage.Value() != nil {
            structMap["percentage"] = c.Percentage.Value()
        } else {
            structMap["percentage"] = nil
        }
    }
    if c.Recurring != nil {
        structMap["recurring"] = c.Recurring
    }
    if c.RecurringScheme != nil {
        structMap["recurring_scheme"] = c.RecurringScheme
    }
    if c.DurationPeriodCount.IsValueSet() {
        if c.DurationPeriodCount.Value() != nil {
            structMap["duration_period_count"] = c.DurationPeriodCount.Value()
        } else {
            structMap["duration_period_count"] = nil
        }
    }
    if c.DurationInterval.IsValueSet() {
        if c.DurationInterval.Value() != nil {
            structMap["duration_interval"] = c.DurationInterval.Value()
        } else {
            structMap["duration_interval"] = nil
        }
    }
    if c.DurationIntervalUnit.IsValueSet() {
        if c.DurationIntervalUnit.Value() != nil {
            structMap["duration_interval_unit"] = c.DurationIntervalUnit.Value()
        } else {
            structMap["duration_interval_unit"] = nil
        }
    }
    if c.DurationIntervalSpan.IsValueSet() {
        if c.DurationIntervalSpan.Value() != nil {
            structMap["duration_interval_span"] = c.DurationIntervalSpan.Value()
        } else {
            structMap["duration_interval_span"] = nil
        }
    }
    if c.AllowNegativeBalance != nil {
        structMap["allow_negative_balance"] = c.AllowNegativeBalance
    }
    if c.ArchivedAt.IsValueSet() {
        var ArchivedAtVal *string = nil
        if c.ArchivedAt.Value() != nil {
            val := c.ArchivedAt.Value().Format(time.RFC3339)
            ArchivedAtVal = &val
        }
        if c.ArchivedAt.Value() != nil {
            structMap["archived_at"] = ArchivedAtVal
        } else {
            structMap["archived_at"] = nil
        }
    }
    if c.ConversionLimit.IsValueSet() {
        if c.ConversionLimit.Value() != nil {
            structMap["conversion_limit"] = c.ConversionLimit.Value()
        } else {
            structMap["conversion_limit"] = nil
        }
    }
    if c.Stackable != nil {
        structMap["stackable"] = c.Stackable
    }
    if c.CompoundingStrategy.IsValueSet() {
        if c.CompoundingStrategy.Value() != nil {
            structMap["compounding_strategy"] = c.CompoundingStrategy.Value()
        } else {
            structMap["compounding_strategy"] = nil
        }
    }
    if c.UseSiteExchangeRate != nil {
        structMap["use_site_exchange_rate"] = c.UseSiteExchangeRate
    }
    if c.CreatedAt != nil {
        structMap["created_at"] = c.CreatedAt.Format(time.RFC3339)
    }
    if c.UpdatedAt != nil {
        structMap["updated_at"] = c.UpdatedAt.Format(time.RFC3339)
    }
    if c.DiscountType != nil {
        structMap["discount_type"] = c.DiscountType
    }
    if c.ExcludeMidPeriodAllocations != nil {
        structMap["exclude_mid_period_allocations"] = c.ExcludeMidPeriodAllocations
    }
    if c.ApplyOnCancelAtEndOfPeriod != nil {
        structMap["apply_on_cancel_at_end_of_period"] = c.ApplyOnCancelAtEndOfPeriod
    }
    if c.ApplyOnSubscriptionExpiration != nil {
        structMap["apply_on_subscription_expiration"] = c.ApplyOnSubscriptionExpiration
    }
    if c.CouponRestrictions != nil {
        structMap["coupon_restrictions"] = c.CouponRestrictions
    }
    if c.CurrencyPrices != nil {
        structMap["currency_prices"] = c.CurrencyPrices
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Coupon.
// It customizes the JSON unmarshaling process for Coupon objects.
func (c *Coupon) UnmarshalJSON(input []byte) error {
    var temp tempCoupon
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "name", "code", "description", "amount", "amount_in_cents", "product_family_id", "product_family_name", "start_date", "end_date", "percentage", "recurring", "recurring_scheme", "duration_period_count", "duration_interval", "duration_interval_unit", "duration_interval_span", "allow_negative_balance", "archived_at", "conversion_limit", "stackable", "compounding_strategy", "use_site_exchange_rate", "created_at", "updated_at", "discount_type", "exclude_mid_period_allocations", "apply_on_cancel_at_end_of_period", "apply_on_subscription_expiration", "coupon_restrictions", "currency_prices")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Id = temp.Id
    c.Name = temp.Name
    c.Code = temp.Code
    c.Description = temp.Description
    c.Amount = temp.Amount
    c.AmountInCents = temp.AmountInCents
    c.ProductFamilyId = temp.ProductFamilyId
    c.ProductFamilyName = temp.ProductFamilyName
    if temp.StartDate != nil {
        StartDateVal, err := time.Parse(time.RFC3339, *temp.StartDate)
        if err != nil {
            log.Fatalf("Cannot Parse start_date as % s format.", time.RFC3339)
        }
        c.StartDate = &StartDateVal
    }
    c.EndDate.ShouldSetValue(temp.EndDate.IsValueSet())
    if temp.EndDate.Value() != nil {
        EndDateVal, err := time.Parse(time.RFC3339, (*temp.EndDate.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse end_date as % s format.", time.RFC3339)
        }
        c.EndDate.SetValue(&EndDateVal)
    }
    c.Percentage = temp.Percentage
    c.Recurring = temp.Recurring
    c.RecurringScheme = temp.RecurringScheme
    c.DurationPeriodCount = temp.DurationPeriodCount
    c.DurationInterval = temp.DurationInterval
    c.DurationIntervalUnit = temp.DurationIntervalUnit
    c.DurationIntervalSpan = temp.DurationIntervalSpan
    c.AllowNegativeBalance = temp.AllowNegativeBalance
    c.ArchivedAt.ShouldSetValue(temp.ArchivedAt.IsValueSet())
    if temp.ArchivedAt.Value() != nil {
        ArchivedAtVal, err := time.Parse(time.RFC3339, (*temp.ArchivedAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse archived_at as % s format.", time.RFC3339)
        }
        c.ArchivedAt.SetValue(&ArchivedAtVal)
    }
    c.ConversionLimit = temp.ConversionLimit
    c.Stackable = temp.Stackable
    c.CompoundingStrategy = temp.CompoundingStrategy
    c.UseSiteExchangeRate = temp.UseSiteExchangeRate
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        c.CreatedAt = &CreatedAtVal
    }
    if temp.UpdatedAt != nil {
        UpdatedAtVal, err := time.Parse(time.RFC3339, *temp.UpdatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse updated_at as % s format.", time.RFC3339)
        }
        c.UpdatedAt = &UpdatedAtVal
    }
    c.DiscountType = temp.DiscountType
    c.ExcludeMidPeriodAllocations = temp.ExcludeMidPeriodAllocations
    c.ApplyOnCancelAtEndOfPeriod = temp.ApplyOnCancelAtEndOfPeriod
    c.ApplyOnSubscriptionExpiration = temp.ApplyOnSubscriptionExpiration
    c.CouponRestrictions = temp.CouponRestrictions
    c.CurrencyPrices = temp.CurrencyPrices
    return nil
}

// tempCoupon is a temporary struct used for validating the fields of Coupon.
type tempCoupon  struct {
    Id                            *int                          `json:"id,omitempty"`
    Name                          *string                       `json:"name,omitempty"`
    Code                          *string                       `json:"code,omitempty"`
    Description                   *string                       `json:"description,omitempty"`
    Amount                        Optional[float64]             `json:"amount"`
    AmountInCents                 Optional[int64]               `json:"amount_in_cents"`
    ProductFamilyId               *int                          `json:"product_family_id,omitempty"`
    ProductFamilyName             Optional[string]              `json:"product_family_name"`
    StartDate                     *string                       `json:"start_date,omitempty"`
    EndDate                       Optional[string]              `json:"end_date"`
    Percentage                    Optional[string]              `json:"percentage"`
    Recurring                     *bool                         `json:"recurring,omitempty"`
    RecurringScheme               *RecurringScheme              `json:"recurring_scheme,omitempty"`
    DurationPeriodCount           Optional[int]                 `json:"duration_period_count"`
    DurationInterval              Optional[int]                 `json:"duration_interval"`
    DurationIntervalUnit          Optional[string]              `json:"duration_interval_unit"`
    DurationIntervalSpan          Optional[string]              `json:"duration_interval_span"`
    AllowNegativeBalance          *bool                         `json:"allow_negative_balance,omitempty"`
    ArchivedAt                    Optional[string]              `json:"archived_at"`
    ConversionLimit               Optional[string]              `json:"conversion_limit"`
    Stackable                     *bool                         `json:"stackable,omitempty"`
    CompoundingStrategy           Optional[CompoundingStrategy] `json:"compounding_strategy"`
    UseSiteExchangeRate           *bool                         `json:"use_site_exchange_rate,omitempty"`
    CreatedAt                     *string                       `json:"created_at,omitempty"`
    UpdatedAt                     *string                       `json:"updated_at,omitempty"`
    DiscountType                  *DiscountType                 `json:"discount_type,omitempty"`
    ExcludeMidPeriodAllocations   *bool                         `json:"exclude_mid_period_allocations,omitempty"`
    ApplyOnCancelAtEndOfPeriod    *bool                         `json:"apply_on_cancel_at_end_of_period,omitempty"`
    ApplyOnSubscriptionExpiration *bool                         `json:"apply_on_subscription_expiration,omitempty"`
    CouponRestrictions            []CouponRestriction           `json:"coupon_restrictions,omitempty"`
    CurrencyPrices                []CouponCurrency              `json:"currency_prices,omitempty"`
}
