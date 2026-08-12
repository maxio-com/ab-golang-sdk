# Referral Codes

```go
referralCodesController := client.ReferralCodesController()
```

## Class Name

`ReferralCodesController`


# Validate Referral Code

Validates whether a referral code is valid and applicable within your site. This method is useful for validating referral codes that are entered by a customer.

For more information, see [Understanding Referrals](https://docs.maxio.com/hc/en-us/articles/24286981223693-Understanding-Referrals) in the product documentation.

```go
ValidateReferralCode(
    ctx context.Context,
    code string) (
    models.ApiResponse[models.ReferralValidationResponse],
    error)
```

## Authentication

This endpoint requires [BasicAuth](../../doc/auth/basic-authentication.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `code` | `string` | Query, Required | The referral code you are trying to validate |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ReferralValidationResponse](../../doc/models/referral-validation-response.md).

## Example Usage

```go
ctx := context.Background()

code := "code8"

apiResponse, err := referralCodesController.ValidateReferralCode(ctx, code)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.SingleStringErrorResponse:
            log.Fatalln("SingleStringErrorResponseException: ", typedErr)
        default:
            log.Fatalln(err)
    }
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

