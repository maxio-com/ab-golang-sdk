
# List Sales Commission Settings Input

Input structure for the method ListSalesCommissionSettings

## Structure

`ListSalesCommissionSettingsInput`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SellerId` | `string` | Required | The Chargify id of your seller account |
| `Authorization` | `*string` | Optional | For authorization use user API key. See details [here](https://developers.chargify.com/docs/developer-docs/ZG9jOjMyNzk5NTg0-2020-04-20-new-api-authentication).<br><br>**Default**: `"Bearer <<apiKey>>"` |
| `LiveMode` | `*bool` | Optional | This parameter indicates if records should be fetched from live mode sites. Default value is true. |
| `Page` | `*int` | Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |
| `PerPage` | `*int` | Optional | This parameter indicates how many records to fetch in each request. Default value is 100.<br><br>**Default**: `100` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    listSalesCommissionSettingsInput := models.ListSalesCommissionSettingsInput{
        SellerId:             "seller_id8",
        Authorization:        models.ToPointer("Bearer <<apiKey>>"),
        LiveMode:             models.ToPointer(false),
        Page:                 models.ToPointer(1),
        PerPage:              models.ToPointer(100),
    }

}
```

