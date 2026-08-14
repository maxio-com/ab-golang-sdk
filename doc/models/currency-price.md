
# Currency Price

## Structure

`CurrencyPrice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `Currency` | `*string` | Optional | - |
| `Price` | `*float64` | Optional | - |
| `FormattedPrice` | `*string` | Optional | - |
| `PriceId` | `*int` | Optional | - |
| `PricePointId` | `*int` | Optional | - |
| `ProductPricePointId` | `*int` | Optional | - |
| `Role` | [`*models.CurrencyPriceRole`](../../doc/models/currency-price-role.md) | Optional | Role for the price. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    currencyPrice := models.CurrencyPrice{
        Id:                   models.ToPointer(208),
        Currency:             models.ToPointer("currency4"),
        Price:                models.ToPointer(float64(70.88)),
        FormattedPrice:       models.ToPointer("formatted_price2"),
        PriceId:              models.ToPointer(214),
    }

}
```

