
# Create Currency Prices Request

## Structure

`CreateCurrencyPricesRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CurrencyPrices` | [`[]models.CreateCurrencyPrice`](../../doc/models/create-currency-price.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createCurrencyPricesRequest := models.CreateCurrencyPricesRequest{
        CurrencyPrices:       []models.CreateCurrencyPrice{
            models.CreateCurrencyPrice{
                Currency:             models.ToPointer("currency8"),
                Price:                models.ToPointer(float64(233.74)),
                PriceId:              models.ToPointer(116),
            },
        },
    }

}
```

