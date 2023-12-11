
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
| `Amount` | `Optional[float64]` | Optional | - |
| `AmountInCents` | `Optional[int]` | Optional | - |
| `ProductFamilyId` | `*int` | Optional | - |
| `ProductFamilyName` | `*string` | Optional | - |
| `StartDate` | `*string` | Optional | - |
| `EndDate` | `Optional[string]` | Optional | - |
| `Percentage` | `Optional[float64]` | Optional | - |
| `Recurring` | `*bool` | Optional | - |
| `RecurringScheme` | [`*models.RecurringSchemeEnum`](recurring-scheme-enum.md) | Optional | - |
| `DurationPeriodCount` | `Optional[int]` | Optional | - |
| `DurationInterval` | `Optional[int]` | Optional | - |
| `DurationIntervalUnit` | `Optional[string]` | Optional | - |
| `DurationIntervalSpan` | `*string` | Optional | - |
| `AllowNegativeBalance` | `*bool` | Optional | - |
| `ArchivedAt` | `Optional[string]` | Optional | - |
| `ConversionLimit` | `Optional[string]` | Optional | - |
| `Stackable` | `*bool` | Optional | - |
| `CompoundingStrategy` | `*interface{}` | Optional | - |
| `UseSiteExchangeRate` | `*bool` | Optional | - |
| `CreatedAt` | `*string` | Optional | - |
| `UpdatedAt` | `*string` | Optional | - |
| `DiscountType` | [`*models.DiscountTypeEnum`](discount-type-enum.md) | Optional | - |
| `ExcludeMidPeriodAllocations` | `*bool` | Optional | - |
| `ApplyOnCancelAtEndOfPeriod` | `*bool` | Optional | - |
| `CouponRestrictions` | [`[]models.CouponRestriction`](coupon-restriction.md) | Optional | - |

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

