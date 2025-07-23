// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"fmt"
)

// SubscriptionGroupSubscriptionError represents a SubscriptionGroupSubscriptionError struct.
// Object which contains subscription errors.
type SubscriptionGroupSubscriptionError struct {
	Product                       []string               `json:"product,omitempty"`
	ProductPricePointId           []string               `json:"product_price_point_id,omitempty"`
	PaymentProfile                []string               `json:"payment_profile,omitempty"`
	PaymentProfileChargifyToken   []string               `json:"payment_profile.chargify_token,omitempty"`
	Base                          []string               `json:"base,omitempty"`
	PaymentProfileExpirationMonth []string               `json:"payment_profile.expiration_month,omitempty"`
	PaymentProfileExpirationYear  []string               `json:"payment_profile.expiration_year,omitempty"`
	PaymentProfileFullNumber      []string               `json:"payment_profile.full_number,omitempty"`
	AdditionalProperties          map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SubscriptionGroupSubscriptionError,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionGroupSubscriptionError) String() string {
	return fmt.Sprintf(
		"SubscriptionGroupSubscriptionError[Product=%v, ProductPricePointId=%v, PaymentProfile=%v, PaymentProfileChargifyToken=%v, Base=%v, PaymentProfileExpirationMonth=%v, PaymentProfileExpirationYear=%v, PaymentProfileFullNumber=%v, AdditionalProperties=%v]",
		s.Product, s.ProductPricePointId, s.PaymentProfile, s.PaymentProfileChargifyToken, s.Base, s.PaymentProfileExpirationMonth, s.PaymentProfileExpirationYear, s.PaymentProfileFullNumber, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupSubscriptionError.
// It customizes the JSON marshaling process for SubscriptionGroupSubscriptionError objects.
func (s SubscriptionGroupSubscriptionError) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"product", "product_price_point_id", "payment_profile", "payment_profile.chargify_token", "base", "payment_profile.expiration_month", "payment_profile.expiration_year", "payment_profile.full_number"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupSubscriptionError object to a map representation for JSON marshaling.
func (s SubscriptionGroupSubscriptionError) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Product != nil {
		structMap["product"] = s.Product
	}
	if s.ProductPricePointId != nil {
		structMap["product_price_point_id"] = s.ProductPricePointId
	}
	if s.PaymentProfile != nil {
		structMap["payment_profile"] = s.PaymentProfile
	}
	if s.PaymentProfileChargifyToken != nil {
		structMap["payment_profile.chargify_token"] = s.PaymentProfileChargifyToken
	}
	if s.Base != nil {
		structMap["base"] = s.Base
	}
	if s.PaymentProfileExpirationMonth != nil {
		structMap["payment_profile.expiration_month"] = s.PaymentProfileExpirationMonth
	}
	if s.PaymentProfileExpirationYear != nil {
		structMap["payment_profile.expiration_year"] = s.PaymentProfileExpirationYear
	}
	if s.PaymentProfileFullNumber != nil {
		structMap["payment_profile.full_number"] = s.PaymentProfileFullNumber
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupSubscriptionError.
// It customizes the JSON unmarshaling process for SubscriptionGroupSubscriptionError objects.
func (s *SubscriptionGroupSubscriptionError) UnmarshalJSON(input []byte) error {
	var temp tempSubscriptionGroupSubscriptionError
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "product", "product_price_point_id", "payment_profile", "payment_profile.chargify_token", "base", "payment_profile.expiration_month", "payment_profile.expiration_year", "payment_profile.full_number")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Product = temp.Product
	s.ProductPricePointId = temp.ProductPricePointId
	s.PaymentProfile = temp.PaymentProfile
	s.PaymentProfileChargifyToken = temp.PaymentProfileChargifyToken
	s.Base = temp.Base
	s.PaymentProfileExpirationMonth = temp.PaymentProfileExpirationMonth
	s.PaymentProfileExpirationYear = temp.PaymentProfileExpirationYear
	s.PaymentProfileFullNumber = temp.PaymentProfileFullNumber
	return nil
}

// tempSubscriptionGroupSubscriptionError is a temporary struct used for validating the fields of SubscriptionGroupSubscriptionError.
type tempSubscriptionGroupSubscriptionError struct {
	Product                       []string `json:"product,omitempty"`
	ProductPricePointId           []string `json:"product_price_point_id,omitempty"`
	PaymentProfile                []string `json:"payment_profile,omitempty"`
	PaymentProfileChargifyToken   []string `json:"payment_profile.chargify_token,omitempty"`
	Base                          []string `json:"base,omitempty"`
	PaymentProfileExpirationMonth []string `json:"payment_profile.expiration_month,omitempty"`
	PaymentProfileExpirationYear  []string `json:"payment_profile.expiration_year,omitempty"`
	PaymentProfileFullNumber      []string `json:"payment_profile.full_number,omitempty"`
}
