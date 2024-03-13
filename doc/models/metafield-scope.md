
# Metafield Scope

Warning: When updating a metafield's scope attribute, all scope attributes must be passed. Partially complete scope attributes will override the existing settings.

## Structure

`MetafieldScope`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Csv` | [`*models.IncludeOption`](../../doc/models/include-option.md) | Optional | Include (1) or exclude (0) metafields from the csv export. |
| `Invoices` | [`*models.IncludeOption`](../../doc/models/include-option.md) | Optional | Include (1) or exclude (0) metafields from invoices. |
| `Statements` | [`*models.IncludeOption`](../../doc/models/include-option.md) | Optional | Include (1) or exclude (0) metafields from statements. |
| `Portal` | [`*models.IncludeOption`](../../doc/models/include-option.md) | Optional | Include (1) or exclude (0) metafields from the portal. |
| `PublicShow` | [`*models.IncludeOption`](../../doc/models/include-option.md) | Optional | Include (1) or exclude (0) metafields from being viewable by your ecosystem. |
| `PublicEdit` | [`*models.IncludeOption`](../../doc/models/include-option.md) | Optional | Include (1) or exclude (0) metafields from being edited by your ecosystem. |
| `Hosted` | `[]string` | Optional | - |

## Example (as JSON)

```json
{
  "csv": "0",
  "invoices": "0",
  "statements": "0",
  "portal": "0",
  "public_show": "0"
}
```

