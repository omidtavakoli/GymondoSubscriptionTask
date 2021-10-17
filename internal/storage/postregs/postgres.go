package postregs

import (
	"Gymondo/internal/subscription"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strconv"
	"time"
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

func (r *Repository) GetProducts() ([]subscription.Product, error) {
	var products []subscription.Product
	result := r.database.Find(&products)
	return products, result.Error
}

func (r *Repository) GetProduct(id int) (product subscription.Product, err error) {
	err = r.database.Where("id = ?", id).First(&product).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return product, err
	}
	return product, nil
}

func (r *Repository) CreatePlan(name string, price, discount, durationDays int, product subscription.Product) (uint64, error) {
	plan := subscription.Plan{
		Name:      name,
		Price:     price,
		Discount:  discount,
		Duration:  durationDays,
		ProductID: product.ID,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	err := r.database.Create(&plan).Error
	if err != nil {
		return 0, errors.Wrap(err, "failed to create a plan")
	}
	return plan.ID, nil
}

func (r *Repository) BuyProduct(bpr subscription.BuyRequest) (subscription.UserPlan, error) {
	userId, err := strconv.Atoi(bpr.UserId)
	if err != nil {
		return subscription.UserPlan{}, err
	}

	var plan subscription.Plan
	pErr := r.database.Joins("inner join products p on plans.product_id = p.id").Where("product_id=? AND plans.duration=-1", bpr.ProductId).Find(&plan).Error
	if pErr != nil {
		return subscription.UserPlan{}, pErr
	}
	//SELECT plans.id FROM "plans" inner join products p on plans.product_id = p.id WHERE product_id='22' AND plans.name='LifeTime' AND plans."deleted_at" IS NULL
	UserPlan := subscription.UserPlan{
		UserId:     userId,
		PlanId:     int(plan.ID),
		PlanStatus: "active",
		//Voucher:  0,
		StartDate: time.Now(),
		EndDate:   time.Now().Add(1000000 * time.Hour), // large number
		DeletedAt: gorm.DeletedAt{},
	}

	resp := r.database.FirstOrCreate(&UserPlan, subscription.UserPlan{UserId: userId, PlanId: int(plan.ID)})
	if resp.Error != nil {
		return UserPlan, resp.Error
	}
	return UserPlan, nil
}

func (r *Repository) FetchPlansByUserId(userId int) (plans []subscription.UserPlan, err error){
	err = r.database.Where("user_id = ?", userId).Find(&plans).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return plans, err
	}
	return plans, nil
}