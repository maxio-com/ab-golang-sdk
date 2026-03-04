
# List Credit Notes Input

Input structure for the method ListCreditNotes

## Structure

`ListCreditNotesInput`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SubscriptionId` | `*int` | Optional | The subscription's Advanced Billing id |
| `Page` | `*int` | Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |
| `PerPage` | `*int` | Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`.<br><br>**Default**: `20`<br><br>**Constraints**: `<= 200` |
| `LineItems` | `*bool` | Optional | Include line items data<br><br>**Default**: `false` |
| `Discounts` | `*bool` | Optional | Include discounts data<br><br>**Default**: `false` |
| `Taxes` | `*bool` | Optional | Include taxes data<br><br>**Default**: `false` |
| `Refunds` | `*bool` | Optional | Include refunds data<br><br>**Default**: `false` |
| `Applications` | `*bool` | Optional | Include applications data<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    listCreditNotesInput := models.ListCreditNotesInput{
        SubscriptionId:       models.ToPointer(222),
        Page:                 models.ToPointer(1),
        PerPage:              models.ToPointer(50),
        LineItems:            models.ToPointer(false),
        Discounts:            models.ToPointer(false),
        Taxes:                models.ToPointer(false),
        Refunds:              models.ToPointer(false),
        Applications:         models.ToPointer(false),
    }

}
```

