
# List Coupons for Product Family Input

Input structure for the method ListCouponsForProductFamily

## Structure

`ListCouponsForProductFamilyInput`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ProductFamilyId` | `int` | Required | The Advanced Billing id of the product family to which the coupon belongs |
| `Page` | `*int` | Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |
| `PerPage` | `*int` | Optional | This parameter indicates how many records to fetch in each request. Default value is 30. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`.<br><br>**Default**: `30`<br><br>**Constraints**: `<= 200` |
| `Filter` | [`*models.ListCouponsFilter`](../../doc/models/list-coupons-filter.md) | Optional | Filter to use for List Coupons operations |
| `CurrencyPrices` | `*bool` | Optional | When fetching coupons, if you have defined multiple currencies at the site level, you can optionally pass the `?currency_prices=true` query param to include an array of currency price data in the response. Use in query `currency_prices=true`. |

## Example

```go
package main

import (
    "log"
    "time"
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    parseTime := func(layout, value string, errCallback func(error)) time.Time {
        dateTime, err := time.Parse(layout, value)
        if err != nil {
            errCallback(err) 
       }
        return dateTime
    }
    listCouponsForProductFamilyInput := models.ListCouponsForProductFamilyInput{
        ProductFamilyId:      140,
        Page:                 models.ToPointer(1),
        PerPage:              models.ToPointer(50),
        Filter:               models.ToPointer(models.ListCouponsFilter{
            DateField:            models.ToPointer(models.BasicDateField_UPDATEDAT),
            StartDate:            models.ToPointer(parseTime(models.DEFAULT_DATE, "2011-12-17", func(err error) { log.Fatalln(err) })),
            EndDate:              models.ToPointer(parseTime(models.DEFAULT_DATE, "2011-12-15", func(err error) { log.Fatalln(err) })),
            StartDatetime:        models.ToPointer(parseTime(time.RFC3339, "2011-12-19T10:15:30+01:00", func(err error) { log.Fatalln(err) })),
            EndDatetime:          models.ToPointer(parseTime(time.RFC3339, "2019-06-07T17:20:06Z", func(err error) { log.Fatalln(err) })),
            Ids:                  []int{
                1,
                2,
                3,
            },
            Codes:                []string{
                "free",
                "free_trial",
            },
        }),
        CurrencyPrices:       models.ToPointer(true),
    }

}
```

