package main

import (
	"fmt"
	"urlshortener/Database"
    "urlshortener/Routs"

	"github.com/labstack/echo/v4"
)

func main() {

	Database.ConnectDatabase()
	Database.ModelMigration()

	e := echo.New()

	Routes.SetupRoutes(e)

	fmt.Println("Server Run On Port 8082...")
	e.Logger.Fatal(e.Start(":8082"))
}