
# Group Settings

## Structure

`GroupSettings`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Target` | [`models.GroupTarget`](../../doc/models/group-target.md) | Required | Attributes of the target customer who will be the responsible payer of the created subscription. Required. |
| `Billing` | [`*models.GroupBilling`](../../doc/models/group-billing.md) | Optional | Optional attributes related to billing date and accrual. Note: Only applicable for new subscriptions. |

## Example (as JSON)

```json
{
  "target": {
    "type": "parent",
    "id": 236
  },
  "billing": {
    "accrue": false,
    "align_date": false,
    "prorate": false
  }
}
```

