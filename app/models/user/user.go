package user

import "goblog/app/models"

type User struct {
	models.BaseModel

	Name     string `gorm:"type:varchar(255);not nul;unique" valid:"name"`
	Email    string `gorm:"type:varchar(255);unique" valid:"email"`
	Password string `gorm:"type:varchar(255)" valid:"password"`

	PasswordConfirm string `gorm:"-" valid:"password_confirm"`
}
