
# Net Terms

## Structure

`NetTerms`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DefaultNetTerms` | `*int` | Optional | **Default**: `0` |
| `AutomaticNetTerms` | `*int` | Optional | **Default**: `0` |
| `RemittanceNetTerms` | `*int` | Optional | **Default**: `0` |
| `NetTermsOnRemittanceSignupsEnabled` | `*bool` | Optional | **Default**: `false` |
| `CustomNetTermsEnabled` | `*bool` | Optional | **Default**: `false` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    netTerms := models.NetTerms{
        DefaultNetTerms:                    models.ToPointer(0),
        AutomaticNetTerms:                  models.ToPointer(0),
        RemittanceNetTerms:                 models.ToPointer(0),
        NetTermsOnRemittanceSignupsEnabled: models.ToPointer(false),
        CustomNetTermsEnabled:              models.ToPointer(false),
    }

}
```

