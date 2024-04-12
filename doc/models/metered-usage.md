
# Metered Usage

## Structure

`MeteredUsage`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PreviousUnitBalance` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `NewUnitBalance` | `int` | Required | - |
| `UsageQuantity` | `int` | Required | - |
| `ComponentId` | `int` | Required | - |
| `ComponentHandle` | `string` | Required | - |
| `Memo` | `string` | Required | - |

## Example (as JSON)

```json
{
  "previous_unit_balance": "previous_unit_balance6",
  "new_unit_balance": 80,
  "usage_quantity": 42,
  "component_id": 4,
  "component_handle": "component_handle8",
  "memo": "memo2"
}
```

