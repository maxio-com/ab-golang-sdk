
# List Components Filter

## Structure

`ListComponentsFilter`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ids` | `[]int` | Optional | Allows fetching components with matching id based on provided value. Use in query `filter[ids]=1,2,3`.<br><br>**Constraints**: *Minimum Items*: `1` |
| `UseSiteExchangeRate` | `*bool` | Optional | Allows fetching components with matching use_site_exchange_rate based on provided value (refers to default price point). Use in query `filter[use_site_exchange_rate]=true`. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    listComponentsFilter := models.ListComponentsFilter{
        Ids:                  []int{
            1,
            2,
            3,
        },
        UseSiteExchangeRate:  models.ToPointer(false),
    }

}
```

