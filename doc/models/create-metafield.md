
# Create Metafield

## Structure

`CreateMetafield`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `*string` | Optional | - |
| `Scope` | [`*models.MetafieldScope`](../../doc/models/metafield-scope.md) | Optional | Warning: When updating a metafield's scope attribute, all scope attributes must be passed. Partially complete scope attributes will override the existing settings. |
| `InputType` | [`*models.MetafieldInput`](../../doc/models/metafield-input.md) | Optional | Indicates the type of metafield. A text metafield allows any string value. Dropdown and radio metafields have a set of values that can be selected.  Defaults to 'text'. |
| `Enum` | `[]string` | Optional | Only applicable when input_type is radio or dropdown. Empty strings will not be submitted. |

## Example (as JSON)

```json
{
  "name": "my_field",
  "scope": {
    "csv": "0",
    "invoices": "0",
    "statements": "0",
    "portal": "0",
    "public_show": "0",
    "public_edit": "0"
  },
  "input_type": "text",
  "enum": [
    "string"
  ]
}
```

