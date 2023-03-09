package product

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var (
	ErrIDNotFound = errors.New("El producto no contiene un ID!")
)

// Model of product
type Model struct {
	ID          uint
	Name        string
	Observation string
	Price       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (m *Model) String() string {
	return fmt.Sprintf("%02d | %-20s | %-20s | %5d | %10s | %10s",
		m.ID, m.Name, m.Observation, m.Price,
		m.CreatedAt.Format("2006-01-02"), m.UpdatedAt.Format("2006-01-02"))
}

type Models []*Model

func (m Models) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("%02s | %-20s | %-20s | %5s | %10s | %10s\n",
		"ID", "Name", "Observations", "Price", "CreatedAt", "UpdatedAt"))
	for _, model := range m {
		builder.WriteString(model.String() + "\n")
	}
	return builder.String()
}

type Storage interface {
	Migrate() error
	Create(*Model) error
	Update(*Model) error
	GetAll() (Models, error)
	GetByID(uint) (*Model, error)
	Delete(uint) error
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

// Create is used for Create a product
func (s *Service) Create(m *Model) error {
	m.CreatedAt = time.Now()
	return s.storage.Create(m)
}

// GetAll is used for get all the products
func (s *Service) GetAll() (Models, error) {
	return s.storage.GetAll()
}

// GetByID is used for get a product
func (s *Service) GetByID(id uint) (*Model, error) {
	return s.storage.GetByID(id)
}

// Update is used for update a product
func (s *Service) Update(m *Model) error {
	if m.ID == 0 {
		return ErrIDNotFound
	}

	m.UpdatedAt = time.Now()

	return s.storage.Update(m)
}

// Delete is used for delete a product
func (s *Service) Delete(id uint) error {
	return s.storage.Delete(id)
}
