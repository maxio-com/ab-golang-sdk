package models

import (
    "encoding/json"
)

// Seller represents a Seller struct.
type Seller struct {
    Name                 *string          `json:"name,omitempty"`
    Address              *InvoiceAddress  `json:"address,omitempty"`
    Phone                *string          `json:"phone,omitempty"`
    LogoUrl              Optional[string] `json:"logo_url"`
    AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Seller.
// It customizes the JSON marshaling process for Seller objects.
func (s Seller) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the Seller object to a map representation for JSON marshaling.
func (s Seller) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Name != nil {
        structMap["name"] = s.Name
    }
    if s.Address != nil {
        structMap["address"] = s.Address.toMap()
    }
    if s.Phone != nil {
        structMap["phone"] = s.Phone
    }
    if s.LogoUrl.IsValueSet() {
        if s.LogoUrl.Value() != nil {
            structMap["logo_url"] = s.LogoUrl.Value()
        } else {
            structMap["logo_url"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Seller.
// It customizes the JSON unmarshaling process for Seller objects.
func (s *Seller) UnmarshalJSON(input []byte) error {
    var temp seller
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "name", "address", "phone", "logo_url")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Name = temp.Name
    s.Address = temp.Address
    s.Phone = temp.Phone
    s.LogoUrl = temp.LogoUrl
    return nil
}

// TODO
type seller  struct {
    Name    *string          `json:"name,omitempty"`
    Address *InvoiceAddress  `json:"address,omitempty"`
    Phone   *string          `json:"phone,omitempty"`
    LogoUrl Optional[string] `json:"logo_url"`
}
