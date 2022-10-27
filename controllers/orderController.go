package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"rest-api/database"
	"rest-api/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Results struct {
	Customer_name string `json:"customerName"`
	Items         []Item `json:"items"`
}

type Item struct {
	Item_ID     int    `json:"lineItemId"`
	Item_code   string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

func CreateOrder(ctx *gin.Context) {
	db := database.GetDB()

	var jsonString = `{
			"customerName": "Ghea",			
			"items": 
				[{					
					"itemCode": "123",
					"description": "Iphone 10X",
					"quantity": 10
				}]
	}
	`
	now := time.Now()
	var result = Results{}

	var err = json.Unmarshal([]byte(jsonString), &result)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	Order := models.Order{
		Customer_name: result.Customer_name,
		Ordered_at:    now,
	}

	errr := db.Create(&Order).Error
	if errr != nil {
		fmt.Println("Error creating order data:", errr)
		return
	}

	fmt.Println("New Order Data:", Order)

	Item := models.Item{
		Item_code:   result.Items[0].Item_code,
		Description: result.Items[0].Description,
		Quantity:    result.Items[0].Quantity,
		Order_ID:    Order.Order_ID,
	}

	err = db.Create(&Item).Error

	if err != nil {
		fmt.Println("Error creating item data:", err.Error())
		return
	}

	fmt.Println("New item Data:", Item)

	orders := models.Order{}
	err = db.Preload("Items").Find(&orders).Error

	if err != nil {
		fmt.Println("Error getting order datas with items:", err.Error())
		return
	}

	fmt.Println("User Orders with Items")
	fmt.Printf("%+v", orders)

	ctx.JSON(http.StatusCreated, gin.H{
		"order": orders,
	})
}

func GetOrderById(ctx *gin.Context) {
	db := database.GetDB()

	order := models.Order{}
	orderId := ctx.Param("orderId")

	err := db.Preload("Items").First(&order, "order_id = ?", orderId).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error_status":  "Data Not Found",
				"error_message": fmt.Sprintf("Order witd id %v not found", orderId),
			})
			return
		}
		print("Error finding order:", err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"order": order,
	})
}

func UpdateOrder(ctx *gin.Context) {
	db := database.GetDB()
	orderId := ctx.Param("orderId")

	var jsonString = `{
			"customerName": "Hanasui",			
			"items": 
				[{
					"lineItemId": 1,
					"itemCode": "123",
					"description": "Samsung X",
					"quantity": 10
				}]
	}
	`

	var result = Results{}

	var err = json.Unmarshal([]byte(jsonString), &result)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	now := time.Now()
	Order := models.Order{
		Customer_name: result.Customer_name,
		Ordered_at:    now,
	}

	orders := models.Order{}

	errr := db.Model(&orders).Where("order_id = ?", orderId).Updates(Order).Error
	if errr != nil {
		fmt.Println("Error updating order data:", errr)
		return
	}

	fmt.Println("Update Order Data:", Order)

	Item := models.Item{
		Item_code:   result.Items[0].Item_code,
		Description: result.Items[0].Description,
		Quantity:    result.Items[0].Quantity,
	}

	items := models.Item{}

	err = db.Model(&items).Where("order_id = ?", orderId).Updates(Item).Error

	if err != nil {
		fmt.Println("Error updating item data:", err.Error())
		return
	}

	fmt.Println("Update item Data:", Item)

	err = db.Preload("Items").Find(&orders).Error

	if err != nil {
		fmt.Println("Error getting order datas with items:", err.Error())
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("order with id %v not found", orderId),
		})
		return
	}

	fmt.Println("Orders with Items")
	fmt.Printf("%+v", orders)

	ctx.JSON(http.StatusOK, gin.H{
		"order":   orders,
		"message": fmt.Sprintf("order with id %v has been successfulyy updated", orderId),
	})
}

func DeleteOrderById(ctx *gin.Context) {
	db := database.GetDB()
	orderId := ctx.Param("orderId")

	item := models.Item{}
	err := db.Where("order_id = ?", orderId).Delete(&item).Error

	if err != nil {
		fmt.Println("Error deleting item:", err.Error())
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("item with order id %v not found", orderId),
		})
		return
	}

	order := models.Order{}
	err = db.Where("order_id = ?", orderId).Delete(&order).Error

	if err != nil {
		fmt.Println("Error deleting order:", err.Error())
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("order with id %v not found", orderId),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("order and item with order id %v has been successfully deleted", orderId),
	})
}
