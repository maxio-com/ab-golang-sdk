package models

import (
    "encoding/json"
    "errors"
    "log"
    "strings"
    "time"
)

// CreateOrUpdatePercentageCoupon represents a CreateOrUpdatePercentageCoupon struct.
type CreateOrUpdatePercentageCoupon struct {
    // the name of the coupon
    Name                          string                                   `json:"name"`
    // may contain uppercase alphanumeric characters and these special characters (which allow for email addresses to be used): “%”, “@”, “+”, “-”, “_”, and “.”
    Code                          string                                   `json:"code"`
    Description                   *string                                  `json:"description,omitempty"`
    Percentage                    CreateOrUpdatePercentageCouponPercentage `json:"percentage"`
    AllowNegativeBalance          *bool                                    `json:"allow_negative_balance,omitempty"`
    Recurring                     *bool                                    `json:"recurring,omitempty"`
    EndDate                       *time.Time                               `json:"end_date,omitempty"`
    ProductFamilyId               *string                                  `json:"product_family_id,omitempty"`
    Stackable                     *bool                                    `json:"stackable,omitempty"`
    CompoundingStrategy           *CompoundingStrategy                     `json:"compounding_strategy,omitempty"`
    ExcludeMidPeriodAllocations   *bool                                    `json:"exclude_mid_period_allocations,omitempty"`
    ApplyOnCancelAtEndOfPeriod    *bool                                    `json:"apply_on_cancel_at_end_of_period,omitempty"`
    ApplyOnSubscriptionExpiration *bool                                    `json:"apply_on_subscription_expiration,omitempty"`
    AdditionalProperties          map[string]any                           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreateOrUpdatePercentageCoupon.
// It customizes the JSON marshaling process for CreateOrUpdatePercentageCoupon objects.
func (c CreateOrUpdatePercentageCoupon) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateOrUpdatePercentageCoupon object to a map representation for JSON marshaling.
func (c CreateOrUpdatePercentageCoupon) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["name"] = c.Name
    structMap["code"] = c.Code
    if c.Description != nil {
        structMap["description"] = c.Description
    }
    structMap["percentage"] = c.Percentage.toMap()
    if c.AllowNegativeBalance != nil {
        structMap["allow_negative_balance"] = c.AllowNegativeBalance
    }
    if c.Recurring != nil {
        structMap["recurring"] = c.Recurring
    }
    if c.EndDate != nil {
        structMap["end_date"] = c.EndDate.Format(time.RFC3339)
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

// UnmarshalJSON implements the json.Unmarshaler interface for CreateOrUpdatePercentageCoupon.
// It customizes the JSON unmarshaling process for CreateOrUpdatePercentageCoupon objects.
func (c *CreateOrUpdatePercentageCoupon) UnmarshalJSON(input []byte) error {
    var temp createOrUpdatePercentageCoupon
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "name", "code", "description", "percentage", "allow_negative_balance", "recurring", "end_date", "product_family_id", "stackable", "compounding_strategy", "exclude_mid_period_allocations", "apply_on_cancel_at_end_of_period", "apply_on_subscription_expiration")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Name = *temp.Name
    c.Code = *temp.Code
    c.Description = temp.Description
    c.Percentage = *temp.Percentage
    c.AllowNegativeBalance = temp.AllowNegativeBalance
    c.Recurring = temp.Recurring
    if temp.EndDate != nil {
        EndDateVal, err := time.Parse(time.RFC3339, *temp.EndDate)
        if err != nil {
            log.Fatalf("Cannot Parse end_date as % s format.", time.RFC3339)
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

// TODO
type createOrUpdatePercentageCoupon  struct {
    Name                          *string                                   `json:"name"`
    Code                          *string                                   `json:"code"`
    Description                   *string                                   `json:"description,omitempty"`
    Percentage                    *CreateOrUpdatePercentageCouponPercentage `json:"percentage"`
    AllowNegativeBalance          *bool                                     `json:"allow_negative_balance,omitempty"`
    Recurring                     *bool                                     `json:"recurring,omitempty"`
    EndDate                       *string                                   `json:"end_date,omitempty"`
    ProductFamilyId               *string                                   `json:"product_family_id,omitempty"`
    Stackable                     *bool                                     `json:"stackable,omitempty"`
    CompoundingStrategy           *CompoundingStrategy                      `json:"compounding_strategy,omitempty"`
    ExcludeMidPeriodAllocations   *bool                                     `json:"exclude_mid_period_allocations,omitempty"`
    ApplyOnCancelAtEndOfPeriod    *bool                                     `json:"apply_on_cancel_at_end_of_period,omitempty"`
    ApplyOnSubscriptionExpiration *bool                                     `json:"apply_on_subscription_expiration,omitempty"`
}

func (c *createOrUpdatePercentageCoupon) validate() error {
    var errs []string
    if c.Name == nil {
        errs = append(errs, "required field `name` is missing for type `Create or Update Percentage Coupon`")
    }
    if c.Code == nil {
        errs = append(errs, "required field `code` is missing for type `Create or Update Percentage Coupon`")
    }
    if c.Percentage == nil {
        errs = append(errs, "required field `percentage` is missing for type `Create or Update Percentage Coupon`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
