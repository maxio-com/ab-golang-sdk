
# Site Summary

## Structure

`SiteSummary`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SellerName` | `*string` | Optional | - |
| `SiteName` | `*string` | Optional | - |
| `SiteId` | `*int` | Optional | - |
| `SiteCurrency` | `*string` | Optional | - |
| `Stats` | [`*models.SiteStatistics`](site-statistics.md) | Optional | - |

## Example (as JSON)

```json
{
  "seller_name": "seller_name0",
  "site_name": "site_name4",
  "site_id": 218,
  "site_currency": "site_currency6",
  "stats": {
    "total_subscriptions": 110,
    "subscriptions_today": 228,
    "total_revenue": "total_revenue6",
    "revenue_today": "revenue_today4",
    "revenue_this_month": "revenue_this_month4"
  }
}
```

