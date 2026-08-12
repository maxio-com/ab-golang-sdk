
# Offer Signup Page

## Structure

`OfferSignupPage`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `Nickname` | `*string` | Optional | - |
| `Enabled` | `*bool` | Optional | - |
| `ReturnUrl` | `*string` | Optional | - |
| `ReturnParams` | `*string` | Optional | - |
| `Url` | `*string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    offerSignupPage := models.OfferSignupPage{
        Id:                   models.ToPointer(102),
        Nickname:             models.ToPointer("nickname6"),
        Enabled:              models.ToPointer(false),
        ReturnUrl:            models.ToPointer("return_url6"),
        ReturnParams:         models.ToPointer("return_params8"),
    }

}
```

