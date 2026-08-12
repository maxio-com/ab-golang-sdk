
# Subscription Preview Response

## Structure

`SubscriptionPreviewResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SubscriptionPreview` | [`models.SubscriptionPreview`](../../doc/models/subscription-preview.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    subscriptionPreviewResponse := models.SubscriptionPreviewResponse{
        SubscriptionPreview:  models.SubscriptionPreview{
            CurrentBillingManifest: models.ToPointer(models.BillingManifest{
                LineItems:              []models.BillingManifestItem{
                    models.BillingManifestItem{
                        TransactionType:       models.ToPointer(models.LineItemTransactionType_CREDIT),
                        Kind:                  models.ToPointer(models.BillingManifestLineItemKind_COMPONENT),
                        AmountInCents:         models.ToPointer(int64(24)),
                        Memo:                  models.ToPointer("memo2"),
                        DiscountAmountInCents: models.ToPointer(int64(172)),
                    },
                },
                TotalInCents:           models.ToPointer(int64(38)),
                TotalDiscountInCents:   models.ToPointer(int64(24)),
                TotalTaxInCents:        models.ToPointer(int64(18)),
                SubtotalInCents:        models.ToPointer(int64(150)),
            }),
            NextBillingManifest:    models.ToPointer(models.BillingManifest{
                LineItems:              []models.BillingManifestItem{
                    models.BillingManifestItem{
                        TransactionType:       models.ToPointer(models.LineItemTransactionType_CREDIT),
                        Kind:                  models.ToPointer(models.BillingManifestLineItemKind_COMPONENT),
                        AmountInCents:         models.ToPointer(int64(24)),
                        Memo:                  models.ToPointer("memo2"),
                        DiscountAmountInCents: models.ToPointer(int64(172)),
                    },
                    models.BillingManifestItem{
                        TransactionType:       models.ToPointer(models.LineItemTransactionType_CREDIT),
                        Kind:                  models.ToPointer(models.BillingManifestLineItemKind_COMPONENT),
                        AmountInCents:         models.ToPointer(int64(24)),
                        Memo:                  models.ToPointer("memo2"),
                        DiscountAmountInCents: models.ToPointer(int64(172)),
                    },
                    models.BillingManifestItem{
                        TransactionType:       models.ToPointer(models.LineItemTransactionType_CREDIT),
                        Kind:                  models.ToPointer(models.BillingManifestLineItemKind_COMPONENT),
                        AmountInCents:         models.ToPointer(int64(24)),
                        Memo:                  models.ToPointer("memo2"),
                        DiscountAmountInCents: models.ToPointer(int64(172)),
                    },
                },
                TotalInCents:           models.ToPointer(int64(62)),
                TotalDiscountInCents:   models.ToPointer(int64(208)),
                TotalTaxInCents:        models.ToPointer(int64(42)),
                SubtotalInCents:        models.ToPointer(int64(174)),
            }),
        },
    }

}
```

