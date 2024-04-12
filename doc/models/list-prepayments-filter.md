
# List Prepayments Filter

## Structure

`ListPrepaymentsFilter`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DateField` | [`*models.ListPrepaymentDateField`](../../doc/models/list-prepayment-date-field.md) | Optional | The type of filter you would like to apply to your search. `created_at` - Time when prepayment was created. `application_at` - Time when prepayment was applied to invoice. Use in query `filter[date_field]=created_at`. |
| `StartDate` | `*time.Time` | Optional | The start date (format YYYY-MM-DD) with which to filter the date_field. Returns prepayments with a timestamp at or after midnight (12:00:00 AM) in your site's time zone on the date specified. Use in query: `filter[start_date]=2011-12-15`. |
| `EndDate` | `*time.Time` | Optional | The end date (format YYYY-MM-DD) with which to filter the date_field. Returns prepayments with a timestamp up to and including 11:59:59PM in your site's time zone on the date specified. Use in query: `filter[end_date]=2011-12-15`. |

## Example (as JSON)

```json
{
  "date_field": "created_at",
  "start_date": "2024-01-01",
  "end_date": "2024-01-31"
}
```

