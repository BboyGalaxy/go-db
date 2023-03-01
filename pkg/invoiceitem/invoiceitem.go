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

type Storage interface {
	Migrate() error
}

// Service of invoiceitem
type Service struct {
	storage Storage
}

// NewService return a pointer of service // Constructor of Service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Migrate is used for migrate product // Create table invoiceitem
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
