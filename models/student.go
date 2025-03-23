package models

import "gorm.io/gorm"

type Student struct {
    ID    uint   `json:"id" gorm:"primaryKey"`
    Name  string `json:"name"`
    Class string `json:"class"`
    Photo string `json:"photo"`
}

func Migrate(db *gorm.DB) {
    db.AutoMigrate(&Student{})
}
