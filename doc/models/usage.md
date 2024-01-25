
# Usage

## Structure

`Usage`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int64` | Optional | **Constraints**: `>= 0` |
| `Memo` | `*string` | Optional | - |
| `CreatedAt` | `*time.Time` | Optional | - |
| `PricePointId` | `*int` | Optional | - |
| `Quantity` | `*interface{}` | Optional | - |
| `OverageQuantity` | `*int` | Optional | - |
| `ComponentId` | `*int` | Optional | - |
| `ComponentHandle` | `*string` | Optional | - |
| `SubscriptionId` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "id": 252,
  "memo": "memo8",
  "created_at": "2016-03-13T12:52:32.123Z",
  "price_point_id": 126,
  "quantity": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

