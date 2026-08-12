
# Tax Configuration

## Structure

`TaxConfiguration`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Kind` | [`*models.TaxConfigurationKind`](../../doc/models/tax-configuration-kind.md) | Optional | **Default**: `"custom"` |
| `DestinationAddress` | [`*models.TaxDestinationAddress`](../../doc/models/tax-destination-address.md) | Optional | - |
| `FullyConfigured` | `*bool` | Optional | Returns `true` when Chargify has been properly configured to charge tax using the specified tax system. More details about taxes: https://maxio.zendesk.com/hc/en-us/articles/24287012608909-Taxes-Overview<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    taxConfiguration := models.TaxConfiguration{
        Kind:                 models.ToPointer(models.TaxConfigurationKind_CUSTOM),
        DestinationAddress:   models.ToPointer(models.TaxDestinationAddress_SHIPPINGONLY),
        FullyConfigured:      models.ToPointer(false),
    }

}
```

