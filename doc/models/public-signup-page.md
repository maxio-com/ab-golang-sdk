
# Public Signup Page

## Structure

`PublicSignupPage`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | The id of the signup page (public_signup_pages only) |
| `ReturnUrl` | `models.Optional[string]` | Optional | The url to which a customer will be returned after a successful signup (public_signup_pages only). |
| `ReturnParams` | `models.Optional[string]` | Optional | The params to be appended to the return_url (public_signup_pages only) |
| `Url` | `*string` | Optional | The url where the signup page can be viewed (public_signup_pages only). |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    publicSignupPage := models.PublicSignupPage{
        Id:                   models.ToPointer(20),
        ReturnUrl:            models.NewOptional(models.ToPointer("return_url0")),
        ReturnParams:         models.NewOptional(models.ToPointer("return_params2")),
        Url:                  models.ToPointer("url8"),
    }

}
```

