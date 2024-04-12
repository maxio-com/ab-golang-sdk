package models

import (
    "encoding/json"
)

// SubscriptionGroupPaymentProfile represents a SubscriptionGroupPaymentProfile struct.
type SubscriptionGroupPaymentProfile struct {
    Id                   *int           `json:"id,omitempty"`
    FirstName            *string        `json:"first_name,omitempty"`
    LastName             *string        `json:"last_name,omitempty"`
    MaskedCardNumber     *string        `json:"masked_card_number,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupPaymentProfile.
// It customizes the JSON marshaling process for SubscriptionGroupPaymentProfile objects.
func (s SubscriptionGroupPaymentProfile) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupPaymentProfile object to a map representation for JSON marshaling.
func (s SubscriptionGroupPaymentProfile) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Id != nil {
        structMap["id"] = s.Id
    }
    if s.FirstName != nil {
        structMap["first_name"] = s.FirstName
    }
    if s.LastName != nil {
        structMap["last_name"] = s.LastName
    }
    if s.MaskedCardNumber != nil {
        structMap["masked_card_number"] = s.MaskedCardNumber
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupPaymentProfile.
// It customizes the JSON unmarshaling process for SubscriptionGroupPaymentProfile objects.
func (s *SubscriptionGroupPaymentProfile) UnmarshalJSON(input []byte) error {
    var temp subscriptionGroupPaymentProfile
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "id", "first_name", "last_name", "masked_card_number")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Id = temp.Id
    s.FirstName = temp.FirstName
    s.LastName = temp.LastName
    s.MaskedCardNumber = temp.MaskedCardNumber
    return nil
}

// TODO
type subscriptionGroupPaymentProfile  struct {
    Id               *int    `json:"id,omitempty"`
    FirstName        *string `json:"first_name,omitempty"`
    LastName         *string `json:"last_name,omitempty"`
    MaskedCardNumber *string `json:"masked_card_number,omitempty"`
}
