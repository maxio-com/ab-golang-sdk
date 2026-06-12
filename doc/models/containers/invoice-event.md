
# Invoice Event

## Class Name

`InvoiceEvent`

## Cases

| Type | Factory Method |
|  --- | --- |
| [`models.ApplyCreditNoteEvent`](../../../doc/models/apply-credit-note-event.md) | models.InvoiceEventContainer.FromApplyCreditNoteEvent(models.ApplyCreditNoteEvent applyCreditNoteEvent) |
| [`models.ApplyDebitNoteEvent`](../../../doc/models/apply-debit-note-event.md) | models.InvoiceEventContainer.FromApplyDebitNoteEvent(models.ApplyDebitNoteEvent applyDebitNoteEvent) |
| [`models.ApplyPaymentEvent`](../../../doc/models/apply-payment-event.md) | models.InvoiceEventContainer.FromApplyPaymentEvent(models.ApplyPaymentEvent applyPaymentEvent) |
| [`models.BackportInvoiceEvent`](../../../doc/models/backport-invoice-event.md) | models.InvoiceEventContainer.FromBackportInvoiceEvent(models.BackportInvoiceEvent backportInvoiceEvent) |
| [`models.ChangeChargebackStatusEvent`](../../../doc/models/change-chargeback-status-event.md) | models.InvoiceEventContainer.FromChangeChargebackStatusEvent(models.ChangeChargebackStatusEvent changeChargebackStatusEvent) |
| [`models.ChangeInvoiceCollectionMethodEvent`](../../../doc/models/change-invoice-collection-method-event.md) | models.InvoiceEventContainer.FromChangeInvoiceCollectionMethodEvent(models.ChangeInvoiceCollectionMethodEvent changeInvoiceCollectionMethodEvent) |
| [`models.ChangeInvoiceStatusEvent`](../../../doc/models/change-invoice-status-event.md) | models.InvoiceEventContainer.FromChangeInvoiceStatusEvent(models.ChangeInvoiceStatusEvent changeInvoiceStatusEvent) |
| [`models.CreateCreditNoteEvent`](../../../doc/models/create-credit-note-event.md) | models.InvoiceEventContainer.FromCreateCreditNoteEvent(models.CreateCreditNoteEvent createCreditNoteEvent) |
| [`models.CreateDebitNoteEvent`](../../../doc/models/create-debit-note-event.md) | models.InvoiceEventContainer.FromCreateDebitNoteEvent(models.CreateDebitNoteEvent createDebitNoteEvent) |
| [`models.FailedPaymentEvent`](../../../doc/models/failed-payment-event.md) | models.InvoiceEventContainer.FromFailedPaymentEvent(models.FailedPaymentEvent failedPaymentEvent) |
| [`models.IssueInvoiceEvent`](../../../doc/models/issue-invoice-event.md) | models.InvoiceEventContainer.FromIssueInvoiceEvent(models.IssueInvoiceEvent issueInvoiceEvent) |
| [`models.RefundInvoiceEvent`](../../../doc/models/refund-invoice-event.md) | models.InvoiceEventContainer.FromRefundInvoiceEvent(models.RefundInvoiceEvent refundInvoiceEvent) |
| [`models.RemovePaymentEvent`](../../../doc/models/remove-payment-event.md) | models.InvoiceEventContainer.FromRemovePaymentEvent(models.RemovePaymentEvent removePaymentEvent) |
| [`models.VoidInvoiceEvent`](../../../doc/models/void-invoice-event.md) | models.InvoiceEventContainer.FromVoidInvoiceEvent(models.VoidInvoiceEvent voidInvoiceEvent) |
| [`models.VoidRemainderEvent`](../../../doc/models/void-remainder-event.md) | models.InvoiceEventContainer.FromVoidRemainderEvent(models.VoidRemainderEvent voidRemainderEvent) |

## models.ApplyCreditNoteEvent

### Initialization Code

#### Example

```go
value := models.InvoiceEventContainer.FromApplyCreditNoteEvent(models.ApplyCreditNoteEvent{
    Id:                   int64(214),
    Timestamp:            parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
    Invoice:              models.Invoice{
        IssueDate:                  models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
        DueDate:                    models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
        PaidDate:                   models.NewOptional(models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) }))),
        PublicUrlExpiresOn:         models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-21", func(err error) { log.Fatalln(err) })),
    },
    EventType:            models.InvoiceEventType_APPLYCREDITNOTE,
    EventData:            models.ApplyCreditNoteEventData{
        Uid:                  "uid6",
        CreditNoteNumber:     "credit_note_number0",
        CreditNoteUid:        "credit_note_uid0",
        OriginalAmount:       "original_amount0",
        AppliedAmount:        "applied_amount2",
    },
})
```

## models.ApplyDebitNoteEvent

### Initialization Code

#### Example

```go
value := models.InvoiceEventContainer.FromApplyDebitNoteEvent(models.ApplyDebitNoteEvent{
    Id:                   int64(164),
    Timestamp:            parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
    Invoice:              models.Invoice{
        IssueDate:                  models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
        DueDate:                    models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
        PaidDate:                   models.NewOptional(models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) }))),
        PublicUrlExpiresOn:         models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-21", func(err error) { log.Fatalln(err) })),
    },
    EventType:            models.InvoiceEventType_APPLYDEBITNOTE,
    EventData:            models.ApplyDebitNoteEventData{
        DebitNoteNumber:      "debit_note_number6",
        DebitNoteUid:         "debit_note_uid2",
        OriginalAmount:       "original_amount0",
        AppliedAmount:        "applied_amount2",
    },
})
```

## models.ApplyPaymentEvent

### Initialization Code

#### Example

```go
value := models.InvoiceEventContainer.FromApplyPaymentEvent(models.ApplyPaymentEvent{
    Id:                   int64(234),
    Timestamp:            parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
    Invoice:              models.Invoice{
        IssueDate:                  models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
        DueDate:                    models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
        PaidDate:                   models.NewOptional(models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) }))),
        PublicUrlExpiresOn:         models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-21", func(err error) { log.Fatalln(err) })),
    },
    EventType:            models.InvoiceEventType_APPLYPAYMENT,
    EventData:            models.ApplyPaymentEventData{
        ConsolidationLevel:        models.InvoiceConsolidationLevel_CHILD,
        Memo:                      "memo0",
        OriginalAmount:            "original_amount0",
        AppliedAmount:             "applied_amount2",
        TransactionTime:           parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
        PaymentMethod:             models.InvoiceEventPaymentContainer.FromPaymentMethodApplePay(models.PaymentMethodApplePay{
            Type:                 models.InvoiceEventPaymentMethod_APPLEPAY,
        }),
    },
})
```

## models.BackportInvoiceEvent

### Initialization Code

#### Example

```go
value := models.InvoiceEventContainer.FromBackportInvoiceEvent(models.BackportInvoiceEvent{
    Id:                   int64(78),
    Timestamp:            parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
    Invoice:              models.Invoice{
        IssueDate:                  models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
        DueDate:                    models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
        PaidDate:                   models.NewOptional(models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) }))),
        PublicUrlExpiresOn:         models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-21", func(err error) { log.Fatalln(err) })),
    },
    EventType:            models.InvoiceEventType_BACKPORTINVOICE,
    EventData:            models.Invoice{
        IssueDate:                  models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
        DueDate:                    models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
        PaidDate:                   models.NewOptional(models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) }))),
        PublicUrlExpiresOn:         models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-21", func(err error) { log.Fatalln(err) })),
    },
})
```

## models.ChangeChargebackStatusEvent

### Initialization Code

#### Example

```go
value := models.InvoiceEventContainer.FromChangeChargebackStatusEvent(models.ChangeChargebackStatusEvent{
    Id:                   int64(214),
    Timestamp:            parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
    Invoice:              models.Invoice{
        IssueDate:                  models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
        DueDate:                    models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
        PaidDate:                   models.NewOptional(models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) }))),
        PublicUrlExpiresOn:         models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-21", func(err error) { log.Fatalln(err) })),
    },
    EventType:            models.InvoiceEventType_CHANGECHARGEBACKSTATUS,
    EventData:            models.ChangeChargebackStatusEventData{
        ChargebackStatus:     models.ChargebackStatus_WON,
    },
})
```

## models.ChangeInvoiceCollectionMethodEvent

### Initialization Code

#### Example

```go
value := models.InvoiceEventContainer.FromChangeInvoiceCollectionMethodEvent(models.ChangeInvoiceCollectionMethodEvent{
    Id:                   int64(246),
    Timestamp:            parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
    Invoice:              models.Invoice{
        IssueDate:                  models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
        DueDate:                    models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
        PaidDate:                   models.NewOptional(models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) }))),
        PublicUrlExpiresOn:         models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-21", func(err error) { log.Fatalln(err) })),
    },
    EventType:            models.InvoiceEventType_CHANGEINVOICECOLLECTIONMETHOD,
    EventData:            models.ChangeInvoiceCollectionMethodEventData{
        FromCollectionMethod: "from_collection_method4",
        ToCollectionMethod:   "to_collection_method8",
    },
})
```

## models.ChangeInvoiceStatusEvent

### Initialization Code

#### Example

```go
value := models.InvoiceEventContainer.FromChangeInvoiceStatusEvent(models.ChangeInvoiceStatusEvent{
    Id:                   int64(92),
    Timestamp:            parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
    Invoice:              models.Invoice{
        IssueDate:                  models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
        DueDate:                    models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
        PaidDate:                   models.NewOptional(models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) }))),
        PublicUrlExpiresOn:         models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-21", func(err error) { log.Fatalln(err) })),
    },
    EventType:            models.InvoiceEventType_CHANGEINVOICESTATUS,
    EventData:            models.ChangeInvoiceStatusEventData{
        FromStatus:           models.InvoiceStatus_OPEN,
        ToStatus:             models.InvoiceStatus_PENDING,
    },
})
```

## models.CreateCreditNoteEvent

### Initialization Code

#### Example

```go
value := models.InvoiceEventContainer.FromCreateCreditNoteEvent(models.CreateCreditNoteEvent{
    Id:                   int64(28),
    Timestamp:            parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
    Invoice:              models.Invoice{
        IssueDate:                  models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
        DueDate:                    models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
        PaidDate:                   models.NewOptional(models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) }))),
        PublicUrlExpiresOn:         models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-21", func(err error) { log.Fatalln(err) })),
    },
    EventType:            models.InvoiceEventType_CREATECREDITNOTE,
    EventData:            models.CreditNote{
    },
})
```

## models.CreateDebitNoteEvent

### Initialization Code

#### Example

```go
value := models.InvoiceEventContainer.FromCreateDebitNoteEvent(models.CreateDebitNoteEvent{
    Id:                   int64(98),
    Timestamp:            parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
    Invoice:              models.Invoice{
        IssueDate:                  models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
        DueDate:                    models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
        PaidDate:                   models.NewOptional(models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) }))),
        PublicUrlExpiresOn:         models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-21", func(err error) { log.Fatalln(err) })),
    },
    EventType:            models.InvoiceEventType_CREATEDEBITNOTE,
    EventData:            models.DebitNote{
    },
})
```

## models.FailedPaymentEvent

### Initialization Code

#### Example

```go
value := models.InvoiceEventContainer.FromFailedPaymentEvent(models.FailedPaymentEvent{
    Id:                   int64(120),
    Timestamp:            parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
    Invoice:              models.Invoice{
        IssueDate:                  models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
        DueDate:                    models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
        PaidDate:                   models.NewOptional(models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) }))),
        PublicUrlExpiresOn:         models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-21", func(err error) { log.Fatalln(err) })),
    },
    EventType:            models.InvoiceEventType_FAILEDPAYMENT,
    EventData:            models.FailedPaymentEventData{
        AmountInCents:        220,
        AppliedAmount:        194,
        PaymentMethod:        models.InvoicePaymentMethodType_CASH,
        TransactionId:        78,
    },
})
```

## models.IssueInvoiceEvent

### Initialization Code

#### Example

```go
value := models.InvoiceEventContainer.FromIssueInvoiceEvent(models.IssueInvoiceEvent{
    Id:                   int64(130),
    Timestamp:            parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
    Invoice:              models.Invoice{
        IssueDate:                  models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
        DueDate:                    models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
        PaidDate:                   models.NewOptional(models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) }))),
        PublicUrlExpiresOn:         models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-21", func(err error) { log.Fatalln(err) })),
    },
    EventType:            models.InvoiceEventType_ISSUEINVOICE,
    EventData:            models.IssueInvoiceEventData{
        ConsolidationLevel:   models.InvoiceConsolidationLevel_CHILD,
        FromStatus:           models.InvoiceStatus_OPEN,
        ToStatus:             models.InvoiceStatus_PENDING,
        DueAmount:            "due_amount8",
        TotalAmount:          "total_amount2",
    },
})
```

## models.RefundInvoiceEvent

### Initialization Code

#### Example

```go
value := models.InvoiceEventContainer.FromRefundInvoiceEvent(models.RefundInvoiceEvent{
    Id:                   int64(54),
    Timestamp:            parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
    Invoice:              models.Invoice{
        IssueDate:                  models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
        DueDate:                    models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
        PaidDate:                   models.NewOptional(models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) }))),
        PublicUrlExpiresOn:         models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-21", func(err error) { log.Fatalln(err) })),
    },
    EventType:            models.InvoiceEventType_REFUNDINVOICE,
    EventData:            models.RefundInvoiceEventData{
        ApplyCredit:          false,
        CreditNoteAttributes: models.CreditNote{
        },
        PaymentId:            204,
        RefundAmount:         "refund_amount8",
        RefundId:             248,
        TransactionTime:      parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
    },
})
```

## models.RemovePaymentEvent

### Initialization Code

#### Example

```go
value := models.InvoiceEventContainer.FromRemovePaymentEvent(models.RemovePaymentEvent{
    Id:                   int64(236),
    Timestamp:            parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
    Invoice:              models.Invoice{
        IssueDate:                  models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
        DueDate:                    models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
        PaidDate:                   models.NewOptional(models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) }))),
        PublicUrlExpiresOn:         models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-21", func(err error) { log.Fatalln(err) })),
    },
    EventType:            models.InvoiceEventType_REMOVEPAYMENT,
    EventData:            models.RemovePaymentEventData{
        TransactionId:        78,
        Memo:                 "memo0",
        AppliedAmount:        "applied_amount2",
        TransactionTime:      parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
        PaymentMethod:        models.InvoiceEventPaymentContainer.FromPaymentMethodApplePay(models.PaymentMethodApplePay{
            Type:                 models.InvoiceEventPaymentMethod_APPLEPAY,
        }),
        Prepayment:           false,
    },
})
```

## models.VoidInvoiceEvent

### Initialization Code

#### Example

```go
value := models.InvoiceEventContainer.FromVoidInvoiceEvent(models.VoidInvoiceEvent{
    Id:                   int64(16),
    Timestamp:            parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
    Invoice:              models.Invoice{
        IssueDate:                  models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
        DueDate:                    models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
        PaidDate:                   models.NewOptional(models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) }))),
        PublicUrlExpiresOn:         models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-21", func(err error) { log.Fatalln(err) })),
    },
    EventType:            models.InvoiceEventType_VOIDINVOICE,
    EventData:            models.VoidInvoiceEventData{
        CreditNoteAttributes: models.ToPointer(models.CreditNote{
        }),
        Memo:                 models.ToPointer("memo0"),
        AppliedAmount:        models.ToPointer("applied_amount2"),
        TransactionTime:      models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
        IsAdvanceInvoice:     false,
        Reason:               "reason2",
    },
})
```

## models.VoidRemainderEvent

### Initialization Code

#### Example

```go
value := models.InvoiceEventContainer.FromVoidRemainderEvent(models.VoidRemainderEvent{
    Id:                   int64(128),
    Timestamp:            parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
    Invoice:              models.Invoice{
        IssueDate:                  models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
        DueDate:                    models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
        PaidDate:                   models.NewOptional(models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) }))),
        PublicUrlExpiresOn:         models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-21", func(err error) { log.Fatalln(err) })),
    },
    EventType:            models.InvoiceEventType_VOIDREMAINDER,
    EventData:            models.VoidRemainderEventData{
        CreditNoteAttributes: models.CreditNote{
        },
        Memo:                 "memo0",
        AppliedAmount:        "applied_amount2",
        TransactionTime:      parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
    },
})
```

