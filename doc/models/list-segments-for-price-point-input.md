
# List Segments for Price Point Input

Input structure for the method ListSegmentsForPricePoint

## Structure

`ListSegmentsForPricePointInput`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ComponentId` | `string` | Required | ID or Handle for the Component |
| `PricePointId` | `string` | Required | ID or Handle for the Price Point belonging to the Component |
| `Page` | `*int` | Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |
| `PerPage` | `*int` | Optional | This parameter indicates how many records to fetch in each request. Default value is 30. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`.<br><br>**Default**: `30`<br><br>**Constraints**: `<= 200` |
| `Filter` | [`*models.ListSegmentsFilter`](../../doc/models/list-segments-filter.md) | Optional | Filter to use for List Segments for a Price Point operation |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    listSegmentsForPricePointInput := models.ListSegmentsForPricePointInput{
        ComponentId:          "component_id8",
        PricePointId:         "price_point_id8",
        Page:                 models.ToPointer(1),
        PerPage:              models.ToPointer(50),
        Filter:               models.ToPointer(models.ListSegmentsFilter{
            SegmentProperty1Value: models.ToPointer("EU"),
            SegmentProperty2Value: models.ToPointer("segment_property_2_value2"),
            SegmentProperty3Value: models.ToPointer("segment_property_3_value0"),
            SegmentProperty4Value: models.ToPointer("segment_property_4_value4"),
        }),
    }

}
```

