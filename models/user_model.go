package models

type User struct {
	Id     int            `json:"id" gorm:"primaryKey"`
	Name   string         `json:"name" gorm:"not null"`
	Locker LockerResponse `json:"locker"`
	Posts  []PostResponse `json:"Posts"`
}

type UserResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (UserResponse) TableName() string {
	return "users"
}