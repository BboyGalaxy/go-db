## Migrar tabla de productos

```go
storage.NewPostgresDB()
storageProduct := storage.NewPsqlProduct(storage.Pool())
serviceProduct := product.NewService(storageProduct)

if err := serviceProduct.Migrate(); err != nil {
	log.Fatalf("product.Migrate: %v", err)
}
```

## Migrar tabla de invoiceheader

```go
storage.NewPostgresDB()
storageInvoiceHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
serviceInvoiceHeader := invoiceheader.NewService(storageInvoiceHeader)

if err := serviceInvoiceHeader.Migrate(); err != nil {
	log.Fatalf("invoiceHeader.Migrate: %v", err)
}
```

## Migrar tabla de invoiceitem

```go
storage.NewPostgresDB()
storageInvoiceItem := storage.NewPsqlInvoiceItem(storage.Pool())
serviceInvoiceItem := invoiceitem.NewService(storageInvoiceItem)

if err := serviceInvoiceItem.Migrate(); err != nil {
	log.Fatalf("invoiceItem.Migrate: %v", err)
}
```

## Agregar productos

```go
m := &product.Model{
		Name:        "Curso de db con Go",
		Price:       70,
		Observation: "On fire",
	}
```

### Buscar prodcutos por ID

```go
	storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	m, err := serviceProduct.GetByID(3)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		fmt.Println("No hay producto con este ID")
	case err != nil:
		log.Fatalf("product.GetById: %v", err)
	default:
		fmt.Println(m)
	}
```

## Actualizar productos

```go
storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	m := &product.Model{
		ID:          3,
		Name:        "Curso de Python",
		Observation: "this is the course of python",
		Price:       120,
	}
	err := serviceProduct.Update(m)

	if err != nil {
		log.Fatalf("product.Update: %v", err)
	}
```