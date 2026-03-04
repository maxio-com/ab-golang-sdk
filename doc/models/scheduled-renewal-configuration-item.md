
# Scheduled Renewal Configuration Item

## Structure

`ScheduledRenewalConfigurationItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `SubscriptionId` | `*int` | Optional | - |
| `SubscriptionRenewalConfigurationId` | `*int` | Optional | - |
| `ItemId` | `*int` | Optional | - |
| `ItemType` | `*string` | Optional | - |
| `ItemSubclass` | `*string` | Optional | - |
| `PricePointId` | `*int` | Optional | - |
| `PricePointType` | `*string` | Optional | - |
| `Quantity` | `*int` | Optional | - |
| `DecimalQuantity` | `*string` | Optional | - |
| `CreatedAt` | `*time.Time` | Optional | - |

## Example (as JSON)

```json
{
  "id": 146,
  "subscription_id": 0,
  "subscription_renewal_configuration_id": 156,
  "item_id": 38,
  "item_type": "item_type4"
}
```

