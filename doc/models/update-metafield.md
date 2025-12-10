
# Update Metafield

## Structure

`UpdateMetafield`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CurrentName` | `*string` | Optional | - |
| `Name` | `*string` | Optional | - |
| `Scope` | [`*models.MetafieldScope`](../../doc/models/metafield-scope.md) | Optional | Warning: When updating a metafield's scope attribute, all scope attributes must be passed. Partially complete scope attributes will override the existing settings. |
| `InputType` | [`*models.MetafieldInput`](../../doc/models/metafield-input.md) | Optional | Indicates the type of metafield. A text metafield allows any string value. Dropdown and radio metafields have a set of values that can be selected.  Defaults to 'text'. |
| `Enum` | `[]string` | Optional | Only applicable when input_type is radio or dropdown. |

## Example (as JSON)

```json
{
  "current_name": "current_name0",
  "name": "name6",
  "scope": {
    "csv": "0",
    "invoices": "0",
    "statements": "0",
    "portal": "0",
    "public_show": "0"
  },
  "input_type": "balance_tracker",
  "enum": [
    "enum2"
  ]
}
```

