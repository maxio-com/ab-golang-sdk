
# Calendar Billing

(Optional). Cannot be used when also specifying next_billing_at

## Structure

`CalendarBilling`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SnapDay` | `*interface{}` | Optional | A day of month that subscription will be processed on. Can be 1 up to 28 or 'end'. |
| `CalendarBillingFirstCharge` | [`*models.FirstChargeType`](../../doc/models/first-charge-type.md) | Optional | - |

## Example (as JSON)

```json
{
  "snap_day": {
    "key1": "val1",
    "key2": "val2"
  },
  "calendar_billing_first_charge": "prorated"
}
```

