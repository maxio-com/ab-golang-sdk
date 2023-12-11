
# Subscription Preview Response

## Structure

`SubscriptionPreviewResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SubscriptionPreview` | [`models.SubscriptionPreview`](subscription-preview.md) | Required | - |

## Example (as JSON)

```json
{
  "subscription_preview": {
    "current_billing_manifest": {
      "line_items": [
        {
          "transaction_type": "transaction_type0",
          "kind": "kind6",
          "amount_in_cents": 24,
          "memo": "memo2",
          "discount_amount_in_cents": 172
        }
      ],
      "total_in_cents": 38,
      "total_discount_in_cents": 24,
      "total_tax_in_cents": 18,
      "subtotal_in_cents": 150
    },
    "next_billing_manifest": {
      "line_items": [
        {
          "transaction_type": "transaction_type0",
          "kind": "kind6",
          "amount_in_cents": 24,
          "memo": "memo2",
          "discount_amount_in_cents": 172
        },
        {
          "transaction_type": "transaction_type0",
          "kind": "kind6",
          "amount_in_cents": 24,
          "memo": "memo2",
          "discount_amount_in_cents": 172
        },
        {
          "transaction_type": "transaction_type0",
          "kind": "kind6",
          "amount_in_cents": 24,
          "memo": "memo2",
          "discount_amount_in_cents": 172
        }
      ],
      "total_in_cents": 62,
      "total_discount_in_cents": 208,
      "total_tax_in_cents": 42,
      "subtotal_in_cents": 174
    }
  }
}
```

