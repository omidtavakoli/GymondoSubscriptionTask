package subscription

import (
	"gorm.io/gorm"
	"time"
)

type Config struct {
	gorm.Model
	Service        string        `yaml:"SERVICE"`
	DefaultTTL     int64         `yaml:"DEFAULT_TTL"`
	CacheDisabled  bool          `yaml:"CACHE_DISABLED"`
	SleepTime      time.Duration `yaml:"SLEEP_TIME"`
	FetchBatchSize int           `yaml:"FETCH_BATCH_SIZE"`
}

type User struct {
	gorm.Model
	ID        uint64 `gorm:"primary_key"`
	Email     string `gorm:"not null"`
	UserName  string
	FullName  string
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"default:current_timestamp"`
	DeletedAt gorm.DeletedAt
}

type Product struct {
	gorm.Model
	ID   uint64 `gorm:"primary_key"`
	Name string
	//Plans []Plan `gorm:"foreignKey:ID"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"default:current_timestamp"`
	DeletedAt gorm.DeletedAt
}

type BuyRequest struct {
	UserId       string `form:"userId" binding:"required"`
	ProductId    string `form:"productId" binding:"required"`
	VoucherId    int    `form:"voucherId"`
	TrialRequest bool   `form:"trialReq"`
}

type ChangeStatus struct {
	UserId string `form:"userId" binding:"required"`
	PlanId string `form:"planId" binding:"required"`
	Status string `form:"status" binding:"oneof=active pause cancel"`
}

type Plan struct {
	gorm.Model
	ID        uint64 `gorm:"primary_key"`
	Name      string
	Price     int
	Discount  int
	Duration  int       //days
	ProductID uint64    `gorm:"foreignKey:ID"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"default:current_timestamp"`
	DeletedAt gorm.DeletedAt
}

type UserPlan struct {
	gorm.Model
	ID                  uint64 `gorm:"primary_key"`
	UserId              int
	PlanId              int
	PlanStatus          string `gorm:"default:active"`
	Voucher             uint64 `gorm:"foreignKey:ID"`
	VoucherDiscount     int
	VoucherDiscountType string
	Tax                 int
	StartDate           time.Time `gorm:"default:current_timestamp"`
	EndDate             time.Time `gorm:"default:current_timestamp"`
	DeletedAt           gorm.DeletedAt
}

type Voucher struct {
	gorm.Model
	ID           uint64 `gorm:"primary_key"`
	Name         string
	Discount     int
	DiscountType string
	valid        bool
	CreatedAt    time.Time `gorm:"default:current_timestamp"`
	UpdatedAt    time.Time `gorm:"default:current_timestamp"`
	DeletedAt    gorm.DeletedAt
	StartDate    time.Time `gorm:"default:current_timestamp"`
	EndDate      time.Time `gorm:"default:current_timestamp"`
}

type VoucherPlan struct {
	gorm.Model
	ID        uint64    `gorm:"primary_key"`
	VoucherID uint64    `gorm:"foreignKey:ID"`
	PlanID    uint64    `gorm:"foreignKey:ID"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"default:current_timestamp"`
	DeletedAt gorm.DeletedAt
}

type Status struct {
	PlanId              int    `json:"plan_id"`
	PlanDuration        int    `json:"plan_duration"`
	PlanProduct         int    `json:"plan_product"`
	PlanDiscount        int    `json:"plan_discount"`
	PlanPrice           int    `json:"plan_price"`
	PlanStatus          string `json:"plan_status"`
	PlanStartDate       string `json:"plan_start_date"`
	PlanEndDate         string `json:"plan_end_date"`
	PlanTax             int    `json:"plan_tax"`
	VoucherDiscount     int    `json:"voucher_discount"`
	VoucherDiscountType string `json:"voucher_discount_type"`
}

type VoucherPlanProduct struct {
	PlanId              int    `json:"plan_id"`
	PlanDuration        int    `json:"plan_duration"`
	PlanProduct         int    `json:"plan_product"`
	PlanDiscount        int    `json:"plan_discount"`
	PlanPrice           int    `json:"plan_price"`
	VoucherId           int    `json:"voucher_id"`
	VoucherDiscount     int    `json:"voucher_discount"`
	VoucherDiscountType string `json:"voucher_discount_type"`
	ProductName         string `json:"product_name"`
}

type CreateVoucherRequest struct {
	Name         string
	PlanId       uint64
	Discount     int
	DiscountType string `form:"type" binding:"oneof=percent amount"`
	StartDate    time.Time
	EndDate      time.Time
}
