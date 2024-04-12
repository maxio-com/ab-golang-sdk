package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CreatePrepayment represents a CreatePrepayment struct.
type CreatePrepayment struct {
    Amount               float64                `json:"amount"`
    Details              string                 `json:"details"`
    Memo                 string                 `json:"memo"`
    // :- When the `method` specified is `"credit_card_on_file"`, the prepayment amount will be collected using the default credit card payment profile and applied to the prepayment account balance. This is especially useful for manual replenishment of prepaid subscriptions.
    Method               CreatePrepaymentMethod `json:"method"`
    PaymentProfileId     *int                   `json:"payment_profile_id,omitempty"`
    AdditionalProperties map[string]any         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreatePrepayment.
// It customizes the JSON marshaling process for CreatePrepayment objects.
func (c CreatePrepayment) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreatePrepayment object to a map representation for JSON marshaling.
func (c CreatePrepayment) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["amount"] = c.Amount
    structMap["details"] = c.Details
    structMap["memo"] = c.Memo
    structMap["method"] = c.Method
    if c.PaymentProfileId != nil {
        structMap["payment_profile_id"] = c.PaymentProfileId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreatePrepayment.
// It customizes the JSON unmarshaling process for CreatePrepayment objects.
func (c *CreatePrepayment) UnmarshalJSON(input []byte) error {
    var temp createPrepayment
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "amount", "details", "memo", "method", "payment_profile_id")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Amount = *temp.Amount
    c.Details = *temp.Details
    c.Memo = *temp.Memo
    c.Method = *temp.Method
    c.PaymentProfileId = temp.PaymentProfileId
    return nil
}

// TODO
type createPrepayment  struct {
    Amount           *float64                `json:"amount"`
    Details          *string                 `json:"details"`
    Memo             *string                 `json:"memo"`
    Method           *CreatePrepaymentMethod `json:"method"`
    PaymentProfileId *int                    `json:"payment_profile_id,omitempty"`
}

func (c *createPrepayment) validate() error {
    var errs []string
    if c.Amount == nil {
        errs = append(errs, "required field `amount` is missing for type `Create Prepayment`")
    }
    if c.Details == nil {
        errs = append(errs, "required field `details` is missing for type `Create Prepayment`")
    }
    if c.Memo == nil {
        errs = append(errs, "required field `memo` is missing for type `Create Prepayment`")
    }
    if c.Method == nil {
        errs = append(errs, "required field `method` is missing for type `Create Prepayment`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
