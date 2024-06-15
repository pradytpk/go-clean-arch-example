package main

import (
	"github.com/labstack/echo/v4"
	"github.com/pradytpk/go-clean-arch/bookstore/controller"
	"github.com/pradytpk/go-clean-arch/bookstore/database"
	"github.com/pradytpk/go-clean-arch/bookstore/services"
)

func main() {
	echoContext := echo.New()
	datalayer := database.NewBookDatalayerImpl()
	service := services.NewBookServiceImpl(datalayer)
	controller.NewBookContoller(echoContext, service)
	echoContext.Logger.Fatal(echoContext.Start(":1323"))
}
