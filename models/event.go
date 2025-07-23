// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

// Event represents a Event struct.
type Event struct {
	Id             int64     `json:"id"`
	Key            EventKey  `json:"key"`
	Message        string    `json:"message"`
	SubscriptionId *int      `json:"subscription_id"`
	CustomerId     *int      `json:"customer_id"`
	CreatedAt      time.Time `json:"created_at"`
	// The schema varies based on the event key. The key-to-event data mapping is as follows:
	// * `subscription_product_change` - SubscriptionProductChange
	// * `subscription_state_change` - SubscriptionStateChange
	// * `signup_success`, `delayed_signup_creation_success`, `payment_success`, `payment_failure`, `renewal_success`, `renewal_failure`, `chargeback_lost`, `chargeback_accepted`, `chargeback_closed` - PaymentRelatedEvents
	// * `refund_success` - RefundSuccess
	// * `component_allocation_change` - ComponentAllocationChange
	// * `metered_usage` - MeteredUsage
	// * `prepaid_usage` - PrepaidUsage
	// * `dunning_step_reached` - DunningStepReached
	// * `invoice_issued` - InvoiceIssued
	// * `pending_cancellation_change` - PendingCancellationChange
	// * `prepaid_subscription_balance_changed` - PrepaidSubscriptionBalanceChanged
	// * `subscription_group_signup_success` and `subscription_group_signup_failure` - SubscriptionGroupSignupEventData
	// * `proforma_invoice_issued` - ProformaInvoiceIssued
	// * `subscription_prepayment_account_balance_changed` - PrepaymentAccountBalanceChanged
	// * `payment_collection_method_changed` - PaymentCollectionMethodChanged
	// * `subscription_service_credit_account_balance_changed` - CreditAccountBalanceChanged
	// * `item_price_point_changed` - ItemPricePointChanged
	// * `custom_field_value_change` - CustomFieldValueChange
	// * The rest, that is `delayed_signup_creation_failure`, `billing_date_change`, `expiration_date_change`, `expiring_card`,
	// `customer_update`, `customer_create`, `customer_delete`, `upgrade_downgrade_success`, `upgrade_downgrade_failure`,
	// `statement_closed`, `statement_settled`, `subscription_card_update`, `subscription_group_card_update`,
	// `subscription_bank_account_update`, `refund_failure`, `upcoming_renewal_notice`, `trial_end_notice`,
	// `direct_debit_payment_paid_out`, `direct_debit_payment_rejected`, `direct_debit_payment_pending`, `pending_payment_created`,
	// `pending_payment_failed`, `pending_payment_completed`,  don't have event_specific_data defined,
	// `renewal_success_recreated`, `renewal_failure_recreated`, `payment_success_recreated`, `payment_failure_recreated`,
	// `subscription_deletion`, `subscription_group_bank_account_update`, `subscription_paypal_account_update`, `subscription_group_paypal_account_update`,
	// `subscription_customer_change`, `account_transaction_changed`, `go_cardless_payment_paid_out`, `go_cardless_payment_rejected`,
	// `go_cardless_payment_pending`, `stripe_direct_debit_payment_paid_out`, `stripe_direct_debit_payment_rejected`, `stripe_direct_debit_payment_pending`,
	// `maxio_payments_direct_debit_payment_paid_out`, `maxio_payments_direct_debit_payment_rejected`, `maxio_payments_direct_debit_payment_pending`,
	// `invoice_in_collections_canceled`, `subscription_added_to_group`, `subscription_removed_from_group`, `chargeback_opened`, `chargeback_lost`,
	// `chargeback_accepted`, `chargeback_closed`, `chargeback_won`, `payment_collection_method_changed`, `component_billing_date_changed`,
	// `subscription_term_renewal_scheduled`, `subscription_term_renewal_pending`, `subscription_term_renewal_activated`, `subscription_term_renewal_removed`
	// they map to `null` instead.
	EventSpecificData    *EventEventSpecificData `json:"event_specific_data"`
	AdditionalProperties map[string]interface{}  `json:"_"`
}

// String implements the fmt.Stringer interface for Event,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (e Event) String() string {
	return fmt.Sprintf(
		"Event[Id=%v, Key=%v, Message=%v, SubscriptionId=%v, CustomerId=%v, CreatedAt=%v, EventSpecificData=%v, AdditionalProperties=%v]",
		e.Id, e.Key, e.Message, e.SubscriptionId, e.CustomerId, e.CreatedAt, e.EventSpecificData, e.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Event.
// It customizes the JSON marshaling process for Event objects.
func (e Event) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(e.AdditionalProperties,
		"id", "key", "message", "subscription_id", "customer_id", "created_at", "event_specific_data"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(e.toMap())
}

// toMap converts the Event object to a map representation for JSON marshaling.
func (e Event) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, e.AdditionalProperties)
	structMap["id"] = e.Id
	structMap["key"] = e.Key
	structMap["message"] = e.Message
	if e.SubscriptionId != nil {
		structMap["subscription_id"] = e.SubscriptionId
	} else {
		structMap["subscription_id"] = nil
	}
	if e.CustomerId != nil {
		structMap["customer_id"] = e.CustomerId
	} else {
		structMap["customer_id"] = nil
	}
	structMap["created_at"] = e.CreatedAt.Format(time.RFC3339)
	if e.EventSpecificData != nil {
		structMap["event_specific_data"] = e.EventSpecificData.toMap()
	} else {
		structMap["event_specific_data"] = nil
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Event.
// It customizes the JSON unmarshaling process for Event objects.
func (e *Event) UnmarshalJSON(input []byte) error {
	var temp tempEvent
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "key", "message", "subscription_id", "customer_id", "created_at", "event_specific_data")
	if err != nil {
		return err
	}
	e.AdditionalProperties = additionalProperties

	e.Id = *temp.Id
	e.Key = *temp.Key
	e.Message = *temp.Message
	e.SubscriptionId = temp.SubscriptionId
	e.CustomerId = temp.CustomerId
	CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
	if err != nil {
		log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
	}
	e.CreatedAt = CreatedAtVal
	e.EventSpecificData = temp.EventSpecificData
	return nil
}

// tempEvent is a temporary struct used for validating the fields of Event.
type tempEvent struct {
	Id                *int64                  `json:"id"`
	Key               *EventKey               `json:"key"`
	Message           *string                 `json:"message"`
	SubscriptionId    *int                    `json:"subscription_id"`
	CustomerId        *int                    `json:"customer_id"`
	CreatedAt         *string                 `json:"created_at"`
	EventSpecificData *EventEventSpecificData `json:"event_specific_data"`
}

func (e *tempEvent) validate() error {
	var errs []string
	if e.Id == nil {
		errs = append(errs, "required field `id` is missing for type `Event`")
	}
	if e.Key == nil {
		errs = append(errs, "required field `key` is missing for type `Event`")
	}
	if e.Message == nil {
		errs = append(errs, "required field `message` is missing for type `Event`")
	}
	if e.CreatedAt == nil {
		errs = append(errs, "required field `created_at` is missing for type `Event`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
