package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:100;not null"`
	Email    string `gorm:"uniqueIndex;size:100;not null"`
	Password string `gorm:"size:255;not null"`
	Role     string `gorm:"type:varchar(20);default:'user'"`
}
