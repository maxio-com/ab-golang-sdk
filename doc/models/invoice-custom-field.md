
# Invoice Custom Field

## Structure

`InvoiceCustomField`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `OwnerId` | `*int` | Optional | - |
| `OwnerType` | [`*models.CustomFieldOwner`](../../doc/models/custom-field-owner.md) | Optional | - |
| `Name` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `Value` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `MetadatumId` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "owner_id": 142,
  "owner_type": "Customer",
  "name": "name0",
  "value": "value2",
  "metadatum_id": 142
}
```

