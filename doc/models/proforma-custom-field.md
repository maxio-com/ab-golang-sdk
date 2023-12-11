
# Proforma Custom Field

## Structure

`ProformaCustomField`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `OwnerId` | `*int` | Optional | - |
| `OwnerType` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `Name` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `Value` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `MetadatumId` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "owner_id": 224,
  "owner_type": "owner_type0",
  "name": "name8",
  "value": "value0",
  "metadatum_id": 224
}
```

