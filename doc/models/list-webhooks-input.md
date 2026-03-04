
# List Webhooks Input

Input structure for the method ListWebhooks

## Structure

`ListWebhooksInput`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Status` | [`*models.WebhookStatus`](../../doc/models/webhook-status.md) | Optional | Webhooks with matching status would be returned. |
| `SinceDate` | `*string` | Optional | Format YYYY-MM-DD. Returns Webhooks with the created_at date greater than or equal to the one specified. |
| `UntilDate` | `*string` | Optional | Format YYYY-MM-DD. Returns Webhooks with the created_at date less than or equal to the one specified. |
| `Page` | `*int` | Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |
| `PerPage` | `*int` | Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`.<br><br>**Default**: `20`<br><br>**Constraints**: `<= 200` |
| `Order` | [`*models.WebhookOrder`](../../doc/models/webhook-order.md) | Optional | The order in which the Webhooks are returned. |
| `Subscription` | `*int` | Optional | The Advanced Billing id of a subscription you'd like to filter for |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    listWebhooksInput := models.ListWebhooksInput{
        Status:               models.ToPointer(models.WebhookStatus_SUCCESSFUL),
        SinceDate:            models.ToPointer("since_date4"),
        UntilDate:            models.ToPointer("until_date6"),
        Page:                 models.ToPointer(1),
        PerPage:              models.ToPointer(50),
        Order:                models.ToPointer(models.WebhookOrder_NEWESTFIRST),
        Subscription:         models.ToPointer(152),
    }

}
```

