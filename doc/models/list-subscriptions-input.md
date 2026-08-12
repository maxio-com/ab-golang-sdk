
# List Subscriptions Input

Input structure for the method ListSubscriptions

## Structure

`ListSubscriptionsInput`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Page` | `*int` | Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |
| `PerPage` | `*int` | Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`.<br><br>**Default**: `20`<br><br>**Constraints**: `<= 200` |
| `Sort` | [`*models.SubscriptionSort`](../../doc/models/subscription-sort.md) | Optional | The attribute by which to sort<br><br>**Default**: `"signup_date"` |
| `Direction` | [`*models.SortingDirection`](../../doc/models/sorting-direction.md) | Optional | Controls the order in which results are returned.<br>Use in query `direction=asc`. |
| `State` | [`*models.SubscriptionStateFilter`](../../doc/models/subscription-state-filter.md) | Optional | The current state of the subscription |
| `Product` | [`*models.ListSubscriptionsInputProduct`](../../doc/models/containers/list-subscriptions-input-product.md) | Optional | This is a container for one-of cases. |
| `Q` | `*string` | Optional | Search string. |
| `QScope` | [`*models.QScope`](../../doc/models/q-scope.md) | Optional | Scope of fields used by the q search. |
| `CustomerId` | `*int` | Optional | The Advanced Billing id of the customer. |
| `ProductPricePointId` | `*int` | Optional | The ID of the product price point. If supplied, product is required. |
| `Coupon` | `*int` | Optional | The numeric id of the coupon currently applied to the subscription. (This can be found in the URL when editing a coupon. Note that the coupon code cannot be used.) |
| `CouponCode` | `*string` | Optional | The coupon code currently applied to the subscription |
| `CollectionMethod` | [`*models.CollectionMethod1`](../../doc/models/collection-method-1.md) | Optional | The collection method for the subscription. |
| `BrandingThemeId` | `*int` | Optional | Filter subscriptions by the ID of an assigned Branding Theme. Branding Themes is a beta feature. See [Understand Branding Themes](https://docs.maxio.com/hc/en-us/articles/43796895662093-Understand-Branding-Themes#understand-branding-themes-0-0) for more information. |
| `DateField` | [`*models.SubscriptionDateField`](../../doc/models/subscription-date-field.md) | Optional | The type of filter you'd like to apply to your search.  Allowed Values: , current_period_ends_at, current_period_starts_at, created_at, activated_at, canceled_at, expires_at, trial_started_at, trial_ended_at, updated_at |
| `StartDate` | `*time.Time` | Optional | The start date (format YYYY-MM-DD) with which to filter the date_field. Returns subscriptions with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified. Use in query `start_date=2022-07-01`. |
| `EndDate` | `*time.Time` | Optional | The end date (format YYYY-MM-DD) with which to filter the date_field. Returns subscriptions with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified. Use in query `end_date=2022-08-01`. |
| `StartDatetime` | `*time.Time` | Optional | The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns subscriptions with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of start_date. Use in query `start_datetime=2022-07-01 09:00:05`. |
| `EndDatetime` | `*time.Time` | Optional | The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns subscriptions with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of end_date. Use in query `end_datetime=2022-08-01 10:00:05`. |
| `Metadata` | `map[string]string` | Optional | The value of the metadata field specified in the parameter. Use in query `metadata[my-field]=value&metadata[other-field]=another_value`. |
| `GroupStatus` | [`*models.GroupStatus`](../../doc/models/group-status.md) | Optional | Filter by whether a subscription is in a group. |
| `DunningExemption` | `*bool` | Optional | Filter by dunning exemption status. |
| `PaymentGateways` | `*string` | Optional | Comma-separated payment gateway identifiers. |
| `Currencies` | `*string` | Optional | Comma-separated currency codes. |
| `Include` | [`[]models.SubscriptionListInclude`](../../doc/models/subscription-list-include.md) | Optional | Allows including additional data in the response. Use in query: `include[]=self_service_page_token`. |

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
    listSubscriptionsInput := models.ListSubscriptionsInput{
        Page:                 models.ToPointer(1),
        PerPage:              models.ToPointer(50),
        Sort:                 models.ToPointer(models.SubscriptionSort_SIGNUPDATE),
        Direction:            models.ToPointer(models.SortingDirection_ASC),
        State:                models.ToPointer(models.SubscriptionStateFilter_PREPAIDDUNNING),
        Product:              models.ToPointer(models.ListSubscriptionsInputProductContainer.FromNumber(200)),
        Q:                    models.ToPointer("q0"),
        QScope:               models.ToPointer(models.QScope_SUBSCRIPTIONREFERENCE),
        CustomerId:           models.ToPointer(150),
        ProductPricePointId:  models.ToPointer(234),
        Coupon:               models.ToPointer(84),
        CouponCode:           models.ToPointer("coupon_code4"),
        CollectionMethod:     models.ToPointer(models.CollectionMethod1_AUTOMATIC),
        BrandingThemeId:      models.ToPointer(76),
        DateField:            models.ToPointer(models.SubscriptionDateField_TRIALSTARTEDAT),
        StartDate:            models.ToPointer(parseTime(models.DEFAULT_DATE, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
        EndDate:              models.ToPointer(parseTime(models.DEFAULT_DATE, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
        StartDatetime:        models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
        EndDatetime:          models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
        Metadata:             map[string]string{
            "key0": "metadata3",
            "key1": "metadata4",
            "key2": "metadata5",
        },
        GroupStatus:          models.ToPointer(models.GroupStatus_UNGROUPED),
        DunningExemption:     models.ToPointer(false),
        PaymentGateways:      models.ToPointer("payment_gateways2"),
        Currencies:           models.ToPointer("currencies6"),
        Include:              []models.SubscriptionListInclude{
            models.SubscriptionListInclude_SELFSERVICEPAGETOKEN,
        },
    }

}
```

