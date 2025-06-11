package internalsql

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func Connect(dataSourceName string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
