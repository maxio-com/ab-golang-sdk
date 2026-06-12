
# Create Component Price Points Request Price Points

## Class Name

`CreateComponentPricePointsRequestPricePoints`

## Cases

| Type | Factory Method |
|  --- | --- |
| [`models.CreateComponentPricePoint`](../../../doc/models/create-component-price-point.md) | models.CreateComponentPricePointsRequestPricePointsContainer.FromCreateComponentPricePoint(models.CreateComponentPricePoint createComponentPricePoint) |
| [`models.CreatePrepaidUsageComponentPricePoint`](../../../doc/models/create-prepaid-usage-component-price-point.md) | models.CreateComponentPricePointsRequestPricePointsContainer.FromCreatePrepaidUsageComponentPricePoint(models.CreatePrepaidUsageComponentPricePoint createPrepaidUsageComponentPricePoint) |

## models.CreateComponentPricePoint

### Initialization Code

#### Example

```go
value := models.CreateComponentPricePointsRequestPricePointsContainer.FromCreateComponentPricePoint(models.CreateComponentPricePoint{
    Name:                 "name0",
    PricingScheme:        models.PricingScheme_PERUNIT,
    Prices:               []models.Price{
        models.Price{
            StartingQuantity:     models.PriceStartingQuantityContainer.FromNumber(242),
            UnitPrice:            models.PriceUnitPriceContainer.FromPrecision(float64(23.26)),
        },
    },
    UseSiteExchangeRate:  models.ToPointer(true),
})
```

## models.CreatePrepaidUsageComponentPricePoint

### Initialization Code

#### Example

```go
value := models.CreateComponentPricePointsRequestPricePointsContainer.FromCreatePrepaidUsageComponentPricePoint(models.CreatePrepaidUsageComponentPricePoint{
    Name:                     "name0",
    PricingScheme:            models.PricingScheme_PERUNIT,
    Prices:                   []models.Price{
        models.Price{
            StartingQuantity:     models.PriceStartingQuantityContainer.FromNumber(242),
            UnitPrice:            models.PriceUnitPriceContainer.FromPrecision(float64(23.26)),
        },
    },
    OveragePricing:           models.OveragePricing{
        PricingScheme:        models.PricingScheme_STAIRSTEP,
    },
    UseSiteExchangeRate:      models.ToPointer(true),
})
```

