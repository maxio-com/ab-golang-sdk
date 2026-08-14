
# Create Product Currency Prices Request

## Structure

`CreateProductCurrencyPricesRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CurrencyPrices` | [`[]models.CreateProductCurrencyPrice`](../../doc/models/create-product-currency-price.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createProductCurrencyPricesRequest := models.CreateProductCurrencyPricesRequest{
        CurrencyPrices:       []models.CreateProductCurrencyPrice{
            models.CreateProductCurrencyPrice{
                Currency:             "currency8",
                Price:                78,
                Role:                 models.CurrencyPriceRole_INITIAL,
            },
        },
    }

}
```

