package dao

import "study_gin/model"

func migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model.User{})
	if err != nil {
		return
	}
}
