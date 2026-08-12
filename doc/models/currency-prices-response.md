
# Currency Prices Response

## Structure

`CurrencyPricesResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CurrencyPrices` | [`[]models.CurrencyPrice`](../../doc/models/currency-price.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    currencyPricesResponse := models.CurrencyPricesResponse{
        CurrencyPrices:       []models.CurrencyPrice{
            models.CurrencyPrice{
                Id:                   models.ToPointer(50),
                Currency:             models.ToPointer("currency8"),
                Price:                models.ToPointer(float64(233.74)),
                FormattedPrice:       models.ToPointer("formatted_price6"),
                PriceId:              models.ToPointer(116),
            },
        },
    }

}
```

