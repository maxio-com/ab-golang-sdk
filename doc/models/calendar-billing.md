
# Calendar Billing

(Optional). Cannot be used when also specifying next_billing_at

## Structure

`CalendarBilling`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SnapDay` | [`*models.CalendarBillingSnapDay`](../../doc/models/containers/calendar-billing-snap-day.md) | Optional | This is a container for one-of cases. |
| `CalendarBillingFirstCharge` | [`*models.FirstChargeType`](../../doc/models/first-charge-type.md) | Optional | - |

## Example (as JSON)

```json
{
  "snap_day": 210,
  "calendar_billing_first_charge": "prorated"
}
```

