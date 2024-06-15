package intf

import (
	"context"

	"github.com/pradytpk/go-clean-arch/bookstore/model"
)

type BookService interface {
	PrintBookTitle(ctx context.Context, book *model.Book)
	TextBookService(ctx context.Context)
}

type BookDatalayer interface {
	GetBookByID(ctx context.Context, bookID int16)
	TextBookDatalayer(ctx context.Context)
}
