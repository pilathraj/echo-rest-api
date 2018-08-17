package main

import (
	"echo-hello/controllers"
	"echo-hello/db"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	validator "gopkg.in/go-playground/validator.v9"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	db := db.Init()
	defer db.Close()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello Pilathraj")
	})
	// Routes
	e.GET("/products", controllers.GetProducts)
	e.GET("/products/:sku", controllers.GetProduct)
	e.PUT("/products/:sku", controllers.UpdateProduct)
	e.POST("/products", controllers.CreateProduct)
	e.DELETE("/products/:sku", controllers.DeleteProduct)

	// Start server
	e.Logger.Fatal(e.Start(":1234"))
}
