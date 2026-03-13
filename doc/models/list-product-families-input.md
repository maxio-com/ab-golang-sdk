
# List Product Families Input

Input structure for the method ListProductFamilies

## Structure

`ListProductFamiliesInput`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DateField` | [`*models.BasicDateField`](../../doc/models/basic-date-field.md) | Optional | The type of filter you would like to apply to your search.<br>Use in query: `date_field=created_at`. |
| `StartDate` | `*time.Time` | Optional | The start date (format YYYY-MM-DD) with which to filter the date_field. Returns products with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified. |
| `EndDate` | `*time.Time` | Optional | The end date (format YYYY-MM-DD) with which to filter the date_field. Returns products with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified. |
| `StartDatetime` | `*time.Time` | Optional | The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns products with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of start_date. |
| `EndDatetime` | `*time.Time` | Optional | The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns products with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of end_date. |

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
    listProductFamiliesInput := models.ListProductFamiliesInput{
        DateField:            models.ToPointer(models.BasicDateField_UPDATEDAT),
        StartDate:            models.ToPointer(parseTime(models.DEFAULT_DATE, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
        EndDate:              models.ToPointer(parseTime(models.DEFAULT_DATE, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
        StartDatetime:        models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
        EndDatetime:          models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
    }

}
```

