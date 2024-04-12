
# Coupon Restriction

## Structure

`CouponRestriction`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `ItemType` | [`*models.RestrictionType`](../../doc/models/restriction-type.md) | Optional | - |
| `ItemId` | `*int` | Optional | - |
| `Name` | `*string` | Optional | - |
| `Handle` | `models.Optional[string]` | Optional | - |

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

