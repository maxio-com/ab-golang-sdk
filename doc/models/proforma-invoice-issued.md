
# Proforma Invoice Issued

## Structure

`ProformaInvoiceIssued`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Uid` | `string` | Required | - |
| `Number` | `string` | Required | - |
| `Role` | `string` | Required | - |
| `DeliveryDate` | `time.Time` | Required | - |
| `CreatedAt` | `time.Time` | Required | - |
| `DueAmount` | `string` | Required | - |
| `PaidAmount` | `string` | Required | - |
| `TaxAmount` | `string` | Required | - |
| `TotalAmount` | `string` | Required | - |
| `ProductName` | `string` | Required | - |
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
    proformaInvoiceIssued := models.ProformaInvoiceIssued{
        Uid:                  "uid6",
        Number:               "number4",
        Role:                 "role0",
        DeliveryDate:         parseTime(models.DEFAULT_DATE, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
        CreatedAt:            parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
        DueAmount:            "due_amount8",
        PaidAmount:           "paid_amount8",
        TaxAmount:            "tax_amount0",
        TotalAmount:          "total_amount2",
        ProductName:          "product_name2",
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

