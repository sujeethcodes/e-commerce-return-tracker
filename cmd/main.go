package main

import (
	"fmt"
	"os"

	"e-commerce-return-tracker/connectors"
	"e-commerce-return-tracker/controller"

	"e-commerce-return-tracker/repository"

	"github.com/labstack/echo/v4"
)

type container struct {
	returnProductInstance controller.ReturnProduct
}

func init() {
	connectors.LoadEnv()
}

func LoadContainer() *container {
	return &container{
		returnProductInstance: controller.ReturnProduct{Mysql: repository.SingleMysqlConnection()},
	}
}

func main() {
	fmt.Println("Main func created")
	containerInstance := LoadContainer()
	e := echo.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "9001"
	}
	e.POST("/return-products", containerInstance.returnProductInstance.ReturnProducts)
	e.Logger.Fatal(e.Start(":" + port))

}
