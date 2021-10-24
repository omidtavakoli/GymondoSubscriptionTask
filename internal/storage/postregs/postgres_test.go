package postgres

import (
	"Gymondo/internal/subscription"
	"Gymondo/platform/postgres"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"reflect"
	"testing"
	"time"
)

var pgCfg = postgres.Config{
	HOST:     "localhost",
	PORT:     5432,
	NAME:     "gymondo_test",
	USER:     "test",
	PASSWORD: "123456",
	SSLMODE:  "disable",
	DEBUG:    true,
}

type postgresSuit struct {
	suite.Suite
	database *gorm.DB
}

func (p *postgresSuit) SetupTest() {
	err := p.database.Migrator().DropTable(models...)
	if err != nil {
		logrus.Fatal(err)
	}

	err = p.database.AutoMigrate(models...)
	if err != nil {
		logrus.Fatal(err)
	}
}

func (p *postgresSuit) TearDownTest() {
	err := p.database.Migrator().DropTable(models...)
	if err != nil {
		logrus.Fatal(err)
	}
}

func TestPostgresRepository(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip test for postgres repository")
	}

	config := &gorm.Config{}

	if pgCfg.DEBUG {
		config.Logger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second,
				Colorful:      true,
				LogLevel:      logger.Info,
			},
		)
	}

	connection := postgres.CreateConnection(pgCfg, "test")

	database, err := connection.OpenGORM()

	if err != nil {
		log.Fatal(err)
	}

	suite.Run(t, &postgresSuit{
		database: database,
	})
}

func (p *postgresSuit) TestRepositoryCreate() {
	email := "omdtvk@gmail.com"
	u := subscription.User{
		ID:        0,
		Email:     email,
		UserName:  "omidtavakoli",
		FullName:  "Omid Tavakoli",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: gorm.DeletedAt{},
	}
	repository, err := CreateRepository(p.database)
	if err != nil {
		p.T().Fatalf("cannout create repository, err: %s", err)
	}

	cu, err := repository.CreateUser(u)
	if err != nil {
		p.T().Fatalf("we dont expected any error but got %s", err)
	}

	got := subscription.User{}
	p.database.Where("id = ?", cu.ID).Find(&got)
	if got.Email != email {
		p.T().Fatalf("we expected %s as registered email by got %s", email, got.Email)
	}
}

func (p *postgresSuit) TestRepository_GetUserByEmail() {
	tests := []struct {
		description   string
		email         string
		expectedUser  subscription.User
		expectedError error
	}{
		{
			"if email address is not found in database, the repository returns a not found error",
			"invalid_address@email.com",
			subscription.User{},
			gorm.ErrRecordNotFound,
		},
		{
			"if email address is found, the repository returns it with no error",
			"t1@email.com",
			subscription.User{
				ID:       2,
				Email:    "t1@email.com",
				FullName: "t1",
			},
			nil,
		},
	}

	defaultFixtures := []subscription.User{
		{
			ID:       2,
			Email:    "t1@email.com",
			FullName: "t1",
		},
		{
			ID:       3,
			Email:    "t2@email.com",
			FullName: "t2",
		},
	}

	for _, f := range defaultFixtures {
		if err := p.database.Create(&f).Error; err != nil {
			p.T().Fatalf("cannot add default fixture %v with error %s", f, err)
		}
	}

	repository, err := CreateRepository(p.database)
	if err != nil {
		p.T().Fatalf("cannout create repository, err: %s", err)
	}

	for _, test := range tests {
		p.Run(test.description, func() {
			got, err := repository.GetUserByEmail(test.email)
			if !p.Equal(err, test.expectedError) {
				p.T().Fatalf("we expected %s as error but got %s", test.expectedError, err)
			}
			got.CreatedAt = test.expectedUser.CreatedAt
			got.UpdatedAt = test.expectedUser.UpdatedAt
			if !reflect.DeepEqual(got, test.expectedUser) {
				p.T().Fatalf("we expected %v as registred user but got %v", test.expectedUser, got)
			}
		})
	}

}
