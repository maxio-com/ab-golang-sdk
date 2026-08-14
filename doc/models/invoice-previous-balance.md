
# Invoice Previous Balance

## Structure

`InvoicePreviousBalance`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CapturedAt` | `*time.Time` | Optional | - |
| `Invoices` | [`[]models.InvoiceBalanceItem`](../../doc/models/invoice-balance-item.md) | Optional | - |

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
    invoicePreviousBalance := models.InvoicePreviousBalance{
        CapturedAt:           models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
        Invoices:             []models.InvoiceBalanceItem{
            models.InvoiceBalanceItem{
                Uid:                  models.ToPointer("uid6"),
                Number:               models.ToPointer("number6"),
                OutstandingAmount:    models.ToPointer("outstanding_amount8"),
            },
            models.InvoiceBalanceItem{
                Uid:                  models.ToPointer("uid6"),
                Number:               models.ToPointer("number6"),
                OutstandingAmount:    models.ToPointer("outstanding_amount8"),
            },
            models.InvoiceBalanceItem{
                Uid:                  models.ToPointer("uid6"),
                Number:               models.ToPointer("number6"),
                OutstandingAmount:    models.ToPointer("outstanding_amount8"),
            },
        },
    }

}
```

