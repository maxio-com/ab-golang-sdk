// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// InvoiceEvent represents a InvoiceEvent struct.
type InvoiceEvent struct {
    value                                any
    isApplyCreditNoteEvent               bool
    isApplyDebitNoteEvent                bool
    isApplyPaymentEvent                  bool
    isBackportInvoiceEvent               bool
    isChangeChargebackStatusEvent        bool
    isChangeInvoiceCollectionMethodEvent bool
    isChangeInvoiceStatusEvent           bool
    isCreateCreditNoteEvent              bool
    isCreateDebitNoteEvent               bool
    isFailedPaymentEvent                 bool
    isIssueInvoiceEvent                  bool
    isRefundInvoiceEvent                 bool
    isRemovePaymentEvent                 bool
    isVoidInvoiceEvent                   bool
    isVoidRemainderEvent                 bool
}

// String implements the fmt.Stringer interface for InvoiceEvent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i InvoiceEvent) String() string {
    return fmt.Sprintf("%v", i.value)
}

// MarshalJSON implements the json.Marshaler interface for InvoiceEvent.
// It customizes the JSON marshaling process for InvoiceEvent objects.
func (i InvoiceEvent) MarshalJSON() (
    []byte,
    error) {
    if i.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.InvoiceEventContainer.From*` functions to initialize the InvoiceEvent object.")
    }
    return json.Marshal(i.toMap())
}

// toMap converts the InvoiceEvent object to a map representation for JSON marshaling.
func (i *InvoiceEvent) toMap() any {
    switch obj := i.value.(type) {
    case *ApplyCreditNoteEvent:
        return obj.toMap()
    case *ApplyDebitNoteEvent:
        return obj.toMap()
    case *ApplyPaymentEvent:
        return obj.toMap()
    case *BackportInvoiceEvent:
        return obj.toMap()
    case *ChangeChargebackStatusEvent:
        return obj.toMap()
    case *ChangeInvoiceCollectionMethodEvent:
        return obj.toMap()
    case *ChangeInvoiceStatusEvent:
        return obj.toMap()
    case *CreateCreditNoteEvent:
        return obj.toMap()
    case *CreateDebitNoteEvent:
        return obj.toMap()
    case *FailedPaymentEvent:
        return obj.toMap()
    case *IssueInvoiceEvent:
        return obj.toMap()
    case *RefundInvoiceEvent:
        return obj.toMap()
    case *RemovePaymentEvent:
        return obj.toMap()
    case *VoidInvoiceEvent:
        return obj.toMap()
    case *VoidRemainderEvent:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceEvent.
// It customizes the JSON unmarshaling process for InvoiceEvent objects.
func (i *InvoiceEvent) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOfWithDiscriminator(input, "event_type",
        NewTypeHolderDiscriminator(&ApplyCreditNoteEvent{}, false, &i.isApplyCreditNoteEvent, "apply_credit_note"),
        NewTypeHolderDiscriminator(&ApplyDebitNoteEvent{}, false, &i.isApplyDebitNoteEvent, "apply_debit_note"),
        NewTypeHolderDiscriminator(&ApplyPaymentEvent{}, false, &i.isApplyPaymentEvent, "apply_payment"),
        NewTypeHolderDiscriminator(&BackportInvoiceEvent{}, false, &i.isBackportInvoiceEvent, "backport_invoice"),
        NewTypeHolderDiscriminator(&ChangeChargebackStatusEvent{}, false, &i.isChangeChargebackStatusEvent, "change_chargeback_status"),
        NewTypeHolderDiscriminator(&ChangeInvoiceCollectionMethodEvent{}, false, &i.isChangeInvoiceCollectionMethodEvent, "change_invoice_collection_method"),
        NewTypeHolderDiscriminator(&ChangeInvoiceStatusEvent{}, false, &i.isChangeInvoiceStatusEvent, "change_invoice_status"),
        NewTypeHolderDiscriminator(&CreateCreditNoteEvent{}, false, &i.isCreateCreditNoteEvent, "create_credit_note"),
        NewTypeHolderDiscriminator(&CreateDebitNoteEvent{}, false, &i.isCreateDebitNoteEvent, "create_debit_note"),
        NewTypeHolderDiscriminator(&FailedPaymentEvent{}, false, &i.isFailedPaymentEvent, "failed_payment"),
        NewTypeHolderDiscriminator(&IssueInvoiceEvent{}, false, &i.isIssueInvoiceEvent, "issue_invoice"),
        NewTypeHolderDiscriminator(&RefundInvoiceEvent{}, false, &i.isRefundInvoiceEvent, "refund_invoice"),
        NewTypeHolderDiscriminator(&RemovePaymentEvent{}, false, &i.isRemovePaymentEvent, "remove_payment"),
        NewTypeHolderDiscriminator(&VoidInvoiceEvent{}, false, &i.isVoidInvoiceEvent, "void_invoice"),
        NewTypeHolderDiscriminator(&VoidRemainderEvent{}, false, &i.isVoidRemainderEvent, "void_remainder"),
    )
    
    i.value = result
    return err
}

func (i *InvoiceEvent) AsApplyCreditNoteEvent() (
    *ApplyCreditNoteEvent,
    bool) {
    if !i.isApplyCreditNoteEvent {
        return nil, false
    }
    return i.value.(*ApplyCreditNoteEvent), true
}

func (i *InvoiceEvent) AsApplyDebitNoteEvent() (
    *ApplyDebitNoteEvent,
    bool) {
    if !i.isApplyDebitNoteEvent {
        return nil, false
    }
    return i.value.(*ApplyDebitNoteEvent), true
}

func (i *InvoiceEvent) AsApplyPaymentEvent() (
    *ApplyPaymentEvent,
    bool) {
    if !i.isApplyPaymentEvent {
        return nil, false
    }
    return i.value.(*ApplyPaymentEvent), true
}

func (i *InvoiceEvent) AsBackportInvoiceEvent() (
    *BackportInvoiceEvent,
    bool) {
    if !i.isBackportInvoiceEvent {
        return nil, false
    }
    return i.value.(*BackportInvoiceEvent), true
}

func (i *InvoiceEvent) AsChangeChargebackStatusEvent() (
    *ChangeChargebackStatusEvent,
    bool) {
    if !i.isChangeChargebackStatusEvent {
        return nil, false
    }
    return i.value.(*ChangeChargebackStatusEvent), true
}

func (i *InvoiceEvent) AsChangeInvoiceCollectionMethodEvent() (
    *ChangeInvoiceCollectionMethodEvent,
    bool) {
    if !i.isChangeInvoiceCollectionMethodEvent {
        return nil, false
    }
    return i.value.(*ChangeInvoiceCollectionMethodEvent), true
}

func (i *InvoiceEvent) AsChangeInvoiceStatusEvent() (
    *ChangeInvoiceStatusEvent,
    bool) {
    if !i.isChangeInvoiceStatusEvent {
        return nil, false
    }
    return i.value.(*ChangeInvoiceStatusEvent), true
}

func (i *InvoiceEvent) AsCreateCreditNoteEvent() (
    *CreateCreditNoteEvent,
    bool) {
    if !i.isCreateCreditNoteEvent {
        return nil, false
    }
    return i.value.(*CreateCreditNoteEvent), true
}

func (i *InvoiceEvent) AsCreateDebitNoteEvent() (
    *CreateDebitNoteEvent,
    bool) {
    if !i.isCreateDebitNoteEvent {
        return nil, false
    }
    return i.value.(*CreateDebitNoteEvent), true
}

func (i *InvoiceEvent) AsFailedPaymentEvent() (
    *FailedPaymentEvent,
    bool) {
    if !i.isFailedPaymentEvent {
        return nil, false
    }
    return i.value.(*FailedPaymentEvent), true
}

func (i *InvoiceEvent) AsIssueInvoiceEvent() (
    *IssueInvoiceEvent,
    bool) {
    if !i.isIssueInvoiceEvent {
        return nil, false
    }
    return i.value.(*IssueInvoiceEvent), true
}

func (i *InvoiceEvent) AsRefundInvoiceEvent() (
    *RefundInvoiceEvent,
    bool) {
    if !i.isRefundInvoiceEvent {
        return nil, false
    }
    return i.value.(*RefundInvoiceEvent), true
}

func (i *InvoiceEvent) AsRemovePaymentEvent() (
    *RemovePaymentEvent,
    bool) {
    if !i.isRemovePaymentEvent {
        return nil, false
    }
    return i.value.(*RemovePaymentEvent), true
}

func (i *InvoiceEvent) AsVoidInvoiceEvent() (
    *VoidInvoiceEvent,
    bool) {
    if !i.isVoidInvoiceEvent {
        return nil, false
    }
    return i.value.(*VoidInvoiceEvent), true
}

func (i *InvoiceEvent) AsVoidRemainderEvent() (
    *VoidRemainderEvent,
    bool) {
    if !i.isVoidRemainderEvent {
        return nil, false
    }
    return i.value.(*VoidRemainderEvent), true
}

// internalInvoiceEvent represents a invoiceEvent struct.
type internalInvoiceEvent struct {}

var InvoiceEventContainer internalInvoiceEvent

// The internalInvoiceEvent instance, wrapping the provided ApplyCreditNoteEvent value.
func (i *internalInvoiceEvent) FromApplyCreditNoteEvent(val ApplyCreditNoteEvent) InvoiceEvent {
    return InvoiceEvent{value: &val}
}

// The internalInvoiceEvent instance, wrapping the provided ApplyDebitNoteEvent value.
func (i *internalInvoiceEvent) FromApplyDebitNoteEvent(val ApplyDebitNoteEvent) InvoiceEvent {
    return InvoiceEvent{value: &val}
}

// The internalInvoiceEvent instance, wrapping the provided ApplyPaymentEvent value.
func (i *internalInvoiceEvent) FromApplyPaymentEvent(val ApplyPaymentEvent) InvoiceEvent {
    return InvoiceEvent{value: &val}
}

// The internalInvoiceEvent instance, wrapping the provided BackportInvoiceEvent value.
func (i *internalInvoiceEvent) FromBackportInvoiceEvent(val BackportInvoiceEvent) InvoiceEvent {
    return InvoiceEvent{value: &val}
}

// The internalInvoiceEvent instance, wrapping the provided ChangeChargebackStatusEvent value.
func (i *internalInvoiceEvent) FromChangeChargebackStatusEvent(val ChangeChargebackStatusEvent) InvoiceEvent {
    return InvoiceEvent{value: &val}
}

// The internalInvoiceEvent instance, wrapping the provided ChangeInvoiceCollectionMethodEvent value.
func (i *internalInvoiceEvent) FromChangeInvoiceCollectionMethodEvent(val ChangeInvoiceCollectionMethodEvent) InvoiceEvent {
    return InvoiceEvent{value: &val}
}

// The internalInvoiceEvent instance, wrapping the provided ChangeInvoiceStatusEvent value.
func (i *internalInvoiceEvent) FromChangeInvoiceStatusEvent(val ChangeInvoiceStatusEvent) InvoiceEvent {
    return InvoiceEvent{value: &val}
}

// The internalInvoiceEvent instance, wrapping the provided CreateCreditNoteEvent value.
func (i *internalInvoiceEvent) FromCreateCreditNoteEvent(val CreateCreditNoteEvent) InvoiceEvent {
    return InvoiceEvent{value: &val}
}

// The internalInvoiceEvent instance, wrapping the provided CreateDebitNoteEvent value.
func (i *internalInvoiceEvent) FromCreateDebitNoteEvent(val CreateDebitNoteEvent) InvoiceEvent {
    return InvoiceEvent{value: &val}
}

// The internalInvoiceEvent instance, wrapping the provided FailedPaymentEvent value.
func (i *internalInvoiceEvent) FromFailedPaymentEvent(val FailedPaymentEvent) InvoiceEvent {
    return InvoiceEvent{value: &val}
}

// The internalInvoiceEvent instance, wrapping the provided IssueInvoiceEvent value.
func (i *internalInvoiceEvent) FromIssueInvoiceEvent(val IssueInvoiceEvent) InvoiceEvent {
    return InvoiceEvent{value: &val}
}

// The internalInvoiceEvent instance, wrapping the provided RefundInvoiceEvent value.
func (i *internalInvoiceEvent) FromRefundInvoiceEvent(val RefundInvoiceEvent) InvoiceEvent {
    return InvoiceEvent{value: &val}
}

// The internalInvoiceEvent instance, wrapping the provided RemovePaymentEvent value.
func (i *internalInvoiceEvent) FromRemovePaymentEvent(val RemovePaymentEvent) InvoiceEvent {
    return InvoiceEvent{value: &val}
}

// The internalInvoiceEvent instance, wrapping the provided VoidInvoiceEvent value.
func (i *internalInvoiceEvent) FromVoidInvoiceEvent(val VoidInvoiceEvent) InvoiceEvent {
    return InvoiceEvent{value: &val}
}

// The internalInvoiceEvent instance, wrapping the provided VoidRemainderEvent value.
func (i *internalInvoiceEvent) FromVoidRemainderEvent(val VoidRemainderEvent) InvoiceEvent {
    return InvoiceEvent{value: &val}
}
