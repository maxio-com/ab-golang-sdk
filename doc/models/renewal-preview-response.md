
# Renewal Preview Response

## Structure

`RenewalPreviewResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `RenewalPreview` | [`models.RenewalPreview`](../../doc/models/renewal-preview.md) | Required | - |

## Example (as JSON)

```json
{
  "renewal_preview": {
    "next_assessment_at": "2016-03-13T12:52:32.123Z",
    "subtotal_in_cents": 132,
    "total_tax_in_cents": 0,
    "total_discount_in_cents": 250,
    "total_in_cents": 20
  }
}
```

