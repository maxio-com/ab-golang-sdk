
# Entity Identifier Kind

The kind of tax or business identifier held by the customer:

- `vat_eu`: an EU VAT number. Requires `vat_country` to be an EU member state code or `GB`.
- `national_tax`: a national tax ID registered outside the EU. Requires `vat_country` to be one of `AL`, `AM`, `AR`, `AU`, `BR`, `CA`, `CH`, `DZ`, `IN`, `MX`, `NO`, `NZ`, or `ZA`.
- `company_reg`: a company registration number, such as a French SIREN. No `vat_country` is required.
- `gln`: a Global Location Number. The value must be 13 digits.
- `duns`: a D-U-N-S Number. The value must be 9 digits.
- `lei`: a Legal Entity Identifier. The value must be 20 characters: 18 letters or digits followed by 2 digits.

A customer holds one identifier at a time. Saving an identifier of a different kind replaces the existing one.

## Enumeration

`EntityIdentifierKind`

## Fields

| Name |
|  --- |
| `VATEU` |
| `NATIONALTAX` |
| `COMPANYREG` |
| `GLN` |
| `DUNS` |
| `LEI` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    entityIdentifierKind := models.EntityIdentifierKind_COMPANYREG

}
```

