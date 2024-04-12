
# Create or Update Flat Amount Coupon

## Structure

`CreateOrUpdateFlatAmountCoupon`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `string` | Required | the name of the coupon |
| `Code` | `string` | Required | may contain uppercase alphanumeric characters and these special characters (which allow for email addresses to be used): “%”, “@”, “+”, “-”, “_”, and “.” |
| `Description` | `*string` | Optional | - |
| `AmountInCents` | `int64` | Required | - |
| `AllowNegativeBalance` | `*bool` | Optional | - |
| `Recurring` | `*bool` | Optional | - |
| `EndDate` | `*time.Time` | Optional | - |
| `ProductFamilyId` | `*string` | Optional | - |
| `Stackable` | `*bool` | Optional | - |
| `CompoundingStrategy` | [`*models.CompoundingStrategy`](../../doc/models/compounding-strategy.md) | Optional | - |
| `ExcludeMidPeriodAllocations` | `*bool` | Optional | - |
| `ApplyOnCancelAtEndOfPeriod` | `*bool` | Optional | - |
| `ApplyOnSubscriptionExpiration` | `*bool` | Optional | - |

## Example (as JSON)

```json
{
  "name": "name0",
  "code": "code8",
  "description": "description0",
  "amount_in_cents": 120,
  "allow_negative_balance": false,
  "recurring": false,
  "end_date": "2016-03-13T12:52:32.123Z",
  "product_family_id": "product_family_id4"
}
```

