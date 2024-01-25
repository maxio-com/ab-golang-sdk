
# Subscription Component

## Structure

`SubscriptionComponent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `Name` | `*string` | Optional | - |
| `Kind` | [`*models.ComponentKind`](../../doc/models/component-kind.md) | Optional | A handle for the component type |
| `UnitName` | `*string` | Optional | - |
| `Enabled` | `*bool` | Optional | (for on/off components) indicates if the component is enabled for the subscription |
| `UnitBalance` | `*int` | Optional | - |
| `Currency` | `*string` | Optional | - |
| `AllocatedQuantity` | `*interface{}` | Optional | For Quantity-based components: The current allocation for the component on the given subscription. For On/Off components: Use 1 for on. Use 0 for off. |
| `PricingScheme` | [`Optional[models.PricingScheme]`](../../doc/models/pricing-scheme.md) | Optional | - |
| `ComponentId` | `*int` | Optional | - |
| `ComponentHandle` | `Optional[string]` | Optional | - |
| `SubscriptionId` | `*int` | Optional | - |
| `Recurring` | `*bool` | Optional | - |
| `UpgradeCharge` | [`Optional[models.CreditType]`](../../doc/models/credit-type.md) | Optional | The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.<br>Available values: `full`, `prorated`, `none`. |
| `DowngradeCredit` | [`Optional[models.CreditType]`](../../doc/models/credit-type.md) | Optional | The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.<br>Available values: `full`, `prorated`, `none`. |
| `ArchivedAt` | `Optional[string]` | Optional | - |
| `PricePointId` | `Optional[int]` | Optional | - |
| `PricePointHandle` | `Optional[string]` | Optional | - |
| `PricePointType` | `*interface{}` | Optional | - |
| `PricePointName` | `Optional[string]` | Optional | - |
| `ProductFamilyId` | `*int` | Optional | - |
| `ProductFamilyHandle` | `*string` | Optional | - |
| `CreatedAt` | `*time.Time` | Optional | - |
| `UpdatedAt` | `*time.Time` | Optional | - |
| `UseSiteExchangeRate` | `Optional[bool]` | Optional | - |
| `Description` | `Optional[string]` | Optional | - |
| `AllowFractionalQuantities` | `*bool` | Optional | - |
| `Subscription` | [`*models.SubscriptionComponentSubscription`](../../doc/models/subscription-component-subscription.md) | Optional | An optional object, will be returned if provided `include=subscription` query param. |
| `DisplayOnHostedPage` | `*bool` | Optional | - |
| `Interval` | `*int` | Optional | The numerical interval. i.e. an interval of '30' coupled with an interval_unit of day would mean this component price point would renew every 30 days. This property is only available for sites with Multifrequency enabled. |
| `IntervalUnit` | [`*models.IntervalUnit`](../../doc/models/interval-unit.md) | Optional | A string representing the interval unit for this component price point, either month or day. This property is only available for sites with Multifrequency enabled. |

## Example (as JSON)

```json
{
  "id": 20,
  "name": "name8",
  "kind": "quantity_based_component",
  "unit_name": "unit_name0",
  "enabled": false
}
```

