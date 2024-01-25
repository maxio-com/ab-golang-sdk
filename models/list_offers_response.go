package models

import (
	"encoding/json"
)

// ListOffersResponse represents a ListOffersResponse struct.
type ListOffersResponse struct {
	Offers []Offer `json:"offers,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ListOffersResponse.
// It customizes the JSON marshaling process for ListOffersResponse objects.
func (l *ListOffersResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

// toMap converts the ListOffersResponse object to a map representation for JSON marshaling.
func (l *ListOffersResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if l.Offers != nil {
		structMap["offers"] = l.Offers
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListOffersResponse.
// It customizes the JSON unmarshaling process for ListOffersResponse objects.
func (l *ListOffersResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Offers []Offer `json:"offers,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	l.Offers = temp.Offers
	return nil
}
