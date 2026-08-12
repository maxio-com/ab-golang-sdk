
# List Price Points Filter

## Structure

`ListPricePointsFilter`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DateField` | [`*models.BasicDateField`](../../doc/models/basic-date-field.md) | Optional | The type of filter you would like to apply to your search. Use in query: `filter[date_field]=created_at`. |
| `StartDate` | `*time.Time` | Optional | The start date (format YYYY-MM-DD) with which to filter the date_field. Returns price points with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified. |
| `EndDate` | `*time.Time` | Optional | The end date (format YYYY-MM-DD) with which to filter the date_field. Returns price points with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified. |
| `StartDatetime` | `*time.Time` | Optional | The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns price points with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of start_date. |
| `EndDatetime` | `*time.Time` | Optional | The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns price points with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of end_date. |
| `Type` | [`[]models.PricePointType`](../../doc/models/price-point-type.md) | Optional | Allows fetching price points with matching type. Use in query: `filter[type]=custom,catalog`. |
| `Ids` | `[]int` | Optional | Allows fetching price points with matching id based on provided values. Use in query: `filter[ids]=1,2,3`. |
| `ArchivedAt` | [`*models.IncludeNullOrNotNull`](../../doc/models/include-null-or-not-null.md) | Optional | Allows fetching price points only if archived_at is present or not. Use in query: `filter[archived_at]=not_null`. |

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
    listPricePointsFilter := models.ListPricePointsFilter{
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
    }

}
```

