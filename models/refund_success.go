/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// RefundSuccess represents a RefundSuccess struct.
type RefundSuccess struct {
    RefundId             int                    `json:"refund_id"`
    GatewayTransactionId int                    `json:"gateway_transaction_id"`
    ProductId            int                    `json:"product_id"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for RefundSuccess,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RefundSuccess) String() string {
    return fmt.Sprintf(
    	"RefundSuccess[RefundId=%v, GatewayTransactionId=%v, ProductId=%v, AdditionalProperties=%v]",
    	r.RefundId, r.GatewayTransactionId, r.ProductId, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for RefundSuccess.
// It customizes the JSON marshaling process for RefundSuccess objects.
func (r RefundSuccess) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "refund_id", "gateway_transaction_id", "product_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the RefundSuccess object to a map representation for JSON marshaling.
func (r RefundSuccess) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["refund_id"] = r.RefundId
    structMap["gateway_transaction_id"] = r.GatewayTransactionId
    structMap["product_id"] = r.ProductId
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RefundSuccess.
// It customizes the JSON unmarshaling process for RefundSuccess objects.
func (r *RefundSuccess) UnmarshalJSON(input []byte) error {
    var temp tempRefundSuccess
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "refund_id", "gateway_transaction_id", "product_id")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.RefundId = *temp.RefundId
    r.GatewayTransactionId = *temp.GatewayTransactionId
    r.ProductId = *temp.ProductId
    return nil
}

// tempRefundSuccess is a temporary struct used for validating the fields of RefundSuccess.
type tempRefundSuccess  struct {
    RefundId             *int `json:"refund_id"`
    GatewayTransactionId *int `json:"gateway_transaction_id"`
    ProductId            *int `json:"product_id"`
}

func (r *tempRefundSuccess) validate() error {
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
    return errors.New(strings.Join (errs, "\n"))
}
