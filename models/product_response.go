package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ProductResponse represents a ProductResponse struct.
type ProductResponse struct {
    Product              Product        `json:"product"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ProductResponse.
// It customizes the JSON marshaling process for ProductResponse objects.
func (p ProductResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the ProductResponse object to a map representation for JSON marshaling.
func (p ProductResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
    structMap["product"] = p.Product.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ProductResponse.
// It customizes the JSON unmarshaling process for ProductResponse objects.
func (p *ProductResponse) UnmarshalJSON(input []byte) error {
    var temp productResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "product")
    if err != nil {
    	return err
    }
    
    p.AdditionalProperties = additionalProperties
    p.Product = *temp.Product
    return nil
}

// TODO
type productResponse  struct {
    Product *Product `json:"product"`
}

func (p *productResponse) validate() error {
    var errs []string
    if p.Product == nil {
        errs = append(errs, "required field `product` is missing for type `Product Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
