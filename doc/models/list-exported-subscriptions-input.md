
# List Exported Subscriptions Input

Input structure for the method ListExportedSubscriptions

## Structure

`ListExportedSubscriptionsInput`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BatchId` | `string` | Required | Id of a Batch Job. |
| `PerPage` | `*int` | Optional | This parameter indicates how many records to fetch in each request.<br>Default value is 100.<br>The maximum allowed values is 10000; any per_page value over 10000 will be changed to 10000.<br><br>**Default**: `100`<br><br>**Constraints**: `>= 1`, `<= 10000` |
| `Page` | `*int` | Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    listExportedSubscriptionsInput := models.ListExportedSubscriptionsInput{
        BatchId:              "batch_id8",
        PerPage:              models.ToPointer(100),
        Page:                 models.ToPointer(1),
    }

}
```

