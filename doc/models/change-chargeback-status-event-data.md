
# Change Chargeback Status Event Data

Example schema for an `change_chargeback_status` event

## Structure

`ChangeChargebackStatusEventData`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ChargebackStatus` | [`models.ChargebackStatus`](../../doc/models/chargeback-status.md) | Required | - |

## Example (as JSON)

```json
{
  "chargeback_status": "open"
}
```

