
# Change Invoice Collection Method Event Data

Example schema for an `change_invoice_collection_method` event

## Structure

`ChangeInvoiceCollectionMethodEventData`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `FromCollectionMethod` | `string` | Required | The previous collection method of the invoice. |
| `ToCollectionMethod` | `string` | Required | The new collection method of the invoice. |

## Example (as JSON)

```json
{
  "from_collection_method": "from_collection_method4",
  "to_collection_method": "to_collection_method2"
}
```

