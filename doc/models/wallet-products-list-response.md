
# Wallet Products List Response

## Structure

`WalletProductsListResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `WalletProducts` | [`[]models.WalletProductListItem`](../../doc/models/wallet-product-list-item.md) | Required | - |
| `TotalCount` | `int` | Required | Total number of Wallet Products matching the filters, across all pages. |
| `ArchivedCount` | `int` | Required | Not a simple count of archived products matching the filters. With `status=archived`, equals `total_count` (there's nothing else to report). With any other `status`, it's `0` whenever `total_count` is non-zero, and only when `total_count` is `0` does it become the number of archived Wallet Products matching the same `q` search term (ignoring `status`): an empty-state hint, not a running total. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    walletProductsListResponse := models.WalletProductsListResponse{
        WalletProducts:       []models.WalletProductListItem{
            models.WalletProductListItem{
                Id:                        models.ToPointer(70),
                Name:                      models.ToPointer("name4"),
                Handle:                    models.NewOptional(models.ToPointer("handle0")),
                PricingScheme:             models.NewOptional(models.ToPointer(models.PricingScheme_STAIRSTEP)),
                UnitName:                  models.ToPointer("unit_name6"),
                ItemCategory:              models.NewOptional(models.ToPointer(models.ItemCategory_ENUMBUSINESSSOFTWARE)),
            },
        },
        TotalCount:           214,
        ArchivedCount:        238,
    }

}
```

