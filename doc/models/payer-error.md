
# Payer Error

## Structure

`PayerError`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `LastName` | `[]string` | Optional | - |
| `FirstName` | `[]string` | Optional | - |
| `Email` | `[]string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    payerError := models.PayerError{
        LastName:             []string{
            "last_name9",
        },
        FirstName:            []string{
            "first_name2",
            "first_name3",
            "first_name4",
        },
        Email:                []string{
            "email6",
            "email7",
            "email8",
        },
    }

}
```

