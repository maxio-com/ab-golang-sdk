
# Subscription Group Customer

## Structure

`SubscriptionGroupCustomer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `FirstName` | `*string` | Optional | - |
| `LastName` | `*string` | Optional | - |
| `Organization` | `*string` | Optional | - |
| `Email` | `*string` | Optional | - |
| `Reference` | `*string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    subscriptionGroupCustomer := models.SubscriptionGroupCustomer{
        FirstName:            models.ToPointer("first_name6"),
        LastName:             models.ToPointer("last_name4"),
        Organization:         models.ToPointer("organization0"),
        Email:                models.ToPointer("email0"),
        Reference:            models.ToPointer("reference8"),
    }

}
```

