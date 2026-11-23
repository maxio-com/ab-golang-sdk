
# Segment Price Point Error Details

Price-point validation error details for segment writes.

## Structure

`SegmentPricePointErrorDetails`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PricePoint` | `[]string` | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    segmentPricePointErrorDetails := models.SegmentPricePointErrorDetails{
        PricePoint:           []string{
            "price_point2",
        },
    }

}
```

