
# Calendar Billing

(Optional). Cannot be used when also specifying next_billing_at.

## Structure

`CalendarBilling`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SnapDay` | [`*models.CalendarBillingSnapDay`](../../doc/models/containers/calendar-billing-snap-day.md) | Optional | This is a container for one-of cases. |
| `CalendarBillingFirstCharge` | [`*models.FirstChargeType`](../../doc/models/first-charge-type.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    calendarBilling := models.CalendarBilling{
        SnapDay:                    models.ToPointer(models.CalendarBillingSnapDayContainer.FromNumber(200)),
        CalendarBillingFirstCharge: models.ToPointer(models.FirstChargeType_DELAYED),
    }

}
```

