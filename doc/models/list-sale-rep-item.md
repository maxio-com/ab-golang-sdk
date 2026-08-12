
# List Sale Rep Item

## Structure

`ListSaleRepItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `FullName` | `*string` | Optional | - |
| `SubscriptionsCount` | `*int` | Optional | - |
| `MrrData` | [`map[string]models.SaleRepItemMrr`](../../doc/models/sale-rep-item-mrr.md) | Optional | - |
| `TestMode` | `*bool` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    listSaleRepItem := models.ListSaleRepItem{
        Id:                   models.ToPointer(54),
        FullName:             models.ToPointer("full_name2"),
        SubscriptionsCount:   models.ToPointer(126),
        MrrData:              map[string]models.SaleRepItemMrr{
            "november_2019": models.SaleRepItemMrr{
                Mrr:                  models.ToPointer("$0.00"),
                Usage:                models.ToPointer("$0.00"),
                Recurring:            models.ToPointer("$0.00"),
            },
            "december_2019": models.SaleRepItemMrr{
                Mrr:                  models.ToPointer("$0.00"),
                Usage:                models.ToPointer("$0.00"),
                Recurring:            models.ToPointer("$0.00"),
            },
            "january_2020": models.SaleRepItemMrr{
                Mrr:                  models.ToPointer("$400.00"),
                Usage:                models.ToPointer("$0.00"),
                Recurring:            models.ToPointer("$400.00"),
            },
        },
        TestMode:             models.ToPointer(false),
    }

}
```

