
# Refund Invoice Request Refund

## Class Name

`RefundInvoiceRequestRefund`

## Cases

| Type | Factory Method |
|  --- | --- |
| [`models.RefundInvoice`](../../../doc/models/refund-invoice.md) | models.RefundInvoiceRequestRefundContainer.FromRefundInvoice(models.RefundInvoice refundInvoice) |
| [`models.RefundConsolidatedInvoice`](../../../doc/models/refund-consolidated-invoice.md) | models.RefundInvoiceRequestRefundContainer.FromRefundConsolidatedInvoice(models.RefundConsolidatedInvoice refundConsolidatedInvoice) |

## models.RefundInvoice

### Initialization Code

#### Example

```go
value := models.RefundInvoiceRequestRefundContainer.FromRefundInvoice(models.RefundInvoice{
    Amount:               "amount8",
    Memo:                 "memo0",
    PaymentId:            0,
})
```

## models.RefundConsolidatedInvoice

### Initialization Code

#### Example

```go
value := models.RefundInvoiceRequestRefundContainer.FromRefundConsolidatedInvoice(models.RefundConsolidatedInvoice{
    Memo:                 "memo0",
    PaymentId:            46,
    SegmentUids:          models.RefundConsolidatedInvoiceSegmentUidsContainer.FromArrayOfString([]string{
        "String0",
        "String1",
    }),
})
```

