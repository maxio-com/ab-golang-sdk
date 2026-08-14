
# Site Response

## Structure

`SiteResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Site` | [`models.Site`](../../doc/models/site.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    siteResponse := models.SiteResponse{
        Site:                 models.Site{
            Id:                                      models.ToPointer(64),
            Name:                                    models.ToPointer("name4"),
            Subdomain:                               models.ToPointer("subdomain0"),
            Currency:                                models.ToPointer("currency4"),
            SellerId:                                models.ToPointer(228),
        },
    }

}
```

