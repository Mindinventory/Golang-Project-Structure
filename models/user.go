package models

//User structure
type User struct {
	Model
	Name  string `gorm:"type:varchar(50)" json:"name" validate:"required"`
	Email string `gorm:"type:varchar(50)" json:"email" validate:"required,email"`
}

//TableName return name of database table
func (u *User) TableName() string {
	return "user"
}
