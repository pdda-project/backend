package utils

import (
	"fmt"

	"gorm.io/gorm"
)

func CountRecord(db *gorm.DB, model any, col string, value string) (int64, error) {
	var count int64
	err := db.Model(model).Where(fmt.Sprintf("%s = ?", col), value).Count(&count).Error
	return count, err
}
