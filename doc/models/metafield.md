
# Metafield

## Structure

`Metafield`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `Name` | `*string` | Optional | - |
| `Scope` | [`*models.MetafieldScope`](../../doc/models/metafield-scope.md) | Optional | Warning: When updating a metafield's scope attribute, all scope attributes must be passed. Partially complete scope attributes will override the existing settings. |
| `DataCount` | `*int` | Optional | The amount of subscriptions this metafield has been applied to in Advanced Billing. |
| `InputType` | [`*models.MetafieldInput`](../../doc/models/metafield-input.md) | Optional | Indicates the type of metafield. A text metafield allows any string value. Dropdown and radio metafields have a set of values that can be selected.  Defaults to 'text'. |
| `Enum` | [`models.Optional[models.MetafieldEnum]`](../../doc/models/containers/metafield-enum.md) | Optional | This is a container for one-of cases. |

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
  "input_type": "radio"
}
```

