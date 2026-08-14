
# List Mrr Filter

## Structure

`ListMrrFilter`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SubscriptionIds` | `[]int` | Optional | Submit ids in order to limit results. Use in query: `filter[subscription_ids]=1,2,3`.<br><br>**Constraints**: *Minimum Items*: `1` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    listMrrFilter := models.ListMrrFilter{
        SubscriptionIds:      []int{
            1,
            2,
            3,
        },
    }

}
```

