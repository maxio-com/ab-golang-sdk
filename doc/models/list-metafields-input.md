
# List Metafields Input

Input structure for the method ListMetafields

## Structure

`ListMetafieldsInput`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ResourceType` | [`models.ResourceType`](../../doc/models/resource-type.md) | Required | The resource type to which the metafields belong. |
| `Name` | `*string` | Optional | Filter by the name of the metafield. |
| `Page` | `*int` | Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |
| `PerPage` | `*int` | Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`.<br><br>**Default**: `20`<br><br>**Constraints**: `<= 200` |
| `Direction` | [`*models.SortingDirection`](../../doc/models/sorting-direction.md) | Optional | Controls the order in which results are returned.<br>Use in query `direction=asc`. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    listMetafieldsInput := models.ListMetafieldsInput{
        ResourceType:         models.ResourceType_SUBSCRIPTIONS,
        Name:                 models.ToPointer("name0"),
        Page:                 models.ToPointer(1),
        PerPage:              models.ToPointer(50),
        Direction:            models.ToPointer(models.SortingDirection_ASC),
    }

}
```

