
# Event Based Billing Segment Error

## Structure

`EventBasedBillingSegmentError`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Segments` | `map[string]interface{}` | Required | The key of the object would be a number (an index in the request array) where the error occurred. In the value object, the key represents the field and the value is an array with error messages. In most cases, this object would contain just one key. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    eventBasedBillingSegmentError := models.EventBasedBillingSegmentError{
        Segments:             map[string]interface{}{
            "key0": interface{}("[key1, val1][key2, val2]"),
            "key1": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

