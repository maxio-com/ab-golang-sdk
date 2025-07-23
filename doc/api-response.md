
# ApiResponse

Represents the result of an API call, including response metadata and the returned data of type `T`.

## Properties

| Name | Type | Description |
|  --- | --- | --- |
| `Response` | `*http.Response` | Returns the complete response information. |
| `Data` | `T` | Returns the response data. |

## Usage Example

```go
package main

import (
    "fmt"
    "log"
)

func main() {
    apiResponse, err := client.ExampleController().GetExampleType(ctx, body)

    if err != nil {
        log.Fatalln("Error: ", err)
    } else {
        fmt.Println("Success! Result: ", apiResponse.Data)
        fmt.Println("Status Code: ", apiResponse.Response.StatusCode)
    }
}
```

