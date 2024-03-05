
# Portal Management Link

## Structure

`PortalManagementLink`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Url` | `*string` | Optional | - |
| `FetchCount` | `*int` | Optional | - |
| `CreatedAt` | `*time.Time` | Optional | - |
| `NewLinkAvailableAt` | `*time.Time` | Optional | - |
| `ExpiresAt` | `*time.Time` | Optional | - |
| `LastInviteSentAt` | `Optional[time.Time]` | Optional | - |

## Example (as JSON)

```json
{
  "url": "url0",
  "fetch_count": 222,
  "created_at": "2016-03-13T12:52:32.123Z",
  "new_link_available_at": "2016-03-13T12:52:32.123Z",
  "expires_at": "2016-03-13T12:52:32.123Z"
}
```

