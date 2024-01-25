
# Movement

## Structure

`Movement`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Timestamp` | `*string` | Optional | - |
| `AmountInCents` | `*int64` | Optional | - |
| `AmountFormatted` | `*string` | Optional | - |
| `Description` | `*string` | Optional | - |
| `Category` | `*string` | Optional | - |
| `Breakouts` | [`*models.Breakouts`](../../doc/models/breakouts.md) | Optional | - |
| `LineItems` | [`[]models.MovementLineItem`](../../doc/models/movement-line-item.md) | Optional | - |
| `SubscriptionId` | `*int` | Optional | - |
| `SubscriberName` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "timestamp": "timestamp0",
  "amount_in_cents": 174,
  "amount_formatted": "amount_formatted4",
  "description": "description2",
  "category": "category0"
}
```

