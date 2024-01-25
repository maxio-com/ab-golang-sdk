package models

import (
	"encoding/json"
)

// OfferResponse represents a OfferResponse struct.
type OfferResponse struct {
	Offer *Offer `json:"offer,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for OfferResponse.
// It customizes the JSON marshaling process for OfferResponse objects.
func (o *OfferResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(o.toMap())
}

// toMap converts the OfferResponse object to a map representation for JSON marshaling.
func (o *OfferResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if o.Offer != nil {
		structMap["offer"] = o.Offer
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OfferResponse.
// It customizes the JSON unmarshaling process for OfferResponse objects.
func (o *OfferResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Offer *Offer `json:"offer,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	o.Offer = temp.Offer
	return nil
}
