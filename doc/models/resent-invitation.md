
# Resent Invitation

## Structure

`ResentInvitation`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `LastSentAt` | `*string` | Optional | - |
| `LastAcceptedAt` | `*string` | Optional | - |
| `SendInviteLinkText` | `*string` | Optional | - |
| `UninvitedCount` | `*int` | Optional | - |
| `LastInviteSentAt` | `*time.Time` | Optional | - |
| `LastInviteAcceptedAt` | `*time.Time` | Optional | - |

## Example (as JSON)

```json
{
  "last_invite_sent_at": "01/01/2024 04:30:00",
  "last_invite_accepted_at": "01/01/2024 04:35:00",
  "last_sent_at": "last_sent_at8",
  "last_accepted_at": "last_accepted_at8",
  "send_invite_link_text": "send_invite_link_text6",
  "uninvited_count": 6
}
```

