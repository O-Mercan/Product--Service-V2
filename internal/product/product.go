package product

import (
	"errors"

	//"github.com/logrusutil/v1/errfield"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	ErrProductNotFound  = errors.New("Product Not Found")
	ErrProductsNotFound = errors.New("Products Not Found")
	ErrProductUpdate    = errors.New("Product Not Update")
	ErrProductExist     = errors.New("Product already exist")
)

/* func GetErrorFields(err error) log.Fields {
	var e *errfield.Error
	if errors.As(err, &e) {
		//e.Fields are map[string]interface{} the same as logrus.Fieldsreturn e.Fields.(logrus.Fields)
	}
	return log.Fields{}
} */

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

/* log.WithFields(log.Fields{
	"file": "product.go",
	"line": "33",
}).Error("ID doesn't exist")  */

// ProductsService - the interface for our product service
type ProductsService interface {
	GetProduct(ID uint) (Product, error)
	PostProduct(product Product) (Product, error)
	PutProduct(ID uint, newProduct Product) (Product, error)
	GetProducts() (Product, error)
	DeleteProduct(ID uint) (Product, error)
}

// GetProduct - retrieves comments by their ID from the database
func (s *Service) GetProduct(ID uint) (Product, error) {
	var product Product

	if result := s.DB.First(&product, ID); result.Error != nil {
		log.WithFields(log.Fields{
			"productID": product.ID,
			"requestID": ID,
			"Function":  "GetProduct",
		}).Error("ID doesn't exist")
		return Product{}, ErrProductNotFound
	}
	return product, nil
}

// PostProduct - adds a new product to the database
func (s *Service) PostProduct(product Product) (Product, error) {
	if result := s.DB.Save(&product); result.Error != nil {
		log.WithFields(log.Fields{
			"Function": "PostProduct",
		}).Error("Product can't sending")
		return Product{}, ErrProductUpdate
	}
	return product, nil
}

// GetProducts
func (s *Service) GetProducts() (Product, error) {
	var product Product
	if result := s.DB.Scan(&product); result.Error != nil {
		log.WithFields(log.Fields{
			"Function": "GetProducts",
		}).Error("Failed to get product")
		return Product{}, ErrProductsNotFound
	}
	return product, nil
}

// UpdateProduct
func (s *Service) PutProduct(ID uint, newProduct Product) (Product, error) {
	var product Product
	pr, err := s.GetProduct(ID)
	if err != nil {
		log.WithFields(log.Fields{
			"productID": product.ID,
			"requestID": ID,
			"Function":  "PutProduct",
		}).Error("ID doesn't exist")
		return Product{}, err
	}
	if result1 := s.DB.Model(&pr).Updates(newProduct); result1.Error != nil {

		log.Error("PutProduct error")
		return Product{}, ErrProductExist
	}
	return product, nil
}

//DeleteProduct - delete a product from the database
func (s *Service) DeleteProduct(ID uint) (Product, error) {
	var product Product
	if result := s.DB.Delete(&product, ID); result.Error != nil {
		log.WithFields(log.Fields{
			"Function": "DeleteProduct",
		}).Error("ID doesn't exist")
		return Product{}, ErrProductNotFound
	}
	return product, nil
}

// NewService - returns a new product service
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}
