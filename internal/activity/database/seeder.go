package database

import (
	"github.com/jinzhu/gorm"
	"github.com/sumelms/microservice-activity/pkg/seed"
)

func Activities() seed.Seed {
	return seed.Seed{
		Name: "CreateActivities",
		Run: func(db *gorm.DB) error {
			u := &Activity{}
			return db.Create(u).Error
		},
	}
}
