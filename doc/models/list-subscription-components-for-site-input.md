
# List Subscription Components for Site Input

Input structure for the method ListSubscriptionComponentsForSite

## Structure

`ListSubscriptionComponentsForSiteInput`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Page` | `*int` | Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |
| `PerPage` | `*int` | Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`.<br><br>**Default**: `20`<br><br>**Constraints**: `<= 200` |
| `Sort` | [`*models.ListSubscriptionComponentsSort`](../../doc/models/list-subscription-components-sort.md) | Optional | The attribute by which to sort. Use in query: `sort=updated_at`. |
| `Direction` | [`*models.SortingDirection`](../../doc/models/sorting-direction.md) | Optional | Controls the order in which results are returned.<br>Use in query `direction=asc`. |
| `Filter` | [`*models.ListSubscriptionComponentsForSiteFilter`](../../doc/models/list-subscription-components-for-site-filter.md) | Optional | Filter to use for List Subscription Components For Site operation |
| `DateField` | [`*models.SubscriptionListDateField`](../../doc/models/subscription-list-date-field.md) | Optional | The type of filter you'd like to apply to your search. Use in query: `date_field=updated_at`. |
| `StartDate` | `*string` | Optional | The start date (format YYYY-MM-DD) with which to filter the date_field. Returns components with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified. Use in query `start_date=2011-12-15`. |
| `StartDatetime` | `*string` | Optional | The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns components with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site''s time zone will be used. If provided, this parameter will be used instead of start_date. Use in query `start_datetime=2022-07-01 09:00:05`. |
| `EndDate` | `*string` | Optional | The end date (format YYYY-MM-DD) with which to filter the date_field. Returns components with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified. Use in query `end_date=2011-12-16`. |
| `EndDatetime` | `*string` | Optional | The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns components with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site''s time zone will be used. If provided, this parameter will be used instead of end_date. Use in query `end_datetime=2022-07-01 09:00:05`. |
| `SubscriptionIds` | `[]int` | Optional | Allows fetching components allocation with matching subscription id based on provided ids. Use in query `subscription_ids=1,2,3`.<br><br>**Constraints**: *Minimum Items*: `1`, *Maximum Items*: `200` |
| `PricePointIds` | [`*models.IncludeNotNull`](../../doc/models/include-not-null.md) | Optional | Allows fetching components allocation only if price point id is present. Use in query `price_point_ids=not_null`. |
| `ProductFamilyIds` | `[]int` | Optional | Allows fetching components allocation with matching product family id based on provided ids. Use in query `product_family_ids=1,2,3`. |
| `Include` | [`*models.ListSubscriptionComponentsInclude`](../../doc/models/list-subscription-components-include.md) | Optional | Allows including additional data in the response. Use in query `include=subscription,historic_usages`. |

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
    listSubscriptionComponentsForSiteInput := models.ListSubscriptionComponentsForSiteInput{
        Page:                 models.ToPointer(1),
        PerPage:              models.ToPointer(50),
        Sort:                 models.ToPointer(models.ListSubscriptionComponentsSort_UPDATEDAT),
        Direction:            models.ToPointer(models.SortingDirection_ASC),
        Filter:               models.ToPointer(models.ListSubscriptionComponentsForSiteFilter{
            Currencies:           []string{
                "EUR",
                "USD",
            },
            UseSiteExchangeRate:  models.ToPointer(false),
            Subscription:         models.ToPointer(models.SubscriptionFilter{
                States:               []models.SubscriptionStateFilter{
                    models.SubscriptionStateFilter_ACTIVE,
                    models.SubscriptionStateFilter_CANCELED,
                    models.SubscriptionStateFilter_EXPIRED,
                },
                DateField:            models.ToPointer(models.SubscriptionListDateField_UPDATEDAT),
                StartDate:            models.ToPointer(parseTime(models.DEFAULT_DATE, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
                EndDate:              models.ToPointer(parseTime(models.DEFAULT_DATE, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
                StartDatetime:        models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
            }),
        }),
        DateField:            models.ToPointer(models.SubscriptionListDateField_UPDATEDAT),
        StartDate:            models.ToPointer("start_date6"),
        StartDatetime:        models.ToPointer("start_datetime0"),
        EndDate:              models.ToPointer("end_date0"),
        EndDatetime:          models.ToPointer("end_datetime8"),
        SubscriptionIds:      []int{
            1,
            2,
            3,
        },
        PricePointIds:        models.ToPointer(models.IncludeNotNull_NOTNULL),
        ProductFamilyIds:     []int{
            1,
            2,
            3,
        },
        Include:              models.ToPointer(models.ListSubscriptionComponentsInclude_SUBSCRIPTION),
    }

}
```

