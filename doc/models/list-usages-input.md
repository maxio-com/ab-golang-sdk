
# List Usages Input

Input structure for the method ListUsages

## Structure

`ListUsagesInput`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SubscriptionIdOrReference` | [`models.ListUsagesInputSubscriptionIdOrReference`](../../doc/models/containers/list-usages-input-subscription-id-or-reference.md) | Required | This is a container for one-of cases. |
| `ComponentId` | [`models.ListUsagesInputComponentId`](../../doc/models/containers/list-usages-input-component-id.md) | Required | This is a container for one-of cases. |
| `SinceId` | `*int64` | Optional | Returns usages with an id greater than or equal to the one specified |
| `MaxId` | `*int64` | Optional | Returns usages with an id less than or equal to the one specified |
| `SinceDate` | `*time.Time` | Optional | Returns usages with a created_at date greater than or equal to midnight (12:00 AM) on the date specified. |
| `UntilDate` | `*time.Time` | Optional | Returns usages with a created_at date less than or equal to midnight (12:00 AM) on the date specified. |
| `Page` | `*int` | Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |
| `PerPage` | `*int` | Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`.<br><br>**Default**: `20`<br><br>**Constraints**: `<= 200` |

## Example

```go
package main

import (
    "log"
    "time"
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    parseTime := func(layout, value string, errCallback func(error)) time.Time {
        dateTime, err := time.Parse(layout, value)
        if err != nil {
            errCallback(err) 
       }
        return dateTime
    }
    listUsagesInput := models.ListUsagesInput{
        SubscriptionIdOrReference: models.ListUsagesInputSubscriptionIdOrReferenceContainer.FromNumber(234),
        ComponentId:               models.ListUsagesInputComponentIdContainer.FromNumber(144),
        SinceId:                   models.ToPointer(int64(104)),
        MaxId:                     models.ToPointer(int64(0)),
        SinceDate:                 models.ToPointer(parseTime(models.DEFAULT_DATE, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
        UntilDate:                 models.ToPointer(parseTime(models.DEFAULT_DATE, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
        Page:                      models.ToPointer(1),
        PerPage:                   models.ToPointer(50),
    }

}
```

