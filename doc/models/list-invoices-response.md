
# List Invoices Response

## Structure

`ListInvoicesResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Invoices` | [`[]models.Invoice`](../../doc/models/invoice.md) | Required | - |

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
    listInvoicesResponse := models.ListInvoicesResponse{
        Invoices:             []models.Invoice{
            models.Invoice{
                Id:                         models.ToPointer(int64(196)),
                Uid:                        models.ToPointer("uid6"),
                SiteId:                     models.ToPointer(122),
                CustomerId:                 models.ToPointer(234),
                SubscriptionId:             models.ToPointer(50),
                IssueDate:                  models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
                DueDate:                    models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
                PaidDate:                   models.NewOptional(models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) }))),
                PublicUrlExpiresOn:         models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-21", func(err error) { log.Fatalln(err) })),
            },
        },
    }

}
```

