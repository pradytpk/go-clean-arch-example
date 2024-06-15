package services

import (
	"context"
	"fmt"

	"github.com/pradytpk/go-clean-arch/bookstore/intf"
	"github.com/pradytpk/go-clean-arch/bookstore/model"
)

type BookServiceImpl struct {
	bookDl intf.BookDatalayer
}

// TextBookService implements intf.BookService.
func (b *BookServiceImpl) TextBookService(ctx context.Context) {
	fmt.Println("Inside book service")
	b.bookDl.TextBookDatalayer(ctx)
}

// printBookTitle implements intf.BookService.
func (b *BookServiceImpl) PrintBookTitle(ctx context.Context, book *model.Book) {
	panic("unimplemented")
}

func NewBookServiceImpl(bookDatalayer intf.BookDatalayer) intf.BookService {
	return &BookServiceImpl{
		bookDl: bookDatalayer,
	}
}
