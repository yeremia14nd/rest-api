package models

type Item struct {
	Item_ID     uint   `gorm:"primaryKey"`
	Item_code   string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	Order_ID    uint   `gorm:"foreignKey:Order_ID"`
}
