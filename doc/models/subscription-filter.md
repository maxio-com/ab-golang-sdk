
# Subscription Filter

Nested filter used for List Subscription Components For Site Filter

## Structure

`SubscriptionFilter`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `States` | [`[]models.SubscriptionStateFilter`](../../doc/models/subscription-state-filter.md) | Optional | Allows fetching components allocations that belong to the subscription with matching states based on provided values. To use this filter you also have to include the following param in the request `include=subscription`. Use in query `filter[subscription][states]=active,canceled&include=subscription`.<br><br>**Constraints**: *Minimum Items*: `1` |
| `DateField` | [`*models.SubscriptionListDateField`](../../doc/models/subscription-list-date-field.md) | Optional | The type of filter you'd like to apply to your search. To use this filter you also have to include the following param in the request `include=subscription`. |
| `StartDate` | `*time.Time` | Optional | The start date (format YYYY-MM-DD) with which to filter the date_field. Returns components that belong to the subscription with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified. To use this filter you also have to include the following param in the request `include=subscription`. |
| `EndDate` | `*time.Time` | Optional | The end date (format YYYY-MM-DD) with which to filter the date_field. Returns components that belong to the subscription with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified. To use this filter you also have to include the following param in the request `include=subscription`. |
| `StartDatetime` | `*time.Time` | Optional | The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns components that belong to the subscription with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site''s time zone will be used. If provided, this parameter will be used instead of start_date. To use this filter you also have to include the following param in the request `include=subscription`. |
| `EndDatetime` | `*time.Time` | Optional | The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns components that belong to the subscription with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site''s time zone will be used. If provided, this parameter will be used instead of end_date. To use this filter you also have to include the following param in the request `include=subscription`. |

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
    subscriptionFilter := models.SubscriptionFilter{
        States:               []models.SubscriptionStateFilter{
            models.SubscriptionStateFilter_ACTIVE,
            models.SubscriptionStateFilter_CANCELED,
        },
        DateField:            models.ToPointer(models.SubscriptionListDateField_UPDATEDAT),
        StartDate:            models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-17", func(err error) { log.Fatalln(err) })),
        EndDate:              models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-31", func(err error) { log.Fatalln(err) })),
        StartDatetime:        models.ToPointer(parseTime(time.RFC3339, "2024-01-17T10:15:30+01:00", func(err error) { log.Fatalln(err) })),
        EndDatetime:          models.ToPointer(parseTime(time.RFC3339, "2024-01-17T17:20:06Z", func(err error) { log.Fatalln(err) })),
    }

}
```

