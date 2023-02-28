package invoiceitem

import "time"

// Model of invoiceitem
type Model struct {
	ID              uint
	InvoiceHeaderID uint
	ProdcutID       uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
