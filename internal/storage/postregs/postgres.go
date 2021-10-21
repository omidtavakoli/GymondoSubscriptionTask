package postgres

import (
	"Gymondo/internal/subscription"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Repository struct {
	database *gorm.DB
}

var models = []interface{}{
	&subscription.User{},
	&subscription.UserPlan{},
	&subscription.Plan{},
	&subscription.Product{},
	&subscription.Voucher{},
	&subscription.VoucherPlan{},
}

func CreateRepository(db *gorm.DB) (*Repository, error) {
	repo := &Repository{
		database: db,
	}
	logrus.Infof("current db name: %s", db.Migrator().CurrentDatabase())
	err := db.AutoMigrate(models...)
	if err != nil {
		return repo, errors.Wrap(err, "failed to auto migrate models")
	}
	return repo, nil
}

//todo: receive user model instead of raw data
func (r *Repository) CreateUser(email, username, fullname string) (subscription.User, error) {
	mu := subscription.User{
		Email:    email,
		UserName: username,
		FullName: fullname,
	}
	err := r.database.Create(&mu).Error
	if err != nil {
		return mu, errors.Wrap(err, "failed to create a user")
	}
	return mu, nil
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

func (r *Repository) CreateVoucher(cvr subscription.CreateVoucherRequest) (subscription.Voucher, error) {
	sv := subscription.Voucher{
		Name:         cvr.Name,
		Discount:     cvr.Discount,
		DiscountType: cvr.DiscountType,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		DeletedAt:    gorm.DeletedAt{},
		StartDate:    cvr.StartDate,
		EndDate:      cvr.EndDate,
	}
	err := r.database.Create(&sv).Error
	if err != nil {
		return sv, errors.Wrap(err, "failed to create a voucher")
	}
	return sv, nil
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

func (r *Repository) GetPlans() ([]subscription.Plan, error) {
	var plans []subscription.Plan
	result := r.database.Find(&plans)
	return plans, result.Error
}
func (r *Repository) GetVouchers() ([]subscription.Voucher, error) {
	var vouchers []subscription.Voucher
	result := r.database.Find(&vouchers)
	return vouchers, result.Error
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

func (r *Repository) CreateVoucherPlan(plan subscription.Plan, voucher subscription.Voucher) (uint64, error) {
	voucherPlan := subscription.VoucherPlan{
		VoucherID: voucher.ID,
		PlanID:    plan.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: gorm.DeletedAt{},
	}
	err := r.database.Create(&voucherPlan).Error
	if err != nil {
		return 0, errors.Wrap(err, "failed to create a voucher plan")
	}
	return voucherPlan.ID, nil
}

func (r *Repository) BuyProduct(bpr subscription.BuyRequest) (subscription.UserPlan, error) {
	//todo: separate this func logics
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

	tax := taxCalculator(userId)

	UserPlan := subscription.UserPlan{
		UserId:     userId,
		PlanId:     int(plan.ID),
		PlanStatus: "pause",
		//Voucher:  0,
		Tax:       tax,
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

func (r *Repository) FetchPlansByUserId(userId int) (status []subscription.Status, err error) {
	return status, r.database.Raw(fmt.Sprintf("SELECT *  FROM plans inner join user_plans p on p.plan_id = plans.id WHERE p.user_id=%d", userId)).Scan(&status).Error
}

func (r *Repository) ChangeUserPlanStatus(status subscription.ChangeStatus) error {
	return r.database.Model(&subscription.UserPlan{}).Where("user_id=? AND plan_id=?", status.UserId, status.PlanId).Update("plan_status", status.Status).Error
}

func taxCalculator(userId int) int {
	//todo: calculate tax based on country
	return 10
}
