/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// EventEventSpecificData represents a EventEventSpecificData struct.
// This is a container for one-of cases.
type EventEventSpecificData struct {
	value                               any
	isSubscriptionProductChange         bool
	isSubscriptionStateChange           bool
	isPaymentRelatedEvents              bool
	isRefundSuccess                     bool
	isComponentAllocationChange         bool
	isMeteredUsage                      bool
	isPrepaidUsage                      bool
	isDunningStepReached                bool
	isInvoiceIssued                     bool
	isPendingCancellationChange         bool
	isPrepaidSubscriptionBalanceChanged bool
	isProformaInvoiceIssued             bool
	isSubscriptionGroupSignupSuccess    bool
	isSubscriptionGroupSignupFailure    bool
	isCreditAccountBalanceChanged       bool
	isPrepaymentAccountBalanceChanged   bool
	isPaymentCollectionMethodChanged    bool
	isItemPricePointChanged             bool
	isCustomFieldValueChange            bool
}

// String converts the EventEventSpecificData object to a string representation.
func (e EventEventSpecificData) String() string {
	if bytes, err := json.Marshal(e.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for EventEventSpecificData.
// It customizes the JSON marshaling process for EventEventSpecificData objects.
func (e *EventEventSpecificData) MarshalJSON() (
	[]byte,
	error) {
	if e.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.EventEventSpecificDataContainer.From*` functions to initialize the EventEventSpecificData object.")
	}
	return json.Marshal(e.toMap())
}

// toMap converts the EventEventSpecificData object to a map representation for JSON marshaling.
func (e *EventEventSpecificData) toMap() any {
	switch obj := e.value.(type) {
	case *SubscriptionProductChange:
		return obj.toMap()
	case *SubscriptionStateChange:
		return obj.toMap()
	case *PaymentRelatedEvents:
		return obj.toMap()
	case *RefundSuccess:
		return obj.toMap()
	case *ComponentAllocationChange:
		return obj.toMap()
	case *MeteredUsage:
		return obj.toMap()
	case *PrepaidUsage:
		return obj.toMap()
	case *DunningStepReached:
		return obj.toMap()
	case *InvoiceIssued:
		return obj.toMap()
	case *PendingCancellationChange:
		return obj.toMap()
	case *PrepaidSubscriptionBalanceChanged:
		return obj.toMap()
	case *ProformaInvoiceIssued:
		return obj.toMap()
	case *SubscriptionGroupSignupSuccess:
		return obj.toMap()
	case *SubscriptionGroupSignupFailure:
		return obj.toMap()
	case *CreditAccountBalanceChanged:
		return obj.toMap()
	case *PrepaymentAccountBalanceChanged:
		return obj.toMap()
	case *PaymentCollectionMethodChanged:
		return obj.toMap()
	case *ItemPricePointChanged:
		return obj.toMap()
	case *CustomFieldValueChange:
		return obj.toMap()
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for EventEventSpecificData.
// It customizes the JSON unmarshaling process for EventEventSpecificData objects.
func (e *EventEventSpecificData) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(&SubscriptionProductChange{}, false, &e.isSubscriptionProductChange),
		NewTypeHolder(&SubscriptionStateChange{}, false, &e.isSubscriptionStateChange),
		NewTypeHolder(&PaymentRelatedEvents{}, false, &e.isPaymentRelatedEvents),
		NewTypeHolder(&RefundSuccess{}, false, &e.isRefundSuccess),
		NewTypeHolder(&ComponentAllocationChange{}, false, &e.isComponentAllocationChange),
		NewTypeHolder(&MeteredUsage{}, false, &e.isMeteredUsage),
		NewTypeHolder(&PrepaidUsage{}, false, &e.isPrepaidUsage),
		NewTypeHolder(&DunningStepReached{}, false, &e.isDunningStepReached),
		NewTypeHolder(&InvoiceIssued{}, false, &e.isInvoiceIssued),
		NewTypeHolder(&PendingCancellationChange{}, false, &e.isPendingCancellationChange),
		NewTypeHolder(&PrepaidSubscriptionBalanceChanged{}, false, &e.isPrepaidSubscriptionBalanceChanged),
		NewTypeHolder(&ProformaInvoiceIssued{}, false, &e.isProformaInvoiceIssued),
		NewTypeHolder(&SubscriptionGroupSignupSuccess{}, false, &e.isSubscriptionGroupSignupSuccess),
		NewTypeHolder(&SubscriptionGroupSignupFailure{}, false, &e.isSubscriptionGroupSignupFailure),
		NewTypeHolder(&CreditAccountBalanceChanged{}, false, &e.isCreditAccountBalanceChanged),
		NewTypeHolder(&PrepaymentAccountBalanceChanged{}, false, &e.isPrepaymentAccountBalanceChanged),
		NewTypeHolder(&PaymentCollectionMethodChanged{}, false, &e.isPaymentCollectionMethodChanged),
		NewTypeHolder(&ItemPricePointChanged{}, false, &e.isItemPricePointChanged),
		NewTypeHolder(&CustomFieldValueChange{}, false, &e.isCustomFieldValueChange),
	)

	e.value = result
	return err
}

func (e *EventEventSpecificData) AsSubscriptionProductChange() (
	*SubscriptionProductChange,
	bool) {
	if !e.isSubscriptionProductChange {
		return nil, false
	}
	return e.value.(*SubscriptionProductChange), true
}

func (e *EventEventSpecificData) AsSubscriptionStateChange() (
	*SubscriptionStateChange,
	bool) {
	if !e.isSubscriptionStateChange {
		return nil, false
	}
	return e.value.(*SubscriptionStateChange), true
}

func (e *EventEventSpecificData) AsPaymentRelatedEvents() (
	*PaymentRelatedEvents,
	bool) {
	if !e.isPaymentRelatedEvents {
		return nil, false
	}
	return e.value.(*PaymentRelatedEvents), true
}

func (e *EventEventSpecificData) AsRefundSuccess() (
	*RefundSuccess,
	bool) {
	if !e.isRefundSuccess {
		return nil, false
	}
	return e.value.(*RefundSuccess), true
}

func (e *EventEventSpecificData) AsComponentAllocationChange() (
	*ComponentAllocationChange,
	bool) {
	if !e.isComponentAllocationChange {
		return nil, false
	}
	return e.value.(*ComponentAllocationChange), true
}

func (e *EventEventSpecificData) AsMeteredUsage() (
	*MeteredUsage,
	bool) {
	if !e.isMeteredUsage {
		return nil, false
	}
	return e.value.(*MeteredUsage), true
}

func (e *EventEventSpecificData) AsPrepaidUsage() (
	*PrepaidUsage,
	bool) {
	if !e.isPrepaidUsage {
		return nil, false
	}
	return e.value.(*PrepaidUsage), true
}

func (e *EventEventSpecificData) AsDunningStepReached() (
	*DunningStepReached,
	bool) {
	if !e.isDunningStepReached {
		return nil, false
	}
	return e.value.(*DunningStepReached), true
}

func (e *EventEventSpecificData) AsInvoiceIssued() (
	*InvoiceIssued,
	bool) {
	if !e.isInvoiceIssued {
		return nil, false
	}
	return e.value.(*InvoiceIssued), true
}

func (e *EventEventSpecificData) AsPendingCancellationChange() (
	*PendingCancellationChange,
	bool) {
	if !e.isPendingCancellationChange {
		return nil, false
	}
	return e.value.(*PendingCancellationChange), true
}

func (e *EventEventSpecificData) AsPrepaidSubscriptionBalanceChanged() (
	*PrepaidSubscriptionBalanceChanged,
	bool) {
	if !e.isPrepaidSubscriptionBalanceChanged {
		return nil, false
	}
	return e.value.(*PrepaidSubscriptionBalanceChanged), true
}

func (e *EventEventSpecificData) AsProformaInvoiceIssued() (
	*ProformaInvoiceIssued,
	bool) {
	if !e.isProformaInvoiceIssued {
		return nil, false
	}
	return e.value.(*ProformaInvoiceIssued), true
}

func (e *EventEventSpecificData) AsSubscriptionGroupSignupSuccess() (
	*SubscriptionGroupSignupSuccess,
	bool) {
	if !e.isSubscriptionGroupSignupSuccess {
		return nil, false
	}
	return e.value.(*SubscriptionGroupSignupSuccess), true
}

func (e *EventEventSpecificData) AsSubscriptionGroupSignupFailure() (
	*SubscriptionGroupSignupFailure,
	bool) {
	if !e.isSubscriptionGroupSignupFailure {
		return nil, false
	}
	return e.value.(*SubscriptionGroupSignupFailure), true
}

func (e *EventEventSpecificData) AsCreditAccountBalanceChanged() (
	*CreditAccountBalanceChanged,
	bool) {
	if !e.isCreditAccountBalanceChanged {
		return nil, false
	}
	return e.value.(*CreditAccountBalanceChanged), true
}

func (e *EventEventSpecificData) AsPrepaymentAccountBalanceChanged() (
	*PrepaymentAccountBalanceChanged,
	bool) {
	if !e.isPrepaymentAccountBalanceChanged {
		return nil, false
	}
	return e.value.(*PrepaymentAccountBalanceChanged), true
}

func (e *EventEventSpecificData) AsPaymentCollectionMethodChanged() (
	*PaymentCollectionMethodChanged,
	bool) {
	if !e.isPaymentCollectionMethodChanged {
		return nil, false
	}
	return e.value.(*PaymentCollectionMethodChanged), true
}

func (e *EventEventSpecificData) AsItemPricePointChanged() (
	*ItemPricePointChanged,
	bool) {
	if !e.isItemPricePointChanged {
		return nil, false
	}
	return e.value.(*ItemPricePointChanged), true
}

func (e *EventEventSpecificData) AsCustomFieldValueChange() (
	*CustomFieldValueChange,
	bool) {
	if !e.isCustomFieldValueChange {
		return nil, false
	}
	return e.value.(*CustomFieldValueChange), true
}

// internalEventEventSpecificData represents a eventEventSpecificData struct.
// This is a container for one-of cases.
type internalEventEventSpecificData struct{}

var EventEventSpecificDataContainer internalEventEventSpecificData

// The internalEventEventSpecificData instance, wrapping the provided SubscriptionProductChange value.
func (e *internalEventEventSpecificData) FromSubscriptionProductChange(val SubscriptionProductChange) EventEventSpecificData {
	return EventEventSpecificData{value: &val}
}

// The internalEventEventSpecificData instance, wrapping the provided SubscriptionStateChange value.
func (e *internalEventEventSpecificData) FromSubscriptionStateChange(val SubscriptionStateChange) EventEventSpecificData {
	return EventEventSpecificData{value: &val}
}

// The internalEventEventSpecificData instance, wrapping the provided PaymentRelatedEvents value.
func (e *internalEventEventSpecificData) FromPaymentRelatedEvents(val PaymentRelatedEvents) EventEventSpecificData {
	return EventEventSpecificData{value: &val}
}

// The internalEventEventSpecificData instance, wrapping the provided RefundSuccess value.
func (e *internalEventEventSpecificData) FromRefundSuccess(val RefundSuccess) EventEventSpecificData {
	return EventEventSpecificData{value: &val}
}

// The internalEventEventSpecificData instance, wrapping the provided ComponentAllocationChange value.
func (e *internalEventEventSpecificData) FromComponentAllocationChange(val ComponentAllocationChange) EventEventSpecificData {
	return EventEventSpecificData{value: &val}
}

// The internalEventEventSpecificData instance, wrapping the provided MeteredUsage value.
func (e *internalEventEventSpecificData) FromMeteredUsage(val MeteredUsage) EventEventSpecificData {
	return EventEventSpecificData{value: &val}
}

// The internalEventEventSpecificData instance, wrapping the provided PrepaidUsage value.
func (e *internalEventEventSpecificData) FromPrepaidUsage(val PrepaidUsage) EventEventSpecificData {
	return EventEventSpecificData{value: &val}
}

// The internalEventEventSpecificData instance, wrapping the provided DunningStepReached value.
func (e *internalEventEventSpecificData) FromDunningStepReached(val DunningStepReached) EventEventSpecificData {
	return EventEventSpecificData{value: &val}
}

// The internalEventEventSpecificData instance, wrapping the provided InvoiceIssued value.
func (e *internalEventEventSpecificData) FromInvoiceIssued(val InvoiceIssued) EventEventSpecificData {
	return EventEventSpecificData{value: &val}
}

// The internalEventEventSpecificData instance, wrapping the provided PendingCancellationChange value.
func (e *internalEventEventSpecificData) FromPendingCancellationChange(val PendingCancellationChange) EventEventSpecificData {
	return EventEventSpecificData{value: &val}
}

// The internalEventEventSpecificData instance, wrapping the provided PrepaidSubscriptionBalanceChanged value.
func (e *internalEventEventSpecificData) FromPrepaidSubscriptionBalanceChanged(val PrepaidSubscriptionBalanceChanged) EventEventSpecificData {
	return EventEventSpecificData{value: &val}
}

// The internalEventEventSpecificData instance, wrapping the provided ProformaInvoiceIssued value.
func (e *internalEventEventSpecificData) FromProformaInvoiceIssued(val ProformaInvoiceIssued) EventEventSpecificData {
	return EventEventSpecificData{value: &val}
}

// The internalEventEventSpecificData instance, wrapping the provided SubscriptionGroupSignupSuccess value.
func (e *internalEventEventSpecificData) FromSubscriptionGroupSignupSuccess(val SubscriptionGroupSignupSuccess) EventEventSpecificData {
	return EventEventSpecificData{value: &val}
}

// The internalEventEventSpecificData instance, wrapping the provided SubscriptionGroupSignupFailure value.
func (e *internalEventEventSpecificData) FromSubscriptionGroupSignupFailure(val SubscriptionGroupSignupFailure) EventEventSpecificData {
	return EventEventSpecificData{value: &val}
}

// The internalEventEventSpecificData instance, wrapping the provided CreditAccountBalanceChanged value.
func (e *internalEventEventSpecificData) FromCreditAccountBalanceChanged(val CreditAccountBalanceChanged) EventEventSpecificData {
	return EventEventSpecificData{value: &val}
}

// The internalEventEventSpecificData instance, wrapping the provided PrepaymentAccountBalanceChanged value.
func (e *internalEventEventSpecificData) FromPrepaymentAccountBalanceChanged(val PrepaymentAccountBalanceChanged) EventEventSpecificData {
	return EventEventSpecificData{value: &val}
}

// The internalEventEventSpecificData instance, wrapping the provided PaymentCollectionMethodChanged value.
func (e *internalEventEventSpecificData) FromPaymentCollectionMethodChanged(val PaymentCollectionMethodChanged) EventEventSpecificData {
	return EventEventSpecificData{value: &val}
}

// The internalEventEventSpecificData instance, wrapping the provided ItemPricePointChanged value.
func (e *internalEventEventSpecificData) FromItemPricePointChanged(val ItemPricePointChanged) EventEventSpecificData {
	return EventEventSpecificData{value: &val}
}

// The internalEventEventSpecificData instance, wrapping the provided CustomFieldValueChange value.
func (e *internalEventEventSpecificData) FromCustomFieldValueChange(val CustomFieldValueChange) EventEventSpecificData {
	return EventEventSpecificData{value: &val}
}
