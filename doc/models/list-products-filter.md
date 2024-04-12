
# List Products Filter

## Structure

`ListProductsFilter`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PrepaidProductPricePoint` | [`*models.PrepaidProductPricePointFilter`](../../doc/models/prepaid-product-price-point-filter.md) | Optional | Allows fetching products only if a prepaid product price point is present or not. To use this filter you also have to include the following param in the request `include=prepaid_product_price_point`. Use in query `filter[prepaid_product_price_point][product_price_point_id]=not_null`. |
| `UseSiteExchangeRate` | `*bool` | Optional | Allows fetching products with matching use_site_exchange_rate based on provided value (refers to default price point). Use in query `filter[use_site_exchange_rate]=true`. |

## Example (as JSON)

```json
{
  "prepaid_product_price_point": {
    "product_price_point_id": "product_price_point_id2"
  },
  "use_site_exchange_rate": false
}
```

