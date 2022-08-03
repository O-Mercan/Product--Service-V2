package product

import "gorm.io/gorm"

// Service - the struct four our product service
type Service struct {
	DB *gorm.DB
}

// Product - defines our product structure
type Product struct {
	gorm.Model
	Category    string
	Description string
	Price       int
}

// ProductsService - the interface for our product service
type ProductsService interface {
	GetProduct(ID uint) (Product, error)
	PostProduct(product Product) (Product, error)
}

// GetProduct - retrieves comments by their ID from the database
func (s *Service) GetProduct(ID uint) (Product, error) {
	var product Product
	if result := s.DB.First(&product, ID); result.Error != nil {
		return Product{}, result.Error
	}
	return product, nil
}

// PostProduct - adds a new product to the database
func (s *Service) PostProduct(product Product) (Product, error) {
	if result := s.DB.Save(&product); result.Error != nil {
		return Product{}, result.Error
	}
	return product, nil
}

// GetProducts
func (s *Service) GetProducts() (Product, error) {
	var product Product
	if result := s.DB.Find(&product); result.Error != nil {
		return Product{}, result.Error
	}
	return product, nil
}

// UpdateProduct
func (s *Service) PutProduct(ID uint) (Product, error) {
	var product Product
	s.DB.Delete(&product, ID)
	if result1 := s.DB.Updates(&product); result1.Error != nil {
		return Product{}, result1.Error
	}
	return product, nil
}

//DeleteProduct - delete a product from the database
func (s *Service) DeleteProduct(ID uint) (Product, error) {
	var product Product
	if result := s.DB.Delete(&product, ID); result.Error != nil {
		return Product{}, result.Error
	}
	return product, nil
}

// NewService - returns a new product service
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}
