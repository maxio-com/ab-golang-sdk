
# List All Component Price Points Input

Input structure for the method ListAllComponentPricePoints

## Structure

`ListAllComponentPricePointsInput`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Include` | [`*models.ListComponentsPricePointsInclude`](../../doc/models/list-components-price-points-include.md) | Optional | Allows including additional data in the response. Use in query: `include=currency_prices`. |
| `Page` | `*int` | Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |
| `PerPage` | `*int` | Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`.<br><br>**Default**: `20`<br><br>**Constraints**: `<= 200` |
| `Direction` | [`*models.SortingDirection`](../../doc/models/sorting-direction.md) | Optional | Controls the order in which results are returned.<br>Use in query `direction=asc`. |
| `Filter` | [`*models.ListPricePointsFilter`](../../doc/models/list-price-points-filter.md) | Optional | Filter to use for List PricePoints operations |

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
    listAllComponentPricePointsInput := models.ListAllComponentPricePointsInput{
        Include:              models.ToPointer(models.ListComponentsPricePointsInclude_CURRENCYPRICES),
        Page:                 models.ToPointer(1),
        PerPage:              models.ToPointer(50),
        Direction:            models.ToPointer(models.SortingDirection_ASC),
        Filter:               models.ToPointer(models.ListPricePointsFilter{
            DateField:            models.ToPointer(models.BasicDateField_UPDATEDAT),
            StartDate:            models.ToPointer(parseTime(models.DEFAULT_DATE, "2011-12-17", func(err error) { log.Fatalln(err) })),
            EndDate:              models.ToPointer(parseTime(models.DEFAULT_DATE, "2011-12-15", func(err error) { log.Fatalln(err) })),
            StartDatetime:        models.ToPointer(parseTime(time.RFC3339, "2011-12-19T10:15:30+01:00", func(err error) { log.Fatalln(err) })),
            EndDatetime:          models.ToPointer(parseTime(time.RFC3339, "2019-06-07T17:20:06Z", func(err error) { log.Fatalln(err) })),
            Type:                 []models.PricePointType{
                models.PricePointType_CATALOG,
                models.PricePointType_ENUMDEFAULT,
                models.PricePointType_CUSTOM,
            },
            Ids:                  []int{
                1,
                2,
                3,
            },
        }),
    }

}
```

