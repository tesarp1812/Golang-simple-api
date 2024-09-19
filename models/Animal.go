package models

import "gorm.io/gorm"

type Animal struct {
	Id    int    `json:"id" form:"primary_key`
	Name  string `json:"name"`
	Class string `json:"class"`
	Legs  int    `json:"legs"`
}

func CheckAnimalExists(db *gorm.DB, name string) (bool, error) {
	var count int64
	err := db.Model(&Animal{}).Where("name = ?", name).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
