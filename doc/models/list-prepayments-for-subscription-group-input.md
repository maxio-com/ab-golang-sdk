
# List Prepayments for Subscription Group Input

Input structure for the method ListPrepaymentsForSubscriptionGroup

## Structure

`ListPrepaymentsForSubscriptionGroupInput`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Uid` | `string` | Required | The uid of the subscription group |
| `Page` | `*int` | Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |
| `PerPage` | `*int` | Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`.<br><br>**Default**: `20`<br><br>**Constraints**: `<= 200` |
| `Filter` | [`*models.ListPrepaymentsFilter`](../../doc/models/list-prepayments-filter.md) | Optional | Filter to use for List Prepayments operations |

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
    listPrepaymentsForSubscriptionGroupInput := models.ListPrepaymentsForSubscriptionGroupInput{
        Uid:                  "uid0",
        Page:                 models.ToPointer(1),
        PerPage:              models.ToPointer(50),
        Filter:               models.ToPointer(models.ListPrepaymentsFilter{
            DateField:            models.ToPointer(models.ListPrepaymentDateField_CREATEDAT),
            StartDate:            models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
            EndDate:              models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-31", func(err error) { log.Fatalln(err) })),
        }),
    }

}
```

