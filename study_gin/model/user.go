package model

// User gorm User结构体映射 users 所以要写TableName()方法来手动映射表名
type User struct {
	//gorm.Model
	Id       int8
	Username string
	Address  string
}

func (User) TableName() string {
	return "tb_user"
}
