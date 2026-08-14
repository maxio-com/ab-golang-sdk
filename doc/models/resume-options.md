
# Resume Options

## Structure

`ResumeOptions`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `RequireResume` | `*bool` | Optional | Chargify will only attempt to resume the subscription's billing period. If not resumable, the subscription will be left in its current state. |
| `ForgiveBalance` | `*bool` | Optional | Indicates whether or not Chargify should clear the subscription's existing balance before attempting to resume the subscription. If subscription cannot be resumed, the balance will remain as it was before the attempt to resume was made. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    resumeOptions := models.ResumeOptions{
        RequireResume:        models.ToPointer(false),
        ForgiveBalance:       models.ToPointer(false),
    }

}
```

