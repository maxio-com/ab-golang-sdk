
# Component Currency Prices Response

## Structure

`ComponentCurrencyPricesResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CurrencyPrices` | [`[]models.ComponentCurrencyPrice`](../../doc/models/component-currency-price.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    componentCurrencyPricesResponse := models.ComponentCurrencyPricesResponse{
        CurrencyPrices:       []models.ComponentCurrencyPrice{
            models.ComponentCurrencyPrice{
                Id:                   models.ToPointer(50),
                Currency:             models.ToPointer("currency8"),
                Price:                models.ToPointer("price4"),
                FormattedPrice:       models.ToPointer("formatted_price6"),
                PriceId:              models.ToPointer(116),
            },
        },
    }

}
```

