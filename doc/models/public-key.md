
# Public Key

## Structure

`PublicKey`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PublicKey` | `*string` | Optional | - |
| `RequiresSecurityToken` | `*bool` | Optional | - |
| `CreatedAt` | `*time.Time` | Optional | - |

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
    publicKey := models.PublicKey{
        PublicKey:             models.ToPointer("public_key2"),
        RequiresSecurityToken: models.ToPointer(false),
        CreatedAt:             models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
    }

}
```

