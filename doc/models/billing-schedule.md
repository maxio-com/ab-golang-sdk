
# Billing Schedule

This attribute is particularly useful when you need to align billing events for different components on distinct schedules within a subscription. This only works for site with Multifrequency enabled.

## Structure

`BillingSchedule`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `InitialBillingAt` | `*time.Time` | Optional | The initial_billing_at attribute in Maxio allows you to specify a custom starting date for billing cycles associated with components that have their own billing frequency set. Only ISO8601 format is supported. |

## Example (as JSON)

```json
{
  "initial_billing_at": "2024-01-01"
}
```

