
# List Subscription Group Proforma Invoices Input

Input structure for the method ListSubscriptionGroupProformaInvoices

## Structure

`ListSubscriptionGroupProformaInvoicesInput`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Uid` | `string` | Required | The uid of the subscription group |
| `LineItems` | `*bool` | Optional | Include line items data<br><br>**Default**: `false` |
| `Discounts` | `*bool` | Optional | Include discounts data<br><br>**Default**: `false` |
| `Taxes` | `*bool` | Optional | Include taxes data<br><br>**Default**: `false` |
| `Credits` | `*bool` | Optional | Include credits data<br><br>**Default**: `false` |
| `Payments` | `*bool` | Optional | Include payments data<br><br>**Default**: `false` |
| `CustomFields` | `*bool` | Optional | Include custom fields data<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    listSubscriptionGroupProformaInvoicesInput := models.ListSubscriptionGroupProformaInvoicesInput{
        Uid:                  "uid0",
        LineItems:            models.ToPointer(false),
        Discounts:            models.ToPointer(false),
        Taxes:                models.ToPointer(false),
        Credits:              models.ToPointer(false),
        Payments:             models.ToPointer(false),
        CustomFields:         models.ToPointer(false),
    }

}
```

