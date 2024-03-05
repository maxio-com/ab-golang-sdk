package models

import (
    "encoding/json"
    "log"
    "time"
)

// Offer represents a Offer struct.
type Offer struct {
    Id                     *int                `json:"id,omitempty"`
    SiteId                 *int                `json:"site_id,omitempty"`
    ProductFamilyId        *int                `json:"product_family_id,omitempty"`
    ProductId              *int                `json:"product_id,omitempty"`
    ProductPricePointId    *int                `json:"product_price_point_id,omitempty"`
    ProductRevisableNumber *int                `json:"product_revisable_number,omitempty"`
    Name                   *string             `json:"name,omitempty"`
    Handle                 *string             `json:"handle,omitempty"`
    Description            Optional[string]    `json:"description"`
    CreatedAt              *time.Time          `json:"created_at,omitempty"`
    UpdatedAt              *time.Time          `json:"updated_at,omitempty"`
    ArchivedAt             Optional[time.Time] `json:"archived_at"`
    OfferItems             []OfferItem         `json:"offer_items,omitempty"`
    OfferDiscounts         []OfferDiscount     `json:"offer_discounts,omitempty"`
    ProductFamilyName      *string             `json:"product_family_name,omitempty"`
    ProductName            *string             `json:"product_name,omitempty"`
    ProductPricePointName  *string             `json:"product_price_point_name,omitempty"`
    ProductPriceInCents    *int64              `json:"product_price_in_cents,omitempty"`
    OfferSignupPages       []OfferSignupPage   `json:"offer_signup_pages,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for Offer.
// It customizes the JSON marshaling process for Offer objects.
func (o *Offer) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the Offer object to a map representation for JSON marshaling.
func (o *Offer) toMap() map[string]any {
    structMap := make(map[string]any)
    if o.Id != nil {
        structMap["id"] = o.Id
    }
    if o.SiteId != nil {
        structMap["site_id"] = o.SiteId
    }
    if o.ProductFamilyId != nil {
        structMap["product_family_id"] = o.ProductFamilyId
    }
    if o.ProductId != nil {
        structMap["product_id"] = o.ProductId
    }
    if o.ProductPricePointId != nil {
        structMap["product_price_point_id"] = o.ProductPricePointId
    }
    if o.ProductRevisableNumber != nil {
        structMap["product_revisable_number"] = o.ProductRevisableNumber
    }
    if o.Name != nil {
        structMap["name"] = o.Name
    }
    if o.Handle != nil {
        structMap["handle"] = o.Handle
    }
    if o.Description.IsValueSet() {
        structMap["description"] = o.Description.Value()
    }
    if o.CreatedAt != nil {
        structMap["created_at"] = o.CreatedAt.Format(time.RFC3339)
    }
    if o.UpdatedAt != nil {
        structMap["updated_at"] = o.UpdatedAt.Format(time.RFC3339)
    }
    if o.ArchivedAt.IsValueSet() {
        var ArchivedAtVal *string = nil
        if o.ArchivedAt.Value() != nil {
            val := o.ArchivedAt.Value().Format(time.RFC3339)
            ArchivedAtVal = &val
        }
        structMap["archived_at"] = ArchivedAtVal
    }
    if o.OfferItems != nil {
        structMap["offer_items"] = o.OfferItems
    }
    if o.OfferDiscounts != nil {
        structMap["offer_discounts"] = o.OfferDiscounts
    }
    if o.ProductFamilyName != nil {
        structMap["product_family_name"] = o.ProductFamilyName
    }
    if o.ProductName != nil {
        structMap["product_name"] = o.ProductName
    }
    if o.ProductPricePointName != nil {
        structMap["product_price_point_name"] = o.ProductPricePointName
    }
    if o.ProductPriceInCents != nil {
        structMap["product_price_in_cents"] = o.ProductPriceInCents
    }
    if o.OfferSignupPages != nil {
        structMap["offer_signup_pages"] = o.OfferSignupPages
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Offer.
// It customizes the JSON unmarshaling process for Offer objects.
func (o *Offer) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id                     *int              `json:"id,omitempty"`
        SiteId                 *int              `json:"site_id,omitempty"`
        ProductFamilyId        *int              `json:"product_family_id,omitempty"`
        ProductId              *int              `json:"product_id,omitempty"`
        ProductPricePointId    *int              `json:"product_price_point_id,omitempty"`
        ProductRevisableNumber *int              `json:"product_revisable_number,omitempty"`
        Name                   *string           `json:"name,omitempty"`
        Handle                 *string           `json:"handle,omitempty"`
        Description            Optional[string]  `json:"description"`
        CreatedAt              *string           `json:"created_at,omitempty"`
        UpdatedAt              *string           `json:"updated_at,omitempty"`
        ArchivedAt             Optional[string]  `json:"archived_at"`
        OfferItems             []OfferItem       `json:"offer_items,omitempty"`
        OfferDiscounts         []OfferDiscount   `json:"offer_discounts,omitempty"`
        ProductFamilyName      *string           `json:"product_family_name,omitempty"`
        ProductName            *string           `json:"product_name,omitempty"`
        ProductPricePointName  *string           `json:"product_price_point_name,omitempty"`
        ProductPriceInCents    *int64            `json:"product_price_in_cents,omitempty"`
        OfferSignupPages       []OfferSignupPage `json:"offer_signup_pages,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    o.Id = temp.Id
    o.SiteId = temp.SiteId
    o.ProductFamilyId = temp.ProductFamilyId
    o.ProductId = temp.ProductId
    o.ProductPricePointId = temp.ProductPricePointId
    o.ProductRevisableNumber = temp.ProductRevisableNumber
    o.Name = temp.Name
    o.Handle = temp.Handle
    o.Description = temp.Description
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        o.CreatedAt = &CreatedAtVal
    }
    if temp.UpdatedAt != nil {
        UpdatedAtVal, err := time.Parse(time.RFC3339, *temp.UpdatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse updated_at as % s format.", time.RFC3339)
        }
        o.UpdatedAt = &UpdatedAtVal
    }
    o.ArchivedAt.ShouldSetValue(temp.ArchivedAt.IsValueSet())
    if temp.ArchivedAt.Value() != nil {
        ArchivedAtVal, err := time.Parse(time.RFC3339, (*temp.ArchivedAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse archived_at as % s format.", time.RFC3339)
        }
        o.ArchivedAt.SetValue(&ArchivedAtVal)
    }
    o.OfferItems = temp.OfferItems
    o.OfferDiscounts = temp.OfferDiscounts
    o.ProductFamilyName = temp.ProductFamilyName
    o.ProductName = temp.ProductName
    o.ProductPricePointName = temp.ProductPricePointName
    o.ProductPriceInCents = temp.ProductPriceInCents
    o.OfferSignupPages = temp.OfferSignupPages
    return nil
}
