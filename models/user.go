package models

type User struct {
	ID       int64  `gorm:"primaryKey" json:"id"`
	FullName string `gorm:"varchar(300)" json:"fullName"`
	Username string `gorm:"varchar(300)" json:"username"`
	Password string `gorm:"varchar(300)" json:"password"`
}
