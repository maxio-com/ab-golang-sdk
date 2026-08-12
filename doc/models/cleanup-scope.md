
# Cleanup Scope

all: Will clear all products, customers, and related subscriptions from the site. customers: Will clear only customers and related subscriptions (leaving the products untouched) for the site. Revenue will also be reset to 0.

## Enumeration

`CleanupScope`

## Fields

| Name |
|  --- |
| `ALL` |
| `CUSTOMERS` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    cleanupScope := models.CleanupScope_ALL

}
```

