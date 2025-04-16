package main

import (
	"fmt"
	"os"

	"e-commerce-return-tracker/connectors"

	"e-commerce-return-tracker/repository"

	"github.com/labstack/echo/v4"
)

func init() {
	connectors.LoadEnv()
}

func main() {
	fmt.Println("Main func created")
	repository.SingleMysqlConnection()
	e := echo.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "9001"
	}
	e.Logger.Fatal(e.Start(":" + port))

}
