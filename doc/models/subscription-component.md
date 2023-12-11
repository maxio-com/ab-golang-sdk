
# Subscription Component

## Structure

`SubscriptionComponent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `Name` | `*string` | Optional | - |
| `Kind` | `*string` | Optional | - |
| `UnitName` | `*string` | Optional | - |
| `Enabled` | `*bool` | Optional | (for on/off components) indicates if the component is enabled for the subscription |
| `UnitBalance` | `*int` | Optional | - |
| `Currency` | `*string` | Optional | - |
| `AllocatedQuantity` | `*int` | Optional | For Quantity-based components: The current allocation for the component on the given subscription. For On/Off components: Use 1 for on. Use 0 for off. |
| `PricingScheme` | [`Optional[models.PricingSchemeEnum]`](pricing-scheme-enum.md) | Optional | - |
| `ComponentId` | `*int` | Optional | - |
| `ComponentHandle` | `Optional[string]` | Optional | - |
| `SubscriptionId` | `*int` | Optional | - |
| `Recurring` | `*bool` | Optional | - |
| `UpgradeCharge` | [`Optional[models.CreditTypeEnum]`](credit-type-enum.md) | Optional | The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.<br>Available values: `full`, `prorated`, `none`. |
| `DowngradeCredit` | [`Optional[models.CreditTypeEnum]`](credit-type-enum.md) | Optional | The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.<br>Available values: `full`, `prorated`, `none`. |
| `ArchivedAt` | `Optional[string]` | Optional | - |
| `PricePointId` | `Optional[int]` | Optional | - |
| `PricePointHandle` | `Optional[string]` | Optional | - |
| `PricePointType` | `*interface{}` | Optional | - |
| `PricePointName` | `Optional[string]` | Optional | - |
| `ProductFamilyId` | `*int` | Optional | - |
| `ProductFamilyHandle` | `*string` | Optional | - |
| `CreatedAt` | `*string` | Optional | - |
| `UpdatedAt` | `*string` | Optional | - |
| `UseSiteExchangeRate` | `Optional[bool]` | Optional | - |
| `Description` | `Optional[string]` | Optional | - |
| `AllowFractionalQuantities` | `*bool` | Optional | - |
| `Subscription` | [`*models.SubscriptionComponentSubscription`](subscription-component-subscription.md) | Optional | An optional object, will be returned if provided `include=subscription` query param. |
| `DisplayOnHostedPage` | `*bool` | Optional | - |

## Example (as JSON)

```json
{
  "id": 20,
  "name": "name8",
  "kind": "kind6",
  "unit_name": "unit_name0",
  "enabled": false
}
```

