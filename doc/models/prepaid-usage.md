
# Prepaid Usage

## Structure

`PrepaidUsage`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PreviousUnitBalance` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `PreviousOverageUnitBalance` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `NewUnitBalance` | `int` | Required | - |
| `NewOverageUnitBalance` | `int` | Required | - |
| `UsageQuantity` | `int` | Required | - |
| `OverageUsageQuantity` | `int` | Required | - |
| `ComponentId` | `int` | Required | - |
| `ComponentHandle` | `string` | Required | - |
| `Memo` | `string` | Required | - |
| `AllocationDetails` | [`[]models.PrepaidUsageAllocationDetail`](../../doc/models/prepaid-usage-allocation-detail.md) | Required | - |

## Example (as JSON)

```json
{
  "previous_unit_balance": "previous_unit_balance0",
  "previous_overage_unit_balance": "previous_overage_unit_balance4",
  "new_unit_balance": 252,
  "new_overage_unit_balance": 224,
  "usage_quantity": 214,
  "overage_usage_quantity": 106,
  "component_id": 176,
  "component_handle": "component_handle4",
  "memo": "memo8",
  "allocation_details": [
    {
      "allocation_id": 18,
      "charge_id": 84,
      "usage_quantity": 10
    }
  ]
}
```

