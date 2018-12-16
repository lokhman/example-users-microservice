package model

const UserEmailUniqueConstraintName = "uix_users_email"

// User model structure.
// We don't use gorm.Model as it doesn't nicely translate JSON fields.
type User struct {
	ID        int    `gorm:"primary_key" json:"id" example:"1"`
	Email     string `gorm:"type:varchar(128); unique_index; not null" json:"email" example:"alex.lokhman@gmail.com"`
	Password  string `gorm:"type:char(60); not null" json:"-" example:"MyPassword"`
	FirstName string `gorm:"type:varchar(72); not null" json:"first_name" example:"Alex"`
	LastName  string `gorm:"type:varchar(72); not null" json:"last_name" example:"Lokhman"`
	Nickname  string `gorm:"type:varchar(32); not null" json:"nickname" example:"VisioN"`
	Country   string `gorm:"type:char(2); not null" json:"country" example:"RU"`
}
