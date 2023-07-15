package models

type Locker struct {
	Id     int          `json:"id" gorm:"primaryKey"`
	Code   string       `json:"code" gorm:"unique,not null"`
	UserId int          `json:"user_id" gorm:"not null"`
	User   UserResponse `json:"user"`
}

type LockerResponse struct {
	Id     int    `json:"id"`
	Code   string `json:"code"`
	UserId int    `json:"-"`
}

func (LockerResponse) TableName() string {
	return "lockers"
}