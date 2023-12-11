
# Subscription Migration Preview Response

## Structure

`SubscriptionMigrationPreviewResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Migration` | [`models.SubscriptionMigrationPreview`](subscription-migration-preview.md) | Required | - |

## Example (as JSON)

```json
{
  "migration": {
    "prorated_adjustment_in_cents": 196,
    "charge_in_cents": 78,
    "payment_due_in_cents": 250,
    "credit_applied_in_cents": 210
  }
}
```

