
# Customer

## Structure

`Customer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `FirstName` | `*string` | Optional | The first name of the customer |
| `LastName` | `*string` | Optional | The last name of the customer |
| `Email` | `*string` | Optional | The email address of the customer |
| `CcEmails` | `Optional[string]` | Optional | A comma-separated list of emails that should be cc’d on all customer communications (i.e. “joe@example.com, sue@example.com”) |
| `Organization` | `Optional[string]` | Optional | The organization of the customer |
| `Reference` | `Optional[string]` | Optional | The unique identifier used within your own application for this customer |
| `Id` | `*int` | Optional | The customer ID in Chargify |
| `CreatedAt` | `*string` | Optional | The timestamp in which the customer object was created in Chargify |
| `UpdatedAt` | `*string` | Optional | The timestamp in which the customer object was last edited |
| `Address` | `Optional[string]` | Optional | The customer’s shipping street address (i.e. “123 Main St.”) |
| `Address2` | `Optional[string]` | Optional | Second line of the customer’s shipping address i.e. “Apt. 100” |
| `City` | `Optional[string]` | Optional | The customer’s shipping address city (i.e. “Boston”) |
| `State` | `Optional[string]` | Optional | The customer’s shipping address state (i.e. “MA”) |
| `StateName` | `Optional[string]` | Optional | The customer's full name of state |
| `Zip` | `Optional[string]` | Optional | The customer’s shipping address zip code (i.e. “12345”) |
| `Country` | `Optional[string]` | Optional | The customer shipping address country |
| `CountryName` | `Optional[string]` | Optional | The customer's full name of country |
| `Phone` | `Optional[string]` | Optional | The phone number of the customer |
| `Verified` | `Optional[bool]` | Optional | Is the customer verified to use ACH as a payment method. Available only on Authorize.Net gateway |
| `PortalCustomerCreatedAt` | `Optional[string]` | Optional | The timestamp of when the Billing Portal entry was created at for the customer |
| `PortalInviteLastSentAt` | `Optional[string]` | Optional | The timestamp of when the Billing Portal invite was last sent at |
| `PortalInviteLastAcceptedAt` | `Optional[string]` | Optional | The timestamp of when the Billing Portal invite was last accepted |
| `TaxExempt` | `*bool` | Optional | The tax exempt status for the customer. Acceptable values are true or 1 for true and false or 0 for false. |
| `VatNumber` | `Optional[string]` | Optional | The VAT business identification number for the customer. This number is used to determine VAT tax opt out rules. It is not validated when added or updated on a customer record. Instead, it is validated via VIES before calculating taxes. Only valid business identification numbers will allow for VAT opt out. |
| `ParentId` | `Optional[int]` | Optional | The parent ID in Chargify if applicable. Parent is another Customer object. |
| `Locale` | `Optional[string]` | Optional | The locale for the customer to identify language-region |
| `DefaultSubscriptionGroupUid` | `Optional[string]` | Optional | - |

## Example (as JSON)

```json
{
  "first_name": "first_name8",
  "last_name": "last_name6",
  "email": "email8",
  "cc_emails": "cc_emails2",
  "organization": "organization8"
}
```

