
# List Product Price Points Input

Input structure for the method ListProductPricePoints

## Structure

`ListProductPricePointsInput`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ProductId` | [`models.ListProductPricePointsInputProductId`](../../doc/models/containers/list-product-price-points-input-product-id.md) | Required | This is a container for one-of cases. |
| `Page` | `*int` | Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |
| `PerPage` | `*int` | Optional | This parameter indicates how many records to fetch in each request. Default value is 10. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br><br>**Default**: `10`<br><br>**Constraints**: `<= 200` |
| `CurrencyPrices` | `*bool` | Optional | When fetching a product's price points, if you have defined multiple currencies at the site level, you can optionally pass the ?currency_prices=true query param to include an array of currency price data in the response. If the product price point is set to use_site_exchange_rate: true, it will return pricing based on the current exchange rate. If the flag is set to false, it will return all of the defined prices for each currency. |
| `FilterType` | [`[]models.PricePointType`](../../doc/models/price-point-type.md) | Optional | Use in query: `filter[type]=catalog,default`. |
| `Archived` | `*bool` | Optional | Set to include archived price points in the response. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    listProductPricePointsInput := models.ListProductPricePointsInput{
        ProductId:            models.ListProductPricePointsInputProductIdContainer.FromNumber(124),
        Page:                 models.ToPointer(1),
        PerPage:              models.ToPointer(10),
        CurrencyPrices:       models.ToPointer(false),
        FilterType:           Liquid error: Value cannot be null. (Parameter 'key'),
        Archived:             models.ToPointer(false),
    }

}
```

