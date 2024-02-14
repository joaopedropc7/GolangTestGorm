package models

type Product struct {
	ProductId    int64 `gorm:"primaryKey"`
	ProductName  string
	Quantity     int32
	CostPrice    float64
	SellPrice    float64
	QuantitySell int32
}

type ProductRequestVO struct {
	ProductName string
	CostPrice   float64
	SellPrice   float64
}
