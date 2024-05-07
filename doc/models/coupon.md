
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
| `AmountInCents` | `models.Optional[int]` | Optional | - |
| `ProductFamilyId` | `*int` | Optional | - |
| `ProductFamilyName` | `models.Optional[string]` | Optional | - |
| `StartDate` | `*time.Time` | Optional | - |
| `EndDate` | `models.Optional[time.Time]` | Optional | - |
| `Percentage` | `models.Optional[string]` | Optional | - |
| `Recurring` | `*bool` | Optional | - |
| `RecurringScheme` | [`*models.RecurringScheme`](../../doc/models/recurring-scheme.md) | Optional | - |
| `DurationPeriodCount` | `models.Optional[int]` | Optional | - |
| `DurationInterval` | `models.Optional[int]` | Optional | - |
| `DurationIntervalUnit` | `models.Optional[string]` | Optional | - |
| `DurationIntervalSpan` | `models.Optional[string]` | Optional | - |
| `AllowNegativeBalance` | `*bool` | Optional | - |
| `ArchivedAt` | `models.Optional[time.Time]` | Optional | - |
| `ConversionLimit` | `models.Optional[string]` | Optional | - |
| `Stackable` | `*bool` | Optional | - |
| `CompoundingStrategy` | [`models.Optional[models.CompoundingStrategy]`](../../doc/models/compounding-strategy.md) | Optional | - |
| `UseSiteExchangeRate` | `*bool` | Optional | - |
| `CreatedAt` | `*time.Time` | Optional | - |
| `UpdatedAt` | `*time.Time` | Optional | - |
| `DiscountType` | [`*models.DiscountType`](../../doc/models/discount-type.md) | Optional | - |
| `ExcludeMidPeriodAllocations` | `*bool` | Optional | - |
| `ApplyOnCancelAtEndOfPeriod` | `*bool` | Optional | - |
| `ApplyOnSubscriptionExpiration` | `*bool` | Optional | - |
| `CouponRestrictions` | [`[]models.CouponRestriction`](../../doc/models/coupon-restriction.md) | Optional | - |

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

