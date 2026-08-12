
# List Public Keys Response

## Structure

`ListPublicKeysResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ChargifyJsKeys` | [`[]models.PublicKey`](../../doc/models/public-key.md) | Optional | - |
| `Meta` | [`*models.ListPublicKeysMeta`](../../doc/models/list-public-keys-meta.md) | Optional | - |

## Example

```go
package main

import (
    "log"
    "time"
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    parseTime := func(layout, value string, errCallback func(error)) time.Time {
        dateTime, err := time.Parse(layout, value)
        if err != nil {
            errCallback(err) 
       }
        return dateTime
    }
    listPublicKeysResponse := models.ListPublicKeysResponse{
        ChargifyJsKeys:       []models.PublicKey{
            models.PublicKey{
                PublicKey:             models.ToPointer("public_key8"),
                RequiresSecurityToken: models.ToPointer(false),
                CreatedAt:             models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
            },
        },
        Meta:                 models.ToPointer(models.ListPublicKeysMeta{
            TotalCount:           models.ToPointer(150),
            CurrentPage:          models.ToPointer(126),
            TotalPages:           models.ToPointer(138),
            PerPage:              models.ToPointer(152),
        }),
    }

}
```

