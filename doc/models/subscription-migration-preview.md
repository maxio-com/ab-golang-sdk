
# Subscription Migration Preview

## Structure

`SubscriptionMigrationPreview`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ProratedAdjustmentInCents` | `*int64` | Optional | The amount of the prorated adjustment that would be issued for the current subscription. |
| `ChargeInCents` | `*int64` | Optional | The amount of the charge that would be created for the new product. |
| `PaymentDueInCents` | `*int64` | Optional | The amount of the payment due in the case of an upgrade. |
| `CreditAppliedInCents` | `*int64` | Optional | Represents a credit in cents that is applied to your subscription as part of a migration process for a specific product, which reduces the amount owed for the subscription. |

## Example (as JSON)

```json
{
  "prorated_adjustment_in_cents": 134,
  "charge_in_cents": 16,
  "payment_due_in_cents": 188,
  "credit_applied_in_cents": 148
}
```

