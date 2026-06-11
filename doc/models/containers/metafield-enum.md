
# Metafield Enum

## Class Name

`MetafieldEnum`

## Cases

| Type | Factory Method |
|  --- | --- |
| `string` | models.MetafieldEnumContainer.FromString(string mString) |
| `[]string` | models.MetafieldEnumContainer.FromArrayOfString([]string arrayOfString) |

## string

### Initialization Code

#### Example

```go
value := models.MetafieldEnumContainer.FromString("String0")
```

## []string

### Initialization Code

#### Example

```go
value := models.MetafieldEnumContainer.FromArrayOfString([]string{
    "String1",
})
```

