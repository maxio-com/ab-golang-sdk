
# List Credit Notes Input

Input structure for the method ListCreditNotes

## Structure

`ListCreditNotesInput`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SubscriptionId` | `*int` | Optional | The subscription's Advanced Billing id |
| `DateField` | [`*models.CreditNoteDateField`](../../doc/models/credit-note-date-field.md) | Optional | The type of filter you would like to apply to your search. Use in query `date_field=issue_date`. If a date range is provided without an explicit `date_field`, it defaults to `issue_date`. If only `start_datetime`/`end_datetime` are provided without an explicit `date_field`, it defaults to `created_at` instead. An unrecognized `date_field` is ignored rather than raising an error. |
| `StartDate` | `*string` | Optional | The start date (format YYYY-MM-DD) with which to filter the date_field. Returns credit notes with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified. |
| `EndDate` | `*string` | Optional | The end date (format YYYY-MM-DD) with which to filter the date_field. Returns credit notes with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified. |
| `StartDatetime` | `*string` | Optional | The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns credit notes with a timestamp at or after exact time provided in query. If provided, this parameter will be used instead of start_date. If no timezone offset is included in the value, it is interpreted as UTC. Allowed to be used only along with date_field set to created_at or updated_at. |
| `EndDatetime` | `*string` | Optional | The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns credit notes with a timestamp at or before exact time provided in query. If provided, this parameter will be used instead of end_date. If no timezone offset is included in the value, it is interpreted as UTC. Allowed to be used only along with date_field set to created_at or updated_at. |
| `Page` | `*int` | Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |
| `PerPage` | `*int` | Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`.<br><br>**Default**: `20`<br><br>**Constraints**: `<= 200` |
| `Direction` | [`*models.Direction`](../../doc/models/direction.md) | Optional | The sort direction of the returned credit notes, sorted by sequence_number.<br><br>**Default**: `"desc"` |
| `LineItems` | `*bool` | Optional | Include line items data.<br><br>**Default**: `false` |
| `Discounts` | `*bool` | Optional | Include discounts data.<br><br>**Default**: `false` |
| `Taxes` | `*bool` | Optional | Include taxes data.<br><br>**Default**: `false` |
| `Refunds` | `*bool` | Optional | Include refunds data.<br><br>**Default**: `false` |
| `Applications` | `*bool` | Optional | Include applications data.<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    listCreditNotesInput := models.ListCreditNotesInput{
        SubscriptionId:       models.ToPointer(222),
        DateField:            models.ToPointer(models.CreditNoteDateField_ISSUEDATE),
        StartDate:            models.ToPointer("start_date6"),
        EndDate:              models.ToPointer("end_date0"),
        StartDatetime:        models.ToPointer("start_datetime0"),
        EndDatetime:          models.ToPointer("end_datetime8"),
        Page:                 models.ToPointer(1),
        PerPage:              models.ToPointer(50),
        Direction:            models.ToPointer(models.Direction_DESC),
        LineItems:            models.ToPointer(false),
        Discounts:            models.ToPointer(false),
        Taxes:                models.ToPointer(false),
        Refunds:              models.ToPointer(false),
        Applications:         models.ToPointer(false),
    }

}
```

