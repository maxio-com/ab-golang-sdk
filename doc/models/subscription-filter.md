
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

## Example (as JSON)

```json
{
  "states": [
    "active",
    "canceled"
  ],
  "start_date": "2024-01-17",
  "end_date": "2024-01-31",
  "start_datetime": "01/17/2024 09:15:30",
  "end_datetime": "01/17/2024 17:20:06",
  "date_field": "updated_at"
}
```

