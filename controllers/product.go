package controllers

import (
	"echo-hello/models"
	"net/http"

	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

//get All Products
func GetProducts(c echo.Context) error {
	products := models.GetAllProducts()
	return c.JSON(http.StatusOK, products)
}

// get Product
func GetProduct(c echo.Context) error {
	sku := c.Param("sku")
	product := models.GetProduct(sku)
	return c.JSON(http.StatusOK, product)
}

// update Product
func UpdateProduct(c echo.Context) (err error) {
	p := new(models.Product)
	response := new(models.ErrorNotice)
	sku := c.Param("sku")

	if err = c.Bind(p); err != nil {
		response.Type = "error"
		response.Message = err.Error()
		return c.JSON(http.StatusOK, response)
	}
	if err = c.Validate(p); err != nil {
		response.Type = "error"
		response.Message = err.Error()
		return c.JSON(http.StatusOK, response)
	}

	p.Sku = uuid.Must(uuid.FromString(sku))
	product := models.UpdateProduct(*p)
	return c.JSON(http.StatusOK, product)
}

// create Product
func CreateProduct(c echo.Context) (err error) {
	p := new(models.Product)
	response := new(models.ErrorNotice)

	if err = c.Bind(p); err != nil {
		response.Type = "error"
		response.Message = err.Error()
		return c.JSON(http.StatusOK, response)
	}
	if err = c.Validate(p); err != nil {
		response.Type = "error"
		response.Message = err.Error()
		return c.JSON(http.StatusOK, response)
	}
	product := models.CreateProduct(*p)
	return c.JSON(http.StatusOK, product)
}

// delete Product
func DeleteProduct(c echo.Context) error {
	response := new(models.ErrorNotice)
	sku := c.Param("sku")
	result := models.DeleteProduct(sku)
	if result {
		response.Message = "Product has been deleted successfully"
		response.Type = "success"
	} else {
		response.Message = "No product found!"
		response.Type = "error"
	}
	return c.JSON(http.StatusOK, response)
}
