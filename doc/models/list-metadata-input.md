
# List Metadata Input

Input structure for the method ListMetadata

## Structure

`ListMetadataInput`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ResourceType` | [`models.ResourceType`](../../doc/models/resource-type.md) | Required | The resource type to which the metafields belong. |
| `ResourceId` | `int` | Required | The Advanced Billing id of the customer or the subscription for which the metadata applies |
| `Page` | `*int` | Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |
| `PerPage` | `*int` | Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`.<br><br>**Default**: `20`<br><br>**Constraints**: `<= 200` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    listMetadataInput := models.ListMetadataInput{
        ResourceType:         models.ResourceType_SUBSCRIPTIONS,
        ResourceId:           60,
        Page:                 models.ToPointer(1),
        PerPage:              models.ToPointer(50),
    }

}
```

