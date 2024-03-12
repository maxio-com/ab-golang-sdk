package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// CreateOrUpdateProductRequest represents a CreateOrUpdateProductRequest struct.
type CreateOrUpdateProductRequest struct {
	Product CreateOrUpdateProduct `json:"product"`
}

// MarshalJSON implements the json.Marshaler interface for CreateOrUpdateProductRequest.
// It customizes the JSON marshaling process for CreateOrUpdateProductRequest objects.
func (c *CreateOrUpdateProductRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateOrUpdateProductRequest object to a map representation for JSON marshaling.
func (c *CreateOrUpdateProductRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["product"] = c.Product.toMap()
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateOrUpdateProductRequest.
// It customizes the JSON unmarshaling process for CreateOrUpdateProductRequest objects.
func (c *CreateOrUpdateProductRequest) UnmarshalJSON(input []byte) error {
	var temp createOrUpdateProductRequest
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}

	c.Product = *temp.Product
	return nil
}

// TODO
type createOrUpdateProductRequest struct {
	Product *CreateOrUpdateProduct `json:"product"`
}

func (c *createOrUpdateProductRequest) validate() error {
	var errs []string
	if c.Product == nil {
		errs = append(errs, "required field `product` is missing for type `Create or Update Product Request`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
