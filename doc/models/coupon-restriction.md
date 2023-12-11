
# Coupon Restriction

## Structure

`CouponRestriction`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `ItemType` | [`*models.RestrictionTypeEnum`](restriction-type-enum.md) | Optional | - |
| `ItemId` | `*int` | Optional | - |
| `Name` | `*string` | Optional | - |
| `Handle` | `Optional[string]` | Optional | - |

## Example (as JSON)

```json
{
  "id": 180,
  "item_type": "Component",
  "item_id": 184,
  "name": "name4",
  "handle": "handle0"
}
```

