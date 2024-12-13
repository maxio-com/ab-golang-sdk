/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "log"
    "time"
)

// CouponPayload represents a CouponPayload struct.
type CouponPayload struct {
    // Required when creating a new coupon. This name is not displayed to customers and is limited to 255 characters.
    Name                          *string                  `json:"name,omitempty"`
    // Required when creating a new coupon. The code is limited to 255 characters. May contain uppercase alphanumeric characters and these special characters (which allow for email addresses to be used): “%”, “@”, “+”, “-”, “_”, and “.”
    Code                          *string                  `json:"code,omitempty"`
    // Required when creating a new coupon. A description of the coupon that can be displayed to customers in transactions and on statements. The description is limited to 255 characters.
    Description                   *string                  `json:"description,omitempty"`
    // Required when creating a new percentage coupon. Can't be used together with amount_in_cents. Percentage discount
    Percentage                    *CouponPayloadPercentage `json:"percentage,omitempty"`
    // Required when creating a new flat amount coupon. Can't be used together with percentage. Flat USD discount
    AmountInCents                 *int64                   `json:"amount_in_cents,omitempty"`
    // If set to true, discount is not limited (credits will carry forward to next billing). Can't be used together with restrictions.
    AllowNegativeBalance          *bool                    `json:"allow_negative_balance,omitempty"`
    Recurring                     *bool                    `json:"recurring,omitempty"`
    // After the end of the given day, this coupon code will be invalid for new signups. Recurring discounts started before this date will continue to recur even after this date.
    EndDate                       *time.Time               `json:"end_date,omitempty"`
    ProductFamilyId               *string                  `json:"product_family_id,omitempty"`
    // A stackable coupon can be combined with other coupons on a Subscription.
    Stackable                     *bool                    `json:"stackable,omitempty"`
    // Applicable only to stackable coupons. For `compound`, Percentage-based discounts will be calculated against the remaining price, after prior discounts have been calculated. For `full-price`, Percentage-based discounts will always be calculated against the original item price, before other discounts are applied.
    CompoundingStrategy           *CompoundingStrategy     `json:"compounding_strategy,omitempty"`
    ExcludeMidPeriodAllocations   *bool                    `json:"exclude_mid_period_allocations,omitempty"`
    ApplyOnCancelAtEndOfPeriod    *bool                    `json:"apply_on_cancel_at_end_of_period,omitempty"`
    ApplyOnSubscriptionExpiration *bool                    `json:"apply_on_subscription_expiration,omitempty"`
    AdditionalProperties          map[string]interface{}   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CouponPayload.
// It customizes the JSON marshaling process for CouponPayload objects.
func (c CouponPayload) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "name", "code", "description", "percentage", "amount_in_cents", "allow_negative_balance", "recurring", "end_date", "product_family_id", "stackable", "compounding_strategy", "exclude_mid_period_allocations", "apply_on_cancel_at_end_of_period", "apply_on_subscription_expiration"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CouponPayload object to a map representation for JSON marshaling.
func (c CouponPayload) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Name != nil {
        structMap["name"] = c.Name
    }
    if c.Code != nil {
        structMap["code"] = c.Code
    }
    if c.Description != nil {
        structMap["description"] = c.Description
    }
    if c.Percentage != nil {
        structMap["percentage"] = c.Percentage.toMap()
    }
    if c.AmountInCents != nil {
        structMap["amount_in_cents"] = c.AmountInCents
    }
    if c.AllowNegativeBalance != nil {
        structMap["allow_negative_balance"] = c.AllowNegativeBalance
    }
    if c.Recurring != nil {
        structMap["recurring"] = c.Recurring
    }
    if c.EndDate != nil {
        structMap["end_date"] = c.EndDate.Format(DEFAULT_DATE)
    }
    if c.ProductFamilyId != nil {
        structMap["product_family_id"] = c.ProductFamilyId
    }
    if c.Stackable != nil {
        structMap["stackable"] = c.Stackable
    }
    if c.CompoundingStrategy != nil {
        structMap["compounding_strategy"] = c.CompoundingStrategy
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
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CouponPayload.
// It customizes the JSON unmarshaling process for CouponPayload objects.
func (c *CouponPayload) UnmarshalJSON(input []byte) error {
    var temp tempCouponPayload
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "name", "code", "description", "percentage", "amount_in_cents", "allow_negative_balance", "recurring", "end_date", "product_family_id", "stackable", "compounding_strategy", "exclude_mid_period_allocations", "apply_on_cancel_at_end_of_period", "apply_on_subscription_expiration")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Name = temp.Name
    c.Code = temp.Code
    c.Description = temp.Description
    c.Percentage = temp.Percentage
    c.AmountInCents = temp.AmountInCents
    c.AllowNegativeBalance = temp.AllowNegativeBalance
    c.Recurring = temp.Recurring
    if temp.EndDate != nil {
        EndDateVal, err := time.Parse(DEFAULT_DATE, *temp.EndDate)
        if err != nil {
            log.Fatalf("Cannot Parse end_date as % s format.", DEFAULT_DATE)
        }
        c.EndDate = &EndDateVal
    }
    c.ProductFamilyId = temp.ProductFamilyId
    c.Stackable = temp.Stackable
    c.CompoundingStrategy = temp.CompoundingStrategy
    c.ExcludeMidPeriodAllocations = temp.ExcludeMidPeriodAllocations
    c.ApplyOnCancelAtEndOfPeriod = temp.ApplyOnCancelAtEndOfPeriod
    c.ApplyOnSubscriptionExpiration = temp.ApplyOnSubscriptionExpiration
    return nil
}

// tempCouponPayload is a temporary struct used for validating the fields of CouponPayload.
type tempCouponPayload  struct {
    Name                          *string                  `json:"name,omitempty"`
    Code                          *string                  `json:"code,omitempty"`
    Description                   *string                  `json:"description,omitempty"`
    Percentage                    *CouponPayloadPercentage `json:"percentage,omitempty"`
    AmountInCents                 *int64                   `json:"amount_in_cents,omitempty"`
    AllowNegativeBalance          *bool                    `json:"allow_negative_balance,omitempty"`
    Recurring                     *bool                    `json:"recurring,omitempty"`
    EndDate                       *string                  `json:"end_date,omitempty"`
    ProductFamilyId               *string                  `json:"product_family_id,omitempty"`
    Stackable                     *bool                    `json:"stackable,omitempty"`
    CompoundingStrategy           *CompoundingStrategy     `json:"compounding_strategy,omitempty"`
    ExcludeMidPeriodAllocations   *bool                    `json:"exclude_mid_period_allocations,omitempty"`
    ApplyOnCancelAtEndOfPeriod    *bool                    `json:"apply_on_cancel_at_end_of_period,omitempty"`
    ApplyOnSubscriptionExpiration *bool                    `json:"apply_on_subscription_expiration,omitempty"`
}
