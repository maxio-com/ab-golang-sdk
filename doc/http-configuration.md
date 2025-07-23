
# HttpConfiguration

The following parameters are configurable for the HttpConfiguration:

## Properties

| Name | Type | Description | Setter | Getter |
|  --- | --- | --- | --- | --- |
| timeout | `float64` | Timeout in milliseconds.<br>*Default*: `120` | `WithTimeout` | `Timeout()` |
| transport | `httpRoundTripper` | Establishes network connection and caches them for reuse.<br>*Default*: `http.DefaultTransport` | `WithTransport` | `Transport()` |
| retryConfiguration | [`advancedbillingRetryConfiguration`](../doc/retry-configuration.md) | Configurations to retry requests.<br>*Default*: `advancedbilling.DefaultRetryConfiguration()` | `WithRetryConfiguration` | `RetryConfiguration()` |

The httpConfiguration can be initialized as follows:

```go
package main

import (
    "advancedbilling"
    "net/http"
)

func main() {
    httpConfiguration := advancedbilling.CreateHttpConfiguration(
        advancedbilling.WithTimeout(120),
        advancedbilling.WithTransport(http.DefaultTransport),
        advancedbilling.WithRetryConfiguration(advancedbilling.DefaultRetryConfiguration()),
    )
}
```

