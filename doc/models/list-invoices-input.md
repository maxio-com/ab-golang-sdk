
# List Invoices Input

Input structure for the method ListInvoices

## Structure

`ListInvoicesInput`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `StartDate` | `*string` | Optional | The start date (format YYYY-MM-DD) with which to filter the date_field. Returns invoices with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified. |
| `EndDate` | `*string` | Optional | The end date (format YYYY-MM-DD) with which to filter the date_field. Returns invoices with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified. |
| `Status` | [`*models.InvoiceStatus`](../../doc/models/invoice-status.md) | Optional | The current status of the invoice.  Allowed Values: draft, open, paid, pending, voided |
| `SubscriptionId` | `*int` | Optional | The subscription's ID. |
| `SubscriptionGroupUid` | `*string` | Optional | The UID of the subscription group you want to fetch consolidated invoices for. This will return a paginated list of consolidated invoices for the specified group. |
| `ConsolidationLevel` | `*string` | Optional | The consolidation level of the invoice. Allowed Values: none, parent, child or comma-separated lists of thereof, e.g. none,parent. |
| `Page` | `*int` | Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |
| `PerPage` | `*int` | Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`.<br><br>**Default**: `20`<br><br>**Constraints**: `<= 200` |
| `Direction` | [`*models.Direction`](../../doc/models/direction.md) | Optional | The sort direction of the returned invoices.<br><br>**Default**: `"desc"` |
| `LineItems` | `*bool` | Optional | Include line items data<br><br>**Default**: `false` |
| `Discounts` | `*bool` | Optional | Include discounts data<br><br>**Default**: `false` |
| `Taxes` | `*bool` | Optional | Include taxes data<br><br>**Default**: `false` |
| `Credits` | `*bool` | Optional | Include credits data<br><br>**Default**: `false` |
| `Payments` | `*bool` | Optional | Include payments data<br><br>**Default**: `false` |
| `CustomFields` | `*bool` | Optional | Include custom fields data<br><br>**Default**: `false` |
| `Refunds` | `*bool` | Optional | Include refunds data<br><br>**Default**: `false` |
| `DateField` | [`*models.InvoiceDateField`](../../doc/models/invoice-date-field.md) | Optional | The type of filter you would like to apply to your search. Use in query `date_field=issue_date`.<br><br>**Default**: `"due_date"` |
| `StartDatetime` | `*string` | Optional | The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns invoices with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of start_date. Allowed to be used only along with date_field set to created_at or updated_at. |
| `EndDatetime` | `*string` | Optional | The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns invoices with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of end_date. Allowed to be used only along with date_field set to created_at or updated_at. |
| `CustomerIds` | `[]int` | Optional | Allows fetching invoices with matching customer id based on provided values. Use in query `customer_ids=1,2,3`. |
| `Number` | `[]string` | Optional | Allows fetching invoices with matching invoice number based on provided values. Use in query `number=1234,1235`. |
| `ProductIds` | `[]int` | Optional | Allows fetching invoices with matching line items product ids based on provided values. Use in query `product_ids=23,34`. |
| `Sort` | [`*models.InvoiceSortField`](../../doc/models/invoice-sort-field.md) | Optional | Allows specification of the order of the returned list. Use in query `sort=total_amount`.<br><br>**Default**: `"number"` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    listInvoicesInput := models.ListInvoicesInput{
        StartDate:            models.ToPointer("start_date6"),
        EndDate:              models.ToPointer("end_date0"),
        Status:               models.ToPointer(models.InvoiceStatus_VOIDED),
        SubscriptionId:       models.ToPointer(222),
        SubscriptionGroupUid: models.ToPointer("subscription_group_uid6"),
        ConsolidationLevel:   models.ToPointer("consolidation_level8"),
        Page:                 models.ToPointer(1),
        PerPage:              models.ToPointer(50),
        Direction:            models.ToPointer(models.Direction_DESC),
        LineItems:            models.ToPointer(false),
        Discounts:            models.ToPointer(false),
        Taxes:                models.ToPointer(false),
        Credits:              models.ToPointer(false),
        Payments:             models.ToPointer(false),
        CustomFields:         models.ToPointer(false),
        Refunds:              models.ToPointer(false),
        DateField:            models.ToPointer(models.InvoiceDateField_ISSUEDATE),
        StartDatetime:        models.ToPointer("start_datetime0"),
        EndDatetime:          models.ToPointer("end_datetime8"),
        CustomerIds:          []int{
            1,
            2,
            3,
        },
        Number:               []string{
            "1234",
            "1235",
        },
        ProductIds:           []int{
            23,
            34,
        },
        Sort:                 models.ToPointer(models.InvoiceSortField_TOTALAMOUNT),
    }

}
```

