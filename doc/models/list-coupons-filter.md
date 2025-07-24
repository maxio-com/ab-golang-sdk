
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
| `UseSiteExchangeRate` | `*bool` | Optional | Allows fetching coupons with matching use_site_exchange_rate based on provided value. Use in query `filter[use_site_exchange_rate]=true`. |

## Example (as JSON)

```json
{
  "start_date": "2011-12-17",
  "end_date": "2011-12-15",
  "start_datetime": "12/19/2011 09:15:30",
  "end_datetime": "06/07/2019 17:20:06",
  "ids": [
    1,
    2,
    3
  ],
  "codes": [
    "free",
    "free_trial"
  ],
  "date_field": "updated_at"
}
```

