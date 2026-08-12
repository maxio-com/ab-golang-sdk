
# List MRR Response

## Structure

`ListMRRResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mrr` | [`models.ListMRRResponseResult`](../../doc/models/list-mrr-response-result.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    listMRRResponse := models.ListMRRResponse{
        Mrr:                  models.ListMRRResponseResult{
            Page:                 models.ToPointer(30),
            PerPage:              models.ToPointer(198),
            TotalPages:           models.ToPointer(92),
            TotalEntries:         models.ToPointer(188),
            Currency:             models.ToPointer("currency4"),
        },
    }

}
```

