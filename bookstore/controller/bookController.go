package controller

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/pradytpk/go-clean-arch/bookstore/intf"
)

type BookController struct {
	bookService intf.BookService
}

func NewBookContoller(echoCtx *echo.Echo, bookServiceObject intf.BookService) {
	bookControllerObject := &BookController{
		bookService: bookServiceObject,
	}
	echoCtx.GET("/printAuthor", bookControllerObject.PrintAuthor)
	echoCtx.GET("/text", bookControllerObject.TextAuthor)
}

func (controllerObj *BookController) PrintAuthor(echoCtx echo.Context) error {

	return nil

}
func (controllerObj *BookController) TextAuthor(echoCtx echo.Context) error {
	fmt.Println("Inside Book controller")
	requestContext := echoCtx.Request().Context()
	controllerObj.bookService.TextBookService(requestContext)

	return nil
}
