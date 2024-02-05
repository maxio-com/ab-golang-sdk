package models

import (
    "encoding/json"
)

// CreateOfferRequest represents a CreateOfferRequest struct.
type CreateOfferRequest struct {
    Offer CreateOffer `json:"offer"`
}

// MarshalJSON implements the json.Marshaler interface for CreateOfferRequest.
// It customizes the JSON marshaling process for CreateOfferRequest objects.
func (c *CreateOfferRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateOfferRequest object to a map representation for JSON marshaling.
func (c *CreateOfferRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["offer"] = c.Offer.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateOfferRequest.
// It customizes the JSON unmarshaling process for CreateOfferRequest objects.
func (c *CreateOfferRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Offer CreateOffer `json:"offer"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.Offer = temp.Offer
    return nil
}
