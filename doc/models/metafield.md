
# Metafield

## Structure

`Metafield`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `Name` | `*string` | Optional | - |
| `Scope` | [`*models.MetafieldScope`](metafield-scope.md) | Optional | Warning: When updating a metafield's scope attribute, all scope attributes must be passed. Partially complete scope attributes will override the existing settings. |
| `DataCount` | `*int` | Optional | the amount of subscriptions this metafield has been applied to in Chargify |
| `InputType` | `*string` | Optional | - |
| `Enum` | `Optional[interface{}]` | Optional | - |

## Example (as JSON)

```json
{
  "id": 52,
  "name": "name8",
  "scope": {
    "csv": "0",
    "invoices": "0",
    "statements": "0",
    "portal": "0",
    "public_show": "0"
  },
  "data_count": 216,
  "input_type": "input_type0"
}
```

