package models

import (
	"encoding/json"
)

// SubscriptionGroupSubscriptionError represents a SubscriptionGroupSubscriptionError struct.
// Object which contains subscription errors.
type SubscriptionGroupSubscriptionError struct {
	Product                     []string `json:"product,omitempty"`
	ProductPricePointId         []string `json:"product_price_point_id,omitempty"`
	PaymentProfile              []string `json:"payment_profile,omitempty"`
	PaymentProfileChargifyToken []string `json:"payment_profile.chargify_token,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupSubscriptionError.
// It customizes the JSON marshaling process for SubscriptionGroupSubscriptionError objects.
func (s *SubscriptionGroupSubscriptionError) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupSubscriptionError object to a map representation for JSON marshaling.
func (s *SubscriptionGroupSubscriptionError) toMap() map[string]any {
	structMap := make(map[string]any)
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
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupSubscriptionError.
// It customizes the JSON unmarshaling process for SubscriptionGroupSubscriptionError objects.
func (s *SubscriptionGroupSubscriptionError) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Product                     []string `json:"product,omitempty"`
		ProductPricePointId         []string `json:"product_price_point_id,omitempty"`
		PaymentProfile              []string `json:"payment_profile,omitempty"`
		PaymentProfileChargifyToken []string `json:"payment_profile.chargify_token,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	s.Product = temp.Product
	s.ProductPricePointId = temp.ProductPricePointId
	s.PaymentProfile = temp.PaymentProfile
	s.PaymentProfileChargifyToken = temp.PaymentProfileChargifyToken
	return nil
}
