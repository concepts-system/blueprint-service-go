package migrations

import (
	"github.com/concepts-system/blueprint-service-go/users"
	"github.com/jinzhu/gorm"
	gormigrate "gopkg.in/gormigrate.v1"
)

var migrationV1 = gormigrate.Migration{
	ID: "1",
	Migrate: func(tx *gorm.DB) error {
		// Users
		if err := tx.AutoMigrate(&users.UserModel{}).Error; err != nil {
			return err
		}

		return nil
	},
	Rollback: func(tx *gorm.DB) error {
		// Users
		if err := tx.DropTable(users.UserModel{}.TableName()).Error; err != nil {
			return err
		}

		return nil
	},
}
