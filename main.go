package main

import (
	"log"

	"github.com/BboyGalaxy/go-db/pkg/invoice"
	"github.com/BboyGalaxy/go-db/pkg/invoiceheader"
	"github.com/BboyGalaxy/go-db/pkg/invoiceitem"
	"github.com/BboyGalaxy/go-db/storage"
)

func main() {
	storage.NewPostgresDB()

	storageHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
	storageItems := storage.NewPsqlInvoiceItem(storage.Pool())
	storageInvoice := storage.NewPsqlInvoice(
		storage.Pool(),
		storageHeader,
		storageItems,
	)

	m := &invoice.Model{
		Header: &invoiceheader.Model{
			Client: "Alexys",
		},
		Items: invoiceitem.Models{
			&invoiceitem.Model{ProdcutID: 1},
		},
	}

	serviceInvoice := invoice.NewService(storageInvoice)
	if err := serviceInvoice.Create(m); err != nil {
		log.Fatalf("invoice.Create: %v", err)
	}

}
