
# Invoice Issued

## Structure

`InvoiceIssued`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Uid` | `string` | Required | - |
| `Number` | `string` | Required | - |
| `Role` | `string` | Required | - |
| `DueDate` | `*time.Time` | Required | - |
| `IssueDate` | `string` | Required | Invoice issue date. Can be an empty string if value is missing. |
| `PaidDate` | `string` | Required | Paid date. Can be an empty string if value is missing. |
| `DueAmount` | `string` | Required | - |
| `PaidAmount` | `string` | Required | - |
| `TaxAmount` | `string` | Required | - |
| `RefundAmount` | `string` | Required | - |
| `TotalAmount` | `string` | Required | - |
| `StatusAmount` | `string` | Required | - |
| `ProductName` | `string` | Required | - |
| `ConsolidationLevel` | `string` | Required | - |
| `LineItems` | [`[]models.InvoiceLineItemEventData`](../../doc/models/invoice-line-item-event-data.md) | Required | - |

## Example

```go
package main

import (
    "log"
    "time"
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    parseTime := func(layout, value string, errCallback func(error)) time.Time {
        dateTime, err := time.Parse(layout, value)
        if err != nil {
            errCallback(err) 
       }
        return dateTime
    }
    invoiceIssued := models.InvoiceIssued{
        Uid:                  "uid8",
        Number:               "number4",
        Role:                 "role8",
        DueDate:              models.ToPointer(parseTime(models.DEFAULT_DATE, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
        IssueDate:            "issue_date4",
        PaidDate:             "paid_date8",
        DueAmount:            "due_amount0",
        PaidAmount:           "paid_amount0",
        TaxAmount:            "tax_amount8",
        RefundAmount:         "refund_amount6",
        TotalAmount:          "total_amount4",
        StatusAmount:         "status_amount8",
        ProductName:          "product_name4",
        ConsolidationLevel:   "consolidation_level0",
        LineItems:            []models.InvoiceLineItemEventData{
            models.InvoiceLineItemEventData{
                Uid:                   models.ToPointer("uid8"),
                Title:                 models.ToPointer("title4"),
                Description:           models.ToPointer("description8"),
                Quantity:              models.ToPointer(102),
                QuantityDelta:         models.NewOptional(models.ToPointer(204)),
            },
        },
    }

}
```

