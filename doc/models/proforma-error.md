
# Proforma Error

## Structure

`ProformaError`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Subscription` | [`*models.BaseStringError`](../../doc/models/base-string-error.md) | Optional | The error is base if it is not directly associated with a single attribute. |

## Example (as JSON)

```json
{
  "subscription": {
    "base": [
      "base3",
      "base4"
    ]
  }
}
```

