
# Public Signup Page

## Structure

`PublicSignupPage`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | The id of the signup page (public_signup_pages only) |
| `ReturnUrl` | `Optional[string]` | Optional | The url to which a customer will be returned after a successful signup (public_signup_pages only) |
| `ReturnParams` | `Optional[string]` | Optional | The params to be appended to the return_url (public_signup_pages only) |
| `Url` | `*string` | Optional | The url where the signup page can be viewed (public_signup_pages only) |

## Example (as JSON)

```json
{
  "id": 6,
  "return_url": "return_url0",
  "return_params": "return_params2",
  "url": "url8"
}
```

