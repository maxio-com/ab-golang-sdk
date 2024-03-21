
# Tax Configuration

## Structure

`TaxConfiguration`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Kind` | [`*models.TaxConfigurationKind`](../../doc/models/tax-configuration-kind.md) | Optional | **Default**: `"custom"` |
| `DestinationAddress` | [`*models.TaxDestinationAddress`](../../doc/models/tax-destination-address.md) | Optional | - |
| `FullyConfigured` | `*bool` | Optional | Returns `true` when Chargify has been properly configured to charge tax using the specified tax system. More details about taxes: https://maxio-chargify.zendesk.com/hc/en-us/articles/5405488905869-Taxes-Introduction<br>**Default**: `false` |

## Example (as JSON)

```json
{
  "kind": "custom",
  "fully_configured": false,
  "destination_address": "shipping_only"
}
```

