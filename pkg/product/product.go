package product

import "time"

// Model of product
type Model struct {
	ID          uint
	Name        string
	Observation string
	Price       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Models []*Model

type Storage interface {
	Migrate() error
	//Create(*Model) error
	//Update(*Model) error
	//GetAll() (Models, error)
	//GetByID(uint) (*Model, error)
	//Delete(uint) error
}

// Service of product
type Service struct {
	storage Storage
}

// NewService return a pointer of service // Constructor of Service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Migrate is used for migrate product // Create table products
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
