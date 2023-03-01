package invoiceheader

import "time"

// Model of invoiceheader
type Model struct {
	ID        uint
	Client    string
	CreatedAt time.Time
	UpdatedAt time.Time
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

// Migrate is used for migrate product // Create table invoiceheader
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
