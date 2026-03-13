
# Scheduled Renewal Component Custom Price

Custom pricing for a component within a scheduled renewal.

## Structure

`ScheduledRenewalComponentCustomPrice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TaxIncluded` | `*bool` | Optional | Whether or not the price point includes tax |
| `PricingScheme` | [`models.PricingScheme`](../../doc/models/pricing-scheme.md) | Required | Omit for On/Off components |
| `Prices` | [`[]models.Price`](../../doc/models/price.md) | Required | On/off components only need one price bracket starting at 1. |

## Example (as JSON)

```json
{
  "pricing_scheme": "stairstep",
  "prices": [
    {
      "starting_quantity": 242,
      "ending_quantity": 40,
      "unit_price": 23.26
    }
  ],
  "tax_included": false
}
```

