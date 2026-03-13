
# List Customers Input

Input structure for the method ListCustomers

## Structure

`ListCustomersInput`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Direction` | [`*models.SortingDirection`](../../doc/models/sorting-direction.md) | Optional | Direction to sort customers by time of creation |
| `Page` | `*int` | Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |
| `PerPage` | `*int` | Optional | This parameter indicates how many records to fetch in each request. Default value is 50. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`.<br><br>**Default**: `50`<br><br>**Constraints**: `<= 200` |
| `DateField` | [`*models.BasicDateField`](../../doc/models/basic-date-field.md) | Optional | The type of filter you would like to apply to your search.<br>Use in query: `date_field=created_at`. |
| `StartDate` | `*string` | Optional | The start date (format YYYY-MM-DD) with which to filter the date_field. Returns subscriptions with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified. |
| `EndDate` | `*string` | Optional | The end date (format YYYY-MM-DD) with which to filter the date_field. Returns subscriptions with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified. |
| `StartDatetime` | `*string` | Optional | The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns subscriptions with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of start_date. |
| `EndDatetime` | `*string` | Optional | The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns subscriptions with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of end_date. |
| `Q` | `*string` | Optional | A search query by which to filter customers (can be an email, an ID, a reference, organization) |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    listCustomersInput := models.ListCustomersInput{
        Direction:            models.ToPointer(models.SortingDirection_ASC),
        Page:                 models.ToPointer(1),
        PerPage:              models.ToPointer(30),
        DateField:            models.ToPointer(models.BasicDateField_UPDATEDAT),
        StartDate:            models.ToPointer("start_date6"),
        EndDate:              models.ToPointer("end_date0"),
        StartDatetime:        models.ToPointer("start_datetime0"),
        EndDatetime:          models.ToPointer("end_datetime8"),
        Q:                    models.ToPointer("q0"),
    }

}
```

