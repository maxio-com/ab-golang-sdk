
# List Subscription Events Input

Input structure for the method ListSubscriptionEvents

## Structure

`ListSubscriptionEventsInput`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SubscriptionId` | `int` | Required | The Chargify id of the subscription. |
| `Page` | `*int` | Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |
| `PerPage` | `*int` | Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`.<br><br>**Default**: `20`<br><br>**Constraints**: `<= 200` |
| `SinceId` | `*int64` | Optional | Returns events with an id greater than or equal to the one specified |
| `MaxId` | `*int64` | Optional | Returns events with an id less than or equal to the one specified |
| `Direction` | [`*models.Direction`](../../doc/models/direction.md) | Optional | The sort direction of the returned events.<br><br>**Default**: `"desc"` |
| `Filter` | [`[]models.EventKey`](../../doc/models/event-key.md) | Optional | You can pass multiple event keys after comma.<br>Use in query `filter=signup_success,payment_success`. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    listSubscriptionEventsInput := models.ListSubscriptionEventsInput{
        SubscriptionId:       222,
        Page:                 models.ToPointer(1),
        PerPage:              models.ToPointer(50),
        SinceId:              models.ToPointer(int64(104)),
        MaxId:                models.ToPointer(int64(0)),
        Direction:            models.ToPointer(models.Direction_DESC),
        Filter:               []models.EventKey{
            models.EventKey_CUSTOMFIELDVALUECHANGE,
            models.EventKey_PAYMENTSUCCESS,
        },
    }

}
```

