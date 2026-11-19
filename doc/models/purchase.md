
# Purchase

## Structure

`Purchase`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Quantity` | [`models.PurchaseQuantity`](../../doc/models/containers/purchase-quantity.md) | Required | This is a container for one-of cases. |
| `PricePointId` | `*int` | Optional | The price point to purchase at. Must belong to the wallet's billing product. Defaults to the `price_point_id` set on the wallet. |
| `Memo` | `models.Optional[string]` | Optional | (Optional) A note recorded on the resulting grant's ledger row.<br><br>**Constraints**: *Maximum Length*: `255` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    purchase := models.Purchase{
        Quantity:             models.PurchaseQuantityContainer.FromString("String5"),
        PricePointId:         models.ToPointer(20),
        Memo:                 models.NewOptional(models.ToPointer("memo2")),
    }

}
```

