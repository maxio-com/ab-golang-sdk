
# Coupon

## Structure

`Coupon`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `Name` | `*string` | Optional | - |
| `Code` | `*string` | Optional | - |
| `Description` | `*string` | Optional | - |
| `Amount` | `models.Optional[float64]` | Optional | - |
| `AmountInCents` | `models.Optional[int64]` | Optional | - |
| `ProductFamilyId` | `*int` | Optional | - |
| `ProductFamilyName` | `models.Optional[string]` | Optional | - |
| `StartDate` | `*time.Time` | Optional | - |
| `EndDate` | `models.Optional[time.Time]` | Optional | After the given time, this coupon code will be invalid for new signups. Recurring discounts started before this date will continue to recur even after this date. |
| `Percentage` | `models.Optional[string]` | Optional | - |
| `Recurring` | `*bool` | Optional | - |
| `RecurringScheme` | [`*models.RecurringScheme`](../../doc/models/recurring-scheme.md) | Optional | - |
| `DurationPeriodCount` | `models.Optional[int]` | Optional | - |
| `DurationInterval` | `models.Optional[int]` | Optional | - |
| `DurationIntervalUnit` | `models.Optional[string]` | Optional | - |
| `DurationIntervalSpan` | `models.Optional[string]` | Optional | - |
| `AllowNegativeBalance` | `*bool` | Optional | If set to true, discount is not limited (credits will carry forward to next billing). |
| `ArchivedAt` | `models.Optional[time.Time]` | Optional | - |
| `ConversionLimit` | `models.Optional[string]` | Optional | - |
| `Stackable` | `*bool` | Optional | A stackable coupon can be combined with other coupons on a Subscription. |
| `CompoundingStrategy` | [`models.Optional[models.CompoundingStrategy]`](../../doc/models/compounding-strategy.md) | Optional | Applicable only to stackable coupons. For `compound`, Percentage-based discounts will be calculated against the remaining price, after prior discounts have been calculated. For `full-price`, Percentage-based discounts will always be calculated against the original item price, before other discounts are applied. |
| `UseSiteExchangeRate` | `*bool` | Optional | - |
| `CreatedAt` | `*time.Time` | Optional | - |
| `UpdatedAt` | `*time.Time` | Optional | - |
| `DiscountType` | [`*models.DiscountType`](../../doc/models/discount-type.md) | Optional | - |
| `ExcludeMidPeriodAllocations` | `*bool` | Optional | - |
| `ApplyOnCancelAtEndOfPeriod` | `*bool` | Optional | - |
| `ApplyOnSubscriptionExpiration` | `*bool` | Optional | - |
| `CouponRestrictions` | [`[]models.CouponRestriction`](../../doc/models/coupon-restriction.md) | Optional | - |
| `CurrencyPrices` | [`[]models.CouponCurrency`](../../doc/models/coupon-currency.md) | Optional | Returned in read, find, and list endpoints if the query parameter is provided. |

## Example (as JSON)

```json
{
  "id": 22,
  "name": "name2",
  "code": "code0",
  "description": "description2",
  "amount": 62.64
}
```

