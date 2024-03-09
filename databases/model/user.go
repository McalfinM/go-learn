package model

type User struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	FullName string `gorm:"varchar(150)" json:"full_name"`
	Username string `gorm:"varchar(15)" json:"username"`
	Password string `gorm:"varchar(15)" json:"password"`
}
