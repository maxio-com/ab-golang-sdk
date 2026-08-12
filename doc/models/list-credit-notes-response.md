
# List Credit Notes Response

## Structure

`ListCreditNotesResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreditNotes` | [`[]models.CreditNote`](../../doc/models/credit-note.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    listCreditNotesResponse := models.ListCreditNotesResponse{
        CreditNotes:          []models.CreditNote{
            models.CreditNote{
                Uid:                  models.ToPointer("uid2"),
                SiteId:               models.ToPointer(112),
                CustomerId:           models.ToPointer(224),
                SubscriptionId:       models.ToPointer(40),
                Number:               models.ToPointer("number0"),
            },
        },
    }

}
```

