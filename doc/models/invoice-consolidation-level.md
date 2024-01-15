
# Invoice Consolidation Level

Consolidation level of the invoice, which is applicable to invoice consolidation.  It will hold one of the following values:

* "none": A normal invoice with no consolidation.
* "child": An invoice segment which has been combined into a consolidated invoice.
* "parent": A consolidated invoice, whose contents are composed of invoice segments.

"Parent" invoices do not have lines of their own, but they have subtotals and totals which aggregate the member invoice segments.

See also the [invoice consolidation documentation](https://chargify.zendesk.com/hc/en-us/articles/4407746391835).

## Enumeration

`InvoiceConsolidationLevel`

## Fields

| Name |
|  --- |
| `NONE` |
| `CHILD` |
| `PARENT` |

