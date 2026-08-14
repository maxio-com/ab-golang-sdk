
# Item Price Point Changed

## Structure

`ItemPricePointChanged`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ItemId` | `int` | Required | - |
| `ItemType` | `string` | Required | - |
| `ItemHandle` | `string` | Required | - |
| `ItemName` | `string` | Required | - |
| `PreviousPricePoint` | [`models.ItemPricePointData`](../../doc/models/item-price-point-data.md) | Required | - |
| `CurrentPricePoint` | [`models.ItemPricePointData`](../../doc/models/item-price-point-data.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    itemPricePointChanged := models.ItemPricePointChanged{
        ItemId:               30,
        ItemType:             "item_type6",
        ItemHandle:           "item_handle4",
        ItemName:             "item_name8",
        PreviousPricePoint:   models.ItemPricePointData{
            Id:                   models.ToPointer(216),
            Handle:               models.ToPointer("handle6"),
            Name:                 models.ToPointer("name0"),
        },
        CurrentPricePoint:    models.ItemPricePointData{
            Id:                   models.ToPointer(218),
            Handle:               models.ToPointer("handle6"),
            Name:                 models.ToPointer("name0"),
        },
    }

}
```

