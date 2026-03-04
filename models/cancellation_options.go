// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// CancellationOptions represents a CancellationOptions struct.
type CancellationOptions struct {
    // An indication as to why the subscription is being canceled. For your internal use.
    CancellationMessage            *string                `json:"cancellation_message,omitempty"`
    // The reason code associated with the cancellation. Use the [List Reason Codes]($e/Reason%20Codes/listReasonCodes) endpoint to retrieve the reason codes associated with your site.
    ReasonCode                     *string                `json:"reason_code,omitempty"`
    // When true, the subscription is cancelled at the current period end instead of immediately. To use this option, the Schedule Subscription Cancellation feature must be enabled on your site.
    CancelAtEndOfPeriod            *bool                  `json:"cancel_at_end_of_period,omitempty"`
    // Schedules the cancellation on the provided date. This is option is not applicable for prepaid subscriptions. To use this option, the Schedule Subscription Cancellation feature must be enabled on your site.
    ScheduledCancellationAt        Optional[time.Time]    `json:"scheduled_cancellation_at"`
    // Applies to prepaid subscriptions. When true, which is the default, the remaining prepaid balance is refunded as part of cancellation processing. When false, prepaid balance is not refunded as part of cancellation processing. To use this option, the Schedule Subscription Cancellation feature must be enabled on your site.
    RefundPrepaymentAccountBalance *bool                  `json:"refund_prepayment_account_balance,omitempty"`
    AdditionalProperties           map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CancellationOptions,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CancellationOptions) String() string {
    return fmt.Sprintf(
    	"CancellationOptions[CancellationMessage=%v, ReasonCode=%v, CancelAtEndOfPeriod=%v, ScheduledCancellationAt=%v, RefundPrepaymentAccountBalance=%v, AdditionalProperties=%v]",
    	c.CancellationMessage, c.ReasonCode, c.CancelAtEndOfPeriod, c.ScheduledCancellationAt, c.RefundPrepaymentAccountBalance, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CancellationOptions.
// It customizes the JSON marshaling process for CancellationOptions objects.
func (c CancellationOptions) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "cancellation_message", "reason_code", "cancel_at_end_of_period", "scheduled_cancellation_at", "refund_prepayment_account_balance"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CancellationOptions object to a map representation for JSON marshaling.
func (c CancellationOptions) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.CancellationMessage != nil {
        structMap["cancellation_message"] = c.CancellationMessage
    }
    if c.ReasonCode != nil {
        structMap["reason_code"] = c.ReasonCode
    }
    if c.CancelAtEndOfPeriod != nil {
        structMap["cancel_at_end_of_period"] = c.CancelAtEndOfPeriod
    }
    if c.ScheduledCancellationAt.IsValueSet() {
        var ScheduledCancellationAtVal *string = nil
        if c.ScheduledCancellationAt.Value() != nil {
            val := c.ScheduledCancellationAt.Value().Format(time.RFC3339)
            ScheduledCancellationAtVal = &val
        }
        if c.ScheduledCancellationAt.Value() != nil {
            structMap["scheduled_cancellation_at"] = ScheduledCancellationAtVal
        } else {
            structMap["scheduled_cancellation_at"] = nil
        }
    }
    if c.RefundPrepaymentAccountBalance != nil {
        structMap["refund_prepayment_account_balance"] = c.RefundPrepaymentAccountBalance
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CancellationOptions.
// It customizes the JSON unmarshaling process for CancellationOptions objects.
func (c *CancellationOptions) UnmarshalJSON(input []byte) error {
    var temp tempCancellationOptions
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "cancellation_message", "reason_code", "cancel_at_end_of_period", "scheduled_cancellation_at", "refund_prepayment_account_balance")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.CancellationMessage = temp.CancellationMessage
    c.ReasonCode = temp.ReasonCode
    c.CancelAtEndOfPeriod = temp.CancelAtEndOfPeriod
    c.ScheduledCancellationAt.ShouldSetValue(temp.ScheduledCancellationAt.IsValueSet())
    if temp.ScheduledCancellationAt.Value() != nil {
        ScheduledCancellationAtVal, err := time.Parse(time.RFC3339, (*temp.ScheduledCancellationAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse scheduled_cancellation_at as % s format.", time.RFC3339)
        }
        c.ScheduledCancellationAt.SetValue(&ScheduledCancellationAtVal)
    }
    c.RefundPrepaymentAccountBalance = temp.RefundPrepaymentAccountBalance
    return nil
}

// tempCancellationOptions is a temporary struct used for validating the fields of CancellationOptions.
type tempCancellationOptions  struct {
    CancellationMessage            *string          `json:"cancellation_message,omitempty"`
    ReasonCode                     *string          `json:"reason_code,omitempty"`
    CancelAtEndOfPeriod            *bool            `json:"cancel_at_end_of_period,omitempty"`
    ScheduledCancellationAt        Optional[string] `json:"scheduled_cancellation_at"`
    RefundPrepaymentAccountBalance *bool            `json:"refund_prepayment_account_balance,omitempty"`
}
