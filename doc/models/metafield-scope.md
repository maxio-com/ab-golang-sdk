
# Metafield Scope

Warning: When updating a metafield's scope attribute, all scope attributes must be passed. Partially complete scope attributes will override the existing settings.

## Structure

`MetafieldScope`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Csv` | [`*models.IncludeOption`](../../doc/models/include-option.md) | Optional | Include (1) or exclude (0) metafields from the csv export. |
| `Invoices` | [`*models.IncludeOption`](../../doc/models/include-option.md) | Optional | Include (1) or exclude (0) metafields from invoices. |
| `Statements` | [`*models.IncludeOption`](../../doc/models/include-option.md) | Optional | Include (1) or exclude (0) metafields from statements. |
| `Portal` | [`*models.IncludeOption`](../../doc/models/include-option.md) | Optional | Include (1) or exclude (0) metafields from the portal. |
| `PublicShow` | [`*models.IncludeOption`](../../doc/models/include-option.md) | Optional | Include (1) or exclude (0) metafields used in [Embeddable Components](http://localhost:8080/go) from being viewable by your ecosystem. |
| `PublicEdit` | [`*models.IncludeOption`](../../doc/models/include-option.md) | Optional | Include (1) or exclude (0) metafields used in [Embeddable Components](http://localhost:8080/go) from being editable by your ecosystem. |
| `Hosted` | `[]string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    metafieldScope := models.MetafieldScope{
        Csv:                  models.ToPointer(models.IncludeOption_EXCLUDE),
        Invoices:             models.ToPointer(models.IncludeOption_EXCLUDE),
        Statements:           models.ToPointer(models.IncludeOption_EXCLUDE),
        Portal:               models.ToPointer(models.IncludeOption_EXCLUDE),
        PublicShow:           models.ToPointer(models.IncludeOption_EXCLUDE),
    }

}
```

