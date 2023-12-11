
# RetryConfiguration

The following parameters are configurable for the RetryConfiguration:

## Properties

| Name | Type | Description |
|  --- | --- | --- |
| `maxRetryAttempts` | `int64` | Maximum number of retries.<br>*Default*: `0` |
| `retryOnTimeout` | `bool` | Whether to retry on request timeout.<br>*Default*: `true` |
| `retryInterval` | `time.Duration` | Interval before next retry. Used in calculation of wait time for next request in case of failure.<br>*Default*: `1` |
| `maximumRetryWaitTime` | `time.Duration` | Overall wait time for the requests getting retried.<br>*Default*: `0` |
| `backoffFactor` | `int64` | Used in calculation of wait time for next request in case of failure.<br>*Default*: `2` |
| `httpStatusCodesToRetry` | `[]int64` | Http status codes to retry against.<br>*Default*: `[]int64{408, 413, 429, 500, 502, 503, 504, 521, 522, 524}` |
| `httpMethodsToRetry` | `[]string` | Http methods to retry against.<br>*Default*: `[]string{"GET", "PUT"}` |

The retryConfiguration can be initialized as follows:

```go
retryConfiguration := CreateRetryConfiguration(
    maxioadvancedbilling.WithMaxRetryAttempts(0),
    maxioadvancedbilling.WithRetryOnTimeout(true),
    maxioadvancedbilling.WithRetryInterval(1),
    maxioadvancedbilling.WithMaximumRetryWaitTime(0),
    maxioadvancedbilling.WithBackoffFactor(2),
    maxioadvancedbilling.WithHttpStatusCodesToRetry([]int64{408, 413, 429, 500, 502, 503, 504, 521, 522, 524}),
    maxioadvancedbilling.WithHttpMethodsToRetry([]string{"GET", "PUT"}),
)
```

