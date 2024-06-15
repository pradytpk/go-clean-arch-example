package database

import (
	"context"
	"fmt"

	"github.com/pradytpk/go-clean-arch/bookstore/intf"
)

type BookDatalayerImpl struct {
}

// TextBookDatalayer implements intf.BookDatalayer.
func (b *BookDatalayerImpl) TextBookDatalayer(ctx context.Context) {
	fmt.Println("Inside book data layer")
}

// GetBookByID implements intf.BookDatalayer.
func (b *BookDatalayerImpl) GetBookByID(ctx context.Context, bookID int16) {
	panic("unimplemented")
}

func NewBookDatalayerImpl() intf.BookDatalayer {
	return &BookDatalayerImpl{}
}
