package main

import (
	"fmt"
	"os"

	"e-commerce-return-tracker/connectors"

	"github.com/labstack/echo/v4"
)

func init() {
	connectors.LoadEnv()
}

func main() {
	fmt.Println("Main func created")
	e := echo.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "9001"
	}
	e.Logger.Fatal(e.Start(":" + port))

}
