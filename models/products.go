package models

import (
	"echo-hello/db"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type Product struct {
	Sku         uuid.UUID `json:"sku" xml:"sku" gorm:"primary_key"`
	Name        string    `json:"name" xml:"name" validate:"required"`
	Description string    `json:"description" xml:"description" validate:"required"`
	Price       float64   `json:"price" xml:"price" validate:"required,numeric"`
}

// get all products
func GetAllProducts() []Product {
	db := db.GetDB()
	var rows []Product
	if db != nil {
		db.Find(&rows)
	}
	return rows
}

// get product
func GetProduct(sku string) Product {
	db := db.GetDB()
	var row Product
	if db != nil {
		db.Where("sku=?", sku).Find(&row)
	}
	return row
}

// update products
func UpdateProduct(data Product) Product {
	db := db.GetDB()
	sku := data.Sku

	var p Product

	fmt.Println(data)
	db.Where("sku=?", sku).Find(&p)
	if p.Sku != EmptySku {
		p = data
		fmt.Println(p)
		db.Save(&p)
	}
	return p
}

// create product
func CreateProduct(data Product) Product {
	db := db.GetDB()
	data.Sku = uuid.Must(uuid.NewV4())
	db.Create(&data)
	return data
}

// delete product
func DeleteProduct(sku string) bool {
	db := db.GetDB()
	p := Product{}
	db.Where("sku=?", sku).Find(&p)
	if p.Sku != EmptySku {
		db.Delete(&p)
		return true
	}
	return false
}
