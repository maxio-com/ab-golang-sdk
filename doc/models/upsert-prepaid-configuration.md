
# Upsert Prepaid Configuration

## Structure

`UpsertPrepaidConfiguration`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `InitialFundingAmountInCents` | `*int64` | Optional | - |
| `ReplenishToAmountInCents` | `*int64` | Optional | - |
| `AutoReplenish` | `*bool` | Optional | - |
| `ReplenishThresholdAmountInCents` | `*int64` | Optional | - |

## Example (as JSON)

```json
{
  "initial_funding_amount_in_cents": 40,
  "replenish_to_amount_in_cents": 42,
  "auto_replenish": false,
  "replenish_threshold_amount_in_cents": 242
}
```

