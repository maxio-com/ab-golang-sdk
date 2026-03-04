
# List Component Price Points Input

Input structure for the method ListComponentPricePoints

## Structure

`ListComponentPricePointsInput`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ComponentId` | `int` | Required | The Advanced Billing id of the component |
| `CurrencyPrices` | `*bool` | Optional | Include an array of currency price data |
| `Page` | `*int` | Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |
| `PerPage` | `*int` | Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`.<br><br>**Default**: `20`<br><br>**Constraints**: `<= 200` |
| `FilterType` | [`[]models.PricePointType`](../../doc/models/price-point-type.md) | Optional | Use in query: `filter[type]=catalog,default`. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    listComponentPricePointsInput := models.ListComponentPricePointsInput{
        ComponentId:          222,
        CurrencyPrices:       models.ToPointer(false),
        Page:                 models.ToPointer(1),
        PerPage:              models.ToPointer(50),
        FilterType:           Liquid error: Value cannot be null. (Parameter 'key'),
    }

}
```

