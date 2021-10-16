package postregs

import (
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

	//rawDB, err := db.DB()
	//if err != nil {
	//	return repo, err
	//}

	// Enable UUID Extensions
	//_, err = rawDB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public`)
	//if err != nil {
	//	return repo, err
	//}

	//err = db.AutoMigrate(models...)
	//if err != nil {
	//	return repo,err
	//}
	return repo, nil
}

func (m *Repository) GetUser(email string) error {
	var u user
	err := m.database.Where("email = ?", email).First(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}
