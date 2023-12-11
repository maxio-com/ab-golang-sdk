
# List Public Keys Response

## Structure

`ListPublicKeysResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ChargifyJsKeys` | [`[]models.PublicKey`](public-key.md) | Optional | - |
| `Meta` | [`*models.ListPublicKeysMeta`](list-public-keys-meta.md) | Optional | - |

## Example (as JSON)

```json
{
  "chargify_js_keys": [
    {
      "public_key": "public_key8",
      "requires_security_token": false,
      "created_at": "created_at6"
    },
    {
      "public_key": "public_key8",
      "requires_security_token": false,
      "created_at": "created_at6"
    }
  ],
  "meta": {
    "total_count": 150,
    "current_page": 126,
    "total_pages": 138,
    "per_page": 152
  }
}
```

