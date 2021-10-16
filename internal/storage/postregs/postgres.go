package postregs

import (
	"Gymondo/internal/subscription"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Repository struct {
	database *gorm.DB
}

func CreateRepository(db *gorm.DB) (*Repository, error) {
	repo := &Repository{
		database: db,
	}
	return repo, nil
}

func (r *Repository) CreateUser(email, username, fullname string) error {
	mu := subscription.User{
		Email:    email,
		UserName: username,
		FullName: fullname,
	}
	err := r.database.Create(&mu).Error
	if err != nil {
		return errors.Wrap(err, "failed to create a user")
	}
	return nil
}

func (r *Repository) CreateProduct(name string) error {
	mu := subscription.Product{
		Name: name,
	}
	err := r.database.Create(&mu).Error
	if err != nil {
		return errors.Wrap(err, "failed to create a product")
	}
	return nil
}

func (r *Repository) GetUserByEmail(email string) (u subscription.User, err error) {
	err = r.database.Where("email = ?", email).First(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return u, err
	}
	return u, nil
}

func (r *Repository) GetProducts() ([]subscription.Product, error){
	var products []subscription.Product
	result := r.database.Find(&products)
	return products, result.Error
}
