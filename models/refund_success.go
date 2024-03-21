package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// RefundSuccess represents a RefundSuccess struct.
type RefundSuccess struct {
	RefundId             int `json:"refund_id"`
	GatewayTransactionId int `json:"gateway_transaction_id"`
	ProductId            int `json:"product_id"`
}

// MarshalJSON implements the json.Marshaler interface for RefundSuccess.
// It customizes the JSON marshaling process for RefundSuccess objects.
func (r *RefundSuccess) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RefundSuccess object to a map representation for JSON marshaling.
func (r *RefundSuccess) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["refund_id"] = r.RefundId
	structMap["gateway_transaction_id"] = r.GatewayTransactionId
	structMap["product_id"] = r.ProductId
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RefundSuccess.
// It customizes the JSON unmarshaling process for RefundSuccess objects.
func (r *RefundSuccess) UnmarshalJSON(input []byte) error {
	var temp refundSuccess
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}

	r.RefundId = *temp.RefundId
	r.GatewayTransactionId = *temp.GatewayTransactionId
	r.ProductId = *temp.ProductId
	return nil
}

// TODO
type refundSuccess struct {
	RefundId             *int `json:"refund_id"`
	GatewayTransactionId *int `json:"gateway_transaction_id"`
	ProductId            *int `json:"product_id"`
}

func (r *refundSuccess) validate() error {
	var errs []string
	if r.RefundId == nil {
		errs = append(errs, "required field `refund_id` is missing for type `Refund Success`")
	}
	if r.GatewayTransactionId == nil {
		errs = append(errs, "required field `gateway_transaction_id` is missing for type `Refund Success`")
	}
	if r.ProductId == nil {
		errs = append(errs, "required field `product_id` is missing for type `Refund Success`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
