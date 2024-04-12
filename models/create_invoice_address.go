package models

import (
    "encoding/json"
)

// CreateInvoiceAddress represents a CreateInvoiceAddress struct.
// Overrides the default address.
type CreateInvoiceAddress struct {
    FirstName            *string        `json:"first_name,omitempty"`
    LastName             *string        `json:"last_name,omitempty"`
    Phone                *string        `json:"phone,omitempty"`
    Address              *string        `json:"address,omitempty"`
    Address2             *string        `json:"address_2,omitempty"`
    City                 *string        `json:"city,omitempty"`
    State                *string        `json:"state,omitempty"`
    Zip                  *string        `json:"zip,omitempty"`
    Country              *string        `json:"country,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreateInvoiceAddress.
// It customizes the JSON marshaling process for CreateInvoiceAddress objects.
func (c CreateInvoiceAddress) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateInvoiceAddress object to a map representation for JSON marshaling.
func (c CreateInvoiceAddress) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.FirstName != nil {
        structMap["first_name"] = c.FirstName
    }
    if c.LastName != nil {
        structMap["last_name"] = c.LastName
    }
    if c.Phone != nil {
        structMap["phone"] = c.Phone
    }
    if c.Address != nil {
        structMap["address"] = c.Address
    }
    if c.Address2 != nil {
        structMap["address_2"] = c.Address2
    }
    if c.City != nil {
        structMap["city"] = c.City
    }
    if c.State != nil {
        structMap["state"] = c.State
    }
    if c.Zip != nil {
        structMap["zip"] = c.Zip
    }
    if c.Country != nil {
        structMap["country"] = c.Country
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateInvoiceAddress.
// It customizes the JSON unmarshaling process for CreateInvoiceAddress objects.
func (c *CreateInvoiceAddress) UnmarshalJSON(input []byte) error {
    var temp createInvoiceAddress
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "first_name", "last_name", "phone", "address", "address_2", "city", "state", "zip", "country")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.FirstName = temp.FirstName
    c.LastName = temp.LastName
    c.Phone = temp.Phone
    c.Address = temp.Address
    c.Address2 = temp.Address2
    c.City = temp.City
    c.State = temp.State
    c.Zip = temp.Zip
    c.Country = temp.Country
    return nil
}

// TODO
type createInvoiceAddress  struct {
    FirstName *string `json:"first_name,omitempty"`
    LastName  *string `json:"last_name,omitempty"`
    Phone     *string `json:"phone,omitempty"`
    Address   *string `json:"address,omitempty"`
    Address2  *string `json:"address_2,omitempty"`
    City      *string `json:"city,omitempty"`
    State     *string `json:"state,omitempty"`
    Zip       *string `json:"zip,omitempty"`
    Country   *string `json:"country,omitempty"`
}
