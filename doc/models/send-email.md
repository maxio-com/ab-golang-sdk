
# Send Email

## Structure

`SendEmail`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CanExecute` | `bool` | Required | - |
| `Url` | `string` | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    sendEmail := models.SendEmail{
        CanExecute:           false,
        Url:                  "url2",
    }

}
```

