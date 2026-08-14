
# List Coupons Filter

## Structure

`ListCouponsFilter`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DateField` | [`*models.BasicDateField`](../../doc/models/basic-date-field.md) | Optional | The type of filter you would like to apply to your search. Use in query `filter[date_field]=created_at`. |
| `StartDate` | `*time.Time` | Optional | The start date (format YYYY-MM-DD) with which to filter the date_field. Returns coupons with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified. Use in query `filter[start_date]=2011-12-17`. |
| `EndDate` | `*time.Time` | Optional | The end date (format YYYY-MM-DD) with which to filter the date_field. Returns coupons with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified. Use in query `filter[end_date]=2011-12-15`. |
| `StartDatetime` | `*time.Time` | Optional | The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns coupons with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of start_date. Use in query `filter[start_datetime]=2011-12-19T10:15:30+01:00`. |
| `EndDatetime` | `*time.Time` | Optional | The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns coupons with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of end_date. Use in query `filter[end_datetime]=2011-12-1T10:15:30+01:00`. |
| `Ids` | `[]int` | Optional | Allows fetching coupons with matching id based on provided values. Use in query `filter[ids]=1,2,3`.<br><br>**Constraints**: *Minimum Items*: `1` |
| `Codes` | `[]string` | Optional | Allows fetching coupons with matching codes based on provided values. Use in query `filter[codes]=free,free_trial`. |
| `UseSiteExchangeRate` | `*bool` | Optional | If true, restricts the list to coupons whose pricing is recalculated from the site’s current exchange rates, so their currency_prices array contains on-the-fly conversions rather than stored price records. If false, restricts the list to coupons that have manually defined amounts for each currency, ensuring the response includes the saved currency_prices entries instead of exchange-rate-derived values. Use in query `filter[use_site_exchange_rate]=true`. |
| `IncludeArchived` | `*bool` | Optional | Controls returning archived coupons. |

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
    listCouponsFilter := models.ListCouponsFilter{
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
    }

}
```

