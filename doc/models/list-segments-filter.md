
# List Segments Filter

## Structure

`ListSegmentsFilter`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SegmentProperty1Value` | `*string` | Optional | The value passed here would be used to filter segments. Pass a value related to `segment_property_1` on attached Metric. If empty string is passed, this filter would be rejected. Use in query `filter[segment_property_1_value]=EU`. |
| `SegmentProperty2Value` | `*string` | Optional | The value passed here would be used to filter segments. Pass a value related to `segment_property_2` on attached Metric. If empty string is passed, this filter would be rejected. |
| `SegmentProperty3Value` | `*string` | Optional | The value passed here would be used to filter segments. Pass a value related to `segment_property_3` on attached Metric. If empty string is passed, this filter would be rejected. |
| `SegmentProperty4Value` | `*string` | Optional | The value passed here would be used to filter segments. Pass a value related to `segment_property_4` on attached Metric. If empty string is passed, this filter would be rejected. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    listSegmentsFilter := models.ListSegmentsFilter{
        SegmentProperty1Value: models.ToPointer("EU"),
        SegmentProperty2Value: models.ToPointer("segment_property_2_value6"),
        SegmentProperty3Value: models.ToPointer("segment_property_3_value6"),
        SegmentProperty4Value: models.ToPointer("segment_property_4_value8"),
    }

}
```

