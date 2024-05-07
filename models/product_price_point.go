package models

import (
    "encoding/json"
    "log"
    "time"
)

// ProductPricePoint represents a ProductPricePoint struct.
type ProductPricePoint struct {
    Id                      *int                   `json:"id,omitempty"`
    // The product price point name
    Name                    *string                `json:"name,omitempty"`
    // The product price point API handle
    Handle                  Optional[string]       `json:"handle"`
    // The product price point price, in integer cents
    PriceInCents            *int64                 `json:"price_in_cents,omitempty"`
    // The numerical interval. i.e. an interval of ‘30’ coupled with an interval_unit of day would mean this product price point would renew every 30 days
    Interval                *int                   `json:"interval,omitempty"`
    // A string representing the interval unit for this product price point, either month or day
    IntervalUnit            *IntervalUnit          `json:"interval_unit,omitempty"`
    // The product price point trial price, in integer cents
    TrialPriceInCents       Optional[int64]        `json:"trial_price_in_cents"`
    // The numerical trial interval. i.e. an interval of ‘30’ coupled with a trial_interval_unit of day would mean this product price point trial would last 30 days
    TrialInterval           Optional[int]          `json:"trial_interval"`
    // A string representing the trial interval unit for this product price point, either month or day
    TrialIntervalUnit       Optional[IntervalUnit] `json:"trial_interval_unit"`
    TrialType               *string                `json:"trial_type,omitempty"`
    // reserved for future use
    IntroductoryOffer       Optional[bool]         `json:"introductory_offer"`
    // The product price point initial charge, in integer cents
    InitialChargeInCents    Optional[int64]        `json:"initial_charge_in_cents"`
    InitialChargeAfterTrial Optional[bool]         `json:"initial_charge_after_trial"`
    // The numerical expiration interval. i.e. an expiration_interval of ‘30’ coupled with an expiration_interval_unit of day would mean this product price point would expire after 30 days
    ExpirationInterval      Optional[int]          `json:"expiration_interval"`
    // A string representing the expiration interval unit for this product price point, either month or day
    ExpirationIntervalUnit  Optional[IntervalUnit] `json:"expiration_interval_unit"`
    // The product id this price point belongs to
    ProductId               *int                   `json:"product_id,omitempty"`
    // Timestamp indicating when this price point was archived
    ArchivedAt              Optional[time.Time]    `json:"archived_at"`
    // Timestamp indicating when this price point was created
    CreatedAt               *time.Time             `json:"created_at,omitempty"`
    // Timestamp indicating when this price point was last updated
    UpdatedAt               *time.Time             `json:"updated_at,omitempty"`
    // Whether or not to use the site's exchange rate or define your own pricing when your site has multiple currencies defined.
    UseSiteExchangeRate     *bool                  `json:"use_site_exchange_rate,omitempty"`
    // The type of price point
    Type                    *PricePointType        `json:"type,omitempty"`
    // Whether or not the price point includes tax
    TaxIncluded             *bool                  `json:"tax_included,omitempty"`
    // The subscription id this price point belongs to
    SubscriptionId          Optional[int]          `json:"subscription_id"`
    // An array of currency pricing data is available when multiple currencies are defined for the site. It varies based on the use_site_exchange_rate setting for the price point. This parameter is present only in the response of read endpoints, after including the appropriate query parameter.
    CurrencyPrices          []CurrencyPrice        `json:"currency_prices,omitempty"`
    AdditionalProperties    map[string]any         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ProductPricePoint.
// It customizes the JSON marshaling process for ProductPricePoint objects.
func (p ProductPricePoint) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the ProductPricePoint object to a map representation for JSON marshaling.
func (p ProductPricePoint) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
    if p.Id != nil {
        structMap["id"] = p.Id
    }
    if p.Name != nil {
        structMap["name"] = p.Name
    }
    if p.Handle.IsValueSet() {
        if p.Handle.Value() != nil {
            structMap["handle"] = p.Handle.Value()
        } else {
            structMap["handle"] = nil
        }
    }
    if p.PriceInCents != nil {
        structMap["price_in_cents"] = p.PriceInCents
    }
    if p.Interval != nil {
        structMap["interval"] = p.Interval
    }
    if p.IntervalUnit != nil {
        structMap["interval_unit"] = p.IntervalUnit
    }
    if p.TrialPriceInCents.IsValueSet() {
        if p.TrialPriceInCents.Value() != nil {
            structMap["trial_price_in_cents"] = p.TrialPriceInCents.Value()
        } else {
            structMap["trial_price_in_cents"] = nil
        }
    }
    if p.TrialInterval.IsValueSet() {
        if p.TrialInterval.Value() != nil {
            structMap["trial_interval"] = p.TrialInterval.Value()
        } else {
            structMap["trial_interval"] = nil
        }
    }
    if p.TrialIntervalUnit.IsValueSet() {
        if p.TrialIntervalUnit.Value() != nil {
            structMap["trial_interval_unit"] = p.TrialIntervalUnit.Value()
        } else {
            structMap["trial_interval_unit"] = nil
        }
    }
    if p.TrialType != nil {
        structMap["trial_type"] = p.TrialType
    }
    if p.IntroductoryOffer.IsValueSet() {
        if p.IntroductoryOffer.Value() != nil {
            structMap["introductory_offer"] = p.IntroductoryOffer.Value()
        } else {
            structMap["introductory_offer"] = nil
        }
    }
    if p.InitialChargeInCents.IsValueSet() {
        if p.InitialChargeInCents.Value() != nil {
            structMap["initial_charge_in_cents"] = p.InitialChargeInCents.Value()
        } else {
            structMap["initial_charge_in_cents"] = nil
        }
    }
    if p.InitialChargeAfterTrial.IsValueSet() {
        if p.InitialChargeAfterTrial.Value() != nil {
            structMap["initial_charge_after_trial"] = p.InitialChargeAfterTrial.Value()
        } else {
            structMap["initial_charge_after_trial"] = nil
        }
    }
    if p.ExpirationInterval.IsValueSet() {
        if p.ExpirationInterval.Value() != nil {
            structMap["expiration_interval"] = p.ExpirationInterval.Value()
        } else {
            structMap["expiration_interval"] = nil
        }
    }
    if p.ExpirationIntervalUnit.IsValueSet() {
        if p.ExpirationIntervalUnit.Value() != nil {
            structMap["expiration_interval_unit"] = p.ExpirationIntervalUnit.Value()
        } else {
            structMap["expiration_interval_unit"] = nil
        }
    }
    if p.ProductId != nil {
        structMap["product_id"] = p.ProductId
    }
    if p.ArchivedAt.IsValueSet() {
        var ArchivedAtVal *string = nil
        if p.ArchivedAt.Value() != nil {
            val := p.ArchivedAt.Value().Format(time.RFC3339)
            ArchivedAtVal = &val
        }
        if p.ArchivedAt.Value() != nil {
            structMap["archived_at"] = ArchivedAtVal
        } else {
            structMap["archived_at"] = nil
        }
    }
    if p.CreatedAt != nil {
        structMap["created_at"] = p.CreatedAt.Format(time.RFC3339)
    }
    if p.UpdatedAt != nil {
        structMap["updated_at"] = p.UpdatedAt.Format(time.RFC3339)
    }
    if p.UseSiteExchangeRate != nil {
        structMap["use_site_exchange_rate"] = p.UseSiteExchangeRate
    }
    if p.Type != nil {
        structMap["type"] = p.Type
    }
    if p.TaxIncluded != nil {
        structMap["tax_included"] = p.TaxIncluded
    }
    if p.SubscriptionId.IsValueSet() {
        if p.SubscriptionId.Value() != nil {
            structMap["subscription_id"] = p.SubscriptionId.Value()
        } else {
            structMap["subscription_id"] = nil
        }
    }
    if p.CurrencyPrices != nil {
        structMap["currency_prices"] = p.CurrencyPrices
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ProductPricePoint.
// It customizes the JSON unmarshaling process for ProductPricePoint objects.
func (p *ProductPricePoint) UnmarshalJSON(input []byte) error {
    var temp productPricePoint
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "id", "name", "handle", "price_in_cents", "interval", "interval_unit", "trial_price_in_cents", "trial_interval", "trial_interval_unit", "trial_type", "introductory_offer", "initial_charge_in_cents", "initial_charge_after_trial", "expiration_interval", "expiration_interval_unit", "product_id", "archived_at", "created_at", "updated_at", "use_site_exchange_rate", "type", "tax_included", "subscription_id", "currency_prices")
    if err != nil {
    	return err
    }
    
    p.AdditionalProperties = additionalProperties
    p.Id = temp.Id
    p.Name = temp.Name
    p.Handle = temp.Handle
    p.PriceInCents = temp.PriceInCents
    p.Interval = temp.Interval
    p.IntervalUnit = temp.IntervalUnit
    p.TrialPriceInCents = temp.TrialPriceInCents
    p.TrialInterval = temp.TrialInterval
    p.TrialIntervalUnit = temp.TrialIntervalUnit
    p.TrialType = temp.TrialType
    p.IntroductoryOffer = temp.IntroductoryOffer
    p.InitialChargeInCents = temp.InitialChargeInCents
    p.InitialChargeAfterTrial = temp.InitialChargeAfterTrial
    p.ExpirationInterval = temp.ExpirationInterval
    p.ExpirationIntervalUnit = temp.ExpirationIntervalUnit
    p.ProductId = temp.ProductId
    p.ArchivedAt.ShouldSetValue(temp.ArchivedAt.IsValueSet())
    if temp.ArchivedAt.Value() != nil {
        ArchivedAtVal, err := time.Parse(time.RFC3339, (*temp.ArchivedAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse archived_at as % s format.", time.RFC3339)
        }
        p.ArchivedAt.SetValue(&ArchivedAtVal)
    }
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        p.CreatedAt = &CreatedAtVal
    }
    if temp.UpdatedAt != nil {
        UpdatedAtVal, err := time.Parse(time.RFC3339, *temp.UpdatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse updated_at as % s format.", time.RFC3339)
        }
        p.UpdatedAt = &UpdatedAtVal
    }
    p.UseSiteExchangeRate = temp.UseSiteExchangeRate
    p.Type = temp.Type
    p.TaxIncluded = temp.TaxIncluded
    p.SubscriptionId = temp.SubscriptionId
    p.CurrencyPrices = temp.CurrencyPrices
    return nil
}

// productPricePoint is a temporary struct used for validating the fields of ProductPricePoint.
type productPricePoint  struct {
    Id                      *int                   `json:"id,omitempty"`
    Name                    *string                `json:"name,omitempty"`
    Handle                  Optional[string]       `json:"handle"`
    PriceInCents            *int64                 `json:"price_in_cents,omitempty"`
    Interval                *int                   `json:"interval,omitempty"`
    IntervalUnit            *IntervalUnit          `json:"interval_unit,omitempty"`
    TrialPriceInCents       Optional[int64]        `json:"trial_price_in_cents"`
    TrialInterval           Optional[int]          `json:"trial_interval"`
    TrialIntervalUnit       Optional[IntervalUnit] `json:"trial_interval_unit"`
    TrialType               *string                `json:"trial_type,omitempty"`
    IntroductoryOffer       Optional[bool]         `json:"introductory_offer"`
    InitialChargeInCents    Optional[int64]        `json:"initial_charge_in_cents"`
    InitialChargeAfterTrial Optional[bool]         `json:"initial_charge_after_trial"`
    ExpirationInterval      Optional[int]          `json:"expiration_interval"`
    ExpirationIntervalUnit  Optional[IntervalUnit] `json:"expiration_interval_unit"`
    ProductId               *int                   `json:"product_id,omitempty"`
    ArchivedAt              Optional[string]       `json:"archived_at"`
    CreatedAt               *string                `json:"created_at,omitempty"`
    UpdatedAt               *string                `json:"updated_at,omitempty"`
    UseSiteExchangeRate     *bool                  `json:"use_site_exchange_rate,omitempty"`
    Type                    *PricePointType        `json:"type,omitempty"`
    TaxIncluded             *bool                  `json:"tax_included,omitempty"`
    SubscriptionId          Optional[int]          `json:"subscription_id"`
    CurrencyPrices          []CurrencyPrice        `json:"currency_prices,omitempty"`
}
