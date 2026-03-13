
# List Invoice Events Input

Input structure for the method ListInvoiceEvents

## Structure

`ListInvoiceEventsInput`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SinceDate` | `*string` | Optional | The timestamp in a format `YYYY-MM-DD T HH:MM:SS Z`, or `YYYY-MM-DD`(in this case, it returns data from the beginning of the day). of the event from which you want to start the search. All the events before the `since_date` timestamp are not returned in the response. |
| `SinceId` | `*int64` | Optional | The ID of the event from which you want to start the search(ID is not included. e.g. if ID is set to 2, then all events with ID 3 and more will be shown) This parameter is not used if since_date is defined. |
| `Page` | `*int` | Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |
| `PerPage` | `*int` | Optional | This parameter indicates how many records to fetch in each request. Default value is 100. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br><br>**Default**: `100` |
| `InvoiceUid` | `*string` | Optional | Providing an invoice_uid allows for scoping of the invoice events to a single invoice or credit note. |
| `WithChangeInvoiceStatus` | `*string` | Optional | Use this parameter if you want to fetch also invoice events with change_invoice_status type. |
| `EventTypes` | [`[]models.InvoiceEventType`](../../doc/models/invoice-event-type.md) | Optional | Filter results by event_type. Supply a comma separated list of event types (listed above). Use in query: `event_types=void_invoice,void_remainder`. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    listInvoiceEventsInput := models.ListInvoiceEventsInput{
        SinceDate:               models.ToPointer("since_date4"),
        SinceId:                 models.ToPointer(int64(104)),
        Page:                    models.ToPointer(1),
        PerPage:                 models.ToPointer(100),
        InvoiceUid:              models.ToPointer("invoice_uid0"),
        WithChangeInvoiceStatus: models.ToPointer("with_change_invoice_status2"),
        EventTypes:              []models.InvoiceEventType{
            models.InvoiceEventType_APPLYPAYMENT,
            models.InvoiceEventType_APPLYDEBITNOTE,
        },
    }

}
```

