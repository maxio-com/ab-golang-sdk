
# Event Event Specific Data

## Class Name

`EventEventSpecificData`

## Cases

| Type | Factory Method |
|  --- | --- |
| [`models.SubscriptionProductChange`](../../../doc/models/subscription-product-change.md) | models.EventEventSpecificDataContainer.FromSubscriptionProductChange(models.SubscriptionProductChange subscriptionProductChange) |
| [`models.SubscriptionStateChange`](../../../doc/models/subscription-state-change.md) | models.EventEventSpecificDataContainer.FromSubscriptionStateChange(models.SubscriptionStateChange subscriptionStateChange) |
| [`models.PaymentRelatedEvents`](../../../doc/models/payment-related-events.md) | models.EventEventSpecificDataContainer.FromPaymentRelatedEvents(models.PaymentRelatedEvents paymentRelatedEvents) |
| [`models.RefundSuccess`](../../../doc/models/refund-success.md) | models.EventEventSpecificDataContainer.FromRefundSuccess(models.RefundSuccess refundSuccess) |
| [`models.ComponentAllocationChange`](../../../doc/models/component-allocation-change.md) | models.EventEventSpecificDataContainer.FromComponentAllocationChange(models.ComponentAllocationChange componentAllocationChange) |
| [`models.MeteredUsage`](../../../doc/models/metered-usage.md) | models.EventEventSpecificDataContainer.FromMeteredUsage(models.MeteredUsage meteredUsage) |
| [`models.PrepaidUsage`](../../../doc/models/prepaid-usage.md) | models.EventEventSpecificDataContainer.FromPrepaidUsage(models.PrepaidUsage prepaidUsage) |
| [`models.DunningStepReached`](../../../doc/models/dunning-step-reached.md) | models.EventEventSpecificDataContainer.FromDunningStepReached(models.DunningStepReached dunningStepReached) |
| [`models.InvoiceIssued`](../../../doc/models/invoice-issued.md) | models.EventEventSpecificDataContainer.FromInvoiceIssued(models.InvoiceIssued invoiceIssued) |
| [`models.PendingCancellationChange`](../../../doc/models/pending-cancellation-change.md) | models.EventEventSpecificDataContainer.FromPendingCancellationChange(models.PendingCancellationChange pendingCancellationChange) |
| [`models.PrepaidSubscriptionBalanceChanged`](../../../doc/models/prepaid-subscription-balance-changed.md) | models.EventEventSpecificDataContainer.FromPrepaidSubscriptionBalanceChanged(models.PrepaidSubscriptionBalanceChanged prepaidSubscriptionBalanceChanged) |
| [`models.ProformaInvoiceIssued`](../../../doc/models/proforma-invoice-issued.md) | models.EventEventSpecificDataContainer.FromProformaInvoiceIssued(models.ProformaInvoiceIssued proformaInvoiceIssued) |
| [`models.SubscriptionGroupSignupEventData`](../../../doc/models/subscription-group-signup-event-data.md) | models.EventEventSpecificDataContainer.FromSubscriptionGroupSignupEventData(models.SubscriptionGroupSignupEventData subscriptionGroupSignupEventData) |
| [`models.CreditAccountBalanceChanged`](../../../doc/models/credit-account-balance-changed.md) | models.EventEventSpecificDataContainer.FromCreditAccountBalanceChanged(models.CreditAccountBalanceChanged creditAccountBalanceChanged) |
| [`models.PrepaymentAccountBalanceChanged`](../../../doc/models/prepayment-account-balance-changed.md) | models.EventEventSpecificDataContainer.FromPrepaymentAccountBalanceChanged(models.PrepaymentAccountBalanceChanged prepaymentAccountBalanceChanged) |
| [`models.PaymentCollectionMethodChanged`](../../../doc/models/payment-collection-method-changed.md) | models.EventEventSpecificDataContainer.FromPaymentCollectionMethodChanged(models.PaymentCollectionMethodChanged paymentCollectionMethodChanged) |
| [`models.ItemPricePointChanged`](../../../doc/models/item-price-point-changed.md) | models.EventEventSpecificDataContainer.FromItemPricePointChanged(models.ItemPricePointChanged itemPricePointChanged) |
| [`models.CustomFieldValueChange`](../../../doc/models/custom-field-value-change.md) | models.EventEventSpecificDataContainer.FromCustomFieldValueChange(models.CustomFieldValueChange customFieldValueChange) |
| [`models.ChjsTokenizationSuccess`](../../../doc/models/chjs-tokenization-success.md) | models.EventEventSpecificDataContainer.FromChjsTokenizationSuccess(models.ChjsTokenizationSuccess chjsTokenizationSuccess) |
| [`models.ChjsTokenizationFailure`](../../../doc/models/chjs-tokenization-failure.md) | models.EventEventSpecificDataContainer.FromChjsTokenizationFailure(models.ChjsTokenizationFailure chjsTokenizationFailure) |

## models.SubscriptionProductChange

### Initialization Code

#### Example

```go
value := models.EventEventSpecificDataContainer.FromSubscriptionProductChange(models.SubscriptionProductChange{
    PreviousProductId:           126,
    NewProductId:                12,
})
```

## models.SubscriptionStateChange

### Initialization Code

#### Example

```go
value := models.EventEventSpecificDataContainer.FromSubscriptionStateChange(models.SubscriptionStateChange{
    PreviousSubscriptionState: "previous_subscription_state2",
    NewSubscriptionState:      "new_subscription_state6",
})
```

## models.PaymentRelatedEvents

### Initialization Code

#### Example

```go
value := models.EventEventSpecificDataContainer.FromPaymentRelatedEvents(models.PaymentRelatedEvents{
    ProductId:            42,
    AccountTransactionId: 58,
})
```

## models.RefundSuccess

### Initialization Code

#### Example

```go
value := models.EventEventSpecificDataContainer.FromRefundSuccess(models.RefundSuccess{
    RefundId:             12,
    GatewayTransactionId: 182,
    ProductId:            168,
})
```

## models.ComponentAllocationChange

### Initialization Code

#### Example

```go
value := models.EventEventSpecificDataContainer.FromComponentAllocationChange(models.ComponentAllocationChange{
    PreviousAllocation:   94,
    NewAllocation:        102,
    ComponentId:          88,
    ComponentHandle:      "component_handle8",
    Memo:                 "memo2",
    AllocationId:         158,
})
```

## models.MeteredUsage

### Initialization Code

#### Example

```go
value := models.EventEventSpecificDataContainer.FromMeteredUsage(models.MeteredUsage{
    PreviousUnitBalance:  "previous_unit_balance6",
    NewUnitBalance:       models.MeteredUsageNewUnitBalanceContainer.FromNumber(2),
    UsageQuantity:        42,
    ComponentId:          4,
    ComponentHandle:      "component_handle8",
    Memo:                 "memo2",
})
```

## models.PrepaidUsage

### Initialization Code

#### Example

```go
value := models.EventEventSpecificDataContainer.FromPrepaidUsage(models.PrepaidUsage{
    PreviousUnitBalance:        "previous_unit_balance0",
    PreviousOverageUnitBalance: "previous_overage_unit_balance4",
    NewUnitBalance:             models.PrepaidUsageNewUnitBalanceContainer.FromNumber(174),
    NewOverageUnitBalance:      models.PrepaidUsageNewOverageUnitBalanceContainer.FromNumber(146),
    UsageQuantity:              214,
    OverageUsageQuantity:       106,
    ComponentId:                176,
    ComponentHandle:            "component_handle4",
    Memo:                       "memo8",
    AllocationDetails:          []models.PrepaidUsageAllocationDetail{
        models.PrepaidUsageAllocationDetail{
        },
    },
})
```

## models.DunningStepReached

### Initialization Code

#### Example

```go
value := models.EventEventSpecificDataContainer.FromDunningStepReached(models.DunningStepReached{
    Dunner:               models.DunnerData{
        State:                "state8",
        SubscriptionId:       194,
        RevenueAtRiskInCents: int64(98),
        CreatedAt:            parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
        Attempts:             42,
        LastAttemptedAt:      parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
    },
    CurrentStep:          models.DunningStepData{
        DayThreshold:         198,
        Action:               "action4",
        SendEmail:            false,
        SendBccEmail:         false,
        SendSms:              false,
    },
    NextStep:             models.DunningStepData{
        DayThreshold:         30,
        Action:               "action4",
        SendEmail:            false,
        SendBccEmail:         false,
        SendSms:              false,
    },
})
```

## models.InvoiceIssued

### Initialization Code

#### Example

```go
value := models.EventEventSpecificDataContainer.FromInvoiceIssued(models.InvoiceIssued{
    Uid:                  "uid4",
    Number:               "number8",
    Role:                 "role2",
    DueDate:              models.ToPointer(parseTime(models.DEFAULT_DATE, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
    IssueDate:            "issue_date0",
    PaidDate:             "paid_date6",
    DueAmount:            "due_amount6",
    PaidAmount:           "paid_amount4",
    TaxAmount:            "tax_amount2",
    RefundAmount:         "refund_amount0",
    TotalAmount:          "total_amount0",
    StatusAmount:         "status_amount4",
    ProductName:          "product_name0",
    ConsolidationLevel:   "consolidation_level4",
    LineItems:            []models.InvoiceLineItemEventData{
        models.InvoiceLineItemEventData{
        },
    },
})
```

## models.PendingCancellationChange

### Initialization Code

#### Example

```go
value := models.EventEventSpecificDataContainer.FromPendingCancellationChange(models.PendingCancellationChange{
    CancellationState:    "cancellation_state8",
    CancelsAt:            parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
})
```

## models.PrepaidSubscriptionBalanceChanged

### Initialization Code

#### Example

```go
value := models.EventEventSpecificDataContainer.FromPrepaidSubscriptionBalanceChanged(models.PrepaidSubscriptionBalanceChanged{
    Reason:                          "reason8",
    CurrentAccountBalanceInCents:    int64(250),
    PrepaymentAccountBalanceInCents: int64(44),
    CurrentUsageAmountInCents:       int64(242),
})
```

## models.ProformaInvoiceIssued

### Initialization Code

#### Example

```go
value := models.EventEventSpecificDataContainer.FromProformaInvoiceIssued(models.ProformaInvoiceIssued{
    Uid:                  "uid0",
    Number:               "number2",
    Role:                 "role6",
    DeliveryDate:         parseTime(models.DEFAULT_DATE, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
    CreatedAt:            parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
    DueAmount:            "due_amount2",
    PaidAmount:           "paid_amount8",
    TaxAmount:            "tax_amount6",
    TotalAmount:          "total_amount6",
    ProductName:          "product_name6",
    LineItems:            []models.InvoiceLineItemEventData{
        models.InvoiceLineItemEventData{
        },
    },
})
```

## models.SubscriptionGroupSignupEventData

### Initialization Code

#### Example

```go
value := models.EventEventSpecificDataContainer.FromSubscriptionGroupSignupEventData(models.SubscriptionGroupSignupEventData{
    SubscriptionGroup:    models.SubscriptionGroupSignupFailureData{
    },
    Customer:             models.ToPointer(models.Customer{
    }),
})
```

## models.CreditAccountBalanceChanged

### Initialization Code

#### Example

```go
value := models.EventEventSpecificDataContainer.FromCreditAccountBalanceChanged(models.CreditAccountBalanceChanged{
    Reason:                             "reason8",
    ServiceCreditAccountBalanceInCents: int64(10),
    ServiceCreditBalanceChangeInCents:  int64(116),
    CurrencyCode:                       "currency_code8",
    AtTime:                             parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
})
```

## models.PrepaymentAccountBalanceChanged

### Initialization Code

#### Example

```go
value := models.EventEventSpecificDataContainer.FromPrepaymentAccountBalanceChanged(models.PrepaymentAccountBalanceChanged{
    Reason:                          "reason4",
    PrepaymentAccountBalanceInCents: int64(182),
    PrepaymentBalanceChangeInCents:  int64(206),
    CurrencyCode:                    "currency_code4",
})
```

## models.PaymentCollectionMethodChanged

### Initialization Code

#### Example

```go
value := models.EventEventSpecificDataContainer.FromPaymentCollectionMethodChanged(models.PaymentCollectionMethodChanged{
    PreviousValue:        "previous_value4",
    CurrentValue:         "current_value2",
})
```

## models.ItemPricePointChanged

### Initialization Code

#### Example

```go
value := models.EventEventSpecificDataContainer.FromItemPricePointChanged(models.ItemPricePointChanged{
    ItemId:               66,
    ItemType:             "item_type6",
    ItemHandle:           "item_handle4",
    ItemName:             "item_name8",
    PreviousPricePoint:   models.ItemPricePointData{
    },
    CurrentPricePoint:    models.ItemPricePointData{
    },
})
```

## models.CustomFieldValueChange

### Initialization Code

#### Example

```go
value := models.EventEventSpecificDataContainer.FromCustomFieldValueChange(models.CustomFieldValueChange{
    EventType:            "event_type2",
    MetafieldName:        "metafield_name6",
    MetafieldId:          78,
    OldValue:             models.ToPointer("old_value2"),
    NewValue:             models.ToPointer("new_value8"),
    ResourceType:         "resource_type2",
    ResourceId:           74,
})
```

## models.ChjsTokenizationSuccess

### Initialization Code

#### Example

```go
value := models.EventEventSpecificDataContainer.FromChjsTokenizationSuccess(models.ChjsTokenizationSuccess{
    PaymentProfile:       models.TokenizedPaymentProfile{
        Id:                   44,
    },
})
```

## models.ChjsTokenizationFailure

### Initialization Code

#### Example

```go
value := models.EventEventSpecificDataContainer.FromChjsTokenizationFailure(models.ChjsTokenizationFailure{
    Errors:               "errors2",
})
```

