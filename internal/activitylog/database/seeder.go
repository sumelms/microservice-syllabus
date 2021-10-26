package database

import (
	"github.com/jinzhu/gorm"
	"github.com/sumelms/microservice-activity/pkg/seed"
)

func ActivityLogs() seed.Seed {
	return seed.Seed{
		Name: "CreateActivityLogs",
		Run: func(db *gorm.DB) error {
			u := &ActivityLog{}
			return db.Create(u).Error
		},
	}
}
