
# Prepaid Usage Allocation Detail

## Structure

`PrepaidUsageAllocationDetail`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllocationId` | `*int` | Optional | - |
| `ChargeId` | `*int` | Optional | - |
| `UsageQuantity` | `*int` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    prepaidUsageAllocationDetail := models.PrepaidUsageAllocationDetail{
        AllocationId:         models.ToPointer(144),
        ChargeId:             models.ToPointer(214),
        UsageQuantity:        models.ToPointer(140),
    }

}
```

