package models

type User struct {
	Model
	Name  string `gorm:"type:varchar(50)" json:"name" validate:"required"`
	Email string `gorm:"type:varchar(50)" json:"email" validate:"required,email"`
}

func (u *User) TableName() string {
	return "user"
}
