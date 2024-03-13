package models

import (
	"encoding/json"
)

// InvoiceSeller represents a InvoiceSeller struct.
// Information about the seller (merchant) listed on the masthead of the invoice.
type InvoiceSeller struct {
	Name    *string          `json:"name,omitempty"`
	Address *InvoiceAddress  `json:"address,omitempty"`
	Phone   *string          `json:"phone,omitempty"`
	LogoUrl Optional[string] `json:"logo_url"`
}

// MarshalJSON implements the json.Marshaler interface for InvoiceSeller.
// It customizes the JSON marshaling process for InvoiceSeller objects.
func (i *InvoiceSeller) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(i.toMap())
}

// toMap converts the InvoiceSeller object to a map representation for JSON marshaling.
func (i *InvoiceSeller) toMap() map[string]any {
	structMap := make(map[string]any)
	if i.Name != nil {
		structMap["name"] = i.Name
	}
	if i.Address != nil {
		structMap["address"] = i.Address.toMap()
	}
	if i.Phone != nil {
		structMap["phone"] = i.Phone
	}
	if i.LogoUrl.IsValueSet() {
		if i.LogoUrl.Value() != nil {
			structMap["logo_url"] = i.LogoUrl.Value()
		} else {
			structMap["logo_url"] = nil
		}
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceSeller.
// It customizes the JSON unmarshaling process for InvoiceSeller objects.
func (i *InvoiceSeller) UnmarshalJSON(input []byte) error {
	var temp invoiceSeller
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	i.Name = temp.Name
	i.Address = temp.Address
	i.Phone = temp.Phone
	i.LogoUrl = temp.LogoUrl
	return nil
}

// TODO
type invoiceSeller struct {
	Name    *string          `json:"name,omitempty"`
	Address *InvoiceAddress  `json:"address,omitempty"`
	Phone   *string          `json:"phone,omitempty"`
	LogoUrl Optional[string] `json:"logo_url"`
}
