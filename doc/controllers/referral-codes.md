# Referral Codes

```go
referralCodesController := client.ReferralCodesController()
```

## Class Name

`ReferralCodesController`


# Validate Referral Code

Use this method to determine if the referral code is valid and applicable within your Site. This method is useful for validating referral codes that are entered by a customer.

## Referrals Documentation

Full documentation on how to use the referrals feature in the Chargify UI can be located [here](https://chargify.zendesk.com/hc/en-us/articles/4407802831643).

## Server Response

If the referral code is valid the status code will be `200` and the referral code will be returned. If the referral code is invalid, a `404` response will be returned.

```go
ValidateReferralCode(
    ctx context.Context,
    code string) (
    models.ApiResponse[models.ReferralValidationResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `code` | `string` | Query, Required | The referral code you are trying to validate |

## Response Type

[`models.ReferralValidationResponse`](../../doc/models/referral-validation-response.md)

## Example Usage

```go
ctx := context.Background()

code := "code8"

apiResponse, err := referralCodesController.ValidateReferralCode(ctx, code)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "referral_code": {
    "id": 1032514,
    "site_id": 31615,
    "subscription_id": 16254270,
    "code": "9b6cdw"
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 404 | Not Found | [`SingleStringErrorResponseException`](../../doc/models/single-string-error-response-exception.md) |

