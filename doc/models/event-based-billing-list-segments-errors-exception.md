
# Event Based Billing List Segments Errors Exception

## Structure

`EventBasedBillingListSegmentsErrorsException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Errors` | [`*models.Errors`](errors.md) | Optional | - |

## Example (as JSON)

```json
{
  "errors": {
    "per_page": [
      "per_page1",
      "per_page2",
      "per_page3"
    ],
    "price_point": [
      "price_point0",
      "price_point9",
      "price_point8"
    ]
  }
}
```

